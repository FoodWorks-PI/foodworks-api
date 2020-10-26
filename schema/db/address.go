package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Float("latitude"),
		field.Float("longitude"),
		field.String("streetLine"),
		field.String("geom").
			Optional().
			Nillable().
			SchemaType(map[string]string{
				dialect.Postgres: "geometry(POINT)",
			}),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return nil
}
