// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/customer"
	"foodworks.ml/m/internal/generated/ent/product"
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

// SetComment sets the comment field.
func (rc *RatingCreate) SetComment(s string) *RatingCreate {
	rc.mutation.SetComment(s)
	return rc
}

// SetRating sets the rating field.
func (rc *RatingCreate) SetRating(i int) *RatingCreate {
	rc.mutation.SetRating(i)
	return rc
}

// SetCustomerID sets the customer edge to Customer by id.
func (rc *RatingCreate) SetCustomerID(id int) *RatingCreate {
	rc.mutation.SetCustomerID(id)
	return rc
}

// SetNillableCustomerID sets the customer edge to Customer by id if the given value is not nil.
func (rc *RatingCreate) SetNillableCustomerID(id *int) *RatingCreate {
	if id != nil {
		rc = rc.SetCustomerID(*id)
	}
	return rc
}

// SetCustomer sets the customer edge to Customer.
func (rc *RatingCreate) SetCustomer(c *Customer) *RatingCreate {
	return rc.SetCustomerID(c.ID)
}

// SetProductID sets the product edge to Product by id.
func (rc *RatingCreate) SetProductID(id int) *RatingCreate {
	rc.mutation.SetProductID(id)
	return rc
}

// SetNillableProductID sets the product edge to Product by id if the given value is not nil.
func (rc *RatingCreate) SetNillableProductID(id *int) *RatingCreate {
	if id != nil {
		rc = rc.SetProductID(*id)
	}
	return rc
}

// SetProduct sets the product edge to Product.
func (rc *RatingCreate) SetProduct(p *Product) *RatingCreate {
	return rc.SetProductID(p.ID)
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
	if _, ok := rc.mutation.Comment(); !ok {
		return &ValidationError{Name: "comment", err: errors.New("ent: missing required field \"comment\"")}
	}
	if _, ok := rc.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New("ent: missing required field \"rating\"")}
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
	if value, ok := rc.mutation.Comment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: rating.FieldComment,
		})
		_node.Comment = value
	}
	if value, ok := rc.mutation.Rating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: rating.FieldRating,
		})
		_node.Rating = value
	}
	if nodes := rc.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rating.CustomerTable,
			Columns: []string{rating.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: customer.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rating.ProductTable,
			Columns: []string{rating.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
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
