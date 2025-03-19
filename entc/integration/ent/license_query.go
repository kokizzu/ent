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
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/license"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// LicenseQuery is the builder for querying License entities.
type LicenseQuery struct {
	config
	ctx        *QueryContext
	order      []license.OrderOption
	inters     []Interceptor
	predicates []predicate.License
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LicenseQuery builder.
func (_q *LicenseQuery) Where(ps ...predicate.License) *LicenseQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *LicenseQuery) Limit(limit int) *LicenseQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *LicenseQuery) Offset(offset int) *LicenseQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *LicenseQuery) Unique(unique bool) *LicenseQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *LicenseQuery) Order(o ...license.OrderOption) *LicenseQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first License entity from the query.
// Returns a *NotFoundError when no License was found.
func (_q *LicenseQuery) First(ctx context.Context) (*License, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{license.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *LicenseQuery) FirstX(ctx context.Context) *License {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first License ID from the query.
// Returns a *NotFoundError when no License ID was found.
func (_q *LicenseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{license.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *LicenseQuery) FirstIDX(ctx context.Context) int {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single License entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one License entity is found.
// Returns a *NotFoundError when no License entities are found.
func (_q *LicenseQuery) Only(ctx context.Context) (*License, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{license.Label}
	default:
		return nil, &NotSingularError{license.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *LicenseQuery) OnlyX(ctx context.Context) *License {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only License ID in the query.
// Returns a *NotSingularError when more than one License ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *LicenseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{license.Label}
	default:
		err = &NotSingularError{license.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *LicenseQuery) OnlyIDX(ctx context.Context) int {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Licenses.
func (_q *LicenseQuery) All(ctx context.Context) ([]*License, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*License, *LicenseQuery]()
	return withInterceptors[[]*License](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *LicenseQuery) AllX(ctx context.Context) []*License {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of License IDs.
func (_q *LicenseQuery) IDs(ctx context.Context) (ids []int, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(license.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *LicenseQuery) IDsX(ctx context.Context) []int {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *LicenseQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*LicenseQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *LicenseQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *LicenseQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *LicenseQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LicenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *LicenseQuery) Clone() *LicenseQuery {
	if _q == nil {
		return nil
	}
	return &LicenseQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]license.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.License{}, _q.predicates...),
		// clone intermediate query.
		sql:       _q.sql.Clone(),
		path:      _q.path,
		modifiers: append([]func(*sql.Selector){}, _q.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.License.Query().
//		GroupBy(license.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *LicenseQuery) GroupBy(field string, fields ...string) *LicenseGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LicenseGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = license.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.License.Query().
//		Select(license.FieldCreateTime).
//		Scan(ctx, &v)
func (_q *LicenseQuery) Select(fields ...string) *LicenseSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &LicenseSelect{LicenseQuery: _q}
	sbuild.label = license.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LicenseSelect configured with the given aggregations.
func (_q *LicenseQuery) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *LicenseQuery) prepareQuery(ctx context.Context) error {
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
	for _, f := range _q.ctx.Fields {
		if !license.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.sql = prev
	}
	return nil
}

func (_q *LicenseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*License, error) {
	var (
		nodes = []*License{}
		_spec = _q.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*License).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &License{config: _q.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, _q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (_q *LicenseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	if len(_q.modifiers) > 0 {
		_spec.Modifiers = _q.modifiers
	}
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *LicenseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(license.Table, license.Columns, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, license.FieldID)
		for i := range fields {
			if fields[i] != license.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := _q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := _q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := _q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := _q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (_q *LicenseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(license.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = license.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range _q.modifiers {
		m(selector)
	}
	for _, p := range _q.predicates {
		p(selector)
	}
	for _, p := range _q.order {
		p(selector)
	}
	if offset := _q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := _q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (_q *LicenseQuery) ForUpdate(opts ...sql.LockOption) *LicenseQuery {
	if _q.driver.Dialect() == dialect.Postgres {
		_q.Unique(false)
	}
	_q.modifiers = append(_q.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return _q
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (_q *LicenseQuery) ForShare(opts ...sql.LockOption) *LicenseQuery {
	if _q.driver.Dialect() == dialect.Postgres {
		_q.Unique(false)
	}
	_q.modifiers = append(_q.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return _q
}

// Modify adds a query modifier for attaching custom logic to queries.
func (_q *LicenseQuery) Modify(modifiers ...func(s *sql.Selector)) *LicenseSelect {
	_q.modifiers = append(_q.modifiers, modifiers...)
	return _q.Select()
}

// LicenseGroupBy is the group-by builder for License entities.
type LicenseGroupBy struct {
	selector
	build *LicenseQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LicenseGroupBy) Aggregate(fns ...AggregateFunc) *LicenseGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LicenseGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, ent.OpQueryGroupBy)
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LicenseQuery, *LicenseGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LicenseGroupBy) sqlScan(ctx context.Context, root *LicenseQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LicenseSelect is the builder for selecting fields of License entities.
type LicenseSelect struct {
	*LicenseQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LicenseSelect) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LicenseSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, ent.OpQuerySelect)
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LicenseQuery, *LicenseSelect](ctx, ls.LicenseQuery, ls, ls.inters, v)
}

func (ls *LicenseSelect) sqlScan(ctx context.Context, root *LicenseQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ls *LicenseSelect) Modify(modifiers ...func(s *sql.Selector)) *LicenseSelect {
	ls.modifiers = append(ls.modifiers, modifiers...)
	return ls
}
