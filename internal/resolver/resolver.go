package resolver

import (
	"context"
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/address"
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
func OrderByDistanceP() func(selector *sql.Selector) {
	return func(s *sql.Selector) {
		s.SetDistinct(false)
		s.Where(P(fmt.Sprint("1=1 ORDER BY geom <-> '010100000017CAF411D9CB58C011B34DE08C5C3340'")))
	}
}

func SelectDistance() func(selector *sql.Selector) {
	return func(s *sql.Selector) {
		s.Select(address.Columns...)
		//s.WriteString("12 as distance")
		//s.Offset()
		//tmp : s.String()
		//v:=reflect.ValueOf(s).Elem()
		//tmp1 ,tmp2 := s.Query()
		//fmt.Println(tmp1,tmp2)
		//fmt.Println(s.Columns(address.Columns...))
		//f := v.FieldByName("columns").Type()
		//f := v.FieldByName("columns")
		//fv := reflect.ValueOf(f)
		//fmt.Println(f.Kind())
		//fmt.Println(fv)
		//y := v.FieldByName("columns")
		//fmt.Println(y.Interface())
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
