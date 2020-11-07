package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("kratos_id"),
		field.String("name"),
		field.String("last_name"),
		field.String("email"),
		field.String("phone"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("address", Address.Type).
			Unique(),
		edge.To("ratings", Rating.Type),
	}
}

// Indexes of the Customer
func (Customer) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("kratos_id", "email").
			Unique(),
	}
}
