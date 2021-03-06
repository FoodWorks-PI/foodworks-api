// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"foodworks.ml/m/internal/generated/ent/imagepath"
	"foodworks.ml/m/internal/generated/ent/order"
	"foodworks.ml/m/internal/generated/ent/predicate"
	"foodworks.ml/m/internal/generated/ent/product"
	"foodworks.ml/m/internal/generated/ent/rating"
	"foodworks.ml/m/internal/generated/ent/restaurant"
	"foodworks.ml/m/internal/generated/ent/tag"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks      []Hook
	mutation   *ProductMutation
	predicates []predicate.Product
}

// Where adds a new predicate for the builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the description field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the description field if the given value is not nil.
func (pu *ProductUpdate) SetNillableDescription(s *string) *ProductUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of description.
func (pu *ProductUpdate) ClearDescription() *ProductUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetCost sets the cost field.
func (pu *ProductUpdate) SetCost(i int) *ProductUpdate {
	pu.mutation.ResetCost()
	pu.mutation.SetCost(i)
	return pu
}

// AddCost adds i to cost.
func (pu *ProductUpdate) AddCost(i int) *ProductUpdate {
	pu.mutation.AddCost(i)
	return pu
}

// SetIsActive sets the is_active field.
func (pu *ProductUpdate) SetIsActive(b bool) *ProductUpdate {
	pu.mutation.SetIsActive(b)
	return pu
}

// AddTagIDs adds the tags edge to Tag by ids.
func (pu *ProductUpdate) AddTagIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddTagIDs(ids...)
	return pu
}

// AddTags adds the tags edges to Tag.
func (pu *ProductUpdate) AddTags(t ...*Tag) *ProductUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTagIDs(ids...)
}

// AddRatingIDs adds the ratings edge to Rating by ids.
func (pu *ProductUpdate) AddRatingIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddRatingIDs(ids...)
	return pu
}

// AddRatings adds the ratings edges to Rating.
func (pu *ProductUpdate) AddRatings(r ...*Rating) *ProductUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRatingIDs(ids...)
}

// AddRestaurantIDs adds the restaurant edge to Restaurant by ids.
func (pu *ProductUpdate) AddRestaurantIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddRestaurantIDs(ids...)
	return pu
}

// AddRestaurant adds the restaurant edges to Restaurant.
func (pu *ProductUpdate) AddRestaurant(r ...*Restaurant) *ProductUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRestaurantIDs(ids...)
}

// AddOrderIDs adds the orders edge to Order by ids.
func (pu *ProductUpdate) AddOrderIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddOrderIDs(ids...)
	return pu
}

// AddOrders adds the orders edges to Order.
func (pu *ProductUpdate) AddOrders(o ...*Order) *ProductUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.AddOrderIDs(ids...)
}

// AddImageIDs adds the images edge to ImagePath by ids.
func (pu *ProductUpdate) AddImageIDs(ids ...int) *ProductUpdate {
	pu.mutation.AddImageIDs(ids...)
	return pu
}

// AddImages adds the images edges to ImagePath.
func (pu *ProductUpdate) AddImages(i ...*ImagePath) *ProductUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.AddImageIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearTags clears all "tags" edges to type Tag.
func (pu *ProductUpdate) ClearTags() *ProductUpdate {
	pu.mutation.ClearTags()
	return pu
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (pu *ProductUpdate) RemoveTagIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveTagIDs(ids...)
	return pu
}

// RemoveTags removes tags edges to Tag.
func (pu *ProductUpdate) RemoveTags(t ...*Tag) *ProductUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTagIDs(ids...)
}

// ClearRatings clears all "ratings" edges to type Rating.
func (pu *ProductUpdate) ClearRatings() *ProductUpdate {
	pu.mutation.ClearRatings()
	return pu
}

// RemoveRatingIDs removes the ratings edge to Rating by ids.
func (pu *ProductUpdate) RemoveRatingIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveRatingIDs(ids...)
	return pu
}

// RemoveRatings removes ratings edges to Rating.
func (pu *ProductUpdate) RemoveRatings(r ...*Rating) *ProductUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRatingIDs(ids...)
}

// ClearRestaurant clears all "restaurant" edges to type Restaurant.
func (pu *ProductUpdate) ClearRestaurant() *ProductUpdate {
	pu.mutation.ClearRestaurant()
	return pu
}

// RemoveRestaurantIDs removes the restaurant edge to Restaurant by ids.
func (pu *ProductUpdate) RemoveRestaurantIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveRestaurantIDs(ids...)
	return pu
}

// RemoveRestaurant removes restaurant edges to Restaurant.
func (pu *ProductUpdate) RemoveRestaurant(r ...*Restaurant) *ProductUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRestaurantIDs(ids...)
}

// ClearOrders clears all "orders" edges to type Order.
func (pu *ProductUpdate) ClearOrders() *ProductUpdate {
	pu.mutation.ClearOrders()
	return pu
}

// RemoveOrderIDs removes the orders edge to Order by ids.
func (pu *ProductUpdate) RemoveOrderIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveOrderIDs(ids...)
	return pu
}

// RemoveOrders removes orders edges to Order.
func (pu *ProductUpdate) RemoveOrders(o ...*Order) *ProductUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pu.RemoveOrderIDs(ids...)
}

// ClearImages clears all "images" edges to type ImagePath.
func (pu *ProductUpdate) ClearImages() *ProductUpdate {
	pu.mutation.ClearImages()
	return pu
}

// RemoveImageIDs removes the images edge to ImagePath by ids.
func (pu *ProductUpdate) RemoveImageIDs(ids ...int) *ProductUpdate {
	pu.mutation.RemoveImageIDs(ids...)
	return pu
}

// RemoveImages removes images edges to ImagePath.
func (pu *ProductUpdate) RemoveImages(i ...*ImagePath) *ProductUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pu.RemoveImageIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDescription,
		})
	}
	if pu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: product.FieldDescription,
		})
	}
	if value, ok := pu.mutation.Cost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldCost,
		})
	}
	if value, ok := pu.mutation.AddedCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldCost,
		})
	}
	if value, ok := pu.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: product.FieldIsActive,
		})
	}
	if pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !pu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.RatingsTable,
			Columns: []string{product.RatingsColumn},
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
	if nodes := pu.mutation.RemovedRatingsIDs(); len(nodes) > 0 && !pu.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.RatingsTable,
			Columns: []string{product.RatingsColumn},
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
	if nodes := pu.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.RatingsTable,
			Columns: []string{product.RatingsColumn},
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
	if pu.mutation.RestaurantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.RestaurantTable,
			Columns: product.RestaurantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: restaurant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRestaurantIDs(); len(nodes) > 0 && !pu.mutation.RestaurantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.RestaurantTable,
			Columns: product.RestaurantPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RestaurantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.RestaurantTable,
			Columns: product.RestaurantPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: product.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !pu.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: product.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: product.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ImagesTable,
			Columns: []string{product.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imagepath.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedImagesIDs(); len(nodes) > 0 && !pu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ImagesTable,
			Columns: []string{product.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imagepath.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ImagesTable,
			Columns: []string{product.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imagepath.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// SetName sets the name field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the description field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the description field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableDescription(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of description.
func (puo *ProductUpdateOne) ClearDescription() *ProductUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetCost sets the cost field.
func (puo *ProductUpdateOne) SetCost(i int) *ProductUpdateOne {
	puo.mutation.ResetCost()
	puo.mutation.SetCost(i)
	return puo
}

// AddCost adds i to cost.
func (puo *ProductUpdateOne) AddCost(i int) *ProductUpdateOne {
	puo.mutation.AddCost(i)
	return puo
}

// SetIsActive sets the is_active field.
func (puo *ProductUpdateOne) SetIsActive(b bool) *ProductUpdateOne {
	puo.mutation.SetIsActive(b)
	return puo
}

// AddTagIDs adds the tags edge to Tag by ids.
func (puo *ProductUpdateOne) AddTagIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddTagIDs(ids...)
	return puo
}

// AddTags adds the tags edges to Tag.
func (puo *ProductUpdateOne) AddTags(t ...*Tag) *ProductUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTagIDs(ids...)
}

// AddRatingIDs adds the ratings edge to Rating by ids.
func (puo *ProductUpdateOne) AddRatingIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddRatingIDs(ids...)
	return puo
}

// AddRatings adds the ratings edges to Rating.
func (puo *ProductUpdateOne) AddRatings(r ...*Rating) *ProductUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRatingIDs(ids...)
}

// AddRestaurantIDs adds the restaurant edge to Restaurant by ids.
func (puo *ProductUpdateOne) AddRestaurantIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddRestaurantIDs(ids...)
	return puo
}

// AddRestaurant adds the restaurant edges to Restaurant.
func (puo *ProductUpdateOne) AddRestaurant(r ...*Restaurant) *ProductUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRestaurantIDs(ids...)
}

// AddOrderIDs adds the orders edge to Order by ids.
func (puo *ProductUpdateOne) AddOrderIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddOrderIDs(ids...)
	return puo
}

// AddOrders adds the orders edges to Order.
func (puo *ProductUpdateOne) AddOrders(o ...*Order) *ProductUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.AddOrderIDs(ids...)
}

// AddImageIDs adds the images edge to ImagePath by ids.
func (puo *ProductUpdateOne) AddImageIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.AddImageIDs(ids...)
	return puo
}

// AddImages adds the images edges to ImagePath.
func (puo *ProductUpdateOne) AddImages(i ...*ImagePath) *ProductUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.AddImageIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearTags clears all "tags" edges to type Tag.
func (puo *ProductUpdateOne) ClearTags() *ProductUpdateOne {
	puo.mutation.ClearTags()
	return puo
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (puo *ProductUpdateOne) RemoveTagIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveTagIDs(ids...)
	return puo
}

// RemoveTags removes tags edges to Tag.
func (puo *ProductUpdateOne) RemoveTags(t ...*Tag) *ProductUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTagIDs(ids...)
}

// ClearRatings clears all "ratings" edges to type Rating.
func (puo *ProductUpdateOne) ClearRatings() *ProductUpdateOne {
	puo.mutation.ClearRatings()
	return puo
}

// RemoveRatingIDs removes the ratings edge to Rating by ids.
func (puo *ProductUpdateOne) RemoveRatingIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveRatingIDs(ids...)
	return puo
}

// RemoveRatings removes ratings edges to Rating.
func (puo *ProductUpdateOne) RemoveRatings(r ...*Rating) *ProductUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRatingIDs(ids...)
}

// ClearRestaurant clears all "restaurant" edges to type Restaurant.
func (puo *ProductUpdateOne) ClearRestaurant() *ProductUpdateOne {
	puo.mutation.ClearRestaurant()
	return puo
}

// RemoveRestaurantIDs removes the restaurant edge to Restaurant by ids.
func (puo *ProductUpdateOne) RemoveRestaurantIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveRestaurantIDs(ids...)
	return puo
}

// RemoveRestaurant removes restaurant edges to Restaurant.
func (puo *ProductUpdateOne) RemoveRestaurant(r ...*Restaurant) *ProductUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRestaurantIDs(ids...)
}

// ClearOrders clears all "orders" edges to type Order.
func (puo *ProductUpdateOne) ClearOrders() *ProductUpdateOne {
	puo.mutation.ClearOrders()
	return puo
}

// RemoveOrderIDs removes the orders edge to Order by ids.
func (puo *ProductUpdateOne) RemoveOrderIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveOrderIDs(ids...)
	return puo
}

// RemoveOrders removes orders edges to Order.
func (puo *ProductUpdateOne) RemoveOrders(o ...*Order) *ProductUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return puo.RemoveOrderIDs(ids...)
}

// ClearImages clears all "images" edges to type ImagePath.
func (puo *ProductUpdateOne) ClearImages() *ProductUpdateOne {
	puo.mutation.ClearImages()
	return puo
}

// RemoveImageIDs removes the images edge to ImagePath by ids.
func (puo *ProductUpdateOne) RemoveImageIDs(ids ...int) *ProductUpdateOne {
	puo.mutation.RemoveImageIDs(ids...)
	return puo
}

// RemoveImages removes images edges to ImagePath.
func (puo *ProductUpdateOne) RemoveImages(i ...*ImagePath) *ProductUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return puo.RemoveImageIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   product.Table,
			Columns: product.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: product.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Product.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldName,
		})
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: product.FieldDescription,
		})
	}
	if puo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: product.FieldDescription,
		})
	}
	if value, ok := puo.mutation.Cost(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldCost,
		})
	}
	if value, ok := puo.mutation.AddedCost(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: product.FieldCost,
		})
	}
	if value, ok := puo.mutation.IsActive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: product.FieldIsActive,
		})
	}
	if puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !puo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.TagsTable,
			Columns: product.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.RatingsTable,
			Columns: []string{product.RatingsColumn},
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
	if nodes := puo.mutation.RemovedRatingsIDs(); len(nodes) > 0 && !puo.mutation.RatingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.RatingsTable,
			Columns: []string{product.RatingsColumn},
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
	if nodes := puo.mutation.RatingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.RatingsTable,
			Columns: []string{product.RatingsColumn},
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
	if puo.mutation.RestaurantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.RestaurantTable,
			Columns: product.RestaurantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: restaurant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRestaurantIDs(); len(nodes) > 0 && !puo.mutation.RestaurantCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.RestaurantTable,
			Columns: product.RestaurantPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RestaurantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   product.RestaurantTable,
			Columns: product.RestaurantPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: product.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedOrdersIDs(); len(nodes) > 0 && !puo.mutation.OrdersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: product.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrdersTable,
			Columns: product.OrdersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ImagesTable,
			Columns: []string{product.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imagepath.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !puo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ImagesTable,
			Columns: []string{product.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imagepath.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ImagesTable,
			Columns: []string{product.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: imagepath.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
