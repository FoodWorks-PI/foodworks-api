// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AddressUpdate is the builder for updating Address entities.
type AddressUpdate struct {
	config
	hooks      []Hook
	mutation   *AddressMutation
	predicates []predicate.Address
}

// Where adds a new predicate for the builder.
func (au *AddressUpdate) Where(ps ...predicate.Address) *AddressUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetLatitude sets the latitude field.
func (au *AddressUpdate) SetLatitude(s string) *AddressUpdate {
	au.mutation.SetLatitude(s)
	return au
}

// SetLongitude sets the longitude field.
func (au *AddressUpdate) SetLongitude(s string) *AddressUpdate {
	au.mutation.SetLongitude(s)
	return au
}

// SetStreetLine sets the streetLine field.
func (au *AddressUpdate) SetStreetLine(s string) *AddressUpdate {
	au.mutation.SetStreetLine(s)
	return au
}

// Mutation returns the AddressMutation object of the builder.
func (au *AddressUpdate) Mutation() *AddressMutation {
	return au.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AddressUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddressUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddressUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddressUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   address.Table,
			Columns: address.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: address.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Latitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLatitude,
		})
	}
	if value, ok := au.mutation.Longitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLongitude,
		})
	}
	if value, ok := au.mutation.StreetLine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldStreetLine,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AddressUpdateOne is the builder for updating a single Address entity.
type AddressUpdateOne struct {
	config
	hooks    []Hook
	mutation *AddressMutation
}

// SetLatitude sets the latitude field.
func (auo *AddressUpdateOne) SetLatitude(s string) *AddressUpdateOne {
	auo.mutation.SetLatitude(s)
	return auo
}

// SetLongitude sets the longitude field.
func (auo *AddressUpdateOne) SetLongitude(s string) *AddressUpdateOne {
	auo.mutation.SetLongitude(s)
	return auo
}

// SetStreetLine sets the streetLine field.
func (auo *AddressUpdateOne) SetStreetLine(s string) *AddressUpdateOne {
	auo.mutation.SetStreetLine(s)
	return auo
}

// Mutation returns the AddressMutation object of the builder.
func (auo *AddressUpdateOne) Mutation() *AddressMutation {
	return auo.mutation
}

// Save executes the query and returns the updated entity.
func (auo *AddressUpdateOne) Save(ctx context.Context) (*Address, error) {
	var (
		err  error
		node *Address
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AddressMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddressUpdateOne) SaveX(ctx context.Context) *Address {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddressUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddressUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AddressUpdateOne) sqlSave(ctx context.Context) (_node *Address, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   address.Table,
			Columns: address.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: address.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Address.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Latitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLatitude,
		})
	}
	if value, ok := auo.mutation.Longitude(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldLongitude,
		})
	}
	if value, ok := auo.mutation.StreetLine(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: address.FieldStreetLine,
		})
	}
	_node = &Address{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
