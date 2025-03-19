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
	"entgo.io/ent/entc/integration/edgeschema/ent/friendship"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/schema/field"
)

// FriendshipUpdate is the builder for updating Friendship entities.
type FriendshipUpdate struct {
	config
	hooks    []Hook
	mutation *FriendshipMutation
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (_u *FriendshipUpdate) Where(ps ...predicate.Friendship) *FriendshipUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetWeight sets the "weight" field.
func (_u *FriendshipUpdate) SetWeight(i int) *FriendshipUpdate {
	_u.mutation.ResetWeight()
	_u.mutation.SetWeight(i)
	return _u
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (_u *FriendshipUpdate) SetNillableWeight(i *int) *FriendshipUpdate {
	if i != nil {
		_u.SetWeight(*i)
	}
	return _u
}

// AddWeight adds i to the "weight" field.
func (_u *FriendshipUpdate) AddWeight(i int) *FriendshipUpdate {
	_u.mutation.AddWeight(i)
	return _u
}

// SetCreatedAt sets the "created_at" field.
func (_u *FriendshipUpdate) SetCreatedAt(t time.Time) *FriendshipUpdate {
	_u.mutation.SetCreatedAt(t)
	return _u
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_u *FriendshipUpdate) SetNillableCreatedAt(t *time.Time) *FriendshipUpdate {
	if t != nil {
		_u.SetCreatedAt(*t)
	}
	return _u
}

// Mutation returns the FriendshipMutation object of the builder.
func (_u *FriendshipUpdate) Mutation() *FriendshipMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *FriendshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *FriendshipUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *FriendshipUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *FriendshipUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *FriendshipUpdate) check() error {
	if _u.mutation.UserCleared() && len(_u.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Friendship.user"`)
	}
	if _u.mutation.FriendCleared() && len(_u.mutation.FriendIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

func (_u *FriendshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := _u.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Weight(); ok {
		_spec.SetField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := _u.mutation.AddedWeight(); ok {
		_spec.AddField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := _u.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return n, nil
}

// FriendshipUpdateOne is the builder for updating a single Friendship entity.
type FriendshipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FriendshipMutation
}

// SetWeight sets the "weight" field.
func (_u *FriendshipUpdateOne) SetWeight(i int) *FriendshipUpdateOne {
	_u.mutation.ResetWeight()
	_u.mutation.SetWeight(i)
	return _u
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (_u *FriendshipUpdateOne) SetNillableWeight(i *int) *FriendshipUpdateOne {
	if i != nil {
		_u.SetWeight(*i)
	}
	return _u
}

// AddWeight adds i to the "weight" field.
func (_u *FriendshipUpdateOne) AddWeight(i int) *FriendshipUpdateOne {
	_u.mutation.AddWeight(i)
	return _u
}

// SetCreatedAt sets the "created_at" field.
func (_u *FriendshipUpdateOne) SetCreatedAt(t time.Time) *FriendshipUpdateOne {
	_u.mutation.SetCreatedAt(t)
	return _u
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_u *FriendshipUpdateOne) SetNillableCreatedAt(t *time.Time) *FriendshipUpdateOne {
	if t != nil {
		_u.SetCreatedAt(*t)
	}
	return _u
}

// Mutation returns the FriendshipMutation object of the builder.
func (_u *FriendshipUpdateOne) Mutation() *FriendshipMutation {
	return _u.mutation
}

// Where appends a list predicates to the FriendshipUpdate builder.
func (_u *FriendshipUpdateOne) Where(ps ...predicate.Friendship) *FriendshipUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *FriendshipUpdateOne) Select(field string, fields ...string) *FriendshipUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Friendship entity.
func (_u *FriendshipUpdateOne) Save(ctx context.Context) (*Friendship, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *FriendshipUpdateOne) SaveX(ctx context.Context) *Friendship {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *FriendshipUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *FriendshipUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *FriendshipUpdateOne) check() error {
	if _u.mutation.UserCleared() && len(_u.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Friendship.user"`)
	}
	if _u.mutation.FriendCleared() && len(_u.mutation.FriendIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Friendship.friend"`)
	}
	return nil
}

func (_u *FriendshipUpdateOne) sqlSave(ctx context.Context) (_node *Friendship, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(friendship.Table, friendship.Columns, sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Friendship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for _, f := range fields {
			if !friendship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != friendship.FieldID {
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
	if value, ok := _u.mutation.Weight(); ok {
		_spec.SetField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := _u.mutation.AddedWeight(); ok {
		_spec.AddField(friendship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := _u.mutation.CreatedAt(); ok {
		_spec.SetField(friendship.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Friendship{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{friendship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
