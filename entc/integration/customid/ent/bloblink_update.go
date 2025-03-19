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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/entc/integration/customid/ent/bloblink"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlobLinkUpdate is the builder for updating BlobLink entities.
type BlobLinkUpdate struct {
	config
	hooks    []Hook
	mutation *BlobLinkMutation
}

// Where appends a list predicates to the BlobLinkUpdate builder.
func (_u *BlobLinkUpdate) Where(ps ...predicate.BlobLink) *BlobLinkUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetCreatedAt sets the "created_at" field.
func (_u *BlobLinkUpdate) SetCreatedAt(v time.Time) *BlobLinkUpdate {
	_u.mutation.SetCreatedAt(v)
	return _u
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_u *BlobLinkUpdate) SetNillableCreatedAt(v *time.Time) *BlobLinkUpdate {
	if v != nil {
		_u.SetCreatedAt(*v)
	}
	return _u
}

// SetBlobID sets the "blob_id" field.
func (_u *BlobLinkUpdate) SetBlobID(v uuid.UUID) *BlobLinkUpdate {
	_u.mutation.SetBlobID(v)
	return _u
}

// SetNillableBlobID sets the "blob_id" field if the given value is not nil.
func (_u *BlobLinkUpdate) SetNillableBlobID(v *uuid.UUID) *BlobLinkUpdate {
	if v != nil {
		_u.SetBlobID(*v)
	}
	return _u
}

// SetLinkID sets the "link_id" field.
func (_u *BlobLinkUpdate) SetLinkID(v uuid.UUID) *BlobLinkUpdate {
	_u.mutation.SetLinkID(v)
	return _u
}

// SetNillableLinkID sets the "link_id" field if the given value is not nil.
func (_u *BlobLinkUpdate) SetNillableLinkID(v *uuid.UUID) *BlobLinkUpdate {
	if v != nil {
		_u.SetLinkID(*v)
	}
	return _u
}

// SetBlob sets the "blob" edge to the Blob entity.
func (_u *BlobLinkUpdate) SetBlob(v *Blob) *BlobLinkUpdate {
	return _u.SetBlobID(v.ID)
}

// SetLink sets the "link" edge to the Blob entity.
func (_u *BlobLinkUpdate) SetLink(v *Blob) *BlobLinkUpdate {
	return _u.SetLinkID(v.ID)
}

// Mutation returns the BlobLinkMutation object of the builder.
func (_u *BlobLinkUpdate) Mutation() *BlobLinkMutation {
	return _u.mutation
}

// ClearBlob clears the "blob" edge to the Blob entity.
func (_u *BlobLinkUpdate) ClearBlob() *BlobLinkUpdate {
	_u.mutation.ClearBlob()
	return _u
}

// ClearLink clears the "link" edge to the Blob entity.
func (_u *BlobLinkUpdate) ClearLink() *BlobLinkUpdate {
	_u.mutation.ClearLink()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *BlobLinkUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *BlobLinkUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *BlobLinkUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *BlobLinkUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *BlobLinkUpdate) check() error {
	if _u.mutation.BlobCleared() && len(_u.mutation.BlobIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "BlobLink.blob"`)
	}
	if _u.mutation.LinkCleared() && len(_u.mutation.LinkIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "BlobLink.link"`)
	}
	return nil
}

func (_u *BlobLinkUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(bloblink.Table, bloblink.Columns, sqlgraph.NewFieldSpec(bloblink.FieldBlobID, field.TypeUUID), sqlgraph.NewFieldSpec(bloblink.FieldLinkID, field.TypeUUID))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.CreatedAt(); ok {
		_spec.SetField(bloblink.FieldCreatedAt, field.TypeTime, value)
	}
	if _u.mutation.BlobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.BlobTable,
			Columns: []string{bloblink.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.BlobTable,
			Columns: []string{bloblink.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.LinkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.LinkTable,
			Columns: []string{bloblink.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.LinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.LinkTable,
			Columns: []string{bloblink.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bloblink.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// BlobLinkUpdateOne is the builder for updating a single BlobLink entity.
type BlobLinkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlobLinkMutation
}

// SetCreatedAt sets the "created_at" field.
func (_u *BlobLinkUpdateOne) SetCreatedAt(v time.Time) *BlobLinkUpdateOne {
	_u.mutation.SetCreatedAt(v)
	return _u
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_u *BlobLinkUpdateOne) SetNillableCreatedAt(v *time.Time) *BlobLinkUpdateOne {
	if v != nil {
		_u.SetCreatedAt(*v)
	}
	return _u
}

// SetBlobID sets the "blob_id" field.
func (_u *BlobLinkUpdateOne) SetBlobID(v uuid.UUID) *BlobLinkUpdateOne {
	_u.mutation.SetBlobID(v)
	return _u
}

// SetNillableBlobID sets the "blob_id" field if the given value is not nil.
func (_u *BlobLinkUpdateOne) SetNillableBlobID(v *uuid.UUID) *BlobLinkUpdateOne {
	if v != nil {
		_u.SetBlobID(*v)
	}
	return _u
}

// SetLinkID sets the "link_id" field.
func (_u *BlobLinkUpdateOne) SetLinkID(v uuid.UUID) *BlobLinkUpdateOne {
	_u.mutation.SetLinkID(v)
	return _u
}

// SetNillableLinkID sets the "link_id" field if the given value is not nil.
func (_u *BlobLinkUpdateOne) SetNillableLinkID(v *uuid.UUID) *BlobLinkUpdateOne {
	if v != nil {
		_u.SetLinkID(*v)
	}
	return _u
}

// SetBlob sets the "blob" edge to the Blob entity.
func (_u *BlobLinkUpdateOne) SetBlob(v *Blob) *BlobLinkUpdateOne {
	return _u.SetBlobID(v.ID)
}

// SetLink sets the "link" edge to the Blob entity.
func (_u *BlobLinkUpdateOne) SetLink(v *Blob) *BlobLinkUpdateOne {
	return _u.SetLinkID(v.ID)
}

// Mutation returns the BlobLinkMutation object of the builder.
func (_u *BlobLinkUpdateOne) Mutation() *BlobLinkMutation {
	return _u.mutation
}

// ClearBlob clears the "blob" edge to the Blob entity.
func (_u *BlobLinkUpdateOne) ClearBlob() *BlobLinkUpdateOne {
	_u.mutation.ClearBlob()
	return _u
}

// ClearLink clears the "link" edge to the Blob entity.
func (_u *BlobLinkUpdateOne) ClearLink() *BlobLinkUpdateOne {
	_u.mutation.ClearLink()
	return _u
}

// Where appends a list predicates to the BlobLinkUpdate builder.
func (_u *BlobLinkUpdateOne) Where(ps ...predicate.BlobLink) *BlobLinkUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *BlobLinkUpdateOne) Select(field string, fields ...string) *BlobLinkUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated BlobLink entity.
func (_u *BlobLinkUpdateOne) Save(ctx context.Context) (*BlobLink, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *BlobLinkUpdateOne) SaveX(ctx context.Context) *BlobLink {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *BlobLinkUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *BlobLinkUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *BlobLinkUpdateOne) check() error {
	if _u.mutation.BlobCleared() && len(_u.mutation.BlobIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "BlobLink.blob"`)
	}
	if _u.mutation.LinkCleared() && len(_u.mutation.LinkIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "BlobLink.link"`)
	}
	return nil
}

func (_u *BlobLinkUpdateOne) sqlSave(ctx context.Context) (_node *BlobLink, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(bloblink.Table, bloblink.Columns, sqlgraph.NewFieldSpec(bloblink.FieldBlobID, field.TypeUUID), sqlgraph.NewFieldSpec(bloblink.FieldLinkID, field.TypeUUID))
	if id, ok := _u.mutation.BlobID(); !ok {
		return nil, &ValidationError{Name: "blob_id", err: errors.New(`ent: missing "BlobLink.blob_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := _u.mutation.LinkID(); !ok {
		return nil, &ValidationError{Name: "link_id", err: errors.New(`ent: missing "BlobLink.link_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !bloblink.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.CreatedAt(); ok {
		_spec.SetField(bloblink.FieldCreatedAt, field.TypeTime, value)
	}
	if _u.mutation.BlobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.BlobTable,
			Columns: []string{bloblink.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.BlobTable,
			Columns: []string{bloblink.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.LinkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.LinkTable,
			Columns: []string{bloblink.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.LinkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   bloblink.LinkTable,
			Columns: []string{bloblink.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &BlobLink{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bloblink.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
