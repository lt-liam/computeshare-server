// Code generated by ent, DO NOT EDIT.

package computeimage

import (
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldName, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldImage, v))
}

// Tag applies equality check predicate on the "tag" field. It's identical to TagEQ.
func Tag(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldTag, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldPort, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldContainsFold(FieldName, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldContainsFold(FieldImage, v))
}

// TagEQ applies the EQ predicate on the "tag" field.
func TagEQ(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldTag, v))
}

// TagNEQ applies the NEQ predicate on the "tag" field.
func TagNEQ(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNEQ(FieldTag, v))
}

// TagIn applies the In predicate on the "tag" field.
func TagIn(vs ...string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldIn(FieldTag, vs...))
}

// TagNotIn applies the NotIn predicate on the "tag" field.
func TagNotIn(vs ...string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNotIn(FieldTag, vs...))
}

// TagGT applies the GT predicate on the "tag" field.
func TagGT(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGT(FieldTag, v))
}

// TagGTE applies the GTE predicate on the "tag" field.
func TagGTE(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGTE(FieldTag, v))
}

// TagLT applies the LT predicate on the "tag" field.
func TagLT(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLT(FieldTag, v))
}

// TagLTE applies the LTE predicate on the "tag" field.
func TagLTE(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLTE(FieldTag, v))
}

// TagContains applies the Contains predicate on the "tag" field.
func TagContains(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldContains(FieldTag, v))
}

// TagHasPrefix applies the HasPrefix predicate on the "tag" field.
func TagHasPrefix(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldHasPrefix(FieldTag, v))
}

// TagHasSuffix applies the HasSuffix predicate on the "tag" field.
func TagHasSuffix(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldHasSuffix(FieldTag, v))
}

// TagEqualFold applies the EqualFold predicate on the "tag" field.
func TagEqualFold(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEqualFold(FieldTag, v))
}

// TagContainsFold applies the ContainsFold predicate on the "tag" field.
func TagContainsFold(v string) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldContainsFold(FieldTag, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v int32) predicate.ComputeImage {
	return predicate.ComputeImage(sql.FieldLTE(FieldPort, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ComputeImage) predicate.ComputeImage {
	return predicate.ComputeImage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ComputeImage) predicate.ComputeImage {
	return predicate.ComputeImage(func(s *sql.Selector) {
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
func Not(p predicate.ComputeImage) predicate.ComputeImage {
	return predicate.ComputeImage(func(s *sql.Selector) {
		p(s.Not())
	})
}
