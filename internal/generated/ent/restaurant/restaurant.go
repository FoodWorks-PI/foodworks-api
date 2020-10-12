// Code generated by entc, DO NOT EDIT.

package restaurant

const (
	// Label holds the string label denoting the restaurant type in the database.
	Label = "restaurant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"

	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"

	// Table holds the table name of the restaurant in the database.
	Table = "restaurants"
	// AddressTable is the table the holds the address relation/edge.
	AddressTable = "restaurants"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// AddressColumn is the table column denoting the address relation/edge.
	AddressColumn = "restaurant_address"
	// TagsTable is the table the holds the tags relation/edge. The primary key declared below.
	TagsTable = "restaurant_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "restaurant_owners"
	// OwnerInverseTable is the table name for the RestaurantOwner entity.
	// It exists in this package in order to avoid circular dependency with the "restaurantowner" package.
	OwnerInverseTable = "restaurant_owners"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "restaurant_owner_restaurant"
	// ProductsTable is the table the holds the products relation/edge. The primary key declared below.
	ProductsTable = "restaurant_products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
)

// Columns holds all SQL columns for restaurant fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Restaurant type.
var ForeignKeys = []string{
	"restaurant_address",
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"restaurant_id", "tag_id"}
	// ProductsPrimaryKey and ProductsColumn2 are the table columns denoting the
	// primary key for the products relation (M2M).
	ProductsPrimaryKey = []string{"restaurant_id", "product_id"}
)

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
