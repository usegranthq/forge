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
	"github.com/usegranthq/backend/ent/predicate"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/ent/userverification"
)

// UserVerificationUpdate is the builder for updating UserVerification entities.
type UserVerificationUpdate struct {
	config
	hooks    []Hook
	mutation *UserVerificationMutation
}

// Where appends a list predicates to the UserVerificationUpdate builder.
func (uvu *UserVerificationUpdate) Where(ps ...predicate.UserVerification) *UserVerificationUpdate {
	uvu.mutation.Where(ps...)
	return uvu
}

// SetCode sets the "code" field.
func (uvu *UserVerificationUpdate) SetCode(s string) *UserVerificationUpdate {
	uvu.mutation.SetCode(s)
	return uvu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (uvu *UserVerificationUpdate) SetNillableCode(s *string) *UserVerificationUpdate {
	if s != nil {
		uvu.SetCode(*s)
	}
	return uvu
}

// SetAttempts sets the "attempts" field.
func (uvu *UserVerificationUpdate) SetAttempts(i int) *UserVerificationUpdate {
	uvu.mutation.ResetAttempts()
	uvu.mutation.SetAttempts(i)
	return uvu
}

// SetNillableAttempts sets the "attempts" field if the given value is not nil.
func (uvu *UserVerificationUpdate) SetNillableAttempts(i *int) *UserVerificationUpdate {
	if i != nil {
		uvu.SetAttempts(*i)
	}
	return uvu
}

// AddAttempts adds i to the "attempts" field.
func (uvu *UserVerificationUpdate) AddAttempts(i int) *UserVerificationUpdate {
	uvu.mutation.AddAttempts(i)
	return uvu
}

// SetExpiresAt sets the "expires_at" field.
func (uvu *UserVerificationUpdate) SetExpiresAt(t time.Time) *UserVerificationUpdate {
	uvu.mutation.SetExpiresAt(t)
	return uvu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (uvu *UserVerificationUpdate) SetNillableExpiresAt(t *time.Time) *UserVerificationUpdate {
	if t != nil {
		uvu.SetExpiresAt(*t)
	}
	return uvu
}

// SetUpdatedAt sets the "updated_at" field.
func (uvu *UserVerificationUpdate) SetUpdatedAt(t time.Time) *UserVerificationUpdate {
	uvu.mutation.SetUpdatedAt(t)
	return uvu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uvu *UserVerificationUpdate) SetUserID(id uuid.UUID) *UserVerificationUpdate {
	uvu.mutation.SetUserID(id)
	return uvu
}

// SetUser sets the "user" edge to the User entity.
func (uvu *UserVerificationUpdate) SetUser(u *User) *UserVerificationUpdate {
	return uvu.SetUserID(u.ID)
}

// Mutation returns the UserVerificationMutation object of the builder.
func (uvu *UserVerificationUpdate) Mutation() *UserVerificationMutation {
	return uvu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uvu *UserVerificationUpdate) ClearUser() *UserVerificationUpdate {
	uvu.mutation.ClearUser()
	return uvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uvu *UserVerificationUpdate) Save(ctx context.Context) (int, error) {
	uvu.defaults()
	return withHooks(ctx, uvu.sqlSave, uvu.mutation, uvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uvu *UserVerificationUpdate) SaveX(ctx context.Context) int {
	affected, err := uvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uvu *UserVerificationUpdate) Exec(ctx context.Context) error {
	_, err := uvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvu *UserVerificationUpdate) ExecX(ctx context.Context) {
	if err := uvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uvu *UserVerificationUpdate) defaults() {
	if _, ok := uvu.mutation.UpdatedAt(); !ok {
		v := userverification.UpdateDefaultUpdatedAt()
		uvu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvu *UserVerificationUpdate) check() error {
	if v, ok := uvu.mutation.Code(); ok {
		if err := userverification.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "UserVerification.code": %w`, err)}
		}
	}
	if uvu.mutation.UserCleared() && len(uvu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserVerification.user"`)
	}
	return nil
}

func (uvu *UserVerificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userverification.Table, userverification.Columns, sqlgraph.NewFieldSpec(userverification.FieldID, field.TypeUUID))
	if ps := uvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvu.mutation.Code(); ok {
		_spec.SetField(userverification.FieldCode, field.TypeString, value)
	}
	if value, ok := uvu.mutation.Attempts(); ok {
		_spec.SetField(userverification.FieldAttempts, field.TypeInt, value)
	}
	if value, ok := uvu.mutation.AddedAttempts(); ok {
		_spec.AddField(userverification.FieldAttempts, field.TypeInt, value)
	}
	if value, ok := uvu.mutation.ExpiresAt(); ok {
		_spec.SetField(userverification.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := uvu.mutation.UpdatedAt(); ok {
		_spec.SetField(userverification.FieldUpdatedAt, field.TypeTime, value)
	}
	if uvu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userverification.UserTable,
			Columns: []string{userverification.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uvu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userverification.UserTable,
			Columns: []string{userverification.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userverification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uvu.mutation.done = true
	return n, nil
}

// UserVerificationUpdateOne is the builder for updating a single UserVerification entity.
type UserVerificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserVerificationMutation
}

// SetCode sets the "code" field.
func (uvuo *UserVerificationUpdateOne) SetCode(s string) *UserVerificationUpdateOne {
	uvuo.mutation.SetCode(s)
	return uvuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (uvuo *UserVerificationUpdateOne) SetNillableCode(s *string) *UserVerificationUpdateOne {
	if s != nil {
		uvuo.SetCode(*s)
	}
	return uvuo
}

// SetAttempts sets the "attempts" field.
func (uvuo *UserVerificationUpdateOne) SetAttempts(i int) *UserVerificationUpdateOne {
	uvuo.mutation.ResetAttempts()
	uvuo.mutation.SetAttempts(i)
	return uvuo
}

// SetNillableAttempts sets the "attempts" field if the given value is not nil.
func (uvuo *UserVerificationUpdateOne) SetNillableAttempts(i *int) *UserVerificationUpdateOne {
	if i != nil {
		uvuo.SetAttempts(*i)
	}
	return uvuo
}

// AddAttempts adds i to the "attempts" field.
func (uvuo *UserVerificationUpdateOne) AddAttempts(i int) *UserVerificationUpdateOne {
	uvuo.mutation.AddAttempts(i)
	return uvuo
}

// SetExpiresAt sets the "expires_at" field.
func (uvuo *UserVerificationUpdateOne) SetExpiresAt(t time.Time) *UserVerificationUpdateOne {
	uvuo.mutation.SetExpiresAt(t)
	return uvuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (uvuo *UserVerificationUpdateOne) SetNillableExpiresAt(t *time.Time) *UserVerificationUpdateOne {
	if t != nil {
		uvuo.SetExpiresAt(*t)
	}
	return uvuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uvuo *UserVerificationUpdateOne) SetUpdatedAt(t time.Time) *UserVerificationUpdateOne {
	uvuo.mutation.SetUpdatedAt(t)
	return uvuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uvuo *UserVerificationUpdateOne) SetUserID(id uuid.UUID) *UserVerificationUpdateOne {
	uvuo.mutation.SetUserID(id)
	return uvuo
}

// SetUser sets the "user" edge to the User entity.
func (uvuo *UserVerificationUpdateOne) SetUser(u *User) *UserVerificationUpdateOne {
	return uvuo.SetUserID(u.ID)
}

// Mutation returns the UserVerificationMutation object of the builder.
func (uvuo *UserVerificationUpdateOne) Mutation() *UserVerificationMutation {
	return uvuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uvuo *UserVerificationUpdateOne) ClearUser() *UserVerificationUpdateOne {
	uvuo.mutation.ClearUser()
	return uvuo
}

// Where appends a list predicates to the UserVerificationUpdate builder.
func (uvuo *UserVerificationUpdateOne) Where(ps ...predicate.UserVerification) *UserVerificationUpdateOne {
	uvuo.mutation.Where(ps...)
	return uvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uvuo *UserVerificationUpdateOne) Select(field string, fields ...string) *UserVerificationUpdateOne {
	uvuo.fields = append([]string{field}, fields...)
	return uvuo
}

// Save executes the query and returns the updated UserVerification entity.
func (uvuo *UserVerificationUpdateOne) Save(ctx context.Context) (*UserVerification, error) {
	uvuo.defaults()
	return withHooks(ctx, uvuo.sqlSave, uvuo.mutation, uvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uvuo *UserVerificationUpdateOne) SaveX(ctx context.Context) *UserVerification {
	node, err := uvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uvuo *UserVerificationUpdateOne) Exec(ctx context.Context) error {
	_, err := uvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvuo *UserVerificationUpdateOne) ExecX(ctx context.Context) {
	if err := uvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uvuo *UserVerificationUpdateOne) defaults() {
	if _, ok := uvuo.mutation.UpdatedAt(); !ok {
		v := userverification.UpdateDefaultUpdatedAt()
		uvuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvuo *UserVerificationUpdateOne) check() error {
	if v, ok := uvuo.mutation.Code(); ok {
		if err := userverification.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "UserVerification.code": %w`, err)}
		}
	}
	if uvuo.mutation.UserCleared() && len(uvuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserVerification.user"`)
	}
	return nil
}

func (uvuo *UserVerificationUpdateOne) sqlSave(ctx context.Context) (_node *UserVerification, err error) {
	if err := uvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userverification.Table, userverification.Columns, sqlgraph.NewFieldSpec(userverification.FieldID, field.TypeUUID))
	id, ok := uvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserVerification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userverification.FieldID)
		for _, f := range fields {
			if !userverification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userverification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvuo.mutation.Code(); ok {
		_spec.SetField(userverification.FieldCode, field.TypeString, value)
	}
	if value, ok := uvuo.mutation.Attempts(); ok {
		_spec.SetField(userverification.FieldAttempts, field.TypeInt, value)
	}
	if value, ok := uvuo.mutation.AddedAttempts(); ok {
		_spec.AddField(userverification.FieldAttempts, field.TypeInt, value)
	}
	if value, ok := uvuo.mutation.ExpiresAt(); ok {
		_spec.SetField(userverification.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := uvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userverification.FieldUpdatedAt, field.TypeTime, value)
	}
	if uvuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userverification.UserTable,
			Columns: []string{userverification.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uvuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userverification.UserTable,
			Columns: []string{userverification.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserVerification{config: uvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userverification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uvuo.mutation.done = true
	return _node, nil
}
