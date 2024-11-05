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
	"github.com/usegranthq/backend/ent/project"
	"github.com/usegranthq/backend/ent/token"
	"github.com/usegranthq/backend/ent/user"
	"github.com/usegranthq/backend/ent/usersession"
	"github.com/usegranthq/backend/ent/verification"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetLastLogin sets the "last_login" field.
func (uu *UserUpdate) SetLastLogin(t time.Time) *UserUpdate {
	uu.mutation.SetLastLogin(t)
	return uu
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastLogin(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastLogin(*t)
	}
	return uu
}

// ClearLastLogin clears the value of the "last_login" field.
func (uu *UserUpdate) ClearLastLogin() *UserUpdate {
	uu.mutation.ClearLastLogin()
	return uu
}

// SetVerifiedAt sets the "verified_at" field.
func (uu *UserUpdate) SetVerifiedAt(t time.Time) *UserUpdate {
	uu.mutation.SetVerifiedAt(t)
	return uu
}

// SetNillableVerifiedAt sets the "verified_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerifiedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetVerifiedAt(*t)
	}
	return uu
}

// ClearVerifiedAt clears the value of the "verified_at" field.
func (uu *UserUpdate) ClearVerifiedAt() *UserUpdate {
	uu.mutation.ClearVerifiedAt()
	return uu
}

// SetProvider sets the "provider" field.
func (uu *UserUpdate) SetProvider(u user.Provider) *UserUpdate {
	uu.mutation.SetProvider(u)
	return uu
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (uu *UserUpdate) SetNillableProvider(u *user.Provider) *UserUpdate {
	if u != nil {
		uu.SetProvider(*u)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// AddUserSessionIDs adds the "user_sessions" edge to the UserSession entity by IDs.
func (uu *UserUpdate) AddUserSessionIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddUserSessionIDs(ids...)
	return uu
}

// AddUserSessions adds the "user_sessions" edges to the UserSession entity.
func (uu *UserUpdate) AddUserSessions(u ...*UserSession) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddUserSessionIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (uu *UserUpdate) AddProjectIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddProjectIDs(ids...)
	return uu
}

// AddProjects adds the "projects" edges to the Project entity.
func (uu *UserUpdate) AddProjects(p ...*Project) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddProjectIDs(ids...)
}

// AddVerificationIDs adds the "verifications" edge to the Verification entity by IDs.
func (uu *UserUpdate) AddVerificationIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddVerificationIDs(ids...)
	return uu
}

// AddVerifications adds the "verifications" edges to the Verification entity.
func (uu *UserUpdate) AddVerifications(v ...*Verification) *UserUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uu.AddVerificationIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uu *UserUpdate) AddTokenIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddTokenIDs(ids...)
	return uu
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uu *UserUpdate) AddTokens(t ...*Token) *UserUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearUserSessions clears all "user_sessions" edges to the UserSession entity.
func (uu *UserUpdate) ClearUserSessions() *UserUpdate {
	uu.mutation.ClearUserSessions()
	return uu
}

// RemoveUserSessionIDs removes the "user_sessions" edge to UserSession entities by IDs.
func (uu *UserUpdate) RemoveUserSessionIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveUserSessionIDs(ids...)
	return uu
}

// RemoveUserSessions removes "user_sessions" edges to UserSession entities.
func (uu *UserUpdate) RemoveUserSessions(u ...*UserSession) *UserUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveUserSessionIDs(ids...)
}

// ClearProjects clears all "projects" edges to the Project entity.
func (uu *UserUpdate) ClearProjects() *UserUpdate {
	uu.mutation.ClearProjects()
	return uu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (uu *UserUpdate) RemoveProjectIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveProjectIDs(ids...)
	return uu
}

// RemoveProjects removes "projects" edges to Project entities.
func (uu *UserUpdate) RemoveProjects(p ...*Project) *UserUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveProjectIDs(ids...)
}

// ClearVerifications clears all "verifications" edges to the Verification entity.
func (uu *UserUpdate) ClearVerifications() *UserUpdate {
	uu.mutation.ClearVerifications()
	return uu
}

// RemoveVerificationIDs removes the "verifications" edge to Verification entities by IDs.
func (uu *UserUpdate) RemoveVerificationIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveVerificationIDs(ids...)
	return uu
}

// RemoveVerifications removes "verifications" edges to Verification entities.
func (uu *UserUpdate) RemoveVerifications(v ...*Verification) *UserUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uu.RemoveVerificationIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (uu *UserUpdate) ClearTokens() *UserUpdate {
	uu.mutation.ClearTokens()
	return uu
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (uu *UserUpdate) RemoveTokenIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveTokenIDs(ids...)
	return uu
}

// RemoveTokens removes "tokens" edges to Token entities.
func (uu *UserUpdate) RemoveTokens(t ...*Token) *UserUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Provider(); ok {
		if err := user.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "User.provider": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if uu.mutation.LastLoginCleared() {
		_spec.ClearField(user.FieldLastLogin, field.TypeTime)
	}
	if value, ok := uu.mutation.VerifiedAt(); ok {
		_spec.SetField(user.FieldVerifiedAt, field.TypeTime, value)
	}
	if uu.mutation.VerifiedAtCleared() {
		_spec.ClearField(user.FieldVerifiedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.UserSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserSessionsTable,
			Columns: []string{user.UserSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedUserSessionsIDs(); len(nodes) > 0 && !uu.mutation.UserSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserSessionsTable,
			Columns: []string{user.UserSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserSessionsTable,
			Columns: []string{user.UserSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !uu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
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
	if uu.mutation.VerificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VerificationsTable,
			Columns: []string{user.VerificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedVerificationsIDs(); len(nodes) > 0 && !uu.mutation.VerificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VerificationsTable,
			Columns: []string{user.VerificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.VerificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VerificationsTable,
			Columns: []string{user.VerificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTokensIDs(); len(nodes) > 0 && !uu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetLastLogin sets the "last_login" field.
func (uuo *UserUpdateOne) SetLastLogin(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLogin(t)
	return uuo
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastLogin(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastLogin(*t)
	}
	return uuo
}

// ClearLastLogin clears the value of the "last_login" field.
func (uuo *UserUpdateOne) ClearLastLogin() *UserUpdateOne {
	uuo.mutation.ClearLastLogin()
	return uuo
}

// SetVerifiedAt sets the "verified_at" field.
func (uuo *UserUpdateOne) SetVerifiedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetVerifiedAt(t)
	return uuo
}

// SetNillableVerifiedAt sets the "verified_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerifiedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetVerifiedAt(*t)
	}
	return uuo
}

// ClearVerifiedAt clears the value of the "verified_at" field.
func (uuo *UserUpdateOne) ClearVerifiedAt() *UserUpdateOne {
	uuo.mutation.ClearVerifiedAt()
	return uuo
}

// SetProvider sets the "provider" field.
func (uuo *UserUpdateOne) SetProvider(u user.Provider) *UserUpdateOne {
	uuo.mutation.SetProvider(u)
	return uuo
}

// SetNillableProvider sets the "provider" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProvider(u *user.Provider) *UserUpdateOne {
	if u != nil {
		uuo.SetProvider(*u)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// AddUserSessionIDs adds the "user_sessions" edge to the UserSession entity by IDs.
func (uuo *UserUpdateOne) AddUserSessionIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddUserSessionIDs(ids...)
	return uuo
}

// AddUserSessions adds the "user_sessions" edges to the UserSession entity.
func (uuo *UserUpdateOne) AddUserSessions(u ...*UserSession) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddUserSessionIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (uuo *UserUpdateOne) AddProjectIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddProjectIDs(ids...)
	return uuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (uuo *UserUpdateOne) AddProjects(p ...*Project) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddProjectIDs(ids...)
}

// AddVerificationIDs adds the "verifications" edge to the Verification entity by IDs.
func (uuo *UserUpdateOne) AddVerificationIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddVerificationIDs(ids...)
	return uuo
}

// AddVerifications adds the "verifications" edges to the Verification entity.
func (uuo *UserUpdateOne) AddVerifications(v ...*Verification) *UserUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uuo.AddVerificationIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uuo *UserUpdateOne) AddTokenIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddTokenIDs(ids...)
	return uuo
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uuo *UserUpdateOne) AddTokens(t ...*Token) *UserUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearUserSessions clears all "user_sessions" edges to the UserSession entity.
func (uuo *UserUpdateOne) ClearUserSessions() *UserUpdateOne {
	uuo.mutation.ClearUserSessions()
	return uuo
}

// RemoveUserSessionIDs removes the "user_sessions" edge to UserSession entities by IDs.
func (uuo *UserUpdateOne) RemoveUserSessionIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveUserSessionIDs(ids...)
	return uuo
}

// RemoveUserSessions removes "user_sessions" edges to UserSession entities.
func (uuo *UserUpdateOne) RemoveUserSessions(u ...*UserSession) *UserUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveUserSessionIDs(ids...)
}

// ClearProjects clears all "projects" edges to the Project entity.
func (uuo *UserUpdateOne) ClearProjects() *UserUpdateOne {
	uuo.mutation.ClearProjects()
	return uuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (uuo *UserUpdateOne) RemoveProjectIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveProjectIDs(ids...)
	return uuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (uuo *UserUpdateOne) RemoveProjects(p ...*Project) *UserUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveProjectIDs(ids...)
}

// ClearVerifications clears all "verifications" edges to the Verification entity.
func (uuo *UserUpdateOne) ClearVerifications() *UserUpdateOne {
	uuo.mutation.ClearVerifications()
	return uuo
}

// RemoveVerificationIDs removes the "verifications" edge to Verification entities by IDs.
func (uuo *UserUpdateOne) RemoveVerificationIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveVerificationIDs(ids...)
	return uuo
}

// RemoveVerifications removes "verifications" edges to Verification entities.
func (uuo *UserUpdateOne) RemoveVerifications(v ...*Verification) *UserUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uuo.RemoveVerificationIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (uuo *UserUpdateOne) ClearTokens() *UserUpdateOne {
	uuo.mutation.ClearTokens()
	return uuo
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (uuo *UserUpdateOne) RemoveTokenIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveTokenIDs(ids...)
	return uuo
}

// RemoveTokens removes "tokens" edges to Token entities.
func (uuo *UserUpdateOne) RemoveTokens(t ...*Token) *UserUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTokenIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Provider(); ok {
		if err := user.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "User.provider": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if uuo.mutation.LastLoginCleared() {
		_spec.ClearField(user.FieldLastLogin, field.TypeTime)
	}
	if value, ok := uuo.mutation.VerifiedAt(); ok {
		_spec.SetField(user.FieldVerifiedAt, field.TypeTime, value)
	}
	if uuo.mutation.VerifiedAtCleared() {
		_spec.ClearField(user.FieldVerifiedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.UserSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserSessionsTable,
			Columns: []string{user.UserSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedUserSessionsIDs(); len(nodes) > 0 && !uuo.mutation.UserSessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserSessionsTable,
			Columns: []string{user.UserSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserSessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.UserSessionsTable,
			Columns: []string{user.UserSessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersession.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !uuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ProjectsTable,
			Columns: []string{user.ProjectsColumn},
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
	if uuo.mutation.VerificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VerificationsTable,
			Columns: []string{user.VerificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedVerificationsIDs(); len(nodes) > 0 && !uuo.mutation.VerificationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VerificationsTable,
			Columns: []string{user.VerificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.VerificationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VerificationsTable,
			Columns: []string{user.VerificationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !uuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
