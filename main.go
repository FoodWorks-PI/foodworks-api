package main

import (
	"log"
	"os"

	"foodworks.ml/m/internal/recommendations"

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
		RecommenderURL:     os.Getenv("RECOMMENDER_URL"),
	}

	elasticClient := platform.NewElasticSearchClient(dataStoreConfig)
	db, client := platform.NewEntClient(dataStoreConfig)

	rdb := platform.NewRedisClient(dataStoreConfig)

	fileHandler := filehandler.NewDiskUploader()
	recommender := recommendations.NewRecommender(dataStoreConfig.RecommenderURL)

	api := api.API{}
	api.SetupRoutes(client, db, rdb, dataStoreConfig, elasticClient, &fileHandler, recommender)
	if os.Getenv("DUMMY_PAYMENTS") == "enabled" {
		api.StartAutoPayment(client, db)
	}
	api.StartServer()
	// Cleanup
	log.Println("Closing")
	db.Close()
	rdb.Close()
	client.Close()
}
