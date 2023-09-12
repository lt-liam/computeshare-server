// Code generated by ent, DO NOT EDIT.

package ent

import (
	"computeshare-server/internal/data/ent/agent"
	"computeshare-server/internal/data/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgentUpdate is the builder for updating Agent entities.
type AgentUpdate struct {
	config
	hooks    []Hook
	mutation *AgentMutation
}

// Where appends a list predicates to the AgentUpdate builder.
func (au *AgentUpdate) Where(ps ...predicate.Agent) *AgentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *AgentUpdate) SetName(s string) *AgentUpdate {
	au.mutation.SetName(s)
	return au
}

// Mutation returns the AgentMutation object of the builder.
func (au *AgentUpdate) Mutation() *AgentMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AgentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AgentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AgentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AgentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AgentUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := agent.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Agent.name": %w`, err)}
		}
	}
	return nil
}

func (au *AgentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(agent.Table, agent.Columns, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(agent.FieldName, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AgentUpdateOne is the builder for updating a single Agent entity.
type AgentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AgentMutation
}

// SetName sets the "name" field.
func (auo *AgentUpdateOne) SetName(s string) *AgentUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// Mutation returns the AgentMutation object of the builder.
func (auo *AgentUpdateOne) Mutation() *AgentMutation {
	return auo.mutation
}

// Where appends a list predicates to the AgentUpdate builder.
func (auo *AgentUpdateOne) Where(ps ...predicate.Agent) *AgentUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AgentUpdateOne) Select(field string, fields ...string) *AgentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Agent entity.
func (auo *AgentUpdateOne) Save(ctx context.Context) (*Agent, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AgentUpdateOne) SaveX(ctx context.Context) *Agent {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AgentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AgentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AgentUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := agent.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Agent.name": %w`, err)}
		}
	}
	return nil
}

func (auo *AgentUpdateOne) sqlSave(ctx context.Context) (_node *Agent, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(agent.Table, agent.Columns, sqlgraph.NewFieldSpec(agent.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Agent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, agent.FieldID)
		for _, f := range fields {
			if !agent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != agent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(agent.FieldName, field.TypeString, value)
	}
	_node = &Agent{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
