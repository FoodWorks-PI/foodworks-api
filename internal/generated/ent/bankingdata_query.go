// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"foodworks.ml/m/internal/generated/ent/bankingdata"
	"foodworks.ml/m/internal/generated/ent/predicate"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// BankingDataQuery is the builder for querying BankingData entities.
type BankingDataQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.BankingData
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (bdq *BankingDataQuery) Where(ps ...predicate.BankingData) *BankingDataQuery {
	bdq.predicates = append(bdq.predicates, ps...)
	return bdq
}

// Limit adds a limit step to the query.
func (bdq *BankingDataQuery) Limit(limit int) *BankingDataQuery {
	bdq.limit = &limit
	return bdq
}

// Offset adds an offset step to the query.
func (bdq *BankingDataQuery) Offset(offset int) *BankingDataQuery {
	bdq.offset = &offset
	return bdq
}

// Order adds an order step to the query.
func (bdq *BankingDataQuery) Order(o ...OrderFunc) *BankingDataQuery {
	bdq.order = append(bdq.order, o...)
	return bdq
}

// First returns the first BankingData entity in the query. Returns *NotFoundError when no bankingdata was found.
func (bdq *BankingDataQuery) First(ctx context.Context) (*BankingData, error) {
	nodes, err := bdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bankingdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bdq *BankingDataQuery) FirstX(ctx context.Context) *BankingData {
	node, err := bdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BankingData id in the query. Returns *NotFoundError when no id was found.
func (bdq *BankingDataQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bankingdata.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bdq *BankingDataQuery) FirstXID(ctx context.Context) int {
	id, err := bdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only BankingData entity in the query, returns an error if not exactly one entity was returned.
func (bdq *BankingDataQuery) Only(ctx context.Context) (*BankingData, error) {
	nodes, err := bdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bankingdata.Label}
	default:
		return nil, &NotSingularError{bankingdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bdq *BankingDataQuery) OnlyX(ctx context.Context) *BankingData {
	node, err := bdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID returns the only BankingData id in the query, returns an error if not exactly one id was returned.
func (bdq *BankingDataQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = &NotSingularError{bankingdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bdq *BankingDataQuery) OnlyIDX(ctx context.Context) int {
	id, err := bdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BankingDataSlice.
func (bdq *BankingDataQuery) All(ctx context.Context) ([]*BankingData, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bdq *BankingDataQuery) AllX(ctx context.Context) []*BankingData {
	nodes, err := bdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BankingData ids.
func (bdq *BankingDataQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bdq.Select(bankingdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bdq *BankingDataQuery) IDsX(ctx context.Context) []int {
	ids, err := bdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bdq *BankingDataQuery) Count(ctx context.Context) (int, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bdq *BankingDataQuery) CountX(ctx context.Context) int {
	count, err := bdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bdq *BankingDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := bdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bdq *BankingDataQuery) ExistX(ctx context.Context) bool {
	exist, err := bdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bdq *BankingDataQuery) Clone() *BankingDataQuery {
	return &BankingDataQuery{
		config:     bdq.config,
		limit:      bdq.limit,
		offset:     bdq.offset,
		order:      append([]OrderFunc{}, bdq.order...),
		unique:     append([]string{}, bdq.unique...),
		predicates: append([]predicate.BankingData{}, bdq.predicates...),
		// clone intermediate query.
		sql:  bdq.sql.Clone(),
		path: bdq.path,
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BankAccount string `json:"bank_account,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BankingData.Query().
//		GroupBy(bankingdata.FieldBankAccount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bdq *BankingDataQuery) GroupBy(field string, fields ...string) *BankingDataGroupBy {
	group := &BankingDataGroupBy{config: bdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bdq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		BankAccount string `json:"bank_account,omitempty"`
//	}
//
//	client.BankingData.Query().
//		Select(bankingdata.FieldBankAccount).
//		Scan(ctx, &v)
//
func (bdq *BankingDataQuery) Select(field string, fields ...string) *BankingDataSelect {
	selector := &BankingDataSelect{config: bdq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bdq.sqlQuery(), nil
	}
	return selector
}

func (bdq *BankingDataQuery) prepareQuery(ctx context.Context) error {
	if bdq.path != nil {
		prev, err := bdq.path(ctx)
		if err != nil {
			return err
		}
		bdq.sql = prev
	}
	return nil
}

func (bdq *BankingDataQuery) sqlAll(ctx context.Context) ([]*BankingData, error) {
	var (
		nodes = []*BankingData{}
		_spec = bdq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &BankingData{config: bdq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, bdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bdq *BankingDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bdq.querySpec()
	return sqlgraph.CountNodes(ctx, bdq.driver, _spec)
}

func (bdq *BankingDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bdq *BankingDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bankingdata.Table,
			Columns: bankingdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bankingdata.FieldID,
			},
		},
		From:   bdq.sql,
		Unique: true,
	}
	if ps := bdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, bankingdata.ValidColumn)
			}
		}
	}
	return _spec
}

func (bdq *BankingDataQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bdq.driver.Dialect())
	t1 := builder.Table(bankingdata.Table)
	selector := builder.Select(t1.Columns(bankingdata.Columns...)...).From(t1)
	if bdq.sql != nil {
		selector = bdq.sql
		selector.Select(selector.Columns(bankingdata.Columns...)...)
	}
	for _, p := range bdq.predicates {
		p(selector)
	}
	for _, p := range bdq.order {
		p(selector, bankingdata.ValidColumn)
	}
	if offset := bdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BankingDataGroupBy is the builder for group-by BankingData entities.
type BankingDataGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bdgb *BankingDataGroupBy) Aggregate(fns ...AggregateFunc) *BankingDataGroupBy {
	bdgb.fns = append(bdgb.fns, fns...)
	return bdgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bdgb *BankingDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bdgb.path(ctx)
	if err != nil {
		return err
	}
	bdgb.sql = query
	return bdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bdgb.fields) > 1 {
		return nil, errors.New("ent: BankingDataGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) StringsX(ctx context.Context) []string {
	v, err := bdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) StringX(ctx context.Context) string {
	v, err := bdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bdgb.fields) > 1 {
		return nil, errors.New("ent: BankingDataGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) IntsX(ctx context.Context) []int {
	v, err := bdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) IntX(ctx context.Context) int {
	v, err := bdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bdgb.fields) > 1 {
		return nil, errors.New("ent: BankingDataGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bdgb.fields) > 1 {
		return nil, errors.New("ent: BankingDataGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (bdgb *BankingDataGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bdgb *BankingDataGroupBy) BoolX(ctx context.Context) bool {
	v, err := bdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bdgb *BankingDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bdgb.fields {
		if !bankingdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bdgb *BankingDataGroupBy) sqlQuery() *sql.Selector {
	selector := bdgb.sql
	columns := make([]string, 0, len(bdgb.fields)+len(bdgb.fns))
	columns = append(columns, bdgb.fields...)
	for _, fn := range bdgb.fns {
		columns = append(columns, fn(selector, bankingdata.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(bdgb.fields...)
}

// BankingDataSelect is the builder for select fields of BankingData entities.
type BankingDataSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (bds *BankingDataSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := bds.path(ctx)
	if err != nil {
		return err
	}
	bds.sql = query
	return bds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bds *BankingDataSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bds.fields) > 1 {
		return nil, errors.New("ent: BankingDataSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bds *BankingDataSelect) StringsX(ctx context.Context) []string {
	v, err := bds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bds *BankingDataSelect) StringX(ctx context.Context) string {
	v, err := bds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bds.fields) > 1 {
		return nil, errors.New("ent: BankingDataSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bds *BankingDataSelect) IntsX(ctx context.Context) []int {
	v, err := bds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bds *BankingDataSelect) IntX(ctx context.Context) int {
	v, err := bds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bds.fields) > 1 {
		return nil, errors.New("ent: BankingDataSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bds *BankingDataSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bds *BankingDataSelect) Float64X(ctx context.Context) float64 {
	v, err := bds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bds.fields) > 1 {
		return nil, errors.New("ent: BankingDataSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bds *BankingDataSelect) BoolsX(ctx context.Context) []bool {
	v, err := bds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (bds *BankingDataSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bankingdata.Label}
	default:
		err = fmt.Errorf("ent: BankingDataSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bds *BankingDataSelect) BoolX(ctx context.Context) bool {
	v, err := bds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bds *BankingDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bds.fields {
		if !bankingdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for selection", f)}
		}
	}
	rows := &sql.Rows{}
	query, args := bds.sqlQuery().Query()
	if err := bds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bds *BankingDataSelect) sqlQuery() sql.Querier {
	selector := bds.sql
	selector.Select(selector.Columns(bds.fields...)...)
	return selector
}
