// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv1/car"
	"entgo.io/ent/entc/integration/migrate/entv1/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (_c *UserCreate) SetAge(v int32) *UserCreate {
	_c.mutation.SetAge(v)
	return _c
}

// SetName sets the "name" field.
func (_c *UserCreate) SetName(v string) *UserCreate {
	_c.mutation.SetName(v)
	return _c
}

// SetDescription sets the "description" field.
func (_c *UserCreate) SetDescription(v string) *UserCreate {
	_c.mutation.SetDescription(v)
	return _c
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (_c *UserCreate) SetNillableDescription(v *string) *UserCreate {
	if v != nil {
		_c.SetDescription(*v)
	}
	return _c
}

// SetNickname sets the "nickname" field.
func (_c *UserCreate) SetNickname(v string) *UserCreate {
	_c.mutation.SetNickname(v)
	return _c
}

// SetAddress sets the "address" field.
func (_c *UserCreate) SetAddress(v string) *UserCreate {
	_c.mutation.SetAddress(v)
	return _c
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (_c *UserCreate) SetNillableAddress(v *string) *UserCreate {
	if v != nil {
		_c.SetAddress(*v)
	}
	return _c
}

// SetRenamed sets the "renamed" field.
func (_c *UserCreate) SetRenamed(v string) *UserCreate {
	_c.mutation.SetRenamed(v)
	return _c
}

// SetNillableRenamed sets the "renamed" field if the given value is not nil.
func (_c *UserCreate) SetNillableRenamed(v *string) *UserCreate {
	if v != nil {
		_c.SetRenamed(*v)
	}
	return _c
}

// SetOldToken sets the "old_token" field.
func (_c *UserCreate) SetOldToken(v string) *UserCreate {
	_c.mutation.SetOldToken(v)
	return _c
}

// SetNillableOldToken sets the "old_token" field if the given value is not nil.
func (_c *UserCreate) SetNillableOldToken(v *string) *UserCreate {
	if v != nil {
		_c.SetOldToken(*v)
	}
	return _c
}

// SetBlob sets the "blob" field.
func (_c *UserCreate) SetBlob(v []byte) *UserCreate {
	_c.mutation.SetBlob(v)
	return _c
}

// SetState sets the "state" field.
func (_c *UserCreate) SetState(v user.State) *UserCreate {
	_c.mutation.SetState(v)
	return _c
}

// SetNillableState sets the "state" field if the given value is not nil.
func (_c *UserCreate) SetNillableState(v *user.State) *UserCreate {
	if v != nil {
		_c.SetState(*v)
	}
	return _c
}

// SetStatus sets the "status" field.
func (_c *UserCreate) SetStatus(v string) *UserCreate {
	_c.mutation.SetStatus(v)
	return _c
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (_c *UserCreate) SetNillableStatus(v *string) *UserCreate {
	if v != nil {
		_c.SetStatus(*v)
	}
	return _c
}

// SetWorkplace sets the "workplace" field.
func (_c *UserCreate) SetWorkplace(v string) *UserCreate {
	_c.mutation.SetWorkplace(v)
	return _c
}

// SetNillableWorkplace sets the "workplace" field if the given value is not nil.
func (_c *UserCreate) SetNillableWorkplace(v *string) *UserCreate {
	if v != nil {
		_c.SetWorkplace(*v)
	}
	return _c
}

// SetDropOptional sets the "drop_optional" field.
func (_c *UserCreate) SetDropOptional(v string) *UserCreate {
	_c.mutation.SetDropOptional(v)
	return _c
}

// SetNillableDropOptional sets the "drop_optional" field if the given value is not nil.
func (_c *UserCreate) SetNillableDropOptional(v *string) *UserCreate {
	if v != nil {
		_c.SetDropOptional(*v)
	}
	return _c
}

// SetID sets the "id" field.
func (_c *UserCreate) SetID(v int) *UserCreate {
	_c.mutation.SetID(v)
	return _c
}

// SetParentID sets the "parent" edge to the User entity by ID.
func (_c *UserCreate) SetParentID(id int) *UserCreate {
	_c.mutation.SetParentID(id)
	return _c
}

// SetNillableParentID sets the "parent" edge to the User entity by ID if the given value is not nil.
func (_c *UserCreate) SetNillableParentID(id *int) *UserCreate {
	if id != nil {
		_c = _c.SetParentID(*id)
	}
	return _c
}

// SetParent sets the "parent" edge to the User entity.
func (_c *UserCreate) SetParent(v *User) *UserCreate {
	return _c.SetParentID(v.ID)
}

// AddChildIDs adds the "children" edge to the User entity by IDs.
func (_c *UserCreate) AddChildIDs(ids ...int) *UserCreate {
	_c.mutation.AddChildIDs(ids...)
	return _c
}

// AddChildren adds the "children" edges to the User entity.
func (_c *UserCreate) AddChildren(v ...*User) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddChildIDs(ids...)
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (_c *UserCreate) SetSpouseID(id int) *UserCreate {
	_c.mutation.SetSpouseID(id)
	return _c
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (_c *UserCreate) SetNillableSpouseID(id *int) *UserCreate {
	if id != nil {
		_c = _c.SetSpouseID(*id)
	}
	return _c
}

// SetSpouse sets the "spouse" edge to the User entity.
func (_c *UserCreate) SetSpouse(v *User) *UserCreate {
	return _c.SetSpouseID(v.ID)
}

// SetCarID sets the "car" edge to the Car entity by ID.
func (_c *UserCreate) SetCarID(id int) *UserCreate {
	_c.mutation.SetCarID(id)
	return _c
}

// SetNillableCarID sets the "car" edge to the Car entity by ID if the given value is not nil.
func (_c *UserCreate) SetNillableCarID(id *int) *UserCreate {
	if id != nil {
		_c = _c.SetCarID(*id)
	}
	return _c
}

// SetCar sets the "car" edge to the Car entity.
func (_c *UserCreate) SetCar(v *Car) *UserCreate {
	return _c.SetCarID(v.ID)
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
	if _, ok := _c.mutation.OldToken(); !ok {
		v := user.DefaultOldToken()
		_c.mutation.SetOldToken(v)
	}
	if _, ok := _c.mutation.State(); !ok {
		v := user.DefaultState
		_c.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *UserCreate) check() error {
	if _, ok := _c.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`entv1: missing required field "User.age"`)}
	}
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`entv1: missing required field "User.name"`)}
	}
	if v, ok := _c.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`entv1: validator failed for field "User.name": %w`, err)}
		}
	}
	if _, ok := _c.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`entv1: missing required field "User.nickname"`)}
	}
	if _, ok := _c.mutation.OldToken(); !ok {
		return &ValidationError{Name: "old_token", err: errors.New(`entv1: missing required field "User.old_token"`)}
	}
	if v, ok := _c.mutation.Blob(); ok {
		if err := user.BlobValidator(v); err != nil {
			return &ValidationError{Name: "blob", err: fmt.Errorf(`entv1: validator failed for field "User.blob": %w`, err)}
		}
	}
	if v, ok := _c.mutation.State(); ok {
		if err := user.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`entv1: validator failed for field "User.state": %w`, err)}
		}
	}
	if v, ok := _c.mutation.Workplace(); ok {
		if err := user.WorkplaceValidator(v); err != nil {
			return &ValidationError{Name: "workplace", err: fmt.Errorf(`entv1: validator failed for field "User.workplace": %w`, err)}
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if id, ok := _c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := _c.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt32, value)
		_node.Age = value
	}
	if value, ok := _c.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := _c.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := _c.mutation.Nickname(); ok {
		_spec.SetField(user.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := _c.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := _c.mutation.Renamed(); ok {
		_spec.SetField(user.FieldRenamed, field.TypeString, value)
		_node.Renamed = value
	}
	if value, ok := _c.mutation.OldToken(); ok {
		_spec.SetField(user.FieldOldToken, field.TypeString, value)
		_node.OldToken = value
	}
	if value, ok := _c.mutation.Blob(); ok {
		_spec.SetField(user.FieldBlob, field.TypeBytes, value)
		_node.Blob = value
	}
	if value, ok := _c.mutation.State(); ok {
		_spec.SetField(user.FieldState, field.TypeEnum, value)
		_node.State = value
	}
	if value, ok := _c.mutation.Status(); ok {
		_spec.SetField(user.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := _c.mutation.Workplace(); ok {
		_spec.SetField(user.FieldWorkplace, field.TypeString, value)
		_node.Workplace = value
	}
	if value, ok := _c.mutation.DropOptional(); ok {
		_spec.SetField(user.FieldDropOptional, field.TypeString, value)
		_node.DropOptional = value
	}
	if nodes := _c.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_spouse = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CarTable,
			Columns: []string{user.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
