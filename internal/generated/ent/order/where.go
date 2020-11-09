// Code generated by entc, DO NOT EDIT.

package order

import (
	"foodworks.ml/m/internal/generated/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OrderState applies equality check predicate on the "order_state" field. It's identical to OrderStateEQ.
func OrderState(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderState), v))
	})
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// OrderStateEQ applies the EQ predicate on the "order_state" field.
func OrderStateEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderState), v))
	})
}

// OrderStateNEQ applies the NEQ predicate on the "order_state" field.
func OrderStateNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderState), v))
	})
}

// OrderStateIn applies the In predicate on the "order_state" field.
func OrderStateIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOrderState), v...))
	})
}

// OrderStateNotIn applies the NotIn predicate on the "order_state" field.
func OrderStateNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOrderState), v...))
	})
}

// OrderStateGT applies the GT predicate on the "order_state" field.
func OrderStateGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderState), v))
	})
}

// OrderStateGTE applies the GTE predicate on the "order_state" field.
func OrderStateGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderState), v))
	})
}

// OrderStateLT applies the LT predicate on the "order_state" field.
func OrderStateLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderState), v))
	})
}

// OrderStateLTE applies the LTE predicate on the "order_state" field.
func OrderStateLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderState), v))
	})
}

// OrderStateContains applies the Contains predicate on the "order_state" field.
func OrderStateContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderState), v))
	})
}

// OrderStateHasPrefix applies the HasPrefix predicate on the "order_state" field.
func OrderStateHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderState), v))
	})
}

// OrderStateHasSuffix applies the HasSuffix predicate on the "order_state" field.
func OrderStateHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderState), v))
	})
}

// OrderStateEqualFold applies the EqualFold predicate on the "order_state" field.
func OrderStateEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderState), v))
	})
}

// OrderStateContainsFold applies the ContainsFold predicate on the "order_state" field.
func OrderStateContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderState), v))
	})
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldQuantity), v))
	})
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldQuantity), v))
	})
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldQuantity), v...))
	})
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldQuantity), v...))
	})
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldQuantity), v))
	})
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldQuantity), v))
	})
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldQuantity), v))
	})
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldQuantity), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCustomer applies the HasEdge predicate on the "customer" edge.
func HasCustomer() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CustomerTable, CustomerPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCustomerWith applies the HasEdge predicate on the "customer" edge with a given conditions (other predicates).
func HasCustomerWith(preds ...predicate.Customer) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CustomerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CustomerTable, CustomerPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		p(s.Not())
	})
}
