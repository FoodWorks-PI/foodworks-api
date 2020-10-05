package db

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// BankingData holds the schema definition for the BankingData entity.
type BankingData struct {
	ent.Schema
}

// Fields of the BankingData.
func (BankingData) Fields() []ent.Field {
	return []ent.Field{
		field.String("bank_account"),
	}
}

// Edges of the BankingData.
func (BankingData) Edges() []ent.Edge {
	return nil
}
