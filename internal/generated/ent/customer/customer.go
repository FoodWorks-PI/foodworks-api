// Code generated by entc, DO NOT EDIT.

package customer

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKratosID holds the string denoting the kratos_id field in the database.
	FieldKratosID = "kratos_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"

	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// EdgeRatings holds the string denoting the ratings edge name in mutations.
	EdgeRatings = "ratings"

	// Table holds the table name of the customer in the database.
	Table = "customers"
	// AddressTable is the table the holds the address relation/edge.
	AddressTable = "customers"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// AddressColumn is the table column denoting the address relation/edge.
	AddressColumn = "customer_address"
	// RatingsTable is the table the holds the ratings relation/edge.
	RatingsTable = "ratings"
	// RatingsInverseTable is the table name for the Rating entity.
	// It exists in this package in order to avoid circular dependency with the "rating" package.
	RatingsInverseTable = "ratings"
	// RatingsColumn is the table column denoting the ratings relation/edge.
	RatingsColumn = "customer_ratings"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldKratosID,
	FieldName,
	FieldLastName,
	FieldEmail,
	FieldPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Customer type.
var ForeignKeys = []string{
	"customer_address",
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
