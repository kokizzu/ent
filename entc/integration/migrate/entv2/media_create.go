// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/media"
	"entgo.io/ent/schema/field"
)

// MediaCreate is the builder for creating a Media entity.
type MediaCreate struct {
	config
	mutation *MediaMutation
	hooks    []Hook
}

// SetSource sets the "source" field.
func (_c *MediaCreate) SetSource(s string) *MediaCreate {
	_c.mutation.SetSource(s)
	return _c
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (_c *MediaCreate) SetNillableSource(s *string) *MediaCreate {
	if s != nil {
		_c.SetSource(*s)
	}
	return _c
}

// SetSourceURI sets the "source_uri" field.
func (_c *MediaCreate) SetSourceURI(s string) *MediaCreate {
	_c.mutation.SetSourceURI(s)
	return _c
}

// SetNillableSourceURI sets the "source_uri" field if the given value is not nil.
func (_c *MediaCreate) SetNillableSourceURI(s *string) *MediaCreate {
	if s != nil {
		_c.SetSourceURI(*s)
	}
	return _c
}

// SetText sets the "text" field.
func (_c *MediaCreate) SetText(s string) *MediaCreate {
	_c.mutation.SetText(s)
	return _c
}

// SetNillableText sets the "text" field if the given value is not nil.
func (_c *MediaCreate) SetNillableText(s *string) *MediaCreate {
	if s != nil {
		_c.SetText(*s)
	}
	return _c
}

// Mutation returns the MediaMutation object of the builder.
func (_c *MediaCreate) Mutation() *MediaMutation {
	return _c.mutation
}

// Save creates the Media in the database.
func (_c *MediaCreate) Save(ctx context.Context) (*Media, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *MediaCreate) SaveX(ctx context.Context) *Media {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *MediaCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *MediaCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *MediaCreate) check() error {
	return nil
}

func (_c *MediaCreate) sqlSave(ctx context.Context) (*Media, error) {
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

func (_c *MediaCreate) createSpec() (*Media, *sqlgraph.CreateSpec) {
	var (
		_node = &Media{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(media.Table, sqlgraph.NewFieldSpec(media.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Source(); ok {
		_spec.SetField(media.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if value, ok := _c.mutation.SourceURI(); ok {
		_spec.SetField(media.FieldSourceURI, field.TypeString, value)
		_node.SourceURI = value
	}
	if value, ok := _c.mutation.Text(); ok {
		_spec.SetField(media.FieldText, field.TypeString, value)
		_node.Text = value
	}
	return _node, _spec
}

// MediaCreateBulk is the builder for creating many Media entities in bulk.
type MediaCreateBulk struct {
	config
	err      error
	builders []*MediaCreate
}

// Save creates the Media entities in the database.
func (_c *MediaCreateBulk) Save(ctx context.Context) ([]*Media, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Media, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MediaMutation)
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
func (_c *MediaCreateBulk) SaveX(ctx context.Context) []*Media {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *MediaCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *MediaCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
