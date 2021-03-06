// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/predicate"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RestaurantOwnerDelete is the builder for deleting a RestaurantOwner entity.
type RestaurantOwnerDelete struct {
	config
	hooks      []Hook
	mutation   *RestaurantOwnerMutation
	predicates []predicate.RestaurantOwner
}

// Where adds a new predicate to the delete builder.
func (rod *RestaurantOwnerDelete) Where(ps ...predicate.RestaurantOwner) *RestaurantOwnerDelete {
	rod.predicates = append(rod.predicates, ps...)
	return rod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rod *RestaurantOwnerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rod.hooks) == 0 {
		affected, err = rod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RestaurantOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rod.mutation = mutation
			affected, err = rod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rod.hooks) - 1; i >= 0; i-- {
			mut = rod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rod *RestaurantOwnerDelete) ExecX(ctx context.Context) int {
	n, err := rod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rod *RestaurantOwnerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: restaurantowner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: restaurantowner.FieldID,
			},
		},
	}
	if ps := rod.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rod.driver, _spec)
}

// RestaurantOwnerDeleteOne is the builder for deleting a single RestaurantOwner entity.
type RestaurantOwnerDeleteOne struct {
	rod *RestaurantOwnerDelete
}

// Exec executes the deletion query.
func (rodo *RestaurantOwnerDeleteOne) Exec(ctx context.Context) error {
	n, err := rodo.rod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{restaurantowner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rodo *RestaurantOwnerDeleteOne) ExecX(ctx context.Context) {
	rodo.rod.ExecX(ctx)
}
