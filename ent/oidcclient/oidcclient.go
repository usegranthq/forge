// Code generated by ent, DO NOT EDIT.

package oidcclient

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the oidcclient type in the database.
	Label = "oidc_client"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAudience holds the string denoting the audience field in the database.
	FieldAudience = "audience"
	// FieldClientRefID holds the string denoting the client_ref_id field in the database.
	FieldClientRefID = "client_ref_id"
	// FieldClientID holds the string denoting the client_id field in the database.
	FieldClientID = "client_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the oidcclient in the database.
	Table = "oidc_clients"
	// ProjectTable is the table that holds the project relation/edge.
	ProjectTable = "oidc_clients"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
	// ProjectColumn is the table column denoting the project relation/edge.
	ProjectColumn = "project_oidc_clients"
)

// Columns holds all SQL columns for oidcclient fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAudience,
	FieldClientRefID,
	FieldClientID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "oidc_clients"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"project_oidc_clients",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AudienceValidator is a validator for the "audience" field. It is called by the builders before save.
	AudienceValidator func(string) error
	// ClientRefIDValidator is a validator for the "client_ref_id" field. It is called by the builders before save.
	ClientRefIDValidator func(string) error
	// ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	ClientIDValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the OidcClient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAudience orders the results by the audience field.
func ByAudience(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAudience, opts...).ToFunc()
}

// ByClientRefID orders the results by the client_ref_id field.
func ByClientRefID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientRefID, opts...).ToFunc()
}

// ByClientID orders the results by the client_id field.
func ByClientID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClientID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProjectField orders the results by project field.
func ByProjectField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProjectStep(), sql.OrderByField(field, opts...))
	}
}
func newProjectStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProjectInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
	)
}
