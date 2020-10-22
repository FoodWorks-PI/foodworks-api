package resolver

import (
	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/platform/filehandler"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EntClient     *ent.Client
	Redis         *redis.Client
	DBClient      *sqlx.DB
	ElasticClient *elasticsearch.Client
	FileHandler   *filehandler.FileHandler
}
