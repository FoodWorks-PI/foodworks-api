package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type ImagePath struct {
	ent.Schema
}

// Fields of the Tag.
func (ImagePath) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
	}
}

// Edges of the Tag.
func (ImagePath) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).Ref("images").Unique(),
		edge.From("restaurant", Restaurant.Type).Ref("images").Unique(),
		edge.From("tag", Tag.Type).Ref("images").Unique(),
	}
}
