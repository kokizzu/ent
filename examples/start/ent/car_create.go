// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/start/ent/car"
	"entgo.io/ent/examples/start/ent/user"
	"entgo.io/ent/schema/field"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
}

// SetModel sets the "model" field.
func (cc *CarCreate) SetModel(s string) *CarCreate {
	cc.mutation.SetModel(s)
	return cc
}

// SetRegisteredAt sets the "registered_at" field.
func (cc *CarCreate) SetRegisteredAt(t time.Time) *CarCreate {
	cc.mutation.SetRegisteredAt(t)
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CarCreate) SetOwnerID(id int) *CarCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (cc *CarCreate) SetNillableOwnerID(id *int) *CarCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CarCreate) SetOwner(u *User) *CarCreate {
	return cc.SetOwnerID(u.ID)
}

// Mutation returns the CarMutation object of the builder.
func (cc *CarCreate) Mutation() *CarMutation {
	return cc.mutation
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	var (
		err  error
		node *Car
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CarMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *CarCreate) check() error {
	if _, ok := cc.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New("ent: missing required field \"model\"")}
	}
	if _, ok := cc.mutation.RegisteredAt(); !ok {
		return &ValidationError{Name: "registered_at", err: errors.New("ent: missing required field \"registered_at\"")}
	}
	return nil
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CarCreate) createSpec() (*Car, *sqlgraph.CreateSpec) {
	var (
		_node = &Car{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: car.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: car.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Model(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: car.FieldModel,
		})
		_node.Model = value
	}
	if value, ok := cc.mutation.RegisteredAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: car.FieldRegisteredAt,
		})
		_node.RegisteredAt = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_cars = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CarCreateBulk is the builder for creating many Car entities in bulk.
type CarCreateBulk struct {
	config
	builders []*CarCreate
}

// Save creates the Car entities in the database.
func (ccb *CarCreateBulk) Save(ctx context.Context) ([]*Car, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Car, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CarCreateBulk) SaveX(ctx context.Context) []*Car {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
