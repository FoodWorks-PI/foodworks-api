// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/address"
	"foodworks.ml/m/internal/generated/ent/customer"
	"foodworks.ml/m/internal/generated/ent/predicate"
	"foodworks.ml/m/internal/generated/ent/rating"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// CustomerUpdate is the builder for updating Customer entities.
type CustomerUpdate struct {
	config
	hooks      []Hook
	mutation   *CustomerMutation
	predicates []predicate.Customer
}

// Where adds a new predicate for the builder.
func (cu *CustomerUpdate) Where(ps ...predicate.Customer) *CustomerUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetKratosID sets the kratos_id field.
func (cu *CustomerUpdate) SetKratosID(s string) *CustomerUpdate {
	cu.mutation.SetKratosID(s)
	return cu
}

// SetName sets the name field.
func (cu *CustomerUpdate) SetName(s string) *CustomerUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetLastName sets the last_name field.
func (cu *CustomerUpdate) SetLastName(s string) *CustomerUpdate {
	cu.mutation.SetLastName(s)
	return cu
}

// SetEmail sets the email field.
func (cu *CustomerUpdate) SetEmail(s string) *CustomerUpdate {
	cu.mutation.SetEmail(s)
	return cu
}

// SetPhone sets the phone field.
func (cu *CustomerUpdate) SetPhone(s string) *CustomerUpdate {
	cu.mutation.SetPhone(s)
	return cu
}

// SetAddressID sets the address edge to Address by id.
func (cu *CustomerUpdate) SetAddressID(id int) *CustomerUpdate {
	cu.mutation.SetAddressID(id)
	return cu
}

// SetNillableAddressID sets the address edge to Address by id if the given value is not nil.
func (cu *CustomerUpdate) SetNillableAddressID(id *int) *CustomerUpdate {
	if id != nil {
		cu = cu.SetAddressID(*id)
	}
	return cu
}

// SetAddress sets the address edge to Address.
func (cu *CustomerUpdate) SetAddress(a *Address) *CustomerUpdate {
	return cu.SetAddressID(a.ID)
}

// AddRatingIDs adds the ratings edge to Rating by ids.
func (cu *CustomerUpdate) AddRatingIDs(ids ...int) *CustomerUpdate {
	cu.mutation.AddRatingIDs(ids...)
	return cu
}

// AddRatings adds the ratings edges to Rating.
func (cu *CustomerUpdate) AddRatings(r ...*Rating) *CustomerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.AddRatingIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cu *CustomerUpdate) Mutation() *CustomerMutation {
	return cu.mutation
}

// ClearAddress clears the "address" edge to type Address.
func (cu *CustomerUpdate) ClearAddress() *CustomerUpdate {
	cu.mutation.ClearAddress()
	return cu
}

// ClearRatings clears all "ratings" edges to type Rating.
func (cu *CustomerUpdate) ClearRatings() *CustomerUpdate {
	cu.mutation.ClearRatings()
	return cu
}

// RemoveRatingIDs removes the ratings edge to Rating by ids.
func (cu *CustomerUpdate) RemoveRatingIDs(ids ...int) *CustomerUpdate {
	cu.mutation.RemoveRatingIDs(ids...)
	return cu
}

// RemoveRatings removes ratings edges to Rating.
func (cu *CustomerUpdate) RemoveRatings(r ...*Rating) *CustomerUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cu.RemoveRatingIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CustomerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CustomerUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CustomerUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CustomerUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CustomerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.KratosID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldKratosID,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if value, ok := cu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldLastName,
		})
	}
	if value, ok := cu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
	}
	if value, ok := cu.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if cu.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.AddressTable,
			Columns: []string{customer.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.AddressTable,
			Columns: []string{customer.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.RatingsTable,
			Columns: []string{customer.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rating.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedRatingsIDs(); len(nodes) > 0 && !cu.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.RatingsTable,
			Columns: []string{customer.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rating.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.RatingsTable,
			Columns: []string{customer.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rating.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CustomerUpdateOne is the builder for updating a single Customer entity.
type CustomerUpdateOne struct {
	config
	hooks    []Hook
	mutation *CustomerMutation
}

// SetKratosID sets the kratos_id field.
func (cuo *CustomerUpdateOne) SetKratosID(s string) *CustomerUpdateOne {
	cuo.mutation.SetKratosID(s)
	return cuo
}

// SetName sets the name field.
func (cuo *CustomerUpdateOne) SetName(s string) *CustomerUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetLastName sets the last_name field.
func (cuo *CustomerUpdateOne) SetLastName(s string) *CustomerUpdateOne {
	cuo.mutation.SetLastName(s)
	return cuo
}

// SetEmail sets the email field.
func (cuo *CustomerUpdateOne) SetEmail(s string) *CustomerUpdateOne {
	cuo.mutation.SetEmail(s)
	return cuo
}

// SetPhone sets the phone field.
func (cuo *CustomerUpdateOne) SetPhone(s string) *CustomerUpdateOne {
	cuo.mutation.SetPhone(s)
	return cuo
}

// SetAddressID sets the address edge to Address by id.
func (cuo *CustomerUpdateOne) SetAddressID(id int) *CustomerUpdateOne {
	cuo.mutation.SetAddressID(id)
	return cuo
}

// SetNillableAddressID sets the address edge to Address by id if the given value is not nil.
func (cuo *CustomerUpdateOne) SetNillableAddressID(id *int) *CustomerUpdateOne {
	if id != nil {
		cuo = cuo.SetAddressID(*id)
	}
	return cuo
}

// SetAddress sets the address edge to Address.
func (cuo *CustomerUpdateOne) SetAddress(a *Address) *CustomerUpdateOne {
	return cuo.SetAddressID(a.ID)
}

// AddRatingIDs adds the ratings edge to Rating by ids.
func (cuo *CustomerUpdateOne) AddRatingIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.AddRatingIDs(ids...)
	return cuo
}

// AddRatings adds the ratings edges to Rating.
func (cuo *CustomerUpdateOne) AddRatings(r ...*Rating) *CustomerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.AddRatingIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cuo *CustomerUpdateOne) Mutation() *CustomerMutation {
	return cuo.mutation
}

// ClearAddress clears the "address" edge to type Address.
func (cuo *CustomerUpdateOne) ClearAddress() *CustomerUpdateOne {
	cuo.mutation.ClearAddress()
	return cuo
}

// ClearRatings clears all "ratings" edges to type Rating.
func (cuo *CustomerUpdateOne) ClearRatings() *CustomerUpdateOne {
	cuo.mutation.ClearRatings()
	return cuo
}

// RemoveRatingIDs removes the ratings edge to Rating by ids.
func (cuo *CustomerUpdateOne) RemoveRatingIDs(ids ...int) *CustomerUpdateOne {
	cuo.mutation.RemoveRatingIDs(ids...)
	return cuo
}

// RemoveRatings removes ratings edges to Rating.
func (cuo *CustomerUpdateOne) RemoveRatings(r ...*Rating) *CustomerUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cuo.RemoveRatingIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CustomerUpdateOne) Save(ctx context.Context) (*Customer, error) {
	var (
		err  error
		node *Customer
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CustomerUpdateOne) SaveX(ctx context.Context) *Customer {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CustomerUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CustomerUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CustomerUpdateOne) sqlSave(ctx context.Context) (_node *Customer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   customer.Table,
			Columns: customer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: customer.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Customer.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.KratosID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldKratosID,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldName,
		})
	}
	if value, ok := cuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldLastName,
		})
	}
	if value, ok := cuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldEmail,
		})
	}
	if value, ok := cuo.mutation.Phone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: customer.FieldPhone,
		})
	}
	if cuo.mutation.AddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.AddressTable,
			Columns: []string{customer.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   customer.AddressTable,
			Columns: []string{customer.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: address.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.RatingsTable,
			Columns: []string{customer.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rating.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedRatingsIDs(); len(nodes) > 0 && !cuo.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.RatingsTable,
			Columns: []string{customer.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rating.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.RatingsTable,
			Columns: []string{customer.RatingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rating.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Customer{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customer.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
