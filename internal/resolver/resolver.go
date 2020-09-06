package resolver

import (
	"foodworks.ml/m/internal/generated/ent"
	"github.com/go-redis/redis/v8"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client *ent.Client
	Redis  *redis.Client
}
