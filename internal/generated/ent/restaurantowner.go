// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent/bankingdata"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"
	"github.com/facebook/ent/dialect/sql"
)

// RestaurantOwner is the model entity for the RestaurantOwner schema.
type RestaurantOwner struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// KratosID holds the value of the "kratos_id" field.
	KratosID string `json:"kratos_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RestaurantOwnerQuery when eager-loading is set.
	Edges                         RestaurantOwnerEdges `json:"edges"`
	restaurant_owner_banking_data *int
}

// RestaurantOwnerEdges holds the relations/edges for other nodes in the graph.
type RestaurantOwnerEdges struct {
	// BankingData holds the value of the banking_data edge.
	BankingData *BankingData
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BankingDataOrErr returns the BankingData value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RestaurantOwnerEdges) BankingDataOrErr() (*BankingData, error) {
	if e.loadedTypes[0] {
		if e.BankingData == nil {
			// The edge banking_data was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bankingdata.Label}
		}
		return e.BankingData, nil
	}
	return nil, &NotLoadedError{edge: "banking_data"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RestaurantOwner) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // kratos_id
		&sql.NullString{}, // name
		&sql.NullString{}, // email
		&sql.NullString{}, // phone
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*RestaurantOwner) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // restaurant_owner_banking_data
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RestaurantOwner fields.
func (ro *RestaurantOwner) assignValues(values ...interface{}) error {
	if m, n := len(values), len(restaurantowner.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ro.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field kratos_id", values[0])
	} else if value.Valid {
		ro.KratosID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		ro.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		ro.Email = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[3])
	} else if value.Valid {
		ro.Phone = value.String
	}
	values = values[4:]
	if len(values) == len(restaurantowner.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field restaurant_owner_banking_data", value)
		} else if value.Valid {
			ro.restaurant_owner_banking_data = new(int)
			*ro.restaurant_owner_banking_data = int(value.Int64)
		}
	}
	return nil
}

// QueryBankingData queries the banking_data edge of the RestaurantOwner.
func (ro *RestaurantOwner) QueryBankingData() *BankingDataQuery {
	return (&RestaurantOwnerClient{config: ro.config}).QueryBankingData(ro)
}

// Update returns a builder for updating this RestaurantOwner.
// Note that, you need to call RestaurantOwner.Unwrap() before calling this method, if this RestaurantOwner
// was returned from a transaction, and the transaction was committed or rolled back.
func (ro *RestaurantOwner) Update() *RestaurantOwnerUpdateOne {
	return (&RestaurantOwnerClient{config: ro.config}).UpdateOne(ro)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ro *RestaurantOwner) Unwrap() *RestaurantOwner {
	tx, ok := ro.config.driver.(*txDriver)
	if !ok {
		panic("ent: RestaurantOwner is not a transactional entity")
	}
	ro.config.driver = tx.drv
	return ro
}

// String implements the fmt.Stringer.
func (ro *RestaurantOwner) String() string {
	var builder strings.Builder
	builder.WriteString("RestaurantOwner(")
	builder.WriteString(fmt.Sprintf("id=%v", ro.ID))
	builder.WriteString(", kratos_id=")
	builder.WriteString(ro.KratosID)
	builder.WriteString(", name=")
	builder.WriteString(ro.Name)
	builder.WriteString(", email=")
	builder.WriteString(ro.Email)
	builder.WriteString(", phone=")
	builder.WriteString(ro.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// RestaurantOwners is a parsable slice of RestaurantOwner.
type RestaurantOwners []*RestaurantOwner

func (ro RestaurantOwners) config(cfg config) {
	for _i := range ro {
		ro[_i].config = cfg
	}
}