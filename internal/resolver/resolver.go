package resolver

import (
	"context"
	"fmt"
	"strings"

	"foodworks.ml/m/internal/generated/ent"
	"foodworks.ml/m/internal/generated/ent/tag"
	"foodworks.ml/m/internal/generated/graphql/model"
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

func GetOrCreateTagId(r *Resolver, tags []*model.TagInput, ctx context.Context) ([]*ent.Tag, error) {
	var sb strings.Builder
	if tags == nil {
		return nil, nil
	}
	// Build a comma separated string of tags enclosed by single quotes
	for i := 0; i < len(tags)-1; i++ {
		sb.WriteString(fmt.Sprintf("dsl.match('name', '%s')", tags[i].Name))
		sb.WriteString(",")
	}
	sb.WriteString(fmt.Sprintf("dsl.match('name', '%s')", tags[len(tags)-1]))
	_ = fmt.Sprintf("SELECT * FROM TAGS WHERE TAGS ==> dsl.or(%s)", sb.String())
	predicate := fmt.Sprintf("TAGS ==> dsl.or(%s)", sb.String())
	values, err := r.EntClient.Tag.Query().
		Where(func(s *sql.Selector) {
			s.Where(P(sql.Table(tag.Table).String(), predicate))
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
		if !existingTags[tag.Name] {
			bulk = append(bulk, r.EntClient.Tag.Create().SetName(tag.Name))
		}
	}
	newTags, err := r.EntClient.Tag.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}
	values = append(values, newTags...)
	return values, nil
}

func P(c1, c2 string) *sql.Predicate {
	return sql.P(func(b *sql.Builder) {
		b.WriteString(c2)
	})
}
