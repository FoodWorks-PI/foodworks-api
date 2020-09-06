package db

import "github.com/facebook/ent"

// Restaurant holds the schema definition for the Restaurant entity.
type Restaurant struct {
	ent.Schema
}

// Fields of the Restaurant.
func (Restaurant) Fields() []ent.Field {
	return nil
}

// Edges of the Restaurant.
func (Restaurant) Edges() []ent.Edge {
	return nil
}
