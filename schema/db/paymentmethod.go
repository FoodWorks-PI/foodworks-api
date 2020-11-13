package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// PaymentMethod holds the schema definition for the PaymentMethod entity.
type PaymentMethod struct {
	ent.Schema
}

// Fields of the PaymentMethod.
func (PaymentMethod) Fields() []ent.Field {
	return []ent.Field{
		field.String("data"),
	}
}

// Edges of the PaymentMethod.
func (PaymentMethod) Edges() []ent.Edge {
	return nil
}
