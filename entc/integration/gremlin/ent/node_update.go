// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/node"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (_u *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetValue sets the "value" field.
func (_u *NodeUpdate) SetValue(v int) *NodeUpdate {
	_u.mutation.ResetValue()
	_u.mutation.SetValue(v)
	return _u
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (_u *NodeUpdate) SetNillableValue(v *int) *NodeUpdate {
	if v != nil {
		_u.SetValue(*v)
	}
	return _u
}

// AddValue adds value to the "value" field.
func (_u *NodeUpdate) AddValue(v int) *NodeUpdate {
	_u.mutation.AddValue(v)
	return _u
}

// ClearValue clears the value of the "value" field.
func (_u *NodeUpdate) ClearValue() *NodeUpdate {
	_u.mutation.ClearValue()
	return _u
}

// SetUpdatedAt sets the "updated_at" field.
func (_u *NodeUpdate) SetUpdatedAt(v time.Time) *NodeUpdate {
	_u.mutation.SetUpdatedAt(v)
	return _u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (_u *NodeUpdate) ClearUpdatedAt() *NodeUpdate {
	_u.mutation.ClearUpdatedAt()
	return _u
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (_u *NodeUpdate) SetPrevID(id string) *NodeUpdate {
	_u.mutation.SetPrevID(id)
	return _u
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (_u *NodeUpdate) SetNillablePrevID(id *string) *NodeUpdate {
	if id != nil {
		_u = _u.SetPrevID(*id)
	}
	return _u
}

// SetPrev sets the "prev" edge to the Node entity.
func (_u *NodeUpdate) SetPrev(v *Node) *NodeUpdate {
	return _u.SetPrevID(v.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (_u *NodeUpdate) SetNextID(id string) *NodeUpdate {
	_u.mutation.SetNextID(id)
	return _u
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (_u *NodeUpdate) SetNillableNextID(id *string) *NodeUpdate {
	if id != nil {
		_u = _u.SetNextID(*id)
	}
	return _u
}

// SetNext sets the "next" edge to the Node entity.
func (_u *NodeUpdate) SetNext(v *Node) *NodeUpdate {
	return _u.SetNextID(v.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (_u *NodeUpdate) Mutation() *NodeMutation {
	return _u.mutation
}

// ClearPrev clears the "prev" edge to the Node entity.
func (_u *NodeUpdate) ClearPrev() *NodeUpdate {
	_u.mutation.ClearPrev()
	return _u
}

// ClearNext clears the "next" edge to the Node entity.
func (_u *NodeUpdate) ClearNext() *NodeUpdate {
	_u.mutation.ClearNext()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *NodeUpdate) Save(ctx context.Context) (int, error) {
	_u.defaults()
	return withHooks(ctx, _u.gremlinSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *NodeUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *NodeUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_u *NodeUpdate) defaults() {
	if _, ok := _u.mutation.UpdatedAt(); !ok && !_u.mutation.UpdatedAtCleared() {
		v := node.UpdateDefaultUpdatedAt()
		_u.mutation.SetUpdatedAt(v)
	}
}

func (_u *NodeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _u.gremlin().Query()
	if err := _u.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	_u.mutation.done = true
	return res.ReadInt()
}

func (_u *NodeUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(node.Label)
	for _, p := range _u.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := _u.mutation.Value(); ok {
		v.Property(dsl.Single, node.FieldValue, value)
	}
	if value, ok := _u.mutation.AddedValue(); ok {
		v.Property(dsl.Single, node.FieldValue, __.Union(__.Values(node.FieldValue), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, node.FieldUpdatedAt, value)
	}
	var properties []any
	if _u.mutation.ValueCleared() {
		properties = append(properties, node.FieldValue)
	}
	if _u.mutation.UpdatedAtCleared() {
		properties = append(properties, node.FieldUpdatedAt)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if _u.mutation.PrevCleared() {
		tr := rv.Clone().InE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.PrevIDs() {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if _u.mutation.NextCleared() {
		tr := rv.Clone().OutE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.NextIDs() {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetValue sets the "value" field.
func (_u *NodeUpdateOne) SetValue(v int) *NodeUpdateOne {
	_u.mutation.ResetValue()
	_u.mutation.SetValue(v)
	return _u
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (_u *NodeUpdateOne) SetNillableValue(v *int) *NodeUpdateOne {
	if v != nil {
		_u.SetValue(*v)
	}
	return _u
}

// AddValue adds value to the "value" field.
func (_u *NodeUpdateOne) AddValue(v int) *NodeUpdateOne {
	_u.mutation.AddValue(v)
	return _u
}

// ClearValue clears the value of the "value" field.
func (_u *NodeUpdateOne) ClearValue() *NodeUpdateOne {
	_u.mutation.ClearValue()
	return _u
}

// SetUpdatedAt sets the "updated_at" field.
func (_u *NodeUpdateOne) SetUpdatedAt(v time.Time) *NodeUpdateOne {
	_u.mutation.SetUpdatedAt(v)
	return _u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (_u *NodeUpdateOne) ClearUpdatedAt() *NodeUpdateOne {
	_u.mutation.ClearUpdatedAt()
	return _u
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (_u *NodeUpdateOne) SetPrevID(id string) *NodeUpdateOne {
	_u.mutation.SetPrevID(id)
	return _u
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (_u *NodeUpdateOne) SetNillablePrevID(id *string) *NodeUpdateOne {
	if id != nil {
		_u = _u.SetPrevID(*id)
	}
	return _u
}

// SetPrev sets the "prev" edge to the Node entity.
func (_u *NodeUpdateOne) SetPrev(v *Node) *NodeUpdateOne {
	return _u.SetPrevID(v.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (_u *NodeUpdateOne) SetNextID(id string) *NodeUpdateOne {
	_u.mutation.SetNextID(id)
	return _u
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (_u *NodeUpdateOne) SetNillableNextID(id *string) *NodeUpdateOne {
	if id != nil {
		_u = _u.SetNextID(*id)
	}
	return _u
}

// SetNext sets the "next" edge to the Node entity.
func (_u *NodeUpdateOne) SetNext(v *Node) *NodeUpdateOne {
	return _u.SetNextID(v.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (_u *NodeUpdateOne) Mutation() *NodeMutation {
	return _u.mutation
}

// ClearPrev clears the "prev" edge to the Node entity.
func (_u *NodeUpdateOne) ClearPrev() *NodeUpdateOne {
	_u.mutation.ClearPrev()
	return _u
}

// ClearNext clears the "next" edge to the Node entity.
func (_u *NodeUpdateOne) ClearNext() *NodeUpdateOne {
	_u.mutation.ClearNext()
	return _u
}

// Where appends a list predicates to the NodeUpdate builder.
func (_u *NodeUpdateOne) Where(ps ...predicate.Node) *NodeUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Node entity.
func (_u *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	_u.defaults()
	return withHooks(ctx, _u.gremlinSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_u *NodeUpdateOne) defaults() {
	if _, ok := _u.mutation.UpdatedAt(); !ok && !_u.mutation.UpdatedAtCleared() {
		v := node.UpdateDefaultUpdatedAt()
		_u.mutation.SetUpdatedAt(v)
	}
}

func (_u *NodeUpdateOne) gremlinSave(ctx context.Context) (*Node, error) {
	res := &gremlin.Response{}
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	query, bindings := _u.gremlin(id).Query()
	if err := _u.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	_u.mutation.done = true
	_m := &Node{config: _u.config}
	if err := _m.FromResponse(res); err != nil {
		return nil, err
	}
	return _m, nil
}

func (_u *NodeUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := _u.mutation.Value(); ok {
		v.Property(dsl.Single, node.FieldValue, value)
	}
	if value, ok := _u.mutation.AddedValue(); ok {
		v.Property(dsl.Single, node.FieldValue, __.Union(__.Values(node.FieldValue), __.Constant(value)).Sum())
	}
	if value, ok := _u.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, node.FieldUpdatedAt, value)
	}
	var properties []any
	if _u.mutation.ValueCleared() {
		properties = append(properties, node.FieldValue)
	}
	if _u.mutation.UpdatedAtCleared() {
		properties = append(properties, node.FieldUpdatedAt)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if _u.mutation.PrevCleared() {
		tr := rv.Clone().InE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.PrevIDs() {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if _u.mutation.NextCleared() {
		tr := rv.Clone().OutE(node.NextLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range _u.mutation.NextIDs() {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if len(_u.fields) > 0 {
		fields := make([]any, 0, len(_u.fields)+1)
		fields = append(fields, true)
		for _, f := range _u.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
