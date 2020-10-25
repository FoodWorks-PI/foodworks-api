package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Rating holds the schema definition for the Rating entity.
type Rating struct {
	ent.Schema
}

// Fields of the Rating.
func (Rating) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ProductRate"),
		field.Int("ProductID"),
		field.Int("CustomerID"),
	}
}

// Edges of the Rating.
func (Rating) Edges() []ent.Edge {
	return nil
}
