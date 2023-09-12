// Code generated by ent, DO NOT EDIT.

package ent

import (
	"computeshare-server/internal/data/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CountryCallCoding holds the value of the "country_call_coding" field.
	CountryCallCoding string `json:"country_call_coding,omitempty"`
	// TelephoneNumber holds the value of the "telephone_number" field.
	TelephoneNumber string `json:"telephone_number,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// CreateDate holds the value of the "create_date" field.
	CreateDate time.Time `json:"create_date,omitempty"`
	// LastLoginDate holds the value of the "last_login_date" field.
	LastLoginDate *time.Time `json:"last_login_date,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldCountryCallCoding, user.FieldTelephoneNumber, user.FieldPassword:
			values[i] = new(sql.NullString)
		case user.FieldCreateDate, user.FieldLastLoginDate:
			values[i] = new(sql.NullTime)
		case user.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldCountryCallCoding:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country_call_coding", values[i])
			} else if value.Valid {
				u.CountryCallCoding = value.String
			}
		case user.FieldTelephoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field telephone_number", values[i])
			} else if value.Valid {
				u.TelephoneNumber = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldCreateDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_date", values[i])
			} else if value.Valid {
				u.CreateDate = value.Time
			}
		case user.FieldLastLoginDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_date", values[i])
			} else if value.Valid {
				u.LastLoginDate = new(time.Time)
				*u.LastLoginDate = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("country_call_coding=")
	builder.WriteString(u.CountryCallCoding)
	builder.WriteString(", ")
	builder.WriteString("telephone_number=")
	builder.WriteString(u.TelephoneNumber)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("create_date=")
	builder.WriteString(u.CreateDate.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.LastLoginDate; v != nil {
		builder.WriteString("last_login_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
