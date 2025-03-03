// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/usegranthq/forge/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID uuid.UUID `json:"uid,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// LastLogin holds the value of the "last_login" field.
	LastLogin time.Time `json:"last_login,omitempty"`
	// VerifiedAt holds the value of the "verified_at" field.
	VerifiedAt time.Time `json:"verified_at,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider user.Provider `json:"provider,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserSessions holds the value of the user_sessions edge.
	UserSessions []*UserSession `json:"user_sessions,omitempty"`
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// Verifications holds the value of the verifications edge.
	Verifications []*Verification `json:"verifications,omitempty"`
	// Tokens holds the value of the tokens edge.
	Tokens []*Token `json:"tokens,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// UserSessionsOrErr returns the UserSessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserSessionsOrErr() ([]*UserSession, error) {
	if e.loadedTypes[0] {
		return e.UserSessions, nil
	}
	return nil, &NotLoadedError{edge: "user_sessions"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// VerificationsOrErr returns the Verifications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VerificationsOrErr() ([]*Verification, error) {
	if e.loadedTypes[2] {
		return e.Verifications, nil
	}
	return nil, &NotLoadedError{edge: "verifications"}
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[3] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldEmail, user.FieldProvider:
			values[i] = new(sql.NullString)
		case user.FieldLastLogin, user.FieldVerifiedAt, user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case user.FieldID, user.FieldUID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value != nil {
				u.UID = *value
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldLastLogin:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login", values[i])
			} else if value.Valid {
				u.LastLogin = value.Time
			}
		case user.FieldVerifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field verified_at", values[i])
			} else if value.Valid {
				u.VerifiedAt = value.Time
			}
		case user.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				u.Provider = user.Provider(value.String)
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserSessions queries the "user_sessions" edge of the User entity.
func (u *User) QueryUserSessions() *UserSessionQuery {
	return NewUserClient(u.config).QueryUserSessions(u)
}

// QueryProjects queries the "projects" edge of the User entity.
func (u *User) QueryProjects() *ProjectQuery {
	return NewUserClient(u.config).QueryProjects(u)
}

// QueryVerifications queries the "verifications" edge of the User entity.
func (u *User) QueryVerifications() *VerificationQuery {
	return NewUserClient(u.config).QueryVerifications(u)
}

// QueryTokens queries the "tokens" edge of the User entity.
func (u *User) QueryTokens() *TokenQuery {
	return NewUserClient(u.config).QueryTokens(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("uid=")
	builder.WriteString(fmt.Sprintf("%v", u.UID))
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("last_login=")
	builder.WriteString(u.LastLogin.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("verified_at=")
	builder.WriteString(u.VerifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("provider=")
	builder.WriteString(fmt.Sprintf("%v", u.Provider))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
