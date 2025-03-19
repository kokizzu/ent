// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/metadata"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// MetadataCreate is the builder for creating a Metadata entity.
type MetadataCreate struct {
	config
	mutation *MetadataMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (_c *MetadataCreate) SetAge(i int) *MetadataCreate {
	_c.mutation.SetAge(i)
	return _c
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_c *MetadataCreate) SetNillableAge(i *int) *MetadataCreate {
	if i != nil {
		_c.SetAge(*i)
	}
	return _c
}

// SetParentID sets the "parent_id" field.
func (_c *MetadataCreate) SetParentID(i int) *MetadataCreate {
	_c.mutation.SetParentID(i)
	return _c
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (_c *MetadataCreate) SetNillableParentID(i *int) *MetadataCreate {
	if i != nil {
		_c.SetParentID(*i)
	}
	return _c
}

// SetID sets the "id" field.
func (_c *MetadataCreate) SetID(i int) *MetadataCreate {
	_c.mutation.SetID(i)
	return _c
}

// SetUserID sets the "user" edge to the User entity by ID.
func (_c *MetadataCreate) SetUserID(id int) *MetadataCreate {
	_c.mutation.SetUserID(id)
	return _c
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (_c *MetadataCreate) SetNillableUserID(id *int) *MetadataCreate {
	if id != nil {
		_c = _c.SetUserID(*id)
	}
	return _c
}

// SetUser sets the "user" edge to the User entity.
func (_c *MetadataCreate) SetUser(u *User) *MetadataCreate {
	return _c.SetUserID(u.ID)
}

// AddChildIDs adds the "children" edge to the Metadata entity by IDs.
func (_c *MetadataCreate) AddChildIDs(ids ...int) *MetadataCreate {
	_c.mutation.AddChildIDs(ids...)
	return _c
}

// AddChildren adds the "children" edges to the Metadata entity.
func (_c *MetadataCreate) AddChildren(m ...*Metadata) *MetadataCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return _c.AddChildIDs(ids...)
}

// SetParent sets the "parent" edge to the Metadata entity.
func (_c *MetadataCreate) SetParent(m *Metadata) *MetadataCreate {
	return _c.SetParentID(m.ID)
}

// Mutation returns the MetadataMutation object of the builder.
func (_c *MetadataCreate) Mutation() *MetadataMutation {
	return _c.mutation
}

// Save creates the Metadata in the database.
func (_c *MetadataCreate) Save(ctx context.Context) (*Metadata, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *MetadataCreate) SaveX(ctx context.Context) *Metadata {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *MetadataCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *MetadataCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *MetadataCreate) defaults() {
	if _, ok := _c.mutation.Age(); !ok {
		v := metadata.DefaultAge
		_c.mutation.SetAge(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *MetadataCreate) check() error {
	if _, ok := _c.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Metadata.age"`)}
	}
	return nil
}

func (_c *MetadataCreate) sqlSave(ctx context.Context) (*Metadata, error) {
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *MetadataCreate) createSpec() (*Metadata, *sqlgraph.CreateSpec) {
	var (
		_node = &Metadata{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(metadata.Table, sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt))
	)
	if id, ok := _c.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := _c.mutation.Age(); ok {
		_spec.SetField(metadata.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if nodes := _c.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   metadata.UserTable,
			Columns: []string{metadata.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   metadata.ChildrenTable,
			Columns: []string{metadata.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   metadata.ParentTable,
			Columns: []string{metadata.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metadata.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MetadataCreateBulk is the builder for creating many Metadata entities in bulk.
type MetadataCreateBulk struct {
	config
	err      error
	builders []*MetadataCreate
}

// Save creates the Metadata entities in the database.
func (_c *MetadataCreateBulk) Save(ctx context.Context) ([]*Metadata, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Metadata, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MetadataMutation)
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
func (_c *MetadataCreateBulk) SaveX(ctx context.Context) []*Metadata {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *MetadataCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *MetadataCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
