package platform

import (
	"context"
	"database/sql"
	"log"

	"foodworks.ml/m/internal/generated/ent"
	"github.com/facebook/ent/dialect"
	entsql "github.com/facebook/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DataStoreConfig struct {
	DatabaseURL string
	RedisAddr   string
	RedisPass   string
}

// Open new db connection
func NewEntClient(config DataStoreConfig) (*sql.DB, *ent.Client) {
	db, err := sql.Open("pgx", config.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Printf("failed creating schema resources: %v", err)
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
