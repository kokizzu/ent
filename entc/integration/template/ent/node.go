// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/template/ent/group"
	"entgo.io/ent/entc/integration/template/ent/pet"
	"entgo.io/ent/entc/integration/template/ent/user"

	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitemty"`      // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (_m *Group) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     _m.ID,
		Type:   "Group",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(_m.MaxUsers); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "MaxUsers",
		Value: string(buf),
	}
	return node, nil
}

func (_m *Pet) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     _m.ID,
		Type:   "Pet",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(_m.Age); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "Age",
		Value: string(buf),
	}
	if buf, err = json.Marshal(_m.LicensedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "LicensedAt",
		Value: string(buf),
	}
	var ids []int
	ids, err = _m.QueryOwner().
		Select(user.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[0] = &Edge{
		IDs:  ids,
		Type: "User",
		Name: "Owner",
	}
	return node, nil
}

func (_m *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     _m.ID,
		Type:   "User",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(_m.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "Name",
		Value: string(buf),
	}
	var ids []int
	ids, err = _m.QueryPets().
		Select(pet.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[0] = &Edge{
		IDs:  ids,
		Type: "Pet",
		Name: "Pets",
	}
	ids, err = _m.QueryFriends().
		Select(user.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		IDs:  ids,
		Type: "User",
		Name: "Friends",
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

func (c *Client) Noder(ctx context.Context, id int) (Noder, error) {
	tables, err := c.tables.Load(ctx, c.driver)
	if err != nil {
		return nil, err
	}
	idx := id / (1<<32 - 1)
	if idx < 0 || idx >= len(tables) {
		return nil, fmt.Errorf("cannot resolve table from id %v: %w", id, &NotFoundError{"invalid/unknown"})
	}
	return c.noder(ctx, tables[idx], id)
}

func (c *Client) noder(ctx context.Context, tbl string, id int) (Noder, error) {
	switch tbl {
	case group.Table:
		n, err := c.Group.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return n, nil
	case pet.Table:
		n, err := c.Pet.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		n, err := c.User.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", tbl, &NotFoundError{"invalid/unknown"})
	}
}

type (
	tables struct {
		once  sync.Once
		sem   *semaphore.Weighted
		value atomic.Value
	}

	querier interface {
		Query(ctx context.Context, query string, args, v any) error
	}
)

func (t *tables) Load(ctx context.Context, querier querier) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, querier)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, querier querier) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := querier.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
