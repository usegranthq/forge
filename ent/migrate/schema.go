// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OidcClientsColumns holds the columns for the "oidc_clients" table.
	OidcClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "audience", Type: field.TypeString},
		{Name: "client_ref_id", Type: field.TypeString},
		{Name: "client_id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "project_oidc_clients", Type: field.TypeUUID},
	}
	// OidcClientsTable holds the schema information for the "oidc_clients" table.
	OidcClientsTable = &schema.Table{
		Name:       "oidc_clients",
		Columns:    OidcClientsColumns,
		PrimaryKey: []*schema.Column{OidcClientsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "oidc_clients_projects_oidc_clients",
				Columns:    []*schema.Column{OidcClientsColumns[7]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "url_id", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_projects", Type: field.TypeUUID},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "projects_users_projects",
				Columns:    []*schema.Column{ProjectsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProjectDomainsColumns holds the columns for the "project_domains" table.
	ProjectDomainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "domain", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "verified_at", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "project_domain", Type: field.TypeUUID},
	}
	// ProjectDomainsTable holds the schema information for the "project_domains" table.
	ProjectDomainsTable = &schema.Table{
		Name:       "project_domains",
		Columns:    ProjectDomainsColumns,
		PrimaryKey: []*schema.Column{ProjectDomainsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_domains_projects_domain",
				Columns:    []*schema.Column{ProjectDomainsColumns[6]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "token", Type: field.TypeString},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "last_used_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_tokens", Type: field.TypeUUID, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_users_tokens",
				Columns:    []*schema.Column{TokensColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "uid", Type: field.TypeUUID},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "last_login", Type: field.TypeTime, Nullable: true},
		{Name: "verified_at", Type: field.TypeTime, Nullable: true},
		{Name: "provider", Type: field.TypeEnum, Enums: []string{"GOOGLE", "GITHUB", "EMAIL"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserSessionsColumns holds the columns for the "user_sessions" table.
	UserSessionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "token", Type: field.TypeString},
		{Name: "expires_at", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_user_sessions", Type: field.TypeUUID},
	}
	// UserSessionsTable holds the schema information for the "user_sessions" table.
	UserSessionsTable = &schema.Table{
		Name:       "user_sessions",
		Columns:    UserSessionsColumns,
		PrimaryKey: []*schema.Column{UserSessionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_sessions_users_user_sessions",
				Columns:    []*schema.Column{UserSessionsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserVerificationsColumns holds the columns for the "user_verifications" table.
	UserVerificationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "attempt_id", Type: field.TypeUUID},
		{Name: "code", Type: field.TypeString},
		{Name: "attempts", Type: field.TypeInt, Default: 0},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "user_user_verifications", Type: field.TypeUUID},
	}
	// UserVerificationsTable holds the schema information for the "user_verifications" table.
	UserVerificationsTable = &schema.Table{
		Name:       "user_verifications",
		Columns:    UserVerificationsColumns,
		PrimaryKey: []*schema.Column{UserVerificationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_verifications_users_user_verifications",
				Columns:    []*schema.Column{UserVerificationsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OidcClientsTable,
		ProjectsTable,
		ProjectDomainsTable,
		TokensTable,
		UsersTable,
		UserSessionsTable,
		UserVerificationsTable,
	}
)

func init() {
	OidcClientsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectsTable.ForeignKeys[0].RefTable = UsersTable
	ProjectDomainsTable.ForeignKeys[0].RefTable = ProjectsTable
	TokensTable.ForeignKeys[0].RefTable = UsersTable
	UserSessionsTable.ForeignKeys[0].RefTable = UsersTable
	UserVerificationsTable.ForeignKeys[0].RefTable = UsersTable
}
