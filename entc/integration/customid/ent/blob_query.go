// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/entc/integration/customid/ent/bloblink"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlobQuery is the builder for querying Blob entities.
type BlobQuery struct {
	config
	ctx           *QueryContext
	order         []blob.OrderOption
	inters        []Interceptor
	predicates    []predicate.Blob
	withParent    *BlobQuery
	withLinks     *BlobQuery
	withBlobLinks *BlobLinkQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlobQuery builder.
func (_q *BlobQuery) Where(ps ...predicate.Blob) *BlobQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *BlobQuery) Limit(limit int) *BlobQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *BlobQuery) Offset(offset int) *BlobQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *BlobQuery) Unique(unique bool) *BlobQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *BlobQuery) Order(o ...blob.OrderOption) *BlobQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// QueryParent chains the current query on the "parent" edge.
func (_q *BlobQuery) QueryParent() *BlobQuery {
	query := (&BlobClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blob.Table, blob.FieldID, selector),
			sqlgraph.To(blob.Table, blob.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, blob.ParentTable, blob.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryLinks chains the current query on the "links" edge.
func (_q *BlobQuery) QueryLinks() *BlobQuery {
	query := (&BlobClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blob.Table, blob.FieldID, selector),
			sqlgraph.To(blob.Table, blob.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, blob.LinksTable, blob.LinksPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBlobLinks chains the current query on the "blob_links" edge.
func (_q *BlobQuery) QueryBlobLinks() *BlobLinkQuery {
	query := (&BlobLinkClient{config: _q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := _q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := _q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blob.Table, blob.FieldID, selector),
			sqlgraph.To(bloblink.Table, bloblink.BlobColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, blob.BlobLinksTable, blob.BlobLinksColumn),
		)
		fromU = sqlgraph.SetNeighbors(_q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Blob entity from the query.
// Returns a *NotFoundError when no Blob was found.
func (_q *BlobQuery) First(ctx context.Context) (*Blob, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blob.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *BlobQuery) FirstX(ctx context.Context) *Blob {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Blob ID from the query.
// Returns a *NotFoundError when no Blob ID was found.
func (_q *BlobQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blob.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *BlobQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Blob entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Blob entity is found.
// Returns a *NotFoundError when no Blob entities are found.
func (_q *BlobQuery) Only(ctx context.Context) (*Blob, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blob.Label}
	default:
		return nil, &NotSingularError{blob.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *BlobQuery) OnlyX(ctx context.Context) *Blob {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Blob ID in the query.
// Returns a *NotSingularError when more than one Blob ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *BlobQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blob.Label}
	default:
		err = &NotSingularError{blob.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *BlobQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Blobs.
func (_q *BlobQuery) All(ctx context.Context) ([]*Blob, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Blob, *BlobQuery]()
	return withInterceptors[[]*Blob](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *BlobQuery) AllX(ctx context.Context) []*Blob {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Blob IDs.
func (_q *BlobQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(blob.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *BlobQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *BlobQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*BlobQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *BlobQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *BlobQuery) Exist(ctx context.Context) (bool, error) {
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
func (_q *BlobQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlobQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *BlobQuery) Clone() *BlobQuery {
	if _q == nil {
		return nil
	}
	return &BlobQuery{
		config:        _q.config,
		ctx:           _q.ctx.Clone(),
		order:         append([]blob.OrderOption{}, _q.order...),
		inters:        append([]Interceptor{}, _q.inters...),
		predicates:    append([]predicate.Blob{}, _q.predicates...),
		withParent:    _q.withParent.Clone(),
		withLinks:     _q.withLinks.Clone(),
		withBlobLinks: _q.withBlobLinks.Clone(),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *BlobQuery) WithParent(opts ...func(*BlobQuery)) *BlobQuery {
	query := (&BlobClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withParent = query
	return _q
}

// WithLinks tells the query-builder to eager-load the nodes that are connected to
// the "links" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *BlobQuery) WithLinks(opts ...func(*BlobQuery)) *BlobQuery {
	query := (&BlobClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withLinks = query
	return _q
}

// WithBlobLinks tells the query-builder to eager-load the nodes that are connected to
// the "blob_links" edge. The optional arguments are used to configure the query builder of the edge.
func (_q *BlobQuery) WithBlobLinks(opts ...func(*BlobLinkQuery)) *BlobQuery {
	query := (&BlobLinkClient{config: _q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	_q.withBlobLinks = query
	return _q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Blob.Query().
//		GroupBy(blob.FieldUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *BlobQuery) GroupBy(field string, fields ...string) *BlobGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BlobGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = blob.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UUID uuid.UUID `json:"uuid,omitempty"`
//	}
//
//	client.Blob.Query().
//		Select(blob.FieldUUID).
//		Scan(ctx, &v)
func (_q *BlobQuery) Select(fields ...string) *BlobSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &BlobSelect{BlobQuery: _q}
	sbuild.label = blob.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BlobSelect configured with the given aggregations.
func (_q *BlobQuery) Aggregate(fns ...AggregateFunc) *BlobSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *BlobQuery) prepareQuery(ctx context.Context) error {
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
		if !blob.ValidColumn(f) {
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

func (_q *BlobQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Blob, error) {
	var (
		nodes       = []*Blob{}
		withFKs     = _q.withFKs
		_spec       = _q.querySpec()
		loadedTypes = [3]bool{
			_q.withParent != nil,
			_q.withLinks != nil,
			_q.withBlobLinks != nil,
		}
	)
	if _q.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, blob.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Blob).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Blob{config: _q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := _q.withParent; query != nil {
		if err := _q.loadParent(ctx, query, nodes, nil,
			func(n *Blob, e *Blob) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := _q.withLinks; query != nil {
		if err := _q.loadLinks(ctx, query, nodes,
			func(n *Blob) { n.Edges.Links = []*Blob{} },
			func(n *Blob, e *Blob) { n.Edges.Links = append(n.Edges.Links, e) }); err != nil {
			return nil, err
		}
	}
	if query := _q.withBlobLinks; query != nil {
		if err := _q.loadBlobLinks(ctx, query, nodes,
			func(n *Blob) { n.Edges.BlobLinks = []*BlobLink{} },
			func(n *Blob, e *BlobLink) { n.Edges.BlobLinks = append(n.Edges.BlobLinks, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (_q *BlobQuery) loadParent(ctx context.Context, query *BlobQuery, nodes []*Blob, init func(*Blob), assign func(*Blob, *Blob)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Blob)
	for i := range nodes {
		if nodes[i].blob_parent == nil {
			continue
		}
		fk := *nodes[i].blob_parent
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(blob.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "blob_parent" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (_q *BlobQuery) loadLinks(ctx context.Context, query *BlobQuery, nodes []*Blob, init func(*Blob), assign func(*Blob, *Blob)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uuid.UUID]*Blob)
	nids := make(map[uuid.UUID]map[*Blob]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(blob.LinksTable)
		s.Join(joinT).On(s.C(blob.FieldID), joinT.C(blob.LinksPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(blob.LinksPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(blob.LinksPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(uuid.UUID)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := *values[0].(*uuid.UUID)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*Blob]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Blob](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "links" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (_q *BlobQuery) loadBlobLinks(ctx context.Context, query *BlobLinkQuery, nodes []*Blob, init func(*Blob), assign func(*Blob, *BlobLink)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Blob)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bloblink.FieldBlobID)
	}
	query.Where(predicate.BlobLink(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(blob.BlobLinksColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BlobID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "blob_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (_q *BlobQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *BlobQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(blob.Table, blob.Columns, sqlgraph.NewFieldSpec(blob.FieldID, field.TypeUUID))
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blob.FieldID)
		for i := range fields {
			if fields[i] != blob.FieldID {
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

func (_q *BlobQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(blob.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = blob.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if _q.sql != nil {
		selector = _q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if _q.ctx.Unique != nil && *_q.ctx.Unique {
		selector.Distinct()
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

// BlobGroupBy is the group-by builder for Blob entities.
type BlobGroupBy struct {
	selector
	build *BlobQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BlobGroupBy) Aggregate(fns ...AggregateFunc) *BlobGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BlobGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, ent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlobQuery, *BlobGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BlobGroupBy) sqlScan(ctx context.Context, root *BlobQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BlobSelect is the builder for selecting fields of Blob entities.
type BlobSelect struct {
	*BlobQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BlobSelect) Aggregate(fns ...AggregateFunc) *BlobSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BlobSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, ent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlobQuery, *BlobSelect](ctx, bs.BlobQuery, bs, bs.inters, v)
}

func (bs *BlobSelect) sqlScan(ctx context.Context, root *BlobQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
