// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/agent"
)

// Agent is the model entity for the Agent schema.
type Agent struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// PeerID holds the value of the "peer_id" field.
	PeerID string `json:"peer_id,omitempty"`
	// 是否活动
	Active bool `json:"active,omitempty"`
	// 最后更新时间
	LastUpdateTime time.Time `json:"last_update_time,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Agent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case agent.FieldActive:
			values[i] = new(sql.NullBool)
		case agent.FieldPeerID:
			values[i] = new(sql.NullString)
		case agent.FieldLastUpdateTime:
			values[i] = new(sql.NullTime)
		case agent.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Agent fields.
func (a *Agent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case agent.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case agent.FieldPeerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field peer_id", values[i])
			} else if value.Valid {
				a.PeerID = value.String
			}
		case agent.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				a.Active = value.Bool
			}
		case agent.FieldLastUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_update_time", values[i])
			} else if value.Valid {
				a.LastUpdateTime = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Agent.
// This includes values selected through modifiers, order, etc.
func (a *Agent) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Agent.
// Note that you need to call Agent.Unwrap() before calling this method if this Agent
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Agent) Update() *AgentUpdateOne {
	return NewAgentClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Agent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Agent) Unwrap() *Agent {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Agent is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Agent) String() string {
	var builder strings.Builder
	builder.WriteString("Agent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("peer_id=")
	builder.WriteString(a.PeerID)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", a.Active))
	builder.WriteString(", ")
	builder.WriteString("last_update_time=")
	builder.WriteString(a.LastUpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Agents is a parsable slice of Agent.
type Agents []*Agent
