// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/version/ent/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (_c *UserCreate) SetVersion(v int64) *UserCreate {
	_c.mutation.SetVersion(v)
	return _c
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (_c *UserCreate) SetNillableVersion(v *int64) *UserCreate {
	if v != nil {
		_c.SetVersion(*v)
	}
	return _c
}

// SetStatus sets the "status" field.
func (_c *UserCreate) SetStatus(v user.Status) *UserCreate {
	_c.mutation.SetStatus(v)
	return _c
}

// Mutation returns the UserMutation object of the builder.
func (_c *UserCreate) Mutation() *UserMutation {
	return _c.mutation
}

// Save creates the User in the database.
func (_c *UserCreate) Save(ctx context.Context) (*User, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *UserCreate) SaveX(ctx context.Context) *User {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *UserCreate) defaults() {
	if _, ok := _c.mutation.Version(); !ok {
		v := user.DefaultVersion()
		_c.mutation.SetVersion(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *UserCreate) check() error {
	if _, ok := _c.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "User.version"`)}
	}
	if _, ok := _c.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "User.status"`)}
	}
	if v, ok := _c.mutation.Status(); ok {
		if err := user.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "User.status": %w`, err)}
		}
	}
	return nil
}

func (_c *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Version(); ok {
		_spec.SetField(user.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := _c.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (_c *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*User, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
