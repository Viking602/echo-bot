// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/predicate"
	"echo/internal/data/ent/subbililive"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubBiliLiveQuery is the builder for querying SubBiliLive entities.
type SubBiliLiveQuery struct {
	config
	ctx        *QueryContext
	order      []subbililive.OrderOption
	inters     []Interceptor
	predicates []predicate.SubBiliLive
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubBiliLiveQuery builder.
func (sblq *SubBiliLiveQuery) Where(ps ...predicate.SubBiliLive) *SubBiliLiveQuery {
	sblq.predicates = append(sblq.predicates, ps...)
	return sblq
}

// Limit the number of records to be returned by this query.
func (sblq *SubBiliLiveQuery) Limit(limit int) *SubBiliLiveQuery {
	sblq.ctx.Limit = &limit
	return sblq
}

// Offset to start from.
func (sblq *SubBiliLiveQuery) Offset(offset int) *SubBiliLiveQuery {
	sblq.ctx.Offset = &offset
	return sblq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sblq *SubBiliLiveQuery) Unique(unique bool) *SubBiliLiveQuery {
	sblq.ctx.Unique = &unique
	return sblq
}

// Order specifies how the records should be ordered.
func (sblq *SubBiliLiveQuery) Order(o ...subbililive.OrderOption) *SubBiliLiveQuery {
	sblq.order = append(sblq.order, o...)
	return sblq
}

// First returns the first SubBiliLive entity from the query.
// Returns a *NotFoundError when no SubBiliLive was found.
func (sblq *SubBiliLiveQuery) First(ctx context.Context) (*SubBiliLive, error) {
	nodes, err := sblq.Limit(1).All(setContextOp(ctx, sblq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subbililive.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) FirstX(ctx context.Context) *SubBiliLive {
	node, err := sblq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SubBiliLive ID from the query.
// Returns a *NotFoundError when no SubBiliLive ID was found.
func (sblq *SubBiliLiveQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sblq.Limit(1).IDs(setContextOp(ctx, sblq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subbililive.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) FirstIDX(ctx context.Context) int64 {
	id, err := sblq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SubBiliLive entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SubBiliLive entity is found.
// Returns a *NotFoundError when no SubBiliLive entities are found.
func (sblq *SubBiliLiveQuery) Only(ctx context.Context) (*SubBiliLive, error) {
	nodes, err := sblq.Limit(2).All(setContextOp(ctx, sblq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subbililive.Label}
	default:
		return nil, &NotSingularError{subbililive.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) OnlyX(ctx context.Context) *SubBiliLive {
	node, err := sblq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SubBiliLive ID in the query.
// Returns a *NotSingularError when more than one SubBiliLive ID is found.
// Returns a *NotFoundError when no entities are found.
func (sblq *SubBiliLiveQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = sblq.Limit(2).IDs(setContextOp(ctx, sblq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subbililive.Label}
	default:
		err = &NotSingularError{subbililive.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := sblq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SubBiliLives.
func (sblq *SubBiliLiveQuery) All(ctx context.Context) ([]*SubBiliLive, error) {
	ctx = setContextOp(ctx, sblq.ctx, ent.OpQueryAll)
	if err := sblq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SubBiliLive, *SubBiliLiveQuery]()
	return withInterceptors[[]*SubBiliLive](ctx, sblq, qr, sblq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) AllX(ctx context.Context) []*SubBiliLive {
	nodes, err := sblq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SubBiliLive IDs.
func (sblq *SubBiliLiveQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if sblq.ctx.Unique == nil && sblq.path != nil {
		sblq.Unique(true)
	}
	ctx = setContextOp(ctx, sblq.ctx, ent.OpQueryIDs)
	if err = sblq.Select(subbililive.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) IDsX(ctx context.Context) []int64 {
	ids, err := sblq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sblq *SubBiliLiveQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sblq.ctx, ent.OpQueryCount)
	if err := sblq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sblq, querierCount[*SubBiliLiveQuery](), sblq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) CountX(ctx context.Context) int {
	count, err := sblq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sblq *SubBiliLiveQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sblq.ctx, ent.OpQueryExist)
	switch _, err := sblq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sblq *SubBiliLiveQuery) ExistX(ctx context.Context) bool {
	exist, err := sblq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubBiliLiveQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sblq *SubBiliLiveQuery) Clone() *SubBiliLiveQuery {
	if sblq == nil {
		return nil
	}
	return &SubBiliLiveQuery{
		config:     sblq.config,
		ctx:        sblq.ctx.Clone(),
		order:      append([]subbililive.OrderOption{}, sblq.order...),
		inters:     append([]Interceptor{}, sblq.inters...),
		predicates: append([]predicate.SubBiliLive{}, sblq.predicates...),
		// clone intermediate query.
		sql:  sblq.sql.Clone(),
		path: sblq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RoomID int64 `json:"room_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SubBiliLive.Query().
//		GroupBy(subbililive.FieldRoomID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sblq *SubBiliLiveQuery) GroupBy(field string, fields ...string) *SubBiliLiveGroupBy {
	sblq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubBiliLiveGroupBy{build: sblq}
	grbuild.flds = &sblq.ctx.Fields
	grbuild.label = subbililive.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RoomID int64 `json:"room_id,omitempty"`
//	}
//
//	client.SubBiliLive.Query().
//		Select(subbililive.FieldRoomID).
//		Scan(ctx, &v)
func (sblq *SubBiliLiveQuery) Select(fields ...string) *SubBiliLiveSelect {
	sblq.ctx.Fields = append(sblq.ctx.Fields, fields...)
	sbuild := &SubBiliLiveSelect{SubBiliLiveQuery: sblq}
	sbuild.label = subbililive.Label
	sbuild.flds, sbuild.scan = &sblq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubBiliLiveSelect configured with the given aggregations.
func (sblq *SubBiliLiveQuery) Aggregate(fns ...AggregateFunc) *SubBiliLiveSelect {
	return sblq.Select().Aggregate(fns...)
}

func (sblq *SubBiliLiveQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sblq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sblq); err != nil {
				return err
			}
		}
	}
	for _, f := range sblq.ctx.Fields {
		if !subbililive.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sblq.path != nil {
		prev, err := sblq.path(ctx)
		if err != nil {
			return err
		}
		sblq.sql = prev
	}
	return nil
}

func (sblq *SubBiliLiveQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SubBiliLive, error) {
	var (
		nodes = []*SubBiliLive{}
		_spec = sblq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SubBiliLive).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SubBiliLive{config: sblq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sblq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sblq *SubBiliLiveQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sblq.querySpec()
	_spec.Node.Columns = sblq.ctx.Fields
	if len(sblq.ctx.Fields) > 0 {
		_spec.Unique = sblq.ctx.Unique != nil && *sblq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sblq.driver, _spec)
}

func (sblq *SubBiliLiveQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subbililive.Table, subbililive.Columns, sqlgraph.NewFieldSpec(subbililive.FieldID, field.TypeInt64))
	_spec.From = sblq.sql
	if unique := sblq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sblq.path != nil {
		_spec.Unique = true
	}
	if fields := sblq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subbililive.FieldID)
		for i := range fields {
			if fields[i] != subbililive.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sblq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sblq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sblq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sblq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sblq *SubBiliLiveQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sblq.driver.Dialect())
	t1 := builder.Table(subbililive.Table)
	columns := sblq.ctx.Fields
	if len(columns) == 0 {
		columns = subbililive.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sblq.sql != nil {
		selector = sblq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sblq.ctx.Unique != nil && *sblq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sblq.predicates {
		p(selector)
	}
	for _, p := range sblq.order {
		p(selector)
	}
	if offset := sblq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sblq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubBiliLiveGroupBy is the group-by builder for SubBiliLive entities.
type SubBiliLiveGroupBy struct {
	selector
	build *SubBiliLiveQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sblgb *SubBiliLiveGroupBy) Aggregate(fns ...AggregateFunc) *SubBiliLiveGroupBy {
	sblgb.fns = append(sblgb.fns, fns...)
	return sblgb
}

// Scan applies the selector query and scans the result into the given value.
func (sblgb *SubBiliLiveGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sblgb.build.ctx, ent.OpQueryGroupBy)
	if err := sblgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubBiliLiveQuery, *SubBiliLiveGroupBy](ctx, sblgb.build, sblgb, sblgb.build.inters, v)
}

func (sblgb *SubBiliLiveGroupBy) sqlScan(ctx context.Context, root *SubBiliLiveQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sblgb.fns))
	for _, fn := range sblgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sblgb.flds)+len(sblgb.fns))
		for _, f := range *sblgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sblgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sblgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubBiliLiveSelect is the builder for selecting fields of SubBiliLive entities.
type SubBiliLiveSelect struct {
	*SubBiliLiveQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sbls *SubBiliLiveSelect) Aggregate(fns ...AggregateFunc) *SubBiliLiveSelect {
	sbls.fns = append(sbls.fns, fns...)
	return sbls
}

// Scan applies the selector query and scans the result into the given value.
func (sbls *SubBiliLiveSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sbls.ctx, ent.OpQuerySelect)
	if err := sbls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubBiliLiveQuery, *SubBiliLiveSelect](ctx, sbls.SubBiliLiveQuery, sbls, sbls.inters, v)
}

func (sbls *SubBiliLiveSelect) sqlScan(ctx context.Context, root *SubBiliLiveQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sbls.fns))
	for _, fn := range sbls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sbls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sbls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
