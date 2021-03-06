// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/customer"
	"github.com/facebook/ent/dialect/sql"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// KratosID holds the value of the "kratos_id" field.
	KratosID string `json:"kratos_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges            CustomerEdges `json:"edges"`
	customer_address *int
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Address holds the value of the address edge.
	Address *Address
	// Ratings holds the value of the ratings edge.
	Ratings []*Rating
	// Orders holds the value of the orders edge.
	Orders []*Order
	// PaymentMethod holds the value of the payment_method edge.
	PaymentMethod []*PaymentMethod
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// AddressOrErr returns the Address value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) AddressOrErr() (*Address, error) {
	if e.loadedTypes[0] {
		if e.Address == nil {
			// The edge address was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: address.Label}
		}
		return e.Address, nil
	}
	return nil, &NotLoadedError{edge: "address"}
}

// RatingsOrErr returns the Ratings value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) RatingsOrErr() ([]*Rating, error) {
	if e.loadedTypes[1] {
		return e.Ratings, nil
	}
	return nil, &NotLoadedError{edge: "ratings"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[2] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// PaymentMethodOrErr returns the PaymentMethod value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) PaymentMethodOrErr() ([]*PaymentMethod, error) {
	if e.loadedTypes[3] {
		return e.PaymentMethod, nil
	}
	return nil, &NotLoadedError{edge: "payment_method"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // kratos_id
		&sql.NullString{}, // name
		&sql.NullString{}, // last_name
		&sql.NullString{}, // email
		&sql.NullString{}, // phone
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Customer) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // customer_address
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(values ...interface{}) error {
	if m, n := len(values), len(customer.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field kratos_id", values[0])
	} else if value.Valid {
		c.KratosID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		c.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field last_name", values[2])
	} else if value.Valid {
		c.LastName = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[3])
	} else if value.Valid {
		c.Email = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[4])
	} else if value.Valid {
		c.Phone = value.String
	}
	values = values[5:]
	if len(values) == len(customer.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customer_address", value)
		} else if value.Valid {
			c.customer_address = new(int)
			*c.customer_address = int(value.Int64)
		}
	}
	return nil
}

// QueryAddress queries the address edge of the Customer.
func (c *Customer) QueryAddress() *AddressQuery {
	return (&CustomerClient{config: c.config}).QueryAddress(c)
}

// QueryRatings queries the ratings edge of the Customer.
func (c *Customer) QueryRatings() *RatingQuery {
	return (&CustomerClient{config: c.config}).QueryRatings(c)
}

// QueryOrders queries the orders edge of the Customer.
func (c *Customer) QueryOrders() *OrderQuery {
	return (&CustomerClient{config: c.config}).QueryOrders(c)
}

// QueryPaymentMethod queries the payment_method edge of the Customer.
func (c *Customer) QueryPaymentMethod() *PaymentMethodQuery {
	return (&CustomerClient{config: c.config}).QueryPaymentMethod(c)
}

// Update returns a builder for updating this Customer.
// Note that, you need to call Customer.Unwrap() before calling this method, if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return (&CustomerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Customer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", kratos_id=")
	builder.WriteString(c.KratosID)
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteString(", last_name=")
	builder.WriteString(c.LastName)
	builder.WriteString(", email=")
	builder.WriteString(c.Email)
	builder.WriteString(", phone=")
	builder.WriteString(c.Phone)
	builder.WriteByte(')')
	return builder.String()
}

// Customers is a parsable slice of Customer.
type Customers []*Customer

func (c Customers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
