// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/rating"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RatingCreate is the builder for creating a Rating entity.
type RatingCreate struct {
	config
	mutation *RatingMutation
	hooks    []Hook
}

// SetProductRate sets the ProductRate field.
func (rc *RatingCreate) SetProductRate(i int) *RatingCreate {
	rc.mutation.SetProductRate(i)
	return rc
}

// SetProductID sets the ProductID field.
func (rc *RatingCreate) SetProductID(i int) *RatingCreate {
	rc.mutation.SetProductID(i)
	return rc
}

// SetCustomerID sets the CustomerID field.
func (rc *RatingCreate) SetCustomerID(i int) *RatingCreate {
	rc.mutation.SetCustomerID(i)
	return rc
}

// Mutation returns the RatingMutation object of the builder.
func (rc *RatingCreate) Mutation() *RatingMutation {
	return rc.mutation
}

// Save creates the Rating in the database.
func (rc *RatingCreate) Save(ctx context.Context) (*Rating, error) {
	var (
		err  error
		node *Rating
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RatingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			node, err = rc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RatingCreate) SaveX(ctx context.Context) *Rating {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (rc *RatingCreate) check() error {
	if _, ok := rc.mutation.ProductRate(); !ok {
		return &ValidationError{Name: "ProductRate", err: errors.New("ent: missing required field \"ProductRate\"")}
	}
	if _, ok := rc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "ProductID", err: errors.New("ent: missing required field \"ProductID\"")}
	}
	if _, ok := rc.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "CustomerID", err: errors.New("ent: missing required field \"CustomerID\"")}
	}
	return nil
}

func (rc *RatingCreate) sqlSave(ctx context.Context) (*Rating, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *RatingCreate) createSpec() (*Rating, *sqlgraph.CreateSpec) {
	var (
		_node = &Rating{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: rating.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rating.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.ProductRate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldProductRate,
		})
		_node.ProductRate = value
	}
	if value, ok := rc.mutation.ProductID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldProductID,
		})
		_node.ProductID = value
	}
	if value, ok := rc.mutation.CustomerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldCustomerID,
		})
		_node.CustomerID = value
	}
	return _node, _spec
}

// RatingCreateBulk is the builder for creating a bulk of Rating entities.
type RatingCreateBulk struct {
	config
	builders []*RatingCreate
}

// Save creates the Rating entities in the database.
func (rcb *RatingCreateBulk) Save(ctx context.Context) ([]*Rating, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Rating, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RatingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (rcb *RatingCreateBulk) SaveX(ctx context.Context) []*Rating {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
