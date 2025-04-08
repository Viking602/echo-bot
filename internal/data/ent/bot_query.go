// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/bot"
	"echo/internal/data/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BotQuery is the builder for querying Bot entities.
type BotQuery struct {
	config
	ctx        *QueryContext
	order      []bot.OrderOption
	inters     []Interceptor
	predicates []predicate.Bot
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BotQuery builder.
func (bq *BotQuery) Where(ps ...predicate.Bot) *BotQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BotQuery) Limit(limit int) *BotQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BotQuery) Offset(offset int) *BotQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BotQuery) Unique(unique bool) *BotQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BotQuery) Order(o ...bot.OrderOption) *BotQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Bot entity from the query.
// Returns a *NotFoundError when no Bot was found.
func (bq *BotQuery) First(ctx context.Context) (*Bot, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bot.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BotQuery) FirstX(ctx context.Context) *Bot {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bot ID from the query.
// Returns a *NotFoundError when no Bot ID was found.
func (bq *BotQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bot.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BotQuery) FirstIDX(ctx context.Context) int64 {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bot entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bot entity is found.
// Returns a *NotFoundError when no Bot entities are found.
func (bq *BotQuery) Only(ctx context.Context) (*Bot, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bot.Label}
	default:
		return nil, &NotSingularError{bot.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BotQuery) OnlyX(ctx context.Context) *Bot {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bot ID in the query.
// Returns a *NotSingularError when more than one Bot ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BotQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bot.Label}
	default:
		err = &NotSingularError{bot.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BotQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bots.
func (bq *BotQuery) All(ctx context.Context) ([]*Bot, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryAll)
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bot, *BotQuery]()
	return withInterceptors[[]*Bot](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BotQuery) AllX(ctx context.Context) []*Bot {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bot IDs.
func (bq *BotQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryIDs)
	if err = bq.Select(bot.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BotQuery) IDsX(ctx context.Context) []int64 {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BotQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryCount)
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BotQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BotQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BotQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryExist)
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BotQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BotQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BotQuery) Clone() *BotQuery {
	if bq == nil {
		return nil
	}
	return &BotQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]bot.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Bot{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BotID int64 `json:"bot_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bot.Query().
//		GroupBy(bot.FieldBotID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BotQuery) GroupBy(field string, fields ...string) *BotGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BotGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bot.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BotID int64 `json:"bot_id,omitempty"`
//	}
//
//	client.Bot.Query().
//		Select(bot.FieldBotID).
//		Scan(ctx, &v)
func (bq *BotQuery) Select(fields ...string) *BotSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BotSelect{BotQuery: bq}
	sbuild.label = bot.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BotSelect configured with the given aggregations.
func (bq *BotQuery) Aggregate(fns ...AggregateFunc) *BotSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BotQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !bot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BotQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bot, error) {
	var (
		nodes = []*Bot{}
		_spec = bq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bot).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bot{config: bq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bq *BotQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BotQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bot.Table, bot.Columns, sqlgraph.NewFieldSpec(bot.FieldID, field.TypeInt64))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bot.FieldID)
		for i := range fields {
			if fields[i] != bot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BotQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bot.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BotGroupBy is the group-by builder for Bot entities.
type BotGroupBy struct {
	selector
	build *BotQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BotGroupBy) Aggregate(fns ...AggregateFunc) *BotGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BotGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, ent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BotQuery, *BotGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BotGroupBy) sqlScan(ctx context.Context, root *BotQuery, v any) error {
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

// BotSelect is the builder for selecting fields of Bot entities.
type BotSelect struct {
	*BotQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BotSelect) Aggregate(fns ...AggregateFunc) *BotSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BotSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, ent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BotQuery, *BotSelect](ctx, bs.BotQuery, bs, bs.inters, v)
}

func (bs *BotSelect) sqlScan(ctx context.Context, root *BotQuery, v any) error {
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
