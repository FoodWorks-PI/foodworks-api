// Code generated by entc, DO NOT EDIT.

package address

import (
	"foodworks.ml/m/internal/generated/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// StreetLine applies equality check predicate on the "streetLine" field. It's identical to StreetLineEQ.
func StreetLine(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStreetLine), v))
	})
}

// Geom applies equality check predicate on the "geom" field. It's identical to GeomEQ.
func Geom(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGeom), v))
	})
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLatitude), v))
	})
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLatitude), v))
	})
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...float64) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLatitude), v...))
	})
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...float64) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLatitude), v...))
	})
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLatitude), v))
	})
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLatitude), v))
	})
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLatitude), v))
	})
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLatitude), v))
	})
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLongitude), v))
	})
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLongitude), v))
	})
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...float64) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLongitude), v...))
	})
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...float64) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLongitude), v...))
	})
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLongitude), v))
	})
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLongitude), v))
	})
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLongitude), v))
	})
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v float64) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLongitude), v))
	})
}

// StreetLineEQ applies the EQ predicate on the "streetLine" field.
func StreetLineEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStreetLine), v))
	})
}

// StreetLineNEQ applies the NEQ predicate on the "streetLine" field.
func StreetLineNEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStreetLine), v))
	})
}

// StreetLineIn applies the In predicate on the "streetLine" field.
func StreetLineIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStreetLine), v...))
	})
}

// StreetLineNotIn applies the NotIn predicate on the "streetLine" field.
func StreetLineNotIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStreetLine), v...))
	})
}

// StreetLineGT applies the GT predicate on the "streetLine" field.
func StreetLineGT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStreetLine), v))
	})
}

// StreetLineGTE applies the GTE predicate on the "streetLine" field.
func StreetLineGTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStreetLine), v))
	})
}

// StreetLineLT applies the LT predicate on the "streetLine" field.
func StreetLineLT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStreetLine), v))
	})
}

// StreetLineLTE applies the LTE predicate on the "streetLine" field.
func StreetLineLTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStreetLine), v))
	})
}

// StreetLineContains applies the Contains predicate on the "streetLine" field.
func StreetLineContains(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStreetLine), v))
	})
}

// StreetLineHasPrefix applies the HasPrefix predicate on the "streetLine" field.
func StreetLineHasPrefix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStreetLine), v))
	})
}

// StreetLineHasSuffix applies the HasSuffix predicate on the "streetLine" field.
func StreetLineHasSuffix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStreetLine), v))
	})
}

// StreetLineEqualFold applies the EqualFold predicate on the "streetLine" field.
func StreetLineEqualFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStreetLine), v))
	})
}

// StreetLineContainsFold applies the ContainsFold predicate on the "streetLine" field.
func StreetLineContainsFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStreetLine), v))
	})
}

// GeomEQ applies the EQ predicate on the "geom" field.
func GeomEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGeom), v))
	})
}

// GeomNEQ applies the NEQ predicate on the "geom" field.
func GeomNEQ(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGeom), v))
	})
}

// GeomIn applies the In predicate on the "geom" field.
func GeomIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGeom), v...))
	})
}

// GeomNotIn applies the NotIn predicate on the "geom" field.
func GeomNotIn(vs ...string) predicate.Address {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Address(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGeom), v...))
	})
}

// GeomGT applies the GT predicate on the "geom" field.
func GeomGT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGeom), v))
	})
}

// GeomGTE applies the GTE predicate on the "geom" field.
func GeomGTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGeom), v))
	})
}

// GeomLT applies the LT predicate on the "geom" field.
func GeomLT(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGeom), v))
	})
}

// GeomLTE applies the LTE predicate on the "geom" field.
func GeomLTE(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGeom), v))
	})
}

// GeomContains applies the Contains predicate on the "geom" field.
func GeomContains(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGeom), v))
	})
}

// GeomHasPrefix applies the HasPrefix predicate on the "geom" field.
func GeomHasPrefix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGeom), v))
	})
}

// GeomHasSuffix applies the HasSuffix predicate on the "geom" field.
func GeomHasSuffix(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGeom), v))
	})
}

// GeomIsNil applies the IsNil predicate on the "geom" field.
func GeomIsNil() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGeom)))
	})
}

// GeomNotNil applies the NotNil predicate on the "geom" field.
func GeomNotNil() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGeom)))
	})
}

// GeomEqualFold applies the EqualFold predicate on the "geom" field.
func GeomEqualFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGeom), v))
	})
}

// GeomContainsFold applies the ContainsFold predicate on the "geom" field.
func GeomContainsFold(v string) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGeom), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		p(s.Not())
	})
}
