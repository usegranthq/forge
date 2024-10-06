// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/usegranthq/backend/ent/oidcclient"
	"github.com/usegranthq/backend/ent/predicate"
	"github.com/usegranthq/backend/ent/project"
)

// OidcClientUpdate is the builder for updating OidcClient entities.
type OidcClientUpdate struct {
	config
	hooks    []Hook
	mutation *OidcClientMutation
}

// Where appends a list predicates to the OidcClientUpdate builder.
func (ocu *OidcClientUpdate) Where(ps ...predicate.OidcClient) *OidcClientUpdate {
	ocu.mutation.Where(ps...)
	return ocu
}

// SetName sets the "name" field.
func (ocu *OidcClientUpdate) SetName(s string) *OidcClientUpdate {
	ocu.mutation.SetName(s)
	return ocu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ocu *OidcClientUpdate) SetNillableName(s *string) *OidcClientUpdate {
	if s != nil {
		ocu.SetName(*s)
	}
	return ocu
}

// SetClientID sets the "client_id" field.
func (ocu *OidcClientUpdate) SetClientID(s string) *OidcClientUpdate {
	ocu.mutation.SetClientID(s)
	return ocu
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (ocu *OidcClientUpdate) SetNillableClientID(s *string) *OidcClientUpdate {
	if s != nil {
		ocu.SetClientID(*s)
	}
	return ocu
}

// SetClientSecret sets the "client_secret" field.
func (ocu *OidcClientUpdate) SetClientSecret(s string) *OidcClientUpdate {
	ocu.mutation.SetClientSecret(s)
	return ocu
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (ocu *OidcClientUpdate) SetNillableClientSecret(s *string) *OidcClientUpdate {
	if s != nil {
		ocu.SetClientSecret(*s)
	}
	return ocu
}

// SetUpdatedAt sets the "updated_at" field.
func (ocu *OidcClientUpdate) SetUpdatedAt(t time.Time) *OidcClientUpdate {
	ocu.mutation.SetUpdatedAt(t)
	return ocu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ocu *OidcClientUpdate) SetProjectID(id uuid.UUID) *OidcClientUpdate {
	ocu.mutation.SetProjectID(id)
	return ocu
}

// SetProject sets the "project" edge to the Project entity.
func (ocu *OidcClientUpdate) SetProject(p *Project) *OidcClientUpdate {
	return ocu.SetProjectID(p.ID)
}

// Mutation returns the OidcClientMutation object of the builder.
func (ocu *OidcClientUpdate) Mutation() *OidcClientMutation {
	return ocu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ocu *OidcClientUpdate) ClearProject() *OidcClientUpdate {
	ocu.mutation.ClearProject()
	return ocu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocu *OidcClientUpdate) Save(ctx context.Context) (int, error) {
	ocu.defaults()
	return withHooks(ctx, ocu.sqlSave, ocu.mutation, ocu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ocu *OidcClientUpdate) SaveX(ctx context.Context) int {
	affected, err := ocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocu *OidcClientUpdate) Exec(ctx context.Context) error {
	_, err := ocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocu *OidcClientUpdate) ExecX(ctx context.Context) {
	if err := ocu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ocu *OidcClientUpdate) defaults() {
	if _, ok := ocu.mutation.UpdatedAt(); !ok {
		v := oidcclient.UpdateDefaultUpdatedAt()
		ocu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ocu *OidcClientUpdate) check() error {
	if v, ok := ocu.mutation.Name(); ok {
		if err := oidcclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OidcClient.name": %w`, err)}
		}
	}
	if v, ok := ocu.mutation.ClientID(); ok {
		if err := oidcclient.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "OidcClient.client_id": %w`, err)}
		}
	}
	if v, ok := ocu.mutation.ClientSecret(); ok {
		if err := oidcclient.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`ent: validator failed for field "OidcClient.client_secret": %w`, err)}
		}
	}
	if ocu.mutation.ProjectCleared() && len(ocu.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OidcClient.project"`)
	}
	return nil
}

func (ocu *OidcClientUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ocu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(oidcclient.Table, oidcclient.Columns, sqlgraph.NewFieldSpec(oidcclient.FieldID, field.TypeUUID))
	if ps := ocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocu.mutation.Name(); ok {
		_spec.SetField(oidcclient.FieldName, field.TypeString, value)
	}
	if value, ok := ocu.mutation.ClientID(); ok {
		_spec.SetField(oidcclient.FieldClientID, field.TypeString, value)
	}
	if value, ok := ocu.mutation.ClientSecret(); ok {
		_spec.SetField(oidcclient.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := ocu.mutation.UpdatedAt(); ok {
		_spec.SetField(oidcclient.FieldUpdatedAt, field.TypeTime, value)
	}
	if ocu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oidcclient.ProjectTable,
			Columns: []string{oidcclient.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oidcclient.ProjectTable,
			Columns: []string{oidcclient.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oidcclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ocu.mutation.done = true
	return n, nil
}

// OidcClientUpdateOne is the builder for updating a single OidcClient entity.
type OidcClientUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OidcClientMutation
}

// SetName sets the "name" field.
func (ocuo *OidcClientUpdateOne) SetName(s string) *OidcClientUpdateOne {
	ocuo.mutation.SetName(s)
	return ocuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ocuo *OidcClientUpdateOne) SetNillableName(s *string) *OidcClientUpdateOne {
	if s != nil {
		ocuo.SetName(*s)
	}
	return ocuo
}

// SetClientID sets the "client_id" field.
func (ocuo *OidcClientUpdateOne) SetClientID(s string) *OidcClientUpdateOne {
	ocuo.mutation.SetClientID(s)
	return ocuo
}

// SetNillableClientID sets the "client_id" field if the given value is not nil.
func (ocuo *OidcClientUpdateOne) SetNillableClientID(s *string) *OidcClientUpdateOne {
	if s != nil {
		ocuo.SetClientID(*s)
	}
	return ocuo
}

// SetClientSecret sets the "client_secret" field.
func (ocuo *OidcClientUpdateOne) SetClientSecret(s string) *OidcClientUpdateOne {
	ocuo.mutation.SetClientSecret(s)
	return ocuo
}

// SetNillableClientSecret sets the "client_secret" field if the given value is not nil.
func (ocuo *OidcClientUpdateOne) SetNillableClientSecret(s *string) *OidcClientUpdateOne {
	if s != nil {
		ocuo.SetClientSecret(*s)
	}
	return ocuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ocuo *OidcClientUpdateOne) SetUpdatedAt(t time.Time) *OidcClientUpdateOne {
	ocuo.mutation.SetUpdatedAt(t)
	return ocuo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (ocuo *OidcClientUpdateOne) SetProjectID(id uuid.UUID) *OidcClientUpdateOne {
	ocuo.mutation.SetProjectID(id)
	return ocuo
}

// SetProject sets the "project" edge to the Project entity.
func (ocuo *OidcClientUpdateOne) SetProject(p *Project) *OidcClientUpdateOne {
	return ocuo.SetProjectID(p.ID)
}

// Mutation returns the OidcClientMutation object of the builder.
func (ocuo *OidcClientUpdateOne) Mutation() *OidcClientMutation {
	return ocuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (ocuo *OidcClientUpdateOne) ClearProject() *OidcClientUpdateOne {
	ocuo.mutation.ClearProject()
	return ocuo
}

// Where appends a list predicates to the OidcClientUpdate builder.
func (ocuo *OidcClientUpdateOne) Where(ps ...predicate.OidcClient) *OidcClientUpdateOne {
	ocuo.mutation.Where(ps...)
	return ocuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocuo *OidcClientUpdateOne) Select(field string, fields ...string) *OidcClientUpdateOne {
	ocuo.fields = append([]string{field}, fields...)
	return ocuo
}

// Save executes the query and returns the updated OidcClient entity.
func (ocuo *OidcClientUpdateOne) Save(ctx context.Context) (*OidcClient, error) {
	ocuo.defaults()
	return withHooks(ctx, ocuo.sqlSave, ocuo.mutation, ocuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ocuo *OidcClientUpdateOne) SaveX(ctx context.Context) *OidcClient {
	node, err := ocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocuo *OidcClientUpdateOne) Exec(ctx context.Context) error {
	_, err := ocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocuo *OidcClientUpdateOne) ExecX(ctx context.Context) {
	if err := ocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ocuo *OidcClientUpdateOne) defaults() {
	if _, ok := ocuo.mutation.UpdatedAt(); !ok {
		v := oidcclient.UpdateDefaultUpdatedAt()
		ocuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ocuo *OidcClientUpdateOne) check() error {
	if v, ok := ocuo.mutation.Name(); ok {
		if err := oidcclient.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "OidcClient.name": %w`, err)}
		}
	}
	if v, ok := ocuo.mutation.ClientID(); ok {
		if err := oidcclient.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`ent: validator failed for field "OidcClient.client_id": %w`, err)}
		}
	}
	if v, ok := ocuo.mutation.ClientSecret(); ok {
		if err := oidcclient.ClientSecretValidator(v); err != nil {
			return &ValidationError{Name: "client_secret", err: fmt.Errorf(`ent: validator failed for field "OidcClient.client_secret": %w`, err)}
		}
	}
	if ocuo.mutation.ProjectCleared() && len(ocuo.mutation.ProjectIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "OidcClient.project"`)
	}
	return nil
}

func (ocuo *OidcClientUpdateOne) sqlSave(ctx context.Context) (_node *OidcClient, err error) {
	if err := ocuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(oidcclient.Table, oidcclient.Columns, sqlgraph.NewFieldSpec(oidcclient.FieldID, field.TypeUUID))
	id, ok := ocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OidcClient.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oidcclient.FieldID)
		for _, f := range fields {
			if !oidcclient.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != oidcclient.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocuo.mutation.Name(); ok {
		_spec.SetField(oidcclient.FieldName, field.TypeString, value)
	}
	if value, ok := ocuo.mutation.ClientID(); ok {
		_spec.SetField(oidcclient.FieldClientID, field.TypeString, value)
	}
	if value, ok := ocuo.mutation.ClientSecret(); ok {
		_spec.SetField(oidcclient.FieldClientSecret, field.TypeString, value)
	}
	if value, ok := ocuo.mutation.UpdatedAt(); ok {
		_spec.SetField(oidcclient.FieldUpdatedAt, field.TypeTime, value)
	}
	if ocuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oidcclient.ProjectTable,
			Columns: []string{oidcclient.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ocuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   oidcclient.ProjectTable,
			Columns: []string{oidcclient.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OidcClient{config: ocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{oidcclient.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ocuo.mutation.done = true
	return _node, nil
}
