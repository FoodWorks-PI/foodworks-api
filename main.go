package main

import (
	"log"
	"os"

	"foodworks.ml/m/internal/api"
	"foodworks.ml/m/internal/platform"
)

func main() {
	platform.SetupLog()

	dataStoreConfig := platform.DataStoreConfig{
		DatabaseURL: os.Getenv("POSTGRES_URL"),
		JWKSURL:     os.Getenv("JWKS_URL"),
		RedisAddr:   os.Getenv("REDIS_ADDR"),
		RedisPass:   os.Getenv("REDIS_PASS"),
	}

	db, client := platform.NewEntClient(dataStoreConfig)

	rdb := platform.NewRedisClient(dataStoreConfig)

	api := api.API{}
	api.SetupRoutes(client, rdb, dataStoreConfig)
	api.StartServer()

	// Cleanup
	log.Println("Closing")
	db.Close()
	rdb.Close()
	client.Close()
}
