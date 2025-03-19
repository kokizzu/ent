// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/start/ent/car"
	"entgo.io/ent/examples/start/ent/predicate"
	"entgo.io/ent/schema/field"
)

// CarDelete is the builder for deleting a Car entity.
type CarDelete struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarDelete builder.
func (_d *CarDelete) Where(ps ...predicate.Car) *CarDelete {
	_d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (_d *CarDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, _d.sqlExec, _d.mutation, _d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *CarDelete) ExecX(ctx context.Context) int {
	n, err := _d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (_d *CarDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(car.Table, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	if ps := _d.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, _d.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	_d.mutation.done = true
	return affected, err
}

// CarDeleteOne is the builder for deleting a single Car entity.
type CarDeleteOne struct {
	_d *CarDelete
}

// Where appends a list predicates to the CarDelete builder.
func (_d *CarDeleteOne) Where(ps ...predicate.Car) *CarDeleteOne {
	_d._d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query.
func (_d *CarDeleteOne) Exec(ctx context.Context) error {
	n, err := _d._d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{car.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *CarDeleteOne) ExecX(ctx context.Context) {
	if err := _d.Exec(ctx); err != nil {
		panic(err)
	}
}
