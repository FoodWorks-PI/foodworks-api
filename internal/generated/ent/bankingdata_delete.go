// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/bankingdata"
	"foodworks.ml/m/internal/generated/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// BankingDataDelete is the builder for deleting a BankingData entity.
type BankingDataDelete struct {
	config
	hooks      []Hook
	mutation   *BankingDataMutation
	predicates []predicate.BankingData
}

// Where adds a new predicate to the delete builder.
func (bdd *BankingDataDelete) Where(ps ...predicate.BankingData) *BankingDataDelete {
	bdd.predicates = append(bdd.predicates, ps...)
	return bdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bdd *BankingDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bdd.hooks) == 0 {
		affected, err = bdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankingDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bdd.mutation = mutation
			affected, err = bdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bdd.hooks) - 1; i >= 0; i-- {
			mut = bdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bdd *BankingDataDelete) ExecX(ctx context.Context) int {
	n, err := bdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bdd *BankingDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bankingdata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bankingdata.FieldID,
			},
		},
	}
	if ps := bdd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bdd.driver, _spec)
}

// BankingDataDeleteOne is the builder for deleting a single BankingData entity.
type BankingDataDeleteOne struct {
	bdd *BankingDataDelete
}

// Exec executes the deletion query.
func (bddo *BankingDataDeleteOne) Exec(ctx context.Context) error {
	n, err := bddo.bdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bankingdata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bddo *BankingDataDeleteOne) ExecX(ctx context.Context) {
	bddo.bdd.ExecX(ctx)
}
