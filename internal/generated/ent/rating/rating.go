// Code generated by entc, DO NOT EDIT.

package rating

const (
	// Label holds the string label denoting the rating type in the database.
	Label = "rating"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldComment holds the string denoting the comment field in the database.
	FieldComment = "comment"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"

	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"

	// Table holds the table name of the rating in the database.
	Table = "ratings"
	// CustomerTable is the table the holds the customer relation/edge.
	CustomerTable = "ratings"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_ratings"
	// ProductTable is the table the holds the product relation/edge.
	ProductTable = "ratings"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_ratings"
)

// Columns holds all SQL columns for rating fields.
var Columns = []string{
	FieldID,
	FieldComment,
	FieldRating,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Rating type.
var ForeignKeys = []string{
	"customer_ratings",
	"product_ratings",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
