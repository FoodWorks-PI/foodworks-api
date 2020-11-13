// Code generated by entc, DO NOT EDIT.

package paymentmethod

import (
	"foodworks.ml/m/internal/generated/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
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
func IDGT(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldData), v))
	})
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldData), v))
	})
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.PaymentMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldData), v...))
	})
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.PaymentMethod {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PaymentMethod(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldData), v...))
	})
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldData), v))
	})
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldData), v))
	})
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldData), v))
	})
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldData), v))
	})
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldData), v))
	})
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldData), v))
	})
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldData), v))
	})
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldData), v))
	})
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldData), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.PaymentMethod) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.PaymentMethod) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
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
func Not(p predicate.PaymentMethod) predicate.PaymentMethod {
	return predicate.PaymentMethod(func(s *sql.Selector) {
		p(s.Not())
	})
}
