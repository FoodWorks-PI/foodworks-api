package resolver

import (
	"context"
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/product"
	"foodworks.ml/m/internal/generated/ent/restaurant"

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
func OrderByDistanceP(table string, distance string) func(selector *sql.Selector) {
	return func(s *sql.Selector) {
		addressTable := sql.Table(address.Table)
		switch table {
		case restaurant.Table:
			fmt.Println("hi")
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

func SelectDistance() func(selector *sql.Selector) {
	return func(s *sql.Selector) {
		columns := make([]string, len(restaurant.Columns))
		addressTable := sql.Table(address.Table)
		for i, column := range restaurant.Columns {
			columns[i] = s.C(column)
		}
		s.Join(addressTable).On(addressTable.C(address.FieldID), s.C(restaurant.AddressColumn))
		columns = append(columns, `st_distancesphere(st_makepoint(19.37991393,-99.17228876),st_makepoint(19.37748,-99.16799)) as "distance"`)
		//columns = append(columns, restaurant.Columns...)
		s.Select(columns...)
		//fmt.Println(tmp)
		//s.Select()
		//s.Select("st_distancesphere(st_makepoint(19.37991393,-99.17228876),st_makepoint(19.37748,-99.16799))")
		//s.Select("12 as distance")
		//s.WriteString("ORDER BY geom")
		//s.OrderBy()
		//s.WriteString("geom")
	}
}

func P(c2 string) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString(c2)
	})
}
