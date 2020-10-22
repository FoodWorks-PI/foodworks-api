// Code generated by entc, DO NOT EDIT.

package product

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"

	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeRestaurant holds the string denoting the restaurant edge name in mutations.
	EdgeRestaurant = "restaurant"

	// Table holds the table name of the product in the database.
	Table = "products"
	// TagsTable is the table the holds the tags relation/edge. The primary key declared below.
	TagsTable = "product_tags"
	// TagsInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagsInverseTable = "tags"
	// RestaurantTable is the table the holds the restaurant relation/edge. The primary key declared below.
	RestaurantTable = "restaurant_products"
	// RestaurantInverseTable is the table name for the Restaurant entity.
	// It exists in this package in order to avoid circular dependency with the "restaurant" package.
	RestaurantInverseTable = "restaurants"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCost,
	FieldIsActive,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"product_id", "tag_id"}
	// RestaurantPrimaryKey and RestaurantColumn2 are the table columns denoting the
	// primary key for the restaurant relation (M2M).
	RestaurantPrimaryKey = []string{"restaurant_id", "product_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}