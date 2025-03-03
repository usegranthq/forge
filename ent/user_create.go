// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/usegranthq/forge/ent/project"
	"github.com/usegranthq/forge/ent/token"
	"github.com/usegranthq/forge/ent/user"
	"github.com/usegranthq/forge/ent/usersession"
	"github.com/usegranthq/forge/ent/verification"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUID sets the "uid" field.
func (uc *UserCreate) SetUID(u uuid.UUID) *UserCreate {
	uc.mutation.SetUID(u)
	return uc
}

// SetNillableUID sets the "uid" field if the given value is not nil.
func (uc *UserCreate) SetNillableUID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetUID(*u)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetLastLogin sets the "last_login" field.
func (uc *UserCreate) SetLastLogin(t time.Time) *UserCreate {
	uc.mutation.SetLastLogin(t)
	return uc
}

// SetNillableLastLogin sets the "last_login" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastLogin(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetLastLogin(*t)
	}
	return uc
}

// SetVerifiedAt sets the "verified_at" field.
func (uc *UserCreate) SetVerifiedAt(t time.Time) *UserCreate {
	uc.mutation.SetVerifiedAt(t)
	return uc
}

// SetNillableVerifiedAt sets the "verified_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableVerifiedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetVerifiedAt(*t)
	}
	return uc
}

// SetProvider sets the "provider" field.
func (uc *UserCreate) SetProvider(u user.Provider) *UserCreate {
	uc.mutation.SetProvider(u)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// AddUserSessionIDs adds the "user_sessions" edge to the UserSession entity by IDs.
func (uc *UserCreate) AddUserSessionIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddUserSessionIDs(ids...)
	return uc
}

// AddUserSessions adds the "user_sessions" edges to the UserSession entity.
func (uc *UserCreate) AddUserSessions(u ...*UserSession) *UserCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddUserSessionIDs(ids...)
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (uc *UserCreate) AddProjectIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddProjectIDs(ids...)
	return uc
}

// AddProjects adds the "projects" edges to the Project entity.
func (uc *UserCreate) AddProjects(p ...*Project) *UserCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddProjectIDs(ids...)
}

// AddVerificationIDs adds the "verifications" edge to the Verification entity by IDs.
func (uc *UserCreate) AddVerificationIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddVerificationIDs(ids...)
	return uc
}

// AddVerifications adds the "verifications" edges to the Verification entity.
func (uc *UserCreate) AddVerifications(v ...*Verification) *UserCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uc.AddVerificationIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uc *UserCreate) AddTokenIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddTokenIDs(ids...)
	return uc
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uc *UserCreate) AddTokens(t ...*Token) *UserCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.UID(); !ok {
		v := user.DefaultUID()
		uc.mutation.SetUID(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.UID(); !ok {
		return &ValidationError{Name: "uid", err: errors.New(`ent: missing required field "User.uid"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if v, ok := uc.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if _, ok := uc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "User.provider"`)}
	}
	if v, ok := uc.mutation.Provider(); ok {
		if err := user.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "User.provider": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.UID(); ok {
		_spec.SetField(user.FieldUID, field.TypeUUID, value)
		_node.UID = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
		_node.LastLogin = value
	}
	if value, ok := uc.mutation.VerifiedAt(); ok {
		_spec.SetField(user.FieldVerifiedAt, field.TypeTime, value)
		_node.VerifiedAt = value
	}
	if value, ok := uc.mutation.Provider(); ok {
		_spec.SetField(user.FieldProvider, field.TypeEnum, value)
		_node.Provider = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := uc.mutation.UserSessionsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ProjectsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.VerificationsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.TokensIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// SetLastLogin sets the "last_login" field.
func (u *UserUpsert) SetLastLogin(v time.Time) *UserUpsert {
	u.Set(user.FieldLastLogin, v)
	return u
}

// UpdateLastLogin sets the "last_login" field to the value that was provided on create.
func (u *UserUpsert) UpdateLastLogin() *UserUpsert {
	u.SetExcluded(user.FieldLastLogin)
	return u
}

// ClearLastLogin clears the value of the "last_login" field.
func (u *UserUpsert) ClearLastLogin() *UserUpsert {
	u.SetNull(user.FieldLastLogin)
	return u
}

// SetVerifiedAt sets the "verified_at" field.
func (u *UserUpsert) SetVerifiedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldVerifiedAt, v)
	return u
}

// UpdateVerifiedAt sets the "verified_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateVerifiedAt() *UserUpsert {
	u.SetExcluded(user.FieldVerifiedAt)
	return u
}

// ClearVerifiedAt clears the value of the "verified_at" field.
func (u *UserUpsert) ClearVerifiedAt() *UserUpsert {
	u.SetNull(user.FieldVerifiedAt)
	return u
}

// SetProvider sets the "provider" field.
func (u *UserUpsert) SetProvider(v user.Provider) *UserUpsert {
	u.Set(user.FieldProvider, v)
	return u
}

// UpdateProvider sets the "provider" field to the value that was provided on create.
func (u *UserUpsert) UpdateProvider() *UserUpsert {
	u.SetExcluded(user.FieldProvider)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsert) SetUpdatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedAt() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
		if _, exists := u.create.mutation.UID(); exists {
			s.SetIgnore(user.FieldUID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(user.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetLastLogin sets the "last_login" field.
func (u *UserUpsertOne) SetLastLogin(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetLastLogin(v)
	})
}

// UpdateLastLogin sets the "last_login" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateLastLogin() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastLogin()
	})
}

// ClearLastLogin clears the value of the "last_login" field.
func (u *UserUpsertOne) ClearLastLogin() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearLastLogin()
	})
}

// SetVerifiedAt sets the "verified_at" field.
func (u *UserUpsertOne) SetVerifiedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetVerifiedAt(v)
	})
}

// UpdateVerifiedAt sets the "verified_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateVerifiedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateVerifiedAt()
	})
}

// ClearVerifiedAt clears the value of the "verified_at" field.
func (u *UserUpsertOne) ClearVerifiedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.ClearVerifiedAt()
	})
}

// SetProvider sets the "provider" field.
func (u *UserUpsertOne) SetProvider(v user.Provider) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetProvider(v)
	})
}

// UpdateProvider sets the "provider" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateProvider() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateProvider()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertOne) SetUpdatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetUID(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
			if _, exists := b.mutation.UID(); exists {
				s.SetIgnore(user.FieldUID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(user.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetLastLogin sets the "last_login" field.
func (u *UserUpsertBulk) SetLastLogin(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetLastLogin(v)
	})
}

// UpdateLastLogin sets the "last_login" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateLastLogin() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateLastLogin()
	})
}

// ClearLastLogin clears the value of the "last_login" field.
func (u *UserUpsertBulk) ClearLastLogin() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearLastLogin()
	})
}

// SetVerifiedAt sets the "verified_at" field.
func (u *UserUpsertBulk) SetVerifiedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetVerifiedAt(v)
	})
}

// UpdateVerifiedAt sets the "verified_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateVerifiedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateVerifiedAt()
	})
}

// ClearVerifiedAt clears the value of the "verified_at" field.
func (u *UserUpsertBulk) ClearVerifiedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.ClearVerifiedAt()
	})
}

// SetProvider sets the "provider" field.
func (u *UserUpsertBulk) SetProvider(v user.Provider) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetProvider(v)
	})
}

// UpdateProvider sets the "provider" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateProvider() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateProvider()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertBulk) SetUpdatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
