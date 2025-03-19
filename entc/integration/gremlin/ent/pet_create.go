// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/pet"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
	"github.com/google/uuid"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (_c *PetCreate) SetAge(f float64) *PetCreate {
	_c.mutation.SetAge(f)
	return _c
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_c *PetCreate) SetNillableAge(f *float64) *PetCreate {
	if f != nil {
		_c.SetAge(*f)
	}
	return _c
}

// SetName sets the "name" field.
func (_c *PetCreate) SetName(s string) *PetCreate {
	_c.mutation.SetName(s)
	return _c
}

// SetUUID sets the "uuid" field.
func (_c *PetCreate) SetUUID(u uuid.UUID) *PetCreate {
	_c.mutation.SetUUID(u)
	return _c
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (_c *PetCreate) SetNillableUUID(u *uuid.UUID) *PetCreate {
	if u != nil {
		_c.SetUUID(*u)
	}
	return _c
}

// SetNickname sets the "nickname" field.
func (_c *PetCreate) SetNickname(s string) *PetCreate {
	_c.mutation.SetNickname(s)
	return _c
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (_c *PetCreate) SetNillableNickname(s *string) *PetCreate {
	if s != nil {
		_c.SetNickname(*s)
	}
	return _c
}

// SetTrained sets the "trained" field.
func (_c *PetCreate) SetTrained(b bool) *PetCreate {
	_c.mutation.SetTrained(b)
	return _c
}

// SetNillableTrained sets the "trained" field if the given value is not nil.
func (_c *PetCreate) SetNillableTrained(b *bool) *PetCreate {
	if b != nil {
		_c.SetTrained(*b)
	}
	return _c
}

// SetOptionalTime sets the "optional_time" field.
func (_c *PetCreate) SetOptionalTime(t time.Time) *PetCreate {
	_c.mutation.SetOptionalTime(t)
	return _c
}

// SetNillableOptionalTime sets the "optional_time" field if the given value is not nil.
func (_c *PetCreate) SetNillableOptionalTime(t *time.Time) *PetCreate {
	if t != nil {
		_c.SetOptionalTime(*t)
	}
	return _c
}

// SetTeamID sets the "team" edge to the User entity by ID.
func (_c *PetCreate) SetTeamID(id string) *PetCreate {
	_c.mutation.SetTeamID(id)
	return _c
}

// SetNillableTeamID sets the "team" edge to the User entity by ID if the given value is not nil.
func (_c *PetCreate) SetNillableTeamID(id *string) *PetCreate {
	if id != nil {
		_c = _c.SetTeamID(*id)
	}
	return _c
}

// SetTeam sets the "team" edge to the User entity.
func (_c *PetCreate) SetTeam(u *User) *PetCreate {
	return _c.SetTeamID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_c *PetCreate) SetOwnerID(id string) *PetCreate {
	_c.mutation.SetOwnerID(id)
	return _c
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_c *PetCreate) SetNillableOwnerID(id *string) *PetCreate {
	if id != nil {
		_c = _c.SetOwnerID(*id)
	}
	return _c
}

// SetOwner sets the "owner" edge to the User entity.
func (_c *PetCreate) SetOwner(u *User) *PetCreate {
	return _c.SetOwnerID(u.ID)
}

// Mutation returns the PetMutation object of the builder.
func (_c *PetCreate) Mutation() *PetMutation {
	return _c.mutation
}

// Save creates the Pet in the database.
func (_c *PetCreate) Save(ctx context.Context) (*Pet, error) {
	_c.defaults()
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *PetCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *PetCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *PetCreate) defaults() {
	if _, ok := _c.mutation.Age(); !ok {
		v := pet.DefaultAge
		_c.mutation.SetAge(v)
	}
	if _, ok := _c.mutation.Trained(); !ok {
		v := pet.DefaultTrained
		_c.mutation.SetTrained(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *PetCreate) check() error {
	if _, ok := _c.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Pet.age"`)}
	}
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Pet.name"`)}
	}
	if _, ok := _c.mutation.Trained(); !ok {
		return &ValidationError{Name: "trained", err: errors.New(`ent: missing required field "Pet.trained"`)}
	}
	return nil
}

func (_c *PetCreate) gremlinSave(ctx context.Context) (*Pet, error) {
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
	rnode := &Pet{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *PetCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(pet.Label)
	if value, ok := _c.mutation.Age(); ok {
		v.Property(dsl.Single, pet.FieldAge, value)
	}
	if value, ok := _c.mutation.Name(); ok {
		v.Property(dsl.Single, pet.FieldName, value)
	}
	if value, ok := _c.mutation.UUID(); ok {
		v.Property(dsl.Single, pet.FieldUUID, value)
	}
	if value, ok := _c.mutation.Nickname(); ok {
		v.Property(dsl.Single, pet.FieldNickname, value)
	}
	if value, ok := _c.mutation.Trained(); ok {
		v.Property(dsl.Single, pet.FieldTrained, value)
	}
	if value, ok := _c.mutation.OptionalTime(); ok {
		v.Property(dsl.Single, pet.FieldOptionalTime, value)
	}
	for _, id := range _c.mutation.TeamIDs() {
		v.AddE(user.TeamLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(pet.Label, user.TeamLabel, id)),
		})
	}
	for _, id := range _c.mutation.OwnerIDs() {
		v.AddE(user.PetsLabel).From(g.V(id)).InV()
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

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	err      error
	builders []*PetCreate
}
