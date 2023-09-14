// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ComputeInstanceDelete is the builder for deleting a ComputeInstance entity.
type ComputeInstanceDelete struct {
	config
	hooks    []Hook
	mutation *ComputeInstanceMutation
}

// Where appends a list predicates to the ComputeInstanceDelete builder.
func (cid *ComputeInstanceDelete) Where(ps ...predicate.ComputeInstance) *ComputeInstanceDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *ComputeInstanceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cid.sqlExec, cid.mutation, cid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *ComputeInstanceDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *ComputeInstanceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(computeinstance.Table, sqlgraph.NewFieldSpec(computeinstance.FieldID, field.TypeUUID))
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cid.mutation.done = true
	return affected, err
}

// ComputeInstanceDeleteOne is the builder for deleting a single ComputeInstance entity.
type ComputeInstanceDeleteOne struct {
	cid *ComputeInstanceDelete
}

// Where appends a list predicates to the ComputeInstanceDelete builder.
func (cido *ComputeInstanceDeleteOne) Where(ps ...predicate.ComputeInstance) *ComputeInstanceDeleteOne {
	cido.cid.mutation.Where(ps...)
	return cido
}

// Exec executes the deletion query.
func (cido *ComputeInstanceDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{computeinstance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *ComputeInstanceDeleteOne) ExecX(ctx context.Context) {
	if err := cido.Exec(ctx); err != nil {
		panic(err)
	}
}
