package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").
			Optional().
			Nillable().
			SchemaType(map[string]string{
				dialect.Postgres: "zdb.fulltext",
			}),
		field.Int("cost"),
		field.Bool("is_active"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tag.Type),
		edge.To("ratings", Rating.Type),
		edge.From("restaurant", Restaurant.Type).Ref("products"),
		edge.To("orders", Order.Type),
		edge.To("images", ImagePath.Type),
	}
}
