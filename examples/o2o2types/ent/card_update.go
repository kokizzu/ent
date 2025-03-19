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
	"entgo.io/ent/examples/o2o2types/ent/card"
	"entgo.io/ent/examples/o2o2types/ent/predicate"
	"entgo.io/ent/examples/o2o2types/ent/user"
	"entgo.io/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (_u *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetExpired sets the "expired" field.
func (_u *CardUpdate) SetExpired(t time.Time) *CardUpdate {
	_u.mutation.SetExpired(t)
	return _u
}

// SetNillableExpired sets the "expired" field if the given value is not nil.
func (_u *CardUpdate) SetNillableExpired(t *time.Time) *CardUpdate {
	if t != nil {
		_u.SetExpired(*t)
	}
	return _u
}

// SetNumber sets the "number" field.
func (_u *CardUpdate) SetNumber(s string) *CardUpdate {
	_u.mutation.SetNumber(s)
	return _u
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (_u *CardUpdate) SetNillableNumber(s *string) *CardUpdate {
	if s != nil {
		_u.SetNumber(*s)
	}
	return _u
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_u *CardUpdate) SetOwnerID(id int) *CardUpdate {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *CardUpdate) SetOwner(u *User) *CardUpdate {
	return _u.SetOwnerID(u.ID)
}

// Mutation returns the CardMutation object of the builder.
func (_u *CardUpdate) Mutation() *CardMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *CardUpdate) ClearOwner() *CardUpdate {
	_u.mutation.ClearOwner()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *CardUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CardUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *CardUpdate) check() error {
	if _u.mutation.OwnerCleared() && len(_u.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Card.owner"`)
	}
	return nil
}

func (_u *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := _u.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Expired(); ok {
		_spec.SetField(card.FieldExpired, field.TypeTime, value)
	}
	if value, ok := _u.mutation.Number(); ok {
		_spec.SetField(card.FieldNumber, field.TypeString, value)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetExpired sets the "expired" field.
func (_u *CardUpdateOne) SetExpired(t time.Time) *CardUpdateOne {
	_u.mutation.SetExpired(t)
	return _u
}

// SetNillableExpired sets the "expired" field if the given value is not nil.
func (_u *CardUpdateOne) SetNillableExpired(t *time.Time) *CardUpdateOne {
	if t != nil {
		_u.SetExpired(*t)
	}
	return _u
}

// SetNumber sets the "number" field.
func (_u *CardUpdateOne) SetNumber(s string) *CardUpdateOne {
	_u.mutation.SetNumber(s)
	return _u
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (_u *CardUpdateOne) SetNillableNumber(s *string) *CardUpdateOne {
	if s != nil {
		_u.SetNumber(*s)
	}
	return _u
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_u *CardUpdateOne) SetOwnerID(id int) *CardUpdateOne {
	_u.mutation.SetOwnerID(id)
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return _u.SetOwnerID(u.ID)
}

// Mutation returns the CardMutation object of the builder.
func (_u *CardUpdateOne) Mutation() *CardMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *CardUpdateOne) ClearOwner() *CardUpdateOne {
	_u.mutation.ClearOwner()
	return _u
}

// Where appends a list predicates to the CardUpdate builder.
func (_u *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Card entity.
func (_u *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CardUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *CardUpdateOne) check() error {
	if _u.mutation.OwnerCleared() && len(_u.mutation.OwnerIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Card.owner"`)
	}
	return nil
}

func (_u *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
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
	if value, ok := _u.mutation.Expired(); ok {
		_spec.SetField(card.FieldExpired, field.TypeTime, value)
	}
	if value, ok := _u.mutation.Number(); ok {
		_spec.SetField(card.FieldNumber, field.TypeString, value)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Card{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
