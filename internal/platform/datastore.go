package platform

import (
	"context"
	"fmt"
	"log"

	"foodworks.ml/m/internal/generated/ent"
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
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error %v", err)
	}
	_, err = tx.Exec(`
		ALTER TABLE addresses
		ADD COLUMN IF NOT EXISTS geom geometry(POINT);
	`)
	if err != nil {
		log.Printf("Error %v", err)
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
		log.Printf("Error %v", err)
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
		log.Printf("Error %v", err)
	}
	var stmt string
	collections := []string{restaurant.Table, product.Table, tag.Table}
	for _, collection := range collections {
		stmt = fmt.Sprintf(`
		CREATE INDEX IF NOT EXISTS %s_zombo_idx
		ON %s
		USING zombodb ((%s.*))
		WITH (url='%s')`, collection, collection, collection, config.ElasticsearchDBURL)
		_, err = tx.Exec(stmt)
		if err != nil {
			log.Printf("Error %v", err)
		}
	}
	err = tx.Commit()
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
