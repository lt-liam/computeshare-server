// Code generated by ent, DO NOT EDIT.

package ent

import (
	"computeshare-server/internal/data/ent/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StorageCreate is the builder for creating a Storage entity.
type StorageCreate struct {
	config
	mutation *StorageMutation
	hooks    []Hook
}

// SetOwner sets the "owner" field.
func (sc *StorageCreate) SetOwner(s string) *StorageCreate {
	sc.mutation.SetOwner(s)
	return sc
}

// SetType sets the "type" field.
func (sc *StorageCreate) SetType(i int32) *StorageCreate {
	sc.mutation.SetType(i)
	return sc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (sc *StorageCreate) SetNillableType(i *int32) *StorageCreate {
	if i != nil {
		sc.SetType(*i)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *StorageCreate) SetName(s string) *StorageCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetCid sets the "cid" field.
func (sc *StorageCreate) SetCid(s string) *StorageCreate {
	sc.mutation.SetCid(s)
	return sc
}

// SetSize sets the "size" field.
func (sc *StorageCreate) SetSize(i int32) *StorageCreate {
	sc.mutation.SetSize(i)
	return sc
}

// SetLastModify sets the "last_modify" field.
func (sc *StorageCreate) SetLastModify(t time.Time) *StorageCreate {
	sc.mutation.SetLastModify(t)
	return sc
}

// SetNillableLastModify sets the "last_modify" field if the given value is not nil.
func (sc *StorageCreate) SetNillableLastModify(t *time.Time) *StorageCreate {
	if t != nil {
		sc.SetLastModify(*t)
	}
	return sc
}

// SetParentID sets the "parent_id" field.
func (sc *StorageCreate) SetParentID(s string) *StorageCreate {
	sc.mutation.SetParentID(s)
	return sc
}

// SetID sets the "id" field.
func (sc *StorageCreate) SetID(u uuid.UUID) *StorageCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StorageCreate) SetNillableID(u *uuid.UUID) *StorageCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the StorageMutation object of the builder.
func (sc *StorageCreate) Mutation() *StorageMutation {
	return sc.mutation
}

// Save creates the Storage in the database.
func (sc *StorageCreate) Save(ctx context.Context) (*Storage, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StorageCreate) SaveX(ctx context.Context) *Storage {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StorageCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StorageCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StorageCreate) defaults() {
	if _, ok := sc.mutation.GetType(); !ok {
		v := storage.DefaultType
		sc.mutation.SetType(v)
	}
	if _, ok := sc.mutation.LastModify(); !ok {
		v := storage.DefaultLastModify()
		sc.mutation.SetLastModify(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := storage.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StorageCreate) check() error {
	if _, ok := sc.mutation.Owner(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required field "Storage.owner"`)}
	}
	if v, ok := sc.mutation.Owner(); ok {
		if err := storage.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf(`ent: validator failed for field "Storage.owner": %w`, err)}
		}
	}
	if _, ok := sc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Storage.type"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Storage.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := storage.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Storage.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Cid(); !ok {
		return &ValidationError{Name: "cid", err: errors.New(`ent: missing required field "Storage.cid"`)}
	}
	if v, ok := sc.mutation.Cid(); ok {
		if err := storage.CidValidator(v); err != nil {
			return &ValidationError{Name: "cid", err: fmt.Errorf(`ent: validator failed for field "Storage.cid": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Storage.size"`)}
	}
	if _, ok := sc.mutation.LastModify(); !ok {
		return &ValidationError{Name: "last_modify", err: errors.New(`ent: missing required field "Storage.last_modify"`)}
	}
	if _, ok := sc.mutation.ParentID(); !ok {
		return &ValidationError{Name: "parent_id", err: errors.New(`ent: missing required field "Storage.parent_id"`)}
	}
	if v, ok := sc.mutation.ParentID(); ok {
		if err := storage.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf(`ent: validator failed for field "Storage.parent_id": %w`, err)}
		}
	}
	return nil
}

func (sc *StorageCreate) sqlSave(ctx context.Context) (*Storage, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StorageCreate) createSpec() (*Storage, *sqlgraph.CreateSpec) {
	var (
		_node = &Storage{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(storage.Table, sqlgraph.NewFieldSpec(storage.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Owner(); ok {
		_spec.SetField(storage.FieldOwner, field.TypeString, value)
		_node.Owner = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.SetField(storage.FieldType, field.TypeInt32, value)
		_node.Type = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(storage.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Cid(); ok {
		_spec.SetField(storage.FieldCid, field.TypeString, value)
		_node.Cid = &value
	}
	if value, ok := sc.mutation.Size(); ok {
		_spec.SetField(storage.FieldSize, field.TypeInt32, value)
		_node.Size = value
	}
	if value, ok := sc.mutation.LastModify(); ok {
		_spec.SetField(storage.FieldLastModify, field.TypeTime, value)
		_node.LastModify = value
	}
	if value, ok := sc.mutation.ParentID(); ok {
		_spec.SetField(storage.FieldParentID, field.TypeString, value)
		_node.ParentID = value
	}
	return _node, _spec
}

// StorageCreateBulk is the builder for creating many Storage entities in bulk.
type StorageCreateBulk struct {
	config
	builders []*StorageCreate
}

// Save creates the Storage entities in the database.
func (scb *StorageCreateBulk) Save(ctx context.Context) ([]*Storage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Storage, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StorageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StorageCreateBulk) SaveX(ctx context.Context) []*Storage {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StorageCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StorageCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
