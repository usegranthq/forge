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
	"github.com/usegranthq/forge/ent/oidcclient"
	"github.com/usegranthq/forge/ent/predicate"
	"github.com/usegranthq/forge/ent/project"
)

// OidcClientQuery is the builder for querying OidcClient entities.
type OidcClientQuery struct {
	config
	ctx         *QueryContext
	order       []oidcclient.OrderOption
	inters      []Interceptor
	predicates  []predicate.OidcClient
	withProject *ProjectQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OidcClientQuery builder.
func (ocq *OidcClientQuery) Where(ps ...predicate.OidcClient) *OidcClientQuery {
	ocq.predicates = append(ocq.predicates, ps...)
	return ocq
}

// Limit the number of records to be returned by this query.
func (ocq *OidcClientQuery) Limit(limit int) *OidcClientQuery {
	ocq.ctx.Limit = &limit
	return ocq
}

// Offset to start from.
func (ocq *OidcClientQuery) Offset(offset int) *OidcClientQuery {
	ocq.ctx.Offset = &offset
	return ocq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ocq *OidcClientQuery) Unique(unique bool) *OidcClientQuery {
	ocq.ctx.Unique = &unique
	return ocq
}

// Order specifies how the records should be ordered.
func (ocq *OidcClientQuery) Order(o ...oidcclient.OrderOption) *OidcClientQuery {
	ocq.order = append(ocq.order, o...)
	return ocq
}

// QueryProject chains the current query on the "project" edge.
func (ocq *OidcClientQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: ocq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ocq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ocq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(oidcclient.Table, oidcclient.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, oidcclient.ProjectTable, oidcclient.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(ocq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OidcClient entity from the query.
// Returns a *NotFoundError when no OidcClient was found.
func (ocq *OidcClientQuery) First(ctx context.Context) (*OidcClient, error) {
	nodes, err := ocq.Limit(1).All(setContextOp(ctx, ocq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oidcclient.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ocq *OidcClientQuery) FirstX(ctx context.Context) *OidcClient {
	node, err := ocq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OidcClient ID from the query.
// Returns a *NotFoundError when no OidcClient ID was found.
func (ocq *OidcClientQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ocq.Limit(1).IDs(setContextOp(ctx, ocq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oidcclient.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ocq *OidcClientQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ocq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OidcClient entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OidcClient entity is found.
// Returns a *NotFoundError when no OidcClient entities are found.
func (ocq *OidcClientQuery) Only(ctx context.Context) (*OidcClient, error) {
	nodes, err := ocq.Limit(2).All(setContextOp(ctx, ocq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oidcclient.Label}
	default:
		return nil, &NotSingularError{oidcclient.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ocq *OidcClientQuery) OnlyX(ctx context.Context) *OidcClient {
	node, err := ocq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OidcClient ID in the query.
// Returns a *NotSingularError when more than one OidcClient ID is found.
// Returns a *NotFoundError when no entities are found.
func (ocq *OidcClientQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ocq.Limit(2).IDs(setContextOp(ctx, ocq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oidcclient.Label}
	default:
		err = &NotSingularError{oidcclient.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ocq *OidcClientQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ocq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OidcClients.
func (ocq *OidcClientQuery) All(ctx context.Context) ([]*OidcClient, error) {
	ctx = setContextOp(ctx, ocq.ctx, ent.OpQueryAll)
	if err := ocq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OidcClient, *OidcClientQuery]()
	return withInterceptors[[]*OidcClient](ctx, ocq, qr, ocq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ocq *OidcClientQuery) AllX(ctx context.Context) []*OidcClient {
	nodes, err := ocq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OidcClient IDs.
func (ocq *OidcClientQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ocq.ctx.Unique == nil && ocq.path != nil {
		ocq.Unique(true)
	}
	ctx = setContextOp(ctx, ocq.ctx, ent.OpQueryIDs)
	if err = ocq.Select(oidcclient.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ocq *OidcClientQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ocq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ocq *OidcClientQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ocq.ctx, ent.OpQueryCount)
	if err := ocq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ocq, querierCount[*OidcClientQuery](), ocq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ocq *OidcClientQuery) CountX(ctx context.Context) int {
	count, err := ocq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ocq *OidcClientQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ocq.ctx, ent.OpQueryExist)
	switch _, err := ocq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ocq *OidcClientQuery) ExistX(ctx context.Context) bool {
	exist, err := ocq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OidcClientQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ocq *OidcClientQuery) Clone() *OidcClientQuery {
	if ocq == nil {
		return nil
	}
	return &OidcClientQuery{
		config:      ocq.config,
		ctx:         ocq.ctx.Clone(),
		order:       append([]oidcclient.OrderOption{}, ocq.order...),
		inters:      append([]Interceptor{}, ocq.inters...),
		predicates:  append([]predicate.OidcClient{}, ocq.predicates...),
		withProject: ocq.withProject.Clone(),
		// clone intermediate query.
		sql:  ocq.sql.Clone(),
		path: ocq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (ocq *OidcClientQuery) WithProject(opts ...func(*ProjectQuery)) *OidcClientQuery {
	query := (&ProjectClient{config: ocq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ocq.withProject = query
	return ocq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OidcClient.Query().
//		GroupBy(oidcclient.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ocq *OidcClientQuery) GroupBy(field string, fields ...string) *OidcClientGroupBy {
	ocq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OidcClientGroupBy{build: ocq}
	grbuild.flds = &ocq.ctx.Fields
	grbuild.label = oidcclient.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.OidcClient.Query().
//		Select(oidcclient.FieldName).
//		Scan(ctx, &v)
func (ocq *OidcClientQuery) Select(fields ...string) *OidcClientSelect {
	ocq.ctx.Fields = append(ocq.ctx.Fields, fields...)
	sbuild := &OidcClientSelect{OidcClientQuery: ocq}
	sbuild.label = oidcclient.Label
	sbuild.flds, sbuild.scan = &ocq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OidcClientSelect configured with the given aggregations.
func (ocq *OidcClientQuery) Aggregate(fns ...AggregateFunc) *OidcClientSelect {
	return ocq.Select().Aggregate(fns...)
}

func (ocq *OidcClientQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ocq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ocq); err != nil {
				return err
			}
		}
	}
	for _, f := range ocq.ctx.Fields {
		if !oidcclient.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ocq.path != nil {
		prev, err := ocq.path(ctx)
		if err != nil {
			return err
		}
		ocq.sql = prev
	}
	return nil
}

func (ocq *OidcClientQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OidcClient, error) {
	var (
		nodes       = []*OidcClient{}
		withFKs     = ocq.withFKs
		_spec       = ocq.querySpec()
		loadedTypes = [1]bool{
			ocq.withProject != nil,
		}
	)
	if ocq.withProject != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, oidcclient.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OidcClient).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OidcClient{config: ocq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ocq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ocq.withProject; query != nil {
		if err := ocq.loadProject(ctx, query, nodes, nil,
			func(n *OidcClient, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ocq *OidcClientQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*OidcClient, init func(*OidcClient), assign func(*OidcClient, *Project)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*OidcClient)
	for i := range nodes {
		if nodes[i].project_oidc_clients == nil {
			continue
		}
		fk := *nodes[i].project_oidc_clients
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_oidc_clients" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ocq *OidcClientQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ocq.querySpec()
	_spec.Node.Columns = ocq.ctx.Fields
	if len(ocq.ctx.Fields) > 0 {
		_spec.Unique = ocq.ctx.Unique != nil && *ocq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ocq.driver, _spec)
}

func (ocq *OidcClientQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oidcclient.Table, oidcclient.Columns, sqlgraph.NewFieldSpec(oidcclient.FieldID, field.TypeUUID))
	_spec.From = ocq.sql
	if unique := ocq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ocq.path != nil {
		_spec.Unique = true
	}
	if fields := ocq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oidcclient.FieldID)
		for i := range fields {
			if fields[i] != oidcclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ocq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ocq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ocq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ocq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ocq *OidcClientQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ocq.driver.Dialect())
	t1 := builder.Table(oidcclient.Table)
	columns := ocq.ctx.Fields
	if len(columns) == 0 {
		columns = oidcclient.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ocq.sql != nil {
		selector = ocq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ocq.ctx.Unique != nil && *ocq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ocq.predicates {
		p(selector)
	}
	for _, p := range ocq.order {
		p(selector)
	}
	if offset := ocq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ocq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OidcClientGroupBy is the group-by builder for OidcClient entities.
type OidcClientGroupBy struct {
	selector
	build *OidcClientQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ocgb *OidcClientGroupBy) Aggregate(fns ...AggregateFunc) *OidcClientGroupBy {
	ocgb.fns = append(ocgb.fns, fns...)
	return ocgb
}

// Scan applies the selector query and scans the result into the given value.
func (ocgb *OidcClientGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ocgb.build.ctx, ent.OpQueryGroupBy)
	if err := ocgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OidcClientQuery, *OidcClientGroupBy](ctx, ocgb.build, ocgb, ocgb.build.inters, v)
}

func (ocgb *OidcClientGroupBy) sqlScan(ctx context.Context, root *OidcClientQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ocgb.fns))
	for _, fn := range ocgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ocgb.flds)+len(ocgb.fns))
		for _, f := range *ocgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ocgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OidcClientSelect is the builder for selecting fields of OidcClient entities.
type OidcClientSelect struct {
	*OidcClientQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ocs *OidcClientSelect) Aggregate(fns ...AggregateFunc) *OidcClientSelect {
	ocs.fns = append(ocs.fns, fns...)
	return ocs
}

// Scan applies the selector query and scans the result into the given value.
func (ocs *OidcClientSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ocs.ctx, ent.OpQuerySelect)
	if err := ocs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OidcClientQuery, *OidcClientSelect](ctx, ocs.OidcClientQuery, ocs, ocs.inters, v)
}

func (ocs *OidcClientSelect) sqlScan(ctx context.Context, root *OidcClientQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ocs.fns))
	for _, fn := range ocs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ocs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ocs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
