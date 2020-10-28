package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Rating holds the schema definition for the Rating entity.
type Rating struct {
	ent.Schema
}

// Fields of the Rating.
func (Rating) Fields() []ent.Field {
	return []ent.Field{
		field.String("comment"),
		field.Int("rating"),
	}
}

// Edges of the Rating.
func (Rating) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("ratings").Unique(),
		edge.From("product", Product.Type).Ref("ratings").Unique(),
	}
}
