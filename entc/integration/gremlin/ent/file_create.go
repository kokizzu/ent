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

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/file"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetSetID sets the "set_id" field.
func (_c *FileCreate) SetSetID(i int) *FileCreate {
	_c.mutation.SetSetID(i)
	return _c
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (_c *FileCreate) SetNillableSetID(i *int) *FileCreate {
	if i != nil {
		_c.SetSetID(*i)
	}
	return _c
}

// SetSize sets the "size" field.
func (_c *FileCreate) SetSize(i int) *FileCreate {
	_c.mutation.SetSize(i)
	return _c
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (_c *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		_c.SetSize(*i)
	}
	return _c
}

// SetName sets the "name" field.
func (_c *FileCreate) SetName(s string) *FileCreate {
	_c.mutation.SetName(s)
	return _c
}

// SetUser sets the "user" field.
func (_c *FileCreate) SetUser(s string) *FileCreate {
	_c.mutation.SetUser(s)
	return _c
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (_c *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		_c.SetUser(*s)
	}
	return _c
}

// SetGroup sets the "group" field.
func (_c *FileCreate) SetGroup(s string) *FileCreate {
	_c.mutation.SetGroup(s)
	return _c
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (_c *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		_c.SetGroup(*s)
	}
	return _c
}

// SetOp sets the "op" field.
func (_c *FileCreate) SetOp(b bool) *FileCreate {
	_c.mutation.SetOpField(b)
	return _c
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (_c *FileCreate) SetNillableOp(b *bool) *FileCreate {
	if b != nil {
		_c.SetOp(*b)
	}
	return _c
}

// SetFieldID sets the "field_id" field.
func (_c *FileCreate) SetFieldID(i int) *FileCreate {
	_c.mutation.SetFieldID(i)
	return _c
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (_c *FileCreate) SetNillableFieldID(i *int) *FileCreate {
	if i != nil {
		_c.SetFieldID(*i)
	}
	return _c
}

// SetCreateTime sets the "create_time" field.
func (_c *FileCreate) SetCreateTime(t time.Time) *FileCreate {
	_c.mutation.SetCreateTime(t)
	return _c
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (_c *FileCreate) SetNillableCreateTime(t *time.Time) *FileCreate {
	if t != nil {
		_c.SetCreateTime(*t)
	}
	return _c
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_c *FileCreate) SetOwnerID(id string) *FileCreate {
	_c.mutation.SetOwnerID(id)
	return _c
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_c *FileCreate) SetNillableOwnerID(id *string) *FileCreate {
	if id != nil {
		_c = _c.SetOwnerID(*id)
	}
	return _c
}

// SetOwner sets the "owner" edge to the User entity.
func (_c *FileCreate) SetOwner(u *User) *FileCreate {
	return _c.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (_c *FileCreate) SetTypeID(id string) *FileCreate {
	_c.mutation.SetTypeID(id)
	return _c
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (_c *FileCreate) SetNillableTypeID(id *string) *FileCreate {
	if id != nil {
		_c = _c.SetTypeID(*id)
	}
	return _c
}

// SetType sets the "type" edge to the FileType entity.
func (_c *FileCreate) SetType(f *FileType) *FileCreate {
	return _c.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (_c *FileCreate) AddFieldIDs(ids ...string) *FileCreate {
	_c.mutation.AddFieldIDs(ids...)
	return _c
}

// AddField adds the "field" edges to the FieldType entity.
func (_c *FileCreate) AddField(f ...*FieldType) *FileCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return _c.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (_c *FileCreate) Mutation() *FileMutation {
	return _c.mutation
}

// Save creates the File in the database.
func (_c *FileCreate) Save(ctx context.Context) (*File, error) {
	_c.defaults()
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *FileCreate) SaveX(ctx context.Context) *File {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *FileCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *FileCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *FileCreate) defaults() {
	if _, ok := _c.mutation.Size(); !ok {
		v := file.DefaultSize
		_c.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *FileCreate) check() error {
	if v, ok := _c.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`ent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if _, ok := _c.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "File.size"`)}
	}
	if v, ok := _c.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	return nil
}

func (_c *FileCreate) gremlinSave(ctx context.Context) (*File, error) {
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
	rnode := &File{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *FileCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(file.Label)
	if value, ok := _c.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := _c.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := _c.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := _c.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := _c.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := _c.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := _c.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := _c.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	for _, id := range _c.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range _c.mutation.TypeIDs() {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range _c.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
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

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	err      error
	builders []*FileCreate
}
