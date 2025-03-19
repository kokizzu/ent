// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/item"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetText sets the "text" field.
func (_c *ItemCreate) SetText(s string) *ItemCreate {
	_c.mutation.SetText(s)
	return _c
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_c *ItemCreate) SetNillableText(s *string) *ItemCreate {
	if s != nil {
		_c.SetText(*s)
	}
	return _c
}

// SetID sets the "id" field.
func (_c *ItemCreate) SetID(s string) *ItemCreate {
	_c.mutation.SetID(s)
	return _c
}

// SetNillableID sets the "id" field if the given value is not nil.
func (_c *ItemCreate) SetNillableID(s *string) *ItemCreate {
	if s != nil {
		_c.SetID(*s)
	}
	return _c
}

// Mutation returns the ItemMutation object of the builder.
func (_c *ItemCreate) Mutation() *ItemMutation {
	return _c.mutation
}

// Save creates the Item in the database.
func (_c *ItemCreate) Save(ctx context.Context) (*Item, error) {
	_c.defaults()
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *ItemCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *ItemCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *ItemCreate) defaults() {
	if _, ok := _c.mutation.ID(); !ok {
		v := item.DefaultID()
		_c.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *ItemCreate) check() error {
	if v, ok := _c.mutation.Text(); ok {
		if err := item.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Item.text": %w`, err)}
		}
	}
	if v, ok := _c.mutation.ID(); ok {
		if err := item.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Item.id": %w`, err)}
		}
	}
	return nil
}

func (_c *ItemCreate) gremlinSave(ctx context.Context) (*Item, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := _c.gremlin().Query()
	if err := _c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Item{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *ItemCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(item.Label)
	if id, ok := _c.mutation.ID(); ok {
		v.Property(dsl.ID, id)
	}
	if value, ok := _c.mutation.Text(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(item.Label, item.FieldText, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(item.Label, item.FieldText, value)),
		})
		v.Property(dsl.Single, item.FieldText, value)
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	err      error
	builders []*ItemCreate
}
