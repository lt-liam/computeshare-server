// Code generated by ent, DO NOT EDIT.

package user

import (
	"computeshare-server/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CountryCallCoding applies equality check predicate on the "country_call_coding" field. It's identical to CountryCallCodingEQ.
func CountryCallCoding(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountryCallCoding, v))
}

// TelephoneNumber applies equality check predicate on the "telephone_number" field. It's identical to TelephoneNumberEQ.
func TelephoneNumber(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTelephoneNumber, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// CreateDate applies equality check predicate on the "create_date" field. It's identical to CreateDateEQ.
func CreateDate(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateDate, v))
}

// LastLoginDate applies equality check predicate on the "last_login_date" field. It's identical to LastLoginDateEQ.
func LastLoginDate(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLoginDate, v))
}

// CountryCallCodingEQ applies the EQ predicate on the "country_call_coding" field.
func CountryCallCodingEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCountryCallCoding, v))
}

// CountryCallCodingNEQ applies the NEQ predicate on the "country_call_coding" field.
func CountryCallCodingNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCountryCallCoding, v))
}

// CountryCallCodingIn applies the In predicate on the "country_call_coding" field.
func CountryCallCodingIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldCountryCallCoding, vs...))
}

// CountryCallCodingNotIn applies the NotIn predicate on the "country_call_coding" field.
func CountryCallCodingNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCountryCallCoding, vs...))
}

// CountryCallCodingGT applies the GT predicate on the "country_call_coding" field.
func CountryCallCodingGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldCountryCallCoding, v))
}

// CountryCallCodingGTE applies the GTE predicate on the "country_call_coding" field.
func CountryCallCodingGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCountryCallCoding, v))
}

// CountryCallCodingLT applies the LT predicate on the "country_call_coding" field.
func CountryCallCodingLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldCountryCallCoding, v))
}

// CountryCallCodingLTE applies the LTE predicate on the "country_call_coding" field.
func CountryCallCodingLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCountryCallCoding, v))
}

// CountryCallCodingContains applies the Contains predicate on the "country_call_coding" field.
func CountryCallCodingContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldCountryCallCoding, v))
}

// CountryCallCodingHasPrefix applies the HasPrefix predicate on the "country_call_coding" field.
func CountryCallCodingHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldCountryCallCoding, v))
}

// CountryCallCodingHasSuffix applies the HasSuffix predicate on the "country_call_coding" field.
func CountryCallCodingHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldCountryCallCoding, v))
}

// CountryCallCodingEqualFold applies the EqualFold predicate on the "country_call_coding" field.
func CountryCallCodingEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldCountryCallCoding, v))
}

// CountryCallCodingContainsFold applies the ContainsFold predicate on the "country_call_coding" field.
func CountryCallCodingContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldCountryCallCoding, v))
}

// TelephoneNumberEQ applies the EQ predicate on the "telephone_number" field.
func TelephoneNumberEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldTelephoneNumber, v))
}

// TelephoneNumberNEQ applies the NEQ predicate on the "telephone_number" field.
func TelephoneNumberNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldTelephoneNumber, v))
}

// TelephoneNumberIn applies the In predicate on the "telephone_number" field.
func TelephoneNumberIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldTelephoneNumber, vs...))
}

// TelephoneNumberNotIn applies the NotIn predicate on the "telephone_number" field.
func TelephoneNumberNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldTelephoneNumber, vs...))
}

// TelephoneNumberGT applies the GT predicate on the "telephone_number" field.
func TelephoneNumberGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldTelephoneNumber, v))
}

// TelephoneNumberGTE applies the GTE predicate on the "telephone_number" field.
func TelephoneNumberGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldTelephoneNumber, v))
}

// TelephoneNumberLT applies the LT predicate on the "telephone_number" field.
func TelephoneNumberLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldTelephoneNumber, v))
}

// TelephoneNumberLTE applies the LTE predicate on the "telephone_number" field.
func TelephoneNumberLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldTelephoneNumber, v))
}

// TelephoneNumberContains applies the Contains predicate on the "telephone_number" field.
func TelephoneNumberContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldTelephoneNumber, v))
}

// TelephoneNumberHasPrefix applies the HasPrefix predicate on the "telephone_number" field.
func TelephoneNumberHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldTelephoneNumber, v))
}

// TelephoneNumberHasSuffix applies the HasSuffix predicate on the "telephone_number" field.
func TelephoneNumberHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldTelephoneNumber, v))
}

// TelephoneNumberEqualFold applies the EqualFold predicate on the "telephone_number" field.
func TelephoneNumberEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldTelephoneNumber, v))
}

// TelephoneNumberContainsFold applies the ContainsFold predicate on the "telephone_number" field.
func TelephoneNumberContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldTelephoneNumber, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// CreateDateEQ applies the EQ predicate on the "create_date" field.
func CreateDateEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreateDate, v))
}

// CreateDateNEQ applies the NEQ predicate on the "create_date" field.
func CreateDateNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreateDate, v))
}

// CreateDateIn applies the In predicate on the "create_date" field.
func CreateDateIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreateDate, vs...))
}

// CreateDateNotIn applies the NotIn predicate on the "create_date" field.
func CreateDateNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreateDate, vs...))
}

// CreateDateGT applies the GT predicate on the "create_date" field.
func CreateDateGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreateDate, v))
}

// CreateDateGTE applies the GTE predicate on the "create_date" field.
func CreateDateGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreateDate, v))
}

// CreateDateLT applies the LT predicate on the "create_date" field.
func CreateDateLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreateDate, v))
}

// CreateDateLTE applies the LTE predicate on the "create_date" field.
func CreateDateLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreateDate, v))
}

// LastLoginDateEQ applies the EQ predicate on the "last_login_date" field.
func LastLoginDateEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLoginDate, v))
}

// LastLoginDateNEQ applies the NEQ predicate on the "last_login_date" field.
func LastLoginDateNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastLoginDate, v))
}

// LastLoginDateIn applies the In predicate on the "last_login_date" field.
func LastLoginDateIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastLoginDate, vs...))
}

// LastLoginDateNotIn applies the NotIn predicate on the "last_login_date" field.
func LastLoginDateNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastLoginDate, vs...))
}

// LastLoginDateGT applies the GT predicate on the "last_login_date" field.
func LastLoginDateGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastLoginDate, v))
}

// LastLoginDateGTE applies the GTE predicate on the "last_login_date" field.
func LastLoginDateGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastLoginDate, v))
}

// LastLoginDateLT applies the LT predicate on the "last_login_date" field.
func LastLoginDateLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastLoginDate, v))
}

// LastLoginDateLTE applies the LTE predicate on the "last_login_date" field.
func LastLoginDateLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastLoginDate, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
