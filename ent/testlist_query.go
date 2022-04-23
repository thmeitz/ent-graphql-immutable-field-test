// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/testlist"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// TestListQuery is the builder for querying TestList entities.
type TestListQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TestList
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*TestList) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestListQuery builder.
func (tlq *TestListQuery) Where(ps ...predicate.TestList) *TestListQuery {
	tlq.predicates = append(tlq.predicates, ps...)
	return tlq
}

// Limit adds a limit step to the query.
func (tlq *TestListQuery) Limit(limit int) *TestListQuery {
	tlq.limit = &limit
	return tlq
}

// Offset adds an offset step to the query.
func (tlq *TestListQuery) Offset(offset int) *TestListQuery {
	tlq.offset = &offset
	return tlq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tlq *TestListQuery) Unique(unique bool) *TestListQuery {
	tlq.unique = &unique
	return tlq
}

// Order adds an order step to the query.
func (tlq *TestListQuery) Order(o ...OrderFunc) *TestListQuery {
	tlq.order = append(tlq.order, o...)
	return tlq
}

// First returns the first TestList entity from the query.
// Returns a *NotFoundError when no TestList was found.
func (tlq *TestListQuery) First(ctx context.Context) (*TestList, error) {
	nodes, err := tlq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testlist.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tlq *TestListQuery) FirstX(ctx context.Context) *TestList {
	node, err := tlq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestList ID from the query.
// Returns a *NotFoundError when no TestList ID was found.
func (tlq *TestListQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = tlq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testlist.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tlq *TestListQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := tlq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestList entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestList entity is found.
// Returns a *NotFoundError when no TestList entities are found.
func (tlq *TestListQuery) Only(ctx context.Context) (*TestList, error) {
	nodes, err := tlq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testlist.Label}
	default:
		return nil, &NotSingularError{testlist.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tlq *TestListQuery) OnlyX(ctx context.Context) *TestList {
	node, err := tlq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestList ID in the query.
// Returns a *NotSingularError when more than one TestList ID is found.
// Returns a *NotFoundError when no entities are found.
func (tlq *TestListQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = tlq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testlist.Label}
	default:
		err = &NotSingularError{testlist.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tlq *TestListQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := tlq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestLists.
func (tlq *TestListQuery) All(ctx context.Context) ([]*TestList, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tlq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tlq *TestListQuery) AllX(ctx context.Context) []*TestList {
	nodes, err := tlq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestList IDs.
func (tlq *TestListQuery) IDs(ctx context.Context) ([]xid.ID, error) {
	var ids []xid.ID
	if err := tlq.Select(testlist.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tlq *TestListQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := tlq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tlq *TestListQuery) Count(ctx context.Context) (int, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tlq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tlq *TestListQuery) CountX(ctx context.Context) int {
	count, err := tlq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tlq *TestListQuery) Exist(ctx context.Context) (bool, error) {
	if err := tlq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tlq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tlq *TestListQuery) ExistX(ctx context.Context) bool {
	exist, err := tlq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestListQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tlq *TestListQuery) Clone() *TestListQuery {
	if tlq == nil {
		return nil
	}
	return &TestListQuery{
		config:     tlq.config,
		limit:      tlq.limit,
		offset:     tlq.offset,
		order:      append([]OrderFunc{}, tlq.order...),
		predicates: append([]predicate.TestList{}, tlq.predicates...),
		// clone intermediate query.
		sql:    tlq.sql.Clone(),
		path:   tlq.path,
		unique: tlq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TestList.Query().
//		GroupBy(testlist.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tlq *TestListQuery) GroupBy(field string, fields ...string) *TestListGroupBy {
	grbuild := &TestListGroupBy{config: tlq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tlq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tlq.sqlQuery(ctx), nil
	}
	grbuild.label = testlist.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.TestList.Query().
//		Select(testlist.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (tlq *TestListQuery) Select(fields ...string) *TestListSelect {
	tlq.fields = append(tlq.fields, fields...)
	selbuild := &TestListSelect{TestListQuery: tlq}
	selbuild.label = testlist.Label
	selbuild.flds, selbuild.scan = &tlq.fields, selbuild.Scan
	return selbuild
}

func (tlq *TestListQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tlq.fields {
		if !testlist.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tlq.path != nil {
		prev, err := tlq.path(ctx)
		if err != nil {
			return err
		}
		tlq.sql = prev
	}
	return nil
}

func (tlq *TestListQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestList, error) {
	var (
		nodes = []*TestList{}
		_spec = tlq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TestList).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TestList{config: tlq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tlq.modifiers) > 0 {
		_spec.Modifiers = tlq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tlq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range tlq.loadTotal {
		if err := tlq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tlq *TestListQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tlq.querySpec()
	if len(tlq.modifiers) > 0 {
		_spec.Modifiers = tlq.modifiers
	}
	_spec.Node.Columns = tlq.fields
	if len(tlq.fields) > 0 {
		_spec.Unique = tlq.unique != nil && *tlq.unique
	}
	return sqlgraph.CountNodes(ctx, tlq.driver, _spec)
}

func (tlq *TestListQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tlq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tlq *TestListQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   testlist.Table,
			Columns: testlist.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: testlist.FieldID,
			},
		},
		From:   tlq.sql,
		Unique: true,
	}
	if unique := tlq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tlq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testlist.FieldID)
		for i := range fields {
			if fields[i] != testlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tlq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tlq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tlq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tlq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tlq *TestListQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tlq.driver.Dialect())
	t1 := builder.Table(testlist.Table)
	columns := tlq.fields
	if len(columns) == 0 {
		columns = testlist.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tlq.sql != nil {
		selector = tlq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tlq.unique != nil && *tlq.unique {
		selector.Distinct()
	}
	for _, p := range tlq.predicates {
		p(selector)
	}
	for _, p := range tlq.order {
		p(selector)
	}
	if offset := tlq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tlq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestListGroupBy is the group-by builder for TestList entities.
type TestListGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlgb *TestListGroupBy) Aggregate(fns ...AggregateFunc) *TestListGroupBy {
	tlgb.fns = append(tlgb.fns, fns...)
	return tlgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tlgb *TestListGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tlgb.path(ctx)
	if err != nil {
		return err
	}
	tlgb.sql = query
	return tlgb.sqlScan(ctx, v)
}

func (tlgb *TestListGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tlgb.fields {
		if !testlist.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tlgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tlgb *TestListGroupBy) sqlQuery() *sql.Selector {
	selector := tlgb.sql.Select()
	aggregation := make([]string, 0, len(tlgb.fns))
	for _, fn := range tlgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tlgb.fields)+len(tlgb.fns))
		for _, f := range tlgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tlgb.fields...)...)
}

// TestListSelect is the builder for selecting fields of TestList entities.
type TestListSelect struct {
	*TestListQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tls *TestListSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tls.prepareQuery(ctx); err != nil {
		return err
	}
	tls.sql = tls.TestListQuery.sqlQuery(ctx)
	return tls.sqlScan(ctx, v)
}

func (tls *TestListSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tls.sql.Query()
	if err := tls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
