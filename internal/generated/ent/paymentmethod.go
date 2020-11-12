// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent/paymentmethod"
	"github.com/facebook/ent/dialect/sql"
)

// PaymentMethod is the model entity for the PaymentMethod schema.
type PaymentMethod struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Data holds the value of the "data" field.
	Data                    string `json:"data,omitempty"`
	customer_payment_method *int
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PaymentMethod) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // data
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*PaymentMethod) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // customer_payment_method
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentMethod fields.
func (pm *PaymentMethod) assignValues(values ...interface{}) error {
	if m, n := len(values), len(paymentmethod.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pm.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field data", values[0])
	} else if value.Valid {
		pm.Data = value.String
	}
	values = values[1:]
	if len(values) == len(paymentmethod.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customer_payment_method", value)
		} else if value.Valid {
			pm.customer_payment_method = new(int)
			*pm.customer_payment_method = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this PaymentMethod.
// Note that, you need to call PaymentMethod.Unwrap() before calling this method, if this PaymentMethod
// was returned from a transaction, and the transaction was committed or rolled back.
func (pm *PaymentMethod) Update() *PaymentMethodUpdateOne {
	return (&PaymentMethodClient{config: pm.config}).UpdateOne(pm)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pm *PaymentMethod) Unwrap() *PaymentMethod {
	tx, ok := pm.config.driver.(*txDriver)
	if !ok {
		panic("ent: PaymentMethod is not a transactional entity")
	}
	pm.config.driver = tx.drv
	return pm
}

// String implements the fmt.Stringer.
func (pm *PaymentMethod) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentMethod(")
	builder.WriteString(fmt.Sprintf("id=%v", pm.ID))
	builder.WriteString(", data=")
	builder.WriteString(pm.Data)
	builder.WriteByte(')')
	return builder.String()
}

// PaymentMethods is a parsable slice of PaymentMethod.
type PaymentMethods []*PaymentMethod

func (pm PaymentMethods) config(cfg config) {
	for _i := range pm {
		pm[_i].config = cfg
	}
}