package main

import (
	"log"
	"os"

	"foodworks.ml/m/internal/platform/filehandler"

	"foodworks.ml/m/internal/api"
	"foodworks.ml/m/internal/platform"
)

func main() {
	platform.SetupLog()

	dataStoreConfig := platform.DataStoreConfig{
		DatabaseURL:        os.Getenv("POSTGRES_URL"),
		JWKSURL:            os.Getenv("JWKS_URL"),
		RedisAddr:          os.Getenv("REDIS_ADDR"),
		RedisPass:          os.Getenv("REDIS_PASS"),
		ElasticsearchDBURL: os.Getenv("ELASTICSEARCH_DB_URL"),
		ElasticsearchURL:   os.Getenv("ELASTICSEARCH_URL"),
	}

	elasticClient := platform.NewElasticSearchClient(dataStoreConfig)
	db, client := platform.NewEntClient(dataStoreConfig)

	rdb := platform.NewRedisClient(dataStoreConfig)

	fileHandler := filehandler.NewDiskUploader()

	api := api.API{}
	api.SetupRoutes(client, db, rdb, dataStoreConfig, elasticClient, &fileHandler)
	api.StartServer()

	// Cleanup
	log.Println("Closing")
	db.Close()
	rdb.Close()
	client.Close()
}
