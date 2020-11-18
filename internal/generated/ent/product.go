// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent/product"
	"github.com/facebook/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Cost holds the value of the "cost" field.
	Cost int `json:"cost,omitempty"`
	// IsActive holds the value of the "is_active" field.
	IsActive bool `json:"is_active,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductQuery when eager-loading is set.
	Edges ProductEdges `json:"edges"`

	// StaticField defined by templates.
	Distance float64 `json:"distance,omitempty"`
}

// ProductEdges holds the relations/edges for other nodes in the graph.
type ProductEdges struct {
	// Tags holds the value of the tags edge.
	Tags []*Tag
	// Ratings holds the value of the ratings edge.
	Ratings []*Rating
	// Restaurant holds the value of the restaurant edge.
	Restaurant []*Restaurant
	// Orders holds the value of the orders edge.
	Orders []*Order
	// Images holds the value of the images edge.
	Images []*ImagePath
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// TagsOrErr returns the Tags value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) TagsOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tags, nil
	}
	return nil, &NotLoadedError{edge: "tags"}
}

// RatingsOrErr returns the Ratings value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) RatingsOrErr() ([]*Rating, error) {
	if e.loadedTypes[1] {
		return e.Ratings, nil
	}
	return nil, &NotLoadedError{edge: "ratings"}
}

// RestaurantOrErr returns the Restaurant value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) RestaurantOrErr() ([]*Restaurant, error) {
	if e.loadedTypes[2] {
		return e.Restaurant, nil
	}
	return nil, &NotLoadedError{edge: "restaurant"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[3] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e ProductEdges) ImagesOrErr() ([]*ImagePath, error) {
	if e.loadedTypes[4] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // description
		&sql.NullInt64{},  // cost
		&sql.NullBool{},   // is_active
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(values ...interface{}) error {
	if m, n := len(values), len(product.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pr.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[1])
	} else if value.Valid {
		pr.Description = new(string)
		*pr.Description = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field cost", values[2])
	} else if value.Valid {
		pr.Cost = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field is_active", values[3])
	} else if value.Valid {
		pr.IsActive = value.Bool
	}
	return nil
}

// QueryTags queries the tags edge of the Product.
func (pr *Product) QueryTags() *TagQuery {
	return (&ProductClient{config: pr.config}).QueryTags(pr)
}

// QueryRatings queries the ratings edge of the Product.
func (pr *Product) QueryRatings() *RatingQuery {
	return (&ProductClient{config: pr.config}).QueryRatings(pr)
}

// QueryRestaurant queries the restaurant edge of the Product.
func (pr *Product) QueryRestaurant() *RestaurantQuery {
	return (&ProductClient{config: pr.config}).QueryRestaurant(pr)
}

// QueryOrders queries the orders edge of the Product.
func (pr *Product) QueryOrders() *OrderQuery {
	return (&ProductClient{config: pr.config}).QueryOrders(pr)
}

// QueryImages queries the images edge of the Product.
func (pr *Product) QueryImages() *ImagePathQuery {
	return (&ProductClient{config: pr.config}).QueryImages(pr)
}

// Update returns a builder for updating this Product.
// Note that, you need to call Product.Unwrap() before calling this method, if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return (&ProductClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", name=")
	builder.WriteString(pr.Name)
	if v := pr.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", cost=")
	builder.WriteString(fmt.Sprintf("%v", pr.Cost))
	builder.WriteString(", is_active=")
	builder.WriteString(fmt.Sprintf("%v", pr.IsActive))
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product

func (pr Products) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
