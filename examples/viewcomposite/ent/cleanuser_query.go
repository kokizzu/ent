// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/viewcomposite/ent/cleanuser"
	"entgo.io/ent/examples/viewcomposite/ent/predicate"
)

// CleanUserQuery is the builder for querying CleanUser entities.
type CleanUserQuery struct {
	config
	ctx        *QueryContext
	order      []cleanuser.OrderOption
	inters     []Interceptor
	predicates []predicate.CleanUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CleanUserQuery builder.
func (_q *CleanUserQuery) Where(ps ...predicate.CleanUser) *CleanUserQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *CleanUserQuery) Limit(limit int) *CleanUserQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *CleanUserQuery) Offset(offset int) *CleanUserQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *CleanUserQuery) Unique(unique bool) *CleanUserQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *CleanUserQuery) Order(o ...cleanuser.OrderOption) *CleanUserQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first CleanUser entity from the query.
// Returns a *NotFoundError when no CleanUser was found.
func (_q *CleanUserQuery) First(ctx context.Context) (*CleanUser, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cleanuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *CleanUserQuery) FirstX(ctx context.Context) *CleanUser {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single CleanUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CleanUser entity is found.
// Returns a *NotFoundError when no CleanUser entities are found.
func (_q *CleanUserQuery) Only(ctx context.Context) (*CleanUser, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cleanuser.Label}
	default:
		return nil, &NotSingularError{cleanuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *CleanUserQuery) OnlyX(ctx context.Context) *CleanUser {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of CleanUsers.
func (_q *CleanUserQuery) All(ctx context.Context) ([]*CleanUser, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CleanUser, *CleanUserQuery]()
	return withInterceptors[[]*CleanUser](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *CleanUserQuery) AllX(ctx context.Context) []*CleanUser {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (_q *CleanUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*CleanUserQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *CleanUserQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *CleanUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *CleanUserQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CleanUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *CleanUserQuery) Clone() *CleanUserQuery {
	if _q == nil {
		return nil
	}
	return &CleanUserQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]cleanuser.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.CleanUser{}, _q.predicates...),
		// clone intermediate query.
		sql:  _q.sql.Clone(),
		path: _q.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ID int `json:"id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CleanUser.Query().
//		GroupBy(cleanuser.FieldID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *CleanUserQuery) GroupBy(field string, fields ...string) *CleanUserGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CleanUserGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = cleanuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ID int `json:"id,omitempty"`
//	}
//
//	client.CleanUser.Query().
//		Select(cleanuser.FieldID).
//		Scan(ctx, &v)
func (_q *CleanUserQuery) Select(fields ...string) *CleanUserSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &CleanUserSelect{CleanUserQuery: _q}
	sbuild.label = cleanuser.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CleanUserSelect configured with the given aggregations.
func (_q *CleanUserQuery) Aggregate(fns ...AggregateFunc) *CleanUserSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *CleanUserQuery) prepareQuery(ctx context.Context) error {
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
		if !cleanuser.ValidColumn(f) {
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

func (_q *CleanUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CleanUser, error) {
	var (
		nodes = []*CleanUser{}
		_spec = _q.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CleanUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CleanUser{config: _q.config}
		nodes = append(nodes, node)
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
	return nodes, nil
}

func (_q *CleanUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := _q.querySpec()
	_spec.Node.Columns = _q.ctx.Fields
	if len(_q.ctx.Fields) > 0 {
		_spec.Unique = _q.ctx.Unique != nil && *_q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, _q.driver, _spec)
}

func (_q *CleanUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cleanuser.Table, cleanuser.Columns, nil)
	_spec.From = _q.sql
	if unique := _q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if _q.path != nil {
		_spec.Unique = true
	}
	if fields := _q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
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

func (_q *CleanUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(_q.driver.Dialect())
	t1 := builder.Table(cleanuser.Table)
	columns := _q.ctx.Fields
	if len(columns) == 0 {
		columns = cleanuser.Columns
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

// CleanUserGroupBy is the group-by builder for CleanUser entities.
type CleanUserGroupBy struct {
	selector
	build *CleanUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cugb *CleanUserGroupBy) Aggregate(fns ...AggregateFunc) *CleanUserGroupBy {
	cugb.fns = append(cugb.fns, fns...)
	return cugb
}

// Scan applies the selector query and scans the result into the given value.
func (cugb *CleanUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cugb.build.ctx, ent.OpQueryGroupBy)
	if err := cugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CleanUserQuery, *CleanUserGroupBy](ctx, cugb.build, cugb, cugb.build.inters, v)
}

func (cugb *CleanUserGroupBy) sqlScan(ctx context.Context, root *CleanUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cugb.fns))
	for _, fn := range cugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cugb.flds)+len(cugb.fns))
		for _, f := range *cugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CleanUserSelect is the builder for selecting fields of CleanUser entities.
type CleanUserSelect struct {
	*CleanUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cus *CleanUserSelect) Aggregate(fns ...AggregateFunc) *CleanUserSelect {
	cus.fns = append(cus.fns, fns...)
	return cus
}

// Scan applies the selector query and scans the result into the given value.
func (cus *CleanUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cus.ctx, ent.OpQuerySelect)
	if err := cus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CleanUserQuery, *CleanUserSelect](ctx, cus.CleanUserQuery, cus, cus.inters, v)
}

func (cus *CleanUserSelect) sqlScan(ctx context.Context, root *CleanUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cus.fns))
	for _, fn := range cus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
