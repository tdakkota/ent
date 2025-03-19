// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RentalCreate is the builder for creating a Rental entity.
type RentalCreate struct {
	config
	mutation *RentalMutation
	hooks    []Hook
}

// SetDate sets the "date" field.
func (_c *RentalCreate) SetDate(v time.Time) *RentalCreate {
	_c.mutation.SetDate(v)
	return _c
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (_c *RentalCreate) SetNillableDate(v *time.Time) *RentalCreate {
	if v != nil {
		_c.SetDate(*v)
	}
	return _c
}

// SetUserID sets the "user_id" field.
func (_c *RentalCreate) SetUserID(v int) *RentalCreate {
	_c.mutation.SetUserID(v)
	return _c
}

// SetCarID sets the "car_id" field.
func (_c *RentalCreate) SetCarID(v uuid.UUID) *RentalCreate {
	_c.mutation.SetCarID(v)
	return _c
}

// SetUser sets the "user" edge to the User entity.
func (_c *RentalCreate) SetUser(v *User) *RentalCreate {
	return _c.SetUserID(v.ID)
}

// SetCar sets the "car" edge to the Car entity.
func (_c *RentalCreate) SetCar(v *Car) *RentalCreate {
	return _c.SetCarID(v.ID)
}

// Mutation returns the RentalMutation object of the builder.
func (_c *RentalCreate) Mutation() *RentalMutation {
	return _c.mutation
}

// Save creates the Rental in the database.
func (_c *RentalCreate) Save(ctx context.Context) (*Rental, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *RentalCreate) SaveX(ctx context.Context) *Rental {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *RentalCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *RentalCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *RentalCreate) defaults() {
	if _, ok := _c.mutation.Date(); !ok {
		v := rental.DefaultDate()
		_c.mutation.SetDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *RentalCreate) check() error {
	if _, ok := _c.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Rental.date"`)}
	}
	if _, ok := _c.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Rental.user_id"`)}
	}
	if _, ok := _c.mutation.CarID(); !ok {
		return &ValidationError{Name: "car_id", err: errors.New(`ent: missing required field "Rental.car_id"`)}
	}
	if len(_c.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Rental.user"`)}
	}
	if len(_c.mutation.CarIDs()) == 0 {
		return &ValidationError{Name: "car", err: errors.New(`ent: missing required edge "Rental.car"`)}
	}
	return nil
}

func (_c *RentalCreate) sqlSave(ctx context.Context) (*Rental, error) {
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

func (_c *RentalCreate) createSpec() (*Rental, *sqlgraph.CreateSpec) {
	var (
		_node = &Rental{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(rental.Table, sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Date(); ok {
		_spec.SetField(rental.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if nodes := _c.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.UserTable,
			Columns: []string{rental.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.CarTable,
			Columns: []string{rental.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CarID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RentalCreateBulk is the builder for creating many Rental entities in bulk.
type RentalCreateBulk struct {
	config
	err      error
	builders []*RentalCreate
}

// Save creates the Rental entities in the database.
func (_c *RentalCreateBulk) Save(ctx context.Context) ([]*Rental, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Rental, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RentalMutation)
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
func (_c *RentalCreateBulk) SaveX(ctx context.Context) []*Rental {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *RentalCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *RentalCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
