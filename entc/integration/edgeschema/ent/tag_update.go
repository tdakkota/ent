// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/grouptag"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (_u *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetValue sets the "value" field.
func (_u *TagUpdate) SetValue(v string) *TagUpdate {
	_u.mutation.SetValue(v)
	return _u
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (_u *TagUpdate) SetNillableValue(v *string) *TagUpdate {
	if v != nil {
		_u.SetValue(*v)
	}
	return _u
}

// AddTweetIDs adds the "tweets" edge to the Tweet entity by IDs.
func (_u *TagUpdate) AddTweetIDs(ids ...int) *TagUpdate {
	_u.mutation.AddTweetIDs(ids...)
	return _u
}

// AddTweets adds the "tweets" edges to the Tweet entity.
func (_u *TagUpdate) AddTweets(v ...*Tweet) *TagUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTweetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_u *TagUpdate) AddGroupIDs(ids ...int) *TagUpdate {
	_u.mutation.AddGroupIDs(ids...)
	return _u
}

// AddGroups adds the "groups" edges to the Group entity.
func (_u *TagUpdate) AddGroups(v ...*Group) *TagUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddGroupIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (_u *TagUpdate) AddTweetTagIDs(ids ...uuid.UUID) *TagUpdate {
	_u.mutation.AddTweetTagIDs(ids...)
	return _u
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (_u *TagUpdate) AddTweetTags(v ...*TweetTag) *TagUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTweetTagIDs(ids...)
}

// AddGroupTagIDs adds the "group_tags" edge to the GroupTag entity by IDs.
func (_u *TagUpdate) AddGroupTagIDs(ids ...int) *TagUpdate {
	_u.mutation.AddGroupTagIDs(ids...)
	return _u
}

// AddGroupTags adds the "group_tags" edges to the GroupTag entity.
func (_u *TagUpdate) AddGroupTags(v ...*GroupTag) *TagUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddGroupTagIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (_u *TagUpdate) Mutation() *TagMutation {
	return _u.mutation
}

// ClearTweets clears all "tweets" edges to the Tweet entity.
func (_u *TagUpdate) ClearTweets() *TagUpdate {
	_u.mutation.ClearTweets()
	return _u
}

// RemoveTweetIDs removes the "tweets" edge to Tweet entities by IDs.
func (_u *TagUpdate) RemoveTweetIDs(ids ...int) *TagUpdate {
	_u.mutation.RemoveTweetIDs(ids...)
	return _u
}

// RemoveTweets removes "tweets" edges to Tweet entities.
func (_u *TagUpdate) RemoveTweets(v ...*Tweet) *TagUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTweetIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (_u *TagUpdate) ClearGroups() *TagUpdate {
	_u.mutation.ClearGroups()
	return _u
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (_u *TagUpdate) RemoveGroupIDs(ids ...int) *TagUpdate {
	_u.mutation.RemoveGroupIDs(ids...)
	return _u
}

// RemoveGroups removes "groups" edges to Group entities.
func (_u *TagUpdate) RemoveGroups(v ...*Group) *TagUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveGroupIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (_u *TagUpdate) ClearTweetTags() *TagUpdate {
	_u.mutation.ClearTweetTags()
	return _u
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (_u *TagUpdate) RemoveTweetTagIDs(ids ...uuid.UUID) *TagUpdate {
	_u.mutation.RemoveTweetTagIDs(ids...)
	return _u
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (_u *TagUpdate) RemoveTweetTags(v ...*TweetTag) *TagUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTweetTagIDs(ids...)
}

// ClearGroupTags clears all "group_tags" edges to the GroupTag entity.
func (_u *TagUpdate) ClearGroupTags() *TagUpdate {
	_u.mutation.ClearGroupTags()
	return _u
}

// RemoveGroupTagIDs removes the "group_tags" edge to GroupTag entities by IDs.
func (_u *TagUpdate) RemoveGroupTagIDs(ids ...int) *TagUpdate {
	_u.mutation.RemoveGroupTagIDs(ids...)
	return _u
}

// RemoveGroupTags removes "group_tags" edges to GroupTag entities.
func (_u *TagUpdate) RemoveGroupTags(v ...*GroupTag) *TagUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveGroupTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *TagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *TagUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *TagUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *TagUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Value(); ok {
		_spec.SetField(tag.FieldValue, field.TypeString, value)
	}
	if _u.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		createE := &TweetTagCreate{config: _u.config, mutation: newTweetTagMutation(_u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTweetsIDs(); len(nodes) > 0 && !_u.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: _u.config, mutation: newTweetTagMutation(_u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TweetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: _u.config, mutation: newTweetTagMutation(_u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !_u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !_u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupTagsIDs(); len(nodes) > 0 && !_u.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetValue sets the "value" field.
func (_u *TagUpdateOne) SetValue(v string) *TagUpdateOne {
	_u.mutation.SetValue(v)
	return _u
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (_u *TagUpdateOne) SetNillableValue(v *string) *TagUpdateOne {
	if v != nil {
		_u.SetValue(*v)
	}
	return _u
}

// AddTweetIDs adds the "tweets" edge to the Tweet entity by IDs.
func (_u *TagUpdateOne) AddTweetIDs(ids ...int) *TagUpdateOne {
	_u.mutation.AddTweetIDs(ids...)
	return _u
}

// AddTweets adds the "tweets" edges to the Tweet entity.
func (_u *TagUpdateOne) AddTweets(v ...*Tweet) *TagUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTweetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_u *TagUpdateOne) AddGroupIDs(ids ...int) *TagUpdateOne {
	_u.mutation.AddGroupIDs(ids...)
	return _u
}

// AddGroups adds the "groups" edges to the Group entity.
func (_u *TagUpdateOne) AddGroups(v ...*Group) *TagUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddGroupIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (_u *TagUpdateOne) AddTweetTagIDs(ids ...uuid.UUID) *TagUpdateOne {
	_u.mutation.AddTweetTagIDs(ids...)
	return _u
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (_u *TagUpdateOne) AddTweetTags(v ...*TweetTag) *TagUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddTweetTagIDs(ids...)
}

// AddGroupTagIDs adds the "group_tags" edge to the GroupTag entity by IDs.
func (_u *TagUpdateOne) AddGroupTagIDs(ids ...int) *TagUpdateOne {
	_u.mutation.AddGroupTagIDs(ids...)
	return _u
}

// AddGroupTags adds the "group_tags" edges to the GroupTag entity.
func (_u *TagUpdateOne) AddGroupTags(v ...*GroupTag) *TagUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddGroupTagIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (_u *TagUpdateOne) Mutation() *TagMutation {
	return _u.mutation
}

// ClearTweets clears all "tweets" edges to the Tweet entity.
func (_u *TagUpdateOne) ClearTweets() *TagUpdateOne {
	_u.mutation.ClearTweets()
	return _u
}

// RemoveTweetIDs removes the "tweets" edge to Tweet entities by IDs.
func (_u *TagUpdateOne) RemoveTweetIDs(ids ...int) *TagUpdateOne {
	_u.mutation.RemoveTweetIDs(ids...)
	return _u
}

// RemoveTweets removes "tweets" edges to Tweet entities.
func (_u *TagUpdateOne) RemoveTweets(v ...*Tweet) *TagUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTweetIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (_u *TagUpdateOne) ClearGroups() *TagUpdateOne {
	_u.mutation.ClearGroups()
	return _u
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (_u *TagUpdateOne) RemoveGroupIDs(ids ...int) *TagUpdateOne {
	_u.mutation.RemoveGroupIDs(ids...)
	return _u
}

// RemoveGroups removes "groups" edges to Group entities.
func (_u *TagUpdateOne) RemoveGroups(v ...*Group) *TagUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveGroupIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (_u *TagUpdateOne) ClearTweetTags() *TagUpdateOne {
	_u.mutation.ClearTweetTags()
	return _u
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (_u *TagUpdateOne) RemoveTweetTagIDs(ids ...uuid.UUID) *TagUpdateOne {
	_u.mutation.RemoveTweetTagIDs(ids...)
	return _u
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (_u *TagUpdateOne) RemoveTweetTags(v ...*TweetTag) *TagUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveTweetTagIDs(ids...)
}

// ClearGroupTags clears all "group_tags" edges to the GroupTag entity.
func (_u *TagUpdateOne) ClearGroupTags() *TagUpdateOne {
	_u.mutation.ClearGroupTags()
	return _u
}

// RemoveGroupTagIDs removes the "group_tags" edge to GroupTag entities by IDs.
func (_u *TagUpdateOne) RemoveGroupTagIDs(ids ...int) *TagUpdateOne {
	_u.mutation.RemoveGroupTagIDs(ids...)
	return _u
}

// RemoveGroupTags removes "group_tags" edges to GroupTag entities.
func (_u *TagUpdateOne) RemoveGroupTags(v ...*GroupTag) *TagUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveGroupTagIDs(ids...)
}

// Where appends a list predicates to the TagUpdate builder.
func (_u *TagUpdateOne) Where(ps ...predicate.Tag) *TagUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Tag entity.
func (_u *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *TagUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Value(); ok {
		_spec.SetField(tag.FieldValue, field.TypeString, value)
	}
	if _u.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		createE := &TweetTagCreate{config: _u.config, mutation: newTweetTagMutation(_u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTweetsIDs(); len(nodes) > 0 && !_u.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: _u.config, mutation: newTweetTagMutation(_u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TweetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: _u.config, mutation: newTweetTagMutation(_u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !_u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !_u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupTagsIDs(); len(nodes) > 0 && !_u.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tag{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
