package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			SchemaType(map[string]string{
				dialect.Postgres: "my_completion",
			}),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("tags"),
		edge.From("restaurant", Restaurant.Type).Ref("tags"),
	}
}
