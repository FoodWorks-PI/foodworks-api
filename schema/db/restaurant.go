package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Restaurant holds the schema definition for the Restaurant entity.
type Restaurant struct {
	ent.Schema
}

// Fields of the Restaurant.
func (Restaurant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Restaurant.
func (Restaurant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type).
			Unique(),
	}
}
