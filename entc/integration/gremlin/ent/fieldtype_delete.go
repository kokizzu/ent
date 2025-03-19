// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/fieldtype"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// FieldTypeDelete is the builder for deleting a FieldType entity.
type FieldTypeDelete struct {
	config
	hooks    []Hook
	mutation *FieldTypeMutation
}

// Where appends a list predicates to the FieldTypeDelete builder.
func (_d *FieldTypeDelete) Where(ps ...predicate.FieldType) *FieldTypeDelete {
	_d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (_d *FieldTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, _d.gremlinExec, _d.mutation, _d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *FieldTypeDelete) ExecX(ctx context.Context) int {
	n, err := _d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (_d *FieldTypeDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _d.gremlin().Query()
	if err := _d.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	_d.mutation.done = true
	return res.ReadInt()
}

func (_d *FieldTypeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(fieldtype.Label)
	for _, p := range _d.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// FieldTypeDeleteOne is the builder for deleting a single FieldType entity.
type FieldTypeDeleteOne struct {
	_d *FieldTypeDelete
}

// Where appends a list predicates to the FieldTypeDelete builder.
func (_d *FieldTypeDeleteOne) Where(ps ...predicate.FieldType) *FieldTypeDeleteOne {
	_d._d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query.
func (_d *FieldTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := _d._d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fieldtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *FieldTypeDeleteOne) ExecX(ctx context.Context) {
	if err := _d.Exec(ctx); err != nil {
		panic(err)
	}
}
