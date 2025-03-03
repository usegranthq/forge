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
	"github.com/usegranthq/forge/ent/user"
	"github.com/usegranthq/forge/ent/verification"
)

// VerificationCreate is the builder for creating a Verification entity.
type VerificationCreate struct {
	config
	mutation *VerificationMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAttemptID sets the "attempt_id" field.
func (vc *VerificationCreate) SetAttemptID(u uuid.UUID) *VerificationCreate {
	vc.mutation.SetAttemptID(u)
	return vc
}

// SetNillableAttemptID sets the "attempt_id" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableAttemptID(u *uuid.UUID) *VerificationCreate {
	if u != nil {
		vc.SetAttemptID(*u)
	}
	return vc
}

// SetType sets the "type" field.
func (vc *VerificationCreate) SetType(v verification.Type) *VerificationCreate {
	vc.mutation.SetType(v)
	return vc
}

// SetCode sets the "code" field.
func (vc *VerificationCreate) SetCode(s string) *VerificationCreate {
	vc.mutation.SetCode(s)
	return vc
}

// SetAttempts sets the "attempts" field.
func (vc *VerificationCreate) SetAttempts(i int) *VerificationCreate {
	vc.mutation.SetAttempts(i)
	return vc
}

// SetNillableAttempts sets the "attempts" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableAttempts(i *int) *VerificationCreate {
	if i != nil {
		vc.SetAttempts(*i)
	}
	return vc
}

// SetExpiresAt sets the "expires_at" field.
func (vc *VerificationCreate) SetExpiresAt(t time.Time) *VerificationCreate {
	vc.mutation.SetExpiresAt(t)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VerificationCreate) SetCreatedAt(t time.Time) *VerificationCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableCreatedAt(t *time.Time) *VerificationCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VerificationCreate) SetUpdatedAt(t time.Time) *VerificationCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableUpdatedAt(t *time.Time) *VerificationCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetID sets the "id" field.
func (vc *VerificationCreate) SetID(u uuid.UUID) *VerificationCreate {
	vc.mutation.SetID(u)
	return vc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vc *VerificationCreate) SetNillableID(u *uuid.UUID) *VerificationCreate {
	if u != nil {
		vc.SetID(*u)
	}
	return vc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vc *VerificationCreate) SetUserID(id uuid.UUID) *VerificationCreate {
	vc.mutation.SetUserID(id)
	return vc
}

// SetUser sets the "user" edge to the User entity.
func (vc *VerificationCreate) SetUser(u *User) *VerificationCreate {
	return vc.SetUserID(u.ID)
}

// Mutation returns the VerificationMutation object of the builder.
func (vc *VerificationCreate) Mutation() *VerificationMutation {
	return vc.mutation
}

// Save creates the Verification in the database.
func (vc *VerificationCreate) Save(ctx context.Context) (*Verification, error) {
	vc.defaults()
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VerificationCreate) SaveX(ctx context.Context) *Verification {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VerificationCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VerificationCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vc *VerificationCreate) defaults() {
	if _, ok := vc.mutation.AttemptID(); !ok {
		v := verification.DefaultAttemptID()
		vc.mutation.SetAttemptID(v)
	}
	if _, ok := vc.mutation.Attempts(); !ok {
		v := verification.DefaultAttempts
		vc.mutation.SetAttempts(v)
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		v := verification.DefaultCreatedAt()
		vc.mutation.SetCreatedAt(v)
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		v := verification.DefaultUpdatedAt()
		vc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vc.mutation.ID(); !ok {
		v := verification.DefaultID()
		vc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VerificationCreate) check() error {
	if _, ok := vc.mutation.AttemptID(); !ok {
		return &ValidationError{Name: "attempt_id", err: errors.New(`ent: missing required field "Verification.attempt_id"`)}
	}
	if _, ok := vc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Verification.type"`)}
	}
	if v, ok := vc.mutation.GetType(); ok {
		if err := verification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Verification.type": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Verification.code"`)}
	}
	if v, ok := vc.mutation.Code(); ok {
		if err := verification.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Verification.code": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Attempts(); !ok {
		return &ValidationError{Name: "attempts", err: errors.New(`ent: missing required field "Verification.attempts"`)}
	}
	if _, ok := vc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`ent: missing required field "Verification.expires_at"`)}
	}
	if _, ok := vc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Verification.created_at"`)}
	}
	if _, ok := vc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Verification.updated_at"`)}
	}
	if len(vc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Verification.user"`)}
	}
	return nil
}

func (vc *VerificationCreate) sqlSave(ctx context.Context) (*Verification, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
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
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VerificationCreate) createSpec() (*Verification, *sqlgraph.CreateSpec) {
	var (
		_node = &Verification{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(verification.Table, sqlgraph.NewFieldSpec(verification.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = vc.conflict
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := vc.mutation.AttemptID(); ok {
		_spec.SetField(verification.FieldAttemptID, field.TypeUUID, value)
		_node.AttemptID = value
	}
	if value, ok := vc.mutation.GetType(); ok {
		_spec.SetField(verification.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := vc.mutation.Code(); ok {
		_spec.SetField(verification.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := vc.mutation.Attempts(); ok {
		_spec.SetField(verification.FieldAttempts, field.TypeInt, value)
		_node.Attempts = value
	}
	if value, ok := vc.mutation.ExpiresAt(); ok {
		_spec.SetField(verification.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(verification.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(verification.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := vc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   verification.UserTable,
			Columns: []string{verification.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_verifications = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Verification.Create().
//		SetAttemptID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VerificationUpsert) {
//			SetAttemptID(v+v).
//		}).
//		Exec(ctx)
func (vc *VerificationCreate) OnConflict(opts ...sql.ConflictOption) *VerificationUpsertOne {
	vc.conflict = opts
	return &VerificationUpsertOne{
		create: vc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Verification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vc *VerificationCreate) OnConflictColumns(columns ...string) *VerificationUpsertOne {
	vc.conflict = append(vc.conflict, sql.ConflictColumns(columns...))
	return &VerificationUpsertOne{
		create: vc,
	}
}

type (
	// VerificationUpsertOne is the builder for "upsert"-ing
	//  one Verification node.
	VerificationUpsertOne struct {
		create *VerificationCreate
	}

	// VerificationUpsert is the "OnConflict" setter.
	VerificationUpsert struct {
		*sql.UpdateSet
	}
)

// SetType sets the "type" field.
func (u *VerificationUpsert) SetType(v verification.Type) *VerificationUpsert {
	u.Set(verification.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *VerificationUpsert) UpdateType() *VerificationUpsert {
	u.SetExcluded(verification.FieldType)
	return u
}

// SetCode sets the "code" field.
func (u *VerificationUpsert) SetCode(v string) *VerificationUpsert {
	u.Set(verification.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *VerificationUpsert) UpdateCode() *VerificationUpsert {
	u.SetExcluded(verification.FieldCode)
	return u
}

// SetAttempts sets the "attempts" field.
func (u *VerificationUpsert) SetAttempts(v int) *VerificationUpsert {
	u.Set(verification.FieldAttempts, v)
	return u
}

// UpdateAttempts sets the "attempts" field to the value that was provided on create.
func (u *VerificationUpsert) UpdateAttempts() *VerificationUpsert {
	u.SetExcluded(verification.FieldAttempts)
	return u
}

// AddAttempts adds v to the "attempts" field.
func (u *VerificationUpsert) AddAttempts(v int) *VerificationUpsert {
	u.Add(verification.FieldAttempts, v)
	return u
}

// SetExpiresAt sets the "expires_at" field.
func (u *VerificationUpsert) SetExpiresAt(v time.Time) *VerificationUpsert {
	u.Set(verification.FieldExpiresAt, v)
	return u
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *VerificationUpsert) UpdateExpiresAt() *VerificationUpsert {
	u.SetExcluded(verification.FieldExpiresAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VerificationUpsert) SetUpdatedAt(v time.Time) *VerificationUpsert {
	u.Set(verification.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VerificationUpsert) UpdateUpdatedAt() *VerificationUpsert {
	u.SetExcluded(verification.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Verification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(verification.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *VerificationUpsertOne) UpdateNewValues() *VerificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(verification.FieldID)
		}
		if _, exists := u.create.mutation.AttemptID(); exists {
			s.SetIgnore(verification.FieldAttemptID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(verification.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Verification.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *VerificationUpsertOne) Ignore() *VerificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VerificationUpsertOne) DoNothing() *VerificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VerificationCreate.OnConflict
// documentation for more info.
func (u *VerificationUpsertOne) Update(set func(*VerificationUpsert)) *VerificationUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VerificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *VerificationUpsertOne) SetType(v verification.Type) *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *VerificationUpsertOne) UpdateType() *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateType()
	})
}

// SetCode sets the "code" field.
func (u *VerificationUpsertOne) SetCode(v string) *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *VerificationUpsertOne) UpdateCode() *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateCode()
	})
}

// SetAttempts sets the "attempts" field.
func (u *VerificationUpsertOne) SetAttempts(v int) *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.SetAttempts(v)
	})
}

// AddAttempts adds v to the "attempts" field.
func (u *VerificationUpsertOne) AddAttempts(v int) *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.AddAttempts(v)
	})
}

// UpdateAttempts sets the "attempts" field to the value that was provided on create.
func (u *VerificationUpsertOne) UpdateAttempts() *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateAttempts()
	})
}

// SetExpiresAt sets the "expires_at" field.
func (u *VerificationUpsertOne) SetExpiresAt(v time.Time) *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.SetExpiresAt(v)
	})
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *VerificationUpsertOne) UpdateExpiresAt() *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateExpiresAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VerificationUpsertOne) SetUpdatedAt(v time.Time) *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VerificationUpsertOne) UpdateUpdatedAt() *VerificationUpsertOne {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *VerificationUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VerificationCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VerificationUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VerificationUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: VerificationUpsertOne.ID is not supported by MySQL driver. Use VerificationUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VerificationUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VerificationCreateBulk is the builder for creating many Verification entities in bulk.
type VerificationCreateBulk struct {
	config
	err      error
	builders []*VerificationCreate
	conflict []sql.ConflictOption
}

// Save creates the Verification entities in the database.
func (vcb *VerificationCreateBulk) Save(ctx context.Context) ([]*Verification, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Verification, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VerificationMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VerificationCreateBulk) SaveX(ctx context.Context) []*Verification {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VerificationCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VerificationCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Verification.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VerificationUpsert) {
//			SetAttemptID(v+v).
//		}).
//		Exec(ctx)
func (vcb *VerificationCreateBulk) OnConflict(opts ...sql.ConflictOption) *VerificationUpsertBulk {
	vcb.conflict = opts
	return &VerificationUpsertBulk{
		create: vcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Verification.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vcb *VerificationCreateBulk) OnConflictColumns(columns ...string) *VerificationUpsertBulk {
	vcb.conflict = append(vcb.conflict, sql.ConflictColumns(columns...))
	return &VerificationUpsertBulk{
		create: vcb,
	}
}

// VerificationUpsertBulk is the builder for "upsert"-ing
// a bulk of Verification nodes.
type VerificationUpsertBulk struct {
	create *VerificationCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Verification.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(verification.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *VerificationUpsertBulk) UpdateNewValues() *VerificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(verification.FieldID)
			}
			if _, exists := b.mutation.AttemptID(); exists {
				s.SetIgnore(verification.FieldAttemptID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(verification.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Verification.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *VerificationUpsertBulk) Ignore() *VerificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VerificationUpsertBulk) DoNothing() *VerificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VerificationCreateBulk.OnConflict
// documentation for more info.
func (u *VerificationUpsertBulk) Update(set func(*VerificationUpsert)) *VerificationUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VerificationUpsert{UpdateSet: update})
	}))
	return u
}

// SetType sets the "type" field.
func (u *VerificationUpsertBulk) SetType(v verification.Type) *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *VerificationUpsertBulk) UpdateType() *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateType()
	})
}

// SetCode sets the "code" field.
func (u *VerificationUpsertBulk) SetCode(v string) *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.SetCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *VerificationUpsertBulk) UpdateCode() *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateCode()
	})
}

// SetAttempts sets the "attempts" field.
func (u *VerificationUpsertBulk) SetAttempts(v int) *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.SetAttempts(v)
	})
}

// AddAttempts adds v to the "attempts" field.
func (u *VerificationUpsertBulk) AddAttempts(v int) *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.AddAttempts(v)
	})
}

// UpdateAttempts sets the "attempts" field to the value that was provided on create.
func (u *VerificationUpsertBulk) UpdateAttempts() *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateAttempts()
	})
}

// SetExpiresAt sets the "expires_at" field.
func (u *VerificationUpsertBulk) SetExpiresAt(v time.Time) *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.SetExpiresAt(v)
	})
}

// UpdateExpiresAt sets the "expires_at" field to the value that was provided on create.
func (u *VerificationUpsertBulk) UpdateExpiresAt() *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateExpiresAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *VerificationUpsertBulk) SetUpdatedAt(v time.Time) *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *VerificationUpsertBulk) UpdateUpdatedAt() *VerificationUpsertBulk {
	return u.Update(func(s *VerificationUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *VerificationUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the VerificationCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VerificationCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VerificationUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
