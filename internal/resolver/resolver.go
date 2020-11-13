package resolver

import (
	"context"
	"fmt"
	"strings"

	"foodworks.ml/m/internal/recommendations"

	"foodworks.ml/m/internal/auth"
	"foodworks.ml/m/internal/generated/ent/customer"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"

	"foodworks.ml/m/internal/generated/graphql/model"
	"github.com/99designs/gqlgen/graphql"

	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/product"
	"foodworks.ml/m/internal/generated/ent/restaurant"
	"foodworks.ml/m/internal/generated/ent/tag"

	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/platform/filehandler"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/facebook/ent/dialect/sql"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EntClient     *ent.Client
	Redis         *redis.Client
	DBClient      *sqlx.DB
	ElasticClient *elasticsearch.Client
	FileHandler   *filehandler.FileHandler
	Recommender   recommendations.Reccomender
}

func GetOrCreateTagId(r *Resolver, tags []string, ctx context.Context) ([]*ent.Tag, error) {
	var sb strings.Builder
	if tags == nil {
		return nil, nil
	}
	// Make lowercase
	for i := 0; i < len(tags); i++ {
		tags[i] = strings.ToLower(tags[i])
	}
	// Build a comma separated string of tags enclosed by single quotes
	for i := 0; i < len(tags)-1; i++ {
		sb.WriteString(fmt.Sprintf("dsl.match('name', '%s')", tags[i]))
		sb.WriteString(",")
	}
	sb.WriteString(fmt.Sprintf("dsl.match('name', '%s')", tags[len(tags)-1]))
	predicate := fmt.Sprintf("TAGS ==> dsl.or(%s)", sb.String())
	values, err := r.EntClient.Tag.Query().
		Where(func(s *sql.Selector) {
			s.Where(P(predicate))
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}
	// Create map of existing tags
	existingTags := make(map[string]bool, len(values))
	for _, tag := range values {
		existingTags[tag.Name] = true
	}
	// Create missing tags
	bulk := make([]*ent.TagCreate, 0, len(tags))
	for _, tag := range tags {
		if !existingTags[tag] {
			bulk = append(bulk, r.EntClient.Tag.Create().SetName(tag))
		}
	}
	newTags, err := r.EntClient.Tag.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	values = append(values, newTags...)
	return values, nil
}
func HasRole(entClient *ent.Client) func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		kratosUser := auth.ForContext(ctx)
		switch role {
		case model.RoleOwner:
			exists, err := entClient.RestaurantOwner.Query().Where(restaurantowner.KratosID(kratosUser.ID)).Exist(ctx)
			if err != nil || !exists {
				return nil, fmt.Errorf("access denied")
			}
		case model.RoleCustomer:
			exists, err := entClient.Customer.Query().Where(customer.KratosID(kratosUser.ID)).Exist(ctx)
			if err != nil || !exists {
				return nil, fmt.Errorf("access denied")
			}
		}
		return next(ctx)
	}
	//switch role {
	//case model.RoleOwner:
	//	break;
	//case model.RoleOwner:
	//	break;
	//}
}

func RemoveDuplicateRestaurant(restaurants []*ent.Restaurant) []*ent.Restaurant {
	occurred := map[int]bool{}
	result := make([]*ent.Restaurant, 0, len(restaurants))
	for _, restaurant := range restaurants {
		if !occurred[restaurant.ID] {
			occurred[restaurant.ID] = true
			result = append(result, restaurant)
		}
	}
	return result
}

func RemoveDuplicateProducts(products []*ent.Product) []*ent.Product {
	occurred := map[int]bool{}
	result := make([]*ent.Product, 0, len(products))
	for _, product := range products {
		if !occurred[product.ID] {
			occurred[product.ID] = true
			result = append(result, product)
		}
	}
	return result
}

func WhereTextMatch(table string, query string) func(selector *sql.Selector) {
	return func(s *sql.Selector) {
		// _ := fmt.Sprintf("%s => %s")
		tagsTable := sql.Table(tag.Table).As("tags")
		var sb strings.Builder
		switch table {
		case restaurant.Table:
			restaurantTagTable := sql.Table(restaurant.TagsTable).As("q0")
			s.Join(restaurantTagTable).On(s.C(restaurant.FieldID), restaurantTagTable.C(restaurant.TagsPrimaryKey[0]))
			s.Join(tagsTable).On(tagsTable.C(tag.FieldID), restaurantTagTable.C(restaurant.TagsPrimaryKey[1]))
		case product.Table:
			productTagTable := sql.Table(product.TagsTable).As("q0")
			s.Join(productTagTable).On(s.C(product.FieldID), productTagTable.C(product.TagsPrimaryKey[0]))
			s.Join(tagsTable).On(tagsTable.C(tag.FieldID), productTagTable.C(product.TagsPrimaryKey[1]))
		}
		sb.WriteString(fmt.Sprintf(`"%s" ==> '%s'`, table, query))
		sb.WriteString("OR ")
		sb.WriteString(fmt.Sprintf(`"%s" ==> '%s'`, "tags", query))
		s.Where(P(sb.String()))
	}
}

func OrderByDistanceP(table string, distance string) func(selector *sql.Selector) {
	return func(s *sql.Selector) {
		addressTable := sql.Table(address.Table)
		switch table {
		case restaurant.Table:
			s.Join(addressTable).On(addressTable.C(address.FieldID), s.C(restaurant.AddressColumn))
		case product.Table:
			restaurantProductTable := sql.Table(product.RestaurantTable).As("t0")
			restaurantTable := sql.Table(restaurant.Table).As("t1")
			addressTable = addressTable.As("t2")
			s.Join(restaurantProductTable).On(restaurantProductTable.C(restaurant.ProductsPrimaryKey[1]),
				s.C(product.FieldID))
			s.Join(restaurantTable).On(restaurantProductTable.C(restaurant.ProductsPrimaryKey[0]),
				restaurantTable.C(restaurant.FieldID))
			s.Join(addressTable).On(addressTable.C(address.FieldID), restaurantTable.C(restaurant.AddressColumn))
		}
		s.Where(P(fmt.Sprintf("1=1 ORDER BY geom <-> '%s'", distance)))
	}
}

func GetColumns(prefix string, columns []string) []string {
	newColumns := make([]string, len(columns))
	for i, column := range columns {
		newColumns[i] = fmt.Sprintf(`"%s"."%s"`, prefix, column)
	}
	return newColumns
}

func P(c2 string) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString(c2)
	})
}
