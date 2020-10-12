package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// RestaurantOwner holds the schema definition for the RestaurantOwner entity.
type RestaurantOwner struct {
	ent.Schema
}

// Fields of the RestaurantOwner.
func (RestaurantOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("kratos_id"),
		field.String("name"),
		field.String("email"),
		field.String("phone"),
	}
}

// Edges of the RestaurantOwner.
func (RestaurantOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("banking_data", BankingData.Type).
			Unique(),
		edge.To("restaurant", Restaurant.Type).Unique(),
	}
}

// Index of RestaurantOwner
func (RestaurantOwner) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("kratos_id", "email").
			Unique(),
	}
}
