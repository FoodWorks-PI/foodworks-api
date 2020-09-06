package db

import "github.com/facebook/ent"

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return nil
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return nil
}
