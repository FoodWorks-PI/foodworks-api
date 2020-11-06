// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/bankingdata"
	"foodworks.ml/m/internal/generated/ent/restaurant"
	"foodworks.ml/m/internal/generated/ent/restaurantowner"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RestaurantOwnerCreate is the builder for creating a RestaurantOwner entity.
type RestaurantOwnerCreate struct {
	config
	mutation *RestaurantOwnerMutation
	hooks    []Hook
}

// SetKratosID sets the kratos_id field.
func (roc *RestaurantOwnerCreate) SetKratosID(s string) *RestaurantOwnerCreate {
	roc.mutation.SetKratosID(s)
	return roc
}

// SetName sets the name field.
func (roc *RestaurantOwnerCreate) SetName(s string) *RestaurantOwnerCreate {
	roc.mutation.SetName(s)
	return roc
}

// SetLastName sets the last_name field.
func (roc *RestaurantOwnerCreate) SetLastName(s string) *RestaurantOwnerCreate {
	roc.mutation.SetLastName(s)
	return roc
}

// SetEmail sets the email field.
func (roc *RestaurantOwnerCreate) SetEmail(s string) *RestaurantOwnerCreate {
	roc.mutation.SetEmail(s)
	return roc
}

// SetPhone sets the phone field.
func (roc *RestaurantOwnerCreate) SetPhone(s string) *RestaurantOwnerCreate {
	roc.mutation.SetPhone(s)
	return roc
}

// SetBankingDataID sets the banking_data edge to BankingData by id.
func (roc *RestaurantOwnerCreate) SetBankingDataID(id int) *RestaurantOwnerCreate {
	roc.mutation.SetBankingDataID(id)
	return roc
}

// SetNillableBankingDataID sets the banking_data edge to BankingData by id if the given value is not nil.
func (roc *RestaurantOwnerCreate) SetNillableBankingDataID(id *int) *RestaurantOwnerCreate {
	if id != nil {
		roc = roc.SetBankingDataID(*id)
	}
	return roc
}

// SetBankingData sets the banking_data edge to BankingData.
func (roc *RestaurantOwnerCreate) SetBankingData(b *BankingData) *RestaurantOwnerCreate {
	return roc.SetBankingDataID(b.ID)
}

// SetRestaurantID sets the restaurant edge to Restaurant by id.
func (roc *RestaurantOwnerCreate) SetRestaurantID(id int) *RestaurantOwnerCreate {
	roc.mutation.SetRestaurantID(id)
	return roc
}

// SetNillableRestaurantID sets the restaurant edge to Restaurant by id if the given value is not nil.
func (roc *RestaurantOwnerCreate) SetNillableRestaurantID(id *int) *RestaurantOwnerCreate {
	if id != nil {
		roc = roc.SetRestaurantID(*id)
	}
	return roc
}

// SetRestaurant sets the restaurant edge to Restaurant.
func (roc *RestaurantOwnerCreate) SetRestaurant(r *Restaurant) *RestaurantOwnerCreate {
	return roc.SetRestaurantID(r.ID)
}

// Mutation returns the RestaurantOwnerMutation object of the builder.
func (roc *RestaurantOwnerCreate) Mutation() *RestaurantOwnerMutation {
	return roc.mutation
}

// Save creates the RestaurantOwner in the database.
func (roc *RestaurantOwnerCreate) Save(ctx context.Context) (*RestaurantOwner, error) {
	var (
		err  error
		node *RestaurantOwner
	)
	if len(roc.hooks) == 0 {
		if err = roc.check(); err != nil {
			return nil, err
		}
		node, err = roc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RestaurantOwnerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = roc.check(); err != nil {
				return nil, err
			}
			roc.mutation = mutation
			node, err = roc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(roc.hooks) - 1; i >= 0; i-- {
			mut = roc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, roc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (roc *RestaurantOwnerCreate) SaveX(ctx context.Context) *RestaurantOwner {
	v, err := roc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (roc *RestaurantOwnerCreate) check() error {
	if _, ok := roc.mutation.KratosID(); !ok {
		return &ValidationError{Name: "kratos_id", err: errors.New("ent: missing required field \"kratos_id\"")}
	}
	if _, ok := roc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := roc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New("ent: missing required field \"last_name\"")}
	}
	if _, ok := roc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if _, ok := roc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New("ent: missing required field \"phone\"")}
	}
	return nil
}

func (roc *RestaurantOwnerCreate) sqlSave(ctx context.Context) (*RestaurantOwner, error) {
	_node, _spec := roc.createSpec()
	if err := sqlgraph.CreateNode(ctx, roc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (roc *RestaurantOwnerCreate) createSpec() (*RestaurantOwner, *sqlgraph.CreateSpec) {
	var (
		_node = &RestaurantOwner{config: roc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: restaurantowner.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: restaurantowner.FieldID,
			},
		}
	)
	if value, ok := roc.mutation.KratosID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurantowner.FieldKratosID,
		})
		_node.KratosID = value
	}
	if value, ok := roc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurantowner.FieldName,
		})
		_node.Name = value
	}
	if value, ok := roc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurantowner.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := roc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurantowner.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := roc.mutation.Phone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: restaurantowner.FieldPhone,
		})
		_node.Phone = value
	}
	if nodes := roc.mutation.BankingDataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   restaurantowner.BankingDataTable,
			Columns: []string{restaurantowner.BankingDataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bankingdata.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := roc.mutation.RestaurantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   restaurantowner.RestaurantTable,
			Columns: []string{restaurantowner.RestaurantColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: restaurant.FieldID,
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

// RestaurantOwnerCreateBulk is the builder for creating a bulk of RestaurantOwner entities.
type RestaurantOwnerCreateBulk struct {
	config
	builders []*RestaurantOwnerCreate
}

// Save creates the RestaurantOwner entities in the database.
func (rocb *RestaurantOwnerCreateBulk) Save(ctx context.Context) ([]*RestaurantOwner, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rocb.builders))
	nodes := make([]*RestaurantOwner, len(rocb.builders))
	mutators := make([]Mutator, len(rocb.builders))
	for i := range rocb.builders {
		func(i int, root context.Context) {
			builder := rocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RestaurantOwnerMutation)
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
					_, err = mutators[i+1].Mutate(root, rocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (rocb *RestaurantOwnerCreateBulk) SaveX(ctx context.Context) []*RestaurantOwner {
	v, err := rocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
