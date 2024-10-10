// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/ent/predicate"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/ent/userverification"
)

// UserVerificationQuery is the builder for querying UserVerification entities.
type UserVerificationQuery struct {
	config
	ctx        *QueryContext
	order      []userverification.OrderOption
	inters     []Interceptor
	predicates []predicate.UserVerification
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserVerificationQuery builder.
func (uvq *UserVerificationQuery) Where(ps ...predicate.UserVerification) *UserVerificationQuery {
	uvq.predicates = append(uvq.predicates, ps...)
	return uvq
}

// Limit the number of records to be returned by this query.
func (uvq *UserVerificationQuery) Limit(limit int) *UserVerificationQuery {
	uvq.ctx.Limit = &limit
	return uvq
}

// Offset to start from.
func (uvq *UserVerificationQuery) Offset(offset int) *UserVerificationQuery {
	uvq.ctx.Offset = &offset
	return uvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uvq *UserVerificationQuery) Unique(unique bool) *UserVerificationQuery {
	uvq.ctx.Unique = &unique
	return uvq
}

// Order specifies how the records should be ordered.
func (uvq *UserVerificationQuery) Order(o ...userverification.OrderOption) *UserVerificationQuery {
	uvq.order = append(uvq.order, o...)
	return uvq
}

// QueryUser chains the current query on the "user" edge.
func (uvq *UserVerificationQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: uvq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uvq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userverification.Table, userverification.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userverification.UserTable, userverification.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(uvq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserVerification entity from the query.
// Returns a *NotFoundError when no UserVerification was found.
func (uvq *UserVerificationQuery) First(ctx context.Context) (*UserVerification, error) {
	nodes, err := uvq.Limit(1).All(setContextOp(ctx, uvq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userverification.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uvq *UserVerificationQuery) FirstX(ctx context.Context) *UserVerification {
	node, err := uvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserVerification ID from the query.
// Returns a *NotFoundError when no UserVerification ID was found.
func (uvq *UserVerificationQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uvq.Limit(1).IDs(setContextOp(ctx, uvq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userverification.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uvq *UserVerificationQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserVerification entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserVerification entity is found.
// Returns a *NotFoundError when no UserVerification entities are found.
func (uvq *UserVerificationQuery) Only(ctx context.Context) (*UserVerification, error) {
	nodes, err := uvq.Limit(2).All(setContextOp(ctx, uvq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userverification.Label}
	default:
		return nil, &NotSingularError{userverification.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uvq *UserVerificationQuery) OnlyX(ctx context.Context) *UserVerification {
	node, err := uvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserVerification ID in the query.
// Returns a *NotSingularError when more than one UserVerification ID is found.
// Returns a *NotFoundError when no entities are found.
func (uvq *UserVerificationQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uvq.Limit(2).IDs(setContextOp(ctx, uvq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userverification.Label}
	default:
		err = &NotSingularError{userverification.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uvq *UserVerificationQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserVerifications.
func (uvq *UserVerificationQuery) All(ctx context.Context) ([]*UserVerification, error) {
	ctx = setContextOp(ctx, uvq.ctx, ent.OpQueryAll)
	if err := uvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserVerification, *UserVerificationQuery]()
	return withInterceptors[[]*UserVerification](ctx, uvq, qr, uvq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uvq *UserVerificationQuery) AllX(ctx context.Context) []*UserVerification {
	nodes, err := uvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserVerification IDs.
func (uvq *UserVerificationQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if uvq.ctx.Unique == nil && uvq.path != nil {
		uvq.Unique(true)
	}
	ctx = setContextOp(ctx, uvq.ctx, ent.OpQueryIDs)
	if err = uvq.Select(userverification.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uvq *UserVerificationQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uvq *UserVerificationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uvq.ctx, ent.OpQueryCount)
	if err := uvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uvq, querierCount[*UserVerificationQuery](), uvq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uvq *UserVerificationQuery) CountX(ctx context.Context) int {
	count, err := uvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uvq *UserVerificationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uvq.ctx, ent.OpQueryExist)
	switch _, err := uvq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uvq *UserVerificationQuery) ExistX(ctx context.Context) bool {
	exist, err := uvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserVerificationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uvq *UserVerificationQuery) Clone() *UserVerificationQuery {
	if uvq == nil {
		return nil
	}
	return &UserVerificationQuery{
		config:     uvq.config,
		ctx:        uvq.ctx.Clone(),
		order:      append([]userverification.OrderOption{}, uvq.order...),
		inters:     append([]Interceptor{}, uvq.inters...),
		predicates: append([]predicate.UserVerification{}, uvq.predicates...),
		withUser:   uvq.withUser.Clone(),
		// clone intermediate query.
		sql:  uvq.sql.Clone(),
		path: uvq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (uvq *UserVerificationQuery) WithUser(opts ...func(*UserQuery)) *UserVerificationQuery {
	query := (&UserClient{config: uvq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uvq.withUser = query
	return uvq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AttemptID uuid.UUID `json:"attempt_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserVerification.Query().
//		GroupBy(userverification.FieldAttemptID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uvq *UserVerificationQuery) GroupBy(field string, fields ...string) *UserVerificationGroupBy {
	uvq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserVerificationGroupBy{build: uvq}
	grbuild.flds = &uvq.ctx.Fields
	grbuild.label = userverification.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AttemptID uuid.UUID `json:"attempt_id,omitempty"`
//	}
//
//	client.UserVerification.Query().
//		Select(userverification.FieldAttemptID).
//		Scan(ctx, &v)
func (uvq *UserVerificationQuery) Select(fields ...string) *UserVerificationSelect {
	uvq.ctx.Fields = append(uvq.ctx.Fields, fields...)
	sbuild := &UserVerificationSelect{UserVerificationQuery: uvq}
	sbuild.label = userverification.Label
	sbuild.flds, sbuild.scan = &uvq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserVerificationSelect configured with the given aggregations.
func (uvq *UserVerificationQuery) Aggregate(fns ...AggregateFunc) *UserVerificationSelect {
	return uvq.Select().Aggregate(fns...)
}

func (uvq *UserVerificationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uvq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uvq); err != nil {
				return err
			}
		}
	}
	for _, f := range uvq.ctx.Fields {
		if !userverification.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uvq.path != nil {
		prev, err := uvq.path(ctx)
		if err != nil {
			return err
		}
		uvq.sql = prev
	}
	return nil
}

func (uvq *UserVerificationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserVerification, error) {
	var (
		nodes       = []*UserVerification{}
		withFKs     = uvq.withFKs
		_spec       = uvq.querySpec()
		loadedTypes = [1]bool{
			uvq.withUser != nil,
		}
	)
	if uvq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, userverification.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserVerification).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserVerification{config: uvq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uvq.withUser; query != nil {
		if err := uvq.loadUser(ctx, query, nodes, nil,
			func(n *UserVerification, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uvq *UserVerificationQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserVerification, init func(*UserVerification), assign func(*UserVerification, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*UserVerification)
	for i := range nodes {
		if nodes[i].user_user_verifications == nil {
			continue
		}
		fk := *nodes[i].user_user_verifications
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_user_verifications" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (uvq *UserVerificationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uvq.querySpec()
	_spec.Node.Columns = uvq.ctx.Fields
	if len(uvq.ctx.Fields) > 0 {
		_spec.Unique = uvq.ctx.Unique != nil && *uvq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uvq.driver, _spec)
}

func (uvq *UserVerificationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userverification.Table, userverification.Columns, sqlgraph.NewFieldSpec(userverification.FieldID, field.TypeUUID))
	_spec.From = uvq.sql
	if unique := uvq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uvq.path != nil {
		_spec.Unique = true
	}
	if fields := uvq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userverification.FieldID)
		for i := range fields {
			if fields[i] != userverification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uvq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uvq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uvq *UserVerificationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uvq.driver.Dialect())
	t1 := builder.Table(userverification.Table)
	columns := uvq.ctx.Fields
	if len(columns) == 0 {
		columns = userverification.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uvq.sql != nil {
		selector = uvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uvq.ctx.Unique != nil && *uvq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uvq.predicates {
		p(selector)
	}
	for _, p := range uvq.order {
		p(selector)
	}
	if offset := uvq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uvq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserVerificationGroupBy is the group-by builder for UserVerification entities.
type UserVerificationGroupBy struct {
	selector
	build *UserVerificationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uvgb *UserVerificationGroupBy) Aggregate(fns ...AggregateFunc) *UserVerificationGroupBy {
	uvgb.fns = append(uvgb.fns, fns...)
	return uvgb
}

// Scan applies the selector query and scans the result into the given value.
func (uvgb *UserVerificationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uvgb.build.ctx, ent.OpQueryGroupBy)
	if err := uvgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserVerificationQuery, *UserVerificationGroupBy](ctx, uvgb.build, uvgb, uvgb.build.inters, v)
}

func (uvgb *UserVerificationGroupBy) sqlScan(ctx context.Context, root *UserVerificationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uvgb.fns))
	for _, fn := range uvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uvgb.flds)+len(uvgb.fns))
		for _, f := range *uvgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uvgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserVerificationSelect is the builder for selecting fields of UserVerification entities.
type UserVerificationSelect struct {
	*UserVerificationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uvs *UserVerificationSelect) Aggregate(fns ...AggregateFunc) *UserVerificationSelect {
	uvs.fns = append(uvs.fns, fns...)
	return uvs
}

// Scan applies the selector query and scans the result into the given value.
func (uvs *UserVerificationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uvs.ctx, ent.OpQuerySelect)
	if err := uvs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserVerificationQuery, *UserVerificationSelect](ctx, uvs.UserVerificationQuery, uvs, uvs.inters, v)
}

func (uvs *UserVerificationSelect) sqlScan(ctx context.Context, root *UserVerificationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uvs.fns))
	for _, fn := range uvs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uvs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
