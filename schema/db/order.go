package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Order struct {
	ent.Schema
}

// Fields of the Product.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.String("order_state"),
		field.Int("quantity"),
		field.Int64("updated_at").
			SchemaType(map[string]string{
				dialect.Postgres: "timestampz",
			}),
	}
}

// Edges of the Product.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("orders"),
		edge.From("customer", Customer.Type).Ref("orders"),
	}
}
