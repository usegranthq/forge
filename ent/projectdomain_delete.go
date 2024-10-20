// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/usegranthq/backend/ent/predicate"
	"github.com/usegranthq/backend/ent/projectdomain"
)

// ProjectDomainDelete is the builder for deleting a ProjectDomain entity.
type ProjectDomainDelete struct {
	config
	hooks    []Hook
	mutation *ProjectDomainMutation
}

// Where appends a list predicates to the ProjectDomainDelete builder.
func (pdd *ProjectDomainDelete) Where(ps ...predicate.ProjectDomain) *ProjectDomainDelete {
	pdd.mutation.Where(ps...)
	return pdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pdd *ProjectDomainDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pdd.sqlExec, pdd.mutation, pdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pdd *ProjectDomainDelete) ExecX(ctx context.Context) int {
	n, err := pdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pdd *ProjectDomainDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(projectdomain.Table, sqlgraph.NewFieldSpec(projectdomain.FieldID, field.TypeUUID))
	if ps := pdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pdd.mutation.done = true
	return affected, err
}

// ProjectDomainDeleteOne is the builder for deleting a single ProjectDomain entity.
type ProjectDomainDeleteOne struct {
	pdd *ProjectDomainDelete
}

// Where appends a list predicates to the ProjectDomainDelete builder.
func (pddo *ProjectDomainDeleteOne) Where(ps ...predicate.ProjectDomain) *ProjectDomainDeleteOne {
	pddo.pdd.mutation.Where(ps...)
	return pddo
}

// Exec executes the deletion query.
func (pddo *ProjectDomainDeleteOne) Exec(ctx context.Context) error {
	n, err := pddo.pdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{projectdomain.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pddo *ProjectDomainDeleteOne) ExecX(ctx context.Context) {
	if err := pddo.Exec(ctx); err != nil {
		panic(err)
	}
}
