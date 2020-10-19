package platform

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/product"
	"foodworks.ml/m/internal/generated/ent/restaurant"
	"foodworks.ml/m/internal/generated/ent/tag"
	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DataStoreConfig struct {
	DatabaseURL        string
	JWKSURL            string
	RedisAddr          string
	RedisPass          string
	ElasticsearchDBURL string
	ElasticsearchURL   string
}

func NewElasticSearchClient(config DataStoreConfig) *elasticsearch6.Client {
	cfg := elasticsearch6.Config{
		Addresses: []string{
			config.ElasticsearchURL,
		},
	}
	client, err := elasticsearch6.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	/*
		jsonObj := gabs.New()
		jsonObj.SetP(1, "suggest.suggestion.completion.fuzzy.fuzziness")
		jsonObj.SetP("name", "suggest.suggestion.completion.field")
		jsonObj.SetP("break", "suggest.suggestion.prefix")
		body := jsonObj.String()

		results, err := client.Search(
			client.Search.WithBody(strings.NewReader(body)),
		)
		if err != nil {
			log.Println(err)
		}
		log.Println(results)
		defer results.Body.Close()
		parsed, err := gabs.ParseJSONBuffer(results.Body)
		if err != nil {
			log.Println(err)
		}
		exists := parsed.ExistsP("suggest.suggestion.0.options")
		tmp := parsed.Path("suggest.suggestion.0.options").Data()
		log.Println(exists)
		log.Println(tmp)
		for _, child := range parsed.Path("suggest.suggestion.0.options").Children() {
			fmt.Println(child.S("_source").String())
			// fmt.Printf("key: %v, value: %v\n", key, child.Data().(float64))
		}
		log.Println("Done")
	*/
	return client
}

func prepareZombo(tx *sqlx.DB) error {
	stmt := `SELECT typname, nspname FROM pg_catalog.pg_type JOIN pg_catalog.pg_namespace ON pg_namespace.oid = pg_type.typnamespace WHERE typtype = 'd' and typname='my_completion';`
	rows, err := tx.Exec(stmt)
	if err != nil {
		return err
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		stmt = `
		create domain my_completion as text;
		SELECT zdb.define_type_mapping('my_completion'::regtype, '{
          "type": "completion"
        }'::json);
		`
		_, err = tx.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func seed(db *sqlx.DB) error {
	var err error
	rows, err := db.Exec(`SELECT * from seed_status`)
	if err != nil && err.Error() != `ERROR: relation "seed_status" does not exist (SQLSTATE 42P01)` {
		return err
	}
	if err == nil {
		rowsAffected, _ := rows.RowsAffected()
		if rowsAffected > 0 {
			return nil
		}
	}
	_, err = db.Exec(`
	CREATE TABLE seed_status(
	id int GENERATED ALWAYS AS IDENTITY
	);
	INSERT INTO seed_status default values;`)
	if err != nil {
		return err
	}
	tables := []string{tag.Table, address.Table, product.Table, restaurant.Table, restaurant.TagsTable, product.TagsTable, restaurant.ProductsTable}
	var stmt string
	for _, table := range tables {
		stmt = fmt.Sprintf(`COPY %s FROM '/foodworks/seed/%s.csv' DELIMITER ',' CSV HEADER;`, table, table)
		_, err := db.Exec(stmt)
		if err != nil {
			fmt.Print(err)
		}
	}
	return err
}

func preparePostgis(tx *sql.Tx) error {
	_, err := tx.Exec(`
		ALTER TABLE addresses
		ADD COLUMN IF NOT EXISTS geom geometry(POINT);
	`)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
		CREATE OR REPLACE FUNCTION insert_coordinates()
		RETURNS TRIGGER
		LANGUAGE PLPGSQL
		AS
		$$
		BEGIN
			IF (NEW.longitude <> OLD.longitude) OR (NEW.latitude <> OLD.latitude) OR (TG_OP = 'INSERT') THEN
				update addresses
					set geom = st_makepoint(new.longitude, new.latitude)
				where id=NEW.id;
			END IF;

			RETURN NEW;
		END;
		$$;
	`)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`
	DO
	$$
	BEGIN
		IF NOT EXISTS(select trigger_name from information_schema.triggers where trigger_name = 'coordinate_change') THEN
			CREATE TRIGGER coordinate_change
			AFTER INSERT OR UPDATE
			ON addresses
			FOR EACH ROW
			EXECUTE PROCEDURE insert_coordinates();
		END IF;
	END;
	$$;
	`)
	if err != nil {
		return err
	}
	return nil

}

func createZomboIndexes(tx *sql.Tx, config DataStoreConfig) error {
	var stmt string
	collections := []string{restaurant.Table, product.Table, tag.Table}
	for _, collection := range collections {
		stmt = fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS %s_zombo_idx
		ON %s
		USING zombodb ((%s.*))
		WITH (url='%s')`, collection, collection, collection, config.ElasticsearchDBURL)
		_, err := tx.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

// Open new db connection
func NewEntClient(config DataStoreConfig) (*sqlx.DB, *ent.Client) {
	//db, err := sql.Open("pgx", config.DatabaseURL)
	db, err := sqlx.Connect("pgx", config.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db.DB)
	client := ent.NewClient(ent.Driver(drv))
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error %v", err)
	}
	err = prepareZombo(db)
	if err != nil {
		log.Printf("Error %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}
	err = preparePostgis(tx)
	if err != nil {
		log.Printf("Error %v", err)
	}

	err = createZomboIndexes(tx, config)
	if err != nil {
		log.Printf("Error %v", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("Error %v", err)
	}
	err = seed(db)

	if err != nil {
		log.Printf("Error %v", err)
	}
	return db, client
}

func NewRedisClient(config DataStoreConfig) *redis.Client {

	// init redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPass, // no password set
		DB:       0,                // use default DB
	})

	return rdb
}
