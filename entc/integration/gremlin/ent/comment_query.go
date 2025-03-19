// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/comment"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// CommentQuery is the builder for querying Comment entities.
type CommentQuery struct {
	config
	ctx        *QueryContext
	order      []comment.OrderOption
	inters     []Interceptor
	predicates []predicate.Comment
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the CommentQuery builder.
func (_q *CommentQuery) Where(ps ...predicate.Comment) *CommentQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *CommentQuery) Limit(limit int) *CommentQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *CommentQuery) Offset(offset int) *CommentQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *CommentQuery) Unique(unique bool) *CommentQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *CommentQuery) Order(o ...comment.OrderOption) *CommentQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first Comment entity from the query.
// Returns a *NotFoundError when no Comment was found.
func (_q *CommentQuery) First(ctx context.Context) (*Comment, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{comment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *CommentQuery) FirstX(ctx context.Context) *Comment {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Comment ID from the query.
// Returns a *NotFoundError when no Comment ID was found.
func (_q *CommentQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *CommentQuery) FirstIDX(ctx context.Context) string {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Comment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Comment entity is found.
// Returns a *NotFoundError when no Comment entities are found.
func (_q *CommentQuery) Only(ctx context.Context) (*Comment, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{comment.Label}
	default:
		return nil, &NotSingularError{comment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *CommentQuery) OnlyX(ctx context.Context) *Comment {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Comment ID in the query.
// Returns a *NotSingularError when more than one Comment ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *CommentQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comment.Label}
	default:
		err = &NotSingularError{comment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *CommentQuery) OnlyIDX(ctx context.Context) string {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Comments.
func (_q *CommentQuery) All(ctx context.Context) ([]*Comment, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Comment, *CommentQuery]()
	return withInterceptors[[]*Comment](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *CommentQuery) AllX(ctx context.Context) []*Comment {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Comment IDs.
func (_q *CommentQuery) IDs(ctx context.Context) (ids []string, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(comment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *CommentQuery) IDsX(ctx context.Context) []string {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *CommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*CommentQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *CommentQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *CommentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *CommentQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *CommentQuery) Clone() *CommentQuery {
	if _q == nil {
		return nil
	}
	return &CommentQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]comment.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Comment{}, _q.predicates...),
		// clone intermediate query.
		gremlin: _q.gremlin.Clone(),
		path:    _q.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UniqueInt int `json:"unique_int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Comment.Query().
//		GroupBy(comment.FieldUniqueInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *CommentQuery) GroupBy(field string, fields ...string) *CommentGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommentGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = comment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UniqueInt int `json:"unique_int,omitempty"`
//	}
//
//	client.Comment.Query().
//		Select(comment.FieldUniqueInt).
//		Scan(ctx, &v)
func (_q *CommentQuery) Select(fields ...string) *CommentSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &CommentSelect{CommentQuery: _q}
	sbuild.label = comment.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentSelect configured with the given aggregations.
func (_q *CommentQuery) Aggregate(fns ...AggregateFunc) *CommentSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *CommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.gremlin = prev
	}
	return nil
}

func (_q *CommentQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Comment, error) {
	res := &gremlin.Response{}
	traversal := _q.gremlinQuery(ctx)
	if len(_q.ctx.Fields) > 0 {
		fields := make([]any, len(_q.ctx.Fields))
		for i, f := range _q.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var _ms Comments
	if err := _ms.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range _ms {
		_ms[i].config = _q.config
	}
	return _ms, nil
}

func (_q *CommentQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _q.gremlinQuery(ctx).Count().Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (_q *CommentQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(comment.Label)
	if _q.gremlin != nil {
		v = _q.gremlin.Clone()
	}
	for _, p := range _q.predicates {
		p(v)
	}
	if len(_q.order) > 0 {
		v.Order()
		for _, p := range _q.order {
			p(v)
		}
	}
	switch limit, offset := _q.ctx.Limit, _q.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := _q.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// CommentGroupBy is the group-by builder for Comment entities.
type CommentGroupBy struct {
	selector
	build *CommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommentGroupBy) Aggregate(fns ...AggregateFunc) *CommentGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CommentGroupBy) gremlinScan(ctx context.Context, root *CommentQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range cgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *cgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*cgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := cgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*cgb.flds)+len(cgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// CommentSelect is the builder for selecting fields of Comment entities.
type CommentSelect struct {
	*CommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommentSelect) Aggregate(fns ...AggregateFunc) *CommentSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentSelect](ctx, cs.CommentQuery, cs, cs.inters, v)
}

func (cs *CommentSelect) gremlinScan(ctx context.Context, root *CommentQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := cs.ctx.Fields; len(fields) == 1 {
		if fields[0] != comment.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(cs.ctx.Fields))
		for i, f := range cs.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := cs.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
