package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("latitude"),
		field.String("longitude"),
		field.String("Street"),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return nil
}
