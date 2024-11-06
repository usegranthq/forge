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
	"github.com/usegranthq/forge/ent/usersession"
)

// UserSession is the model entity for the UserSession schema.
type UserSession struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSessionQuery when eager-loading is set.
	Edges              UserSessionEdges `json:"edges"`
	user_user_sessions *uuid.UUID
	selectValues       sql.SelectValues
}

// UserSessionEdges holds the relations/edges for other nodes in the graph.
type UserSessionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSessionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSession) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersession.FieldToken:
			values[i] = new(sql.NullString)
		case usersession.FieldExpiresAt, usersession.FieldCreatedAt, usersession.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case usersession.FieldID:
			values[i] = new(uuid.UUID)
		case usersession.ForeignKeys[0]: // user_user_sessions
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSession fields.
func (us *UserSession) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersession.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				us.ID = *value
			}
		case usersession.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				us.Token = value.String
			}
		case usersession.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				us.ExpiresAt = value.Time
			}
		case usersession.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				us.CreatedAt = value.Time
			}
		case usersession.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				us.UpdatedAt = value.Time
			}
		case usersession.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_user_sessions", values[i])
			} else if value.Valid {
				us.user_user_sessions = new(uuid.UUID)
				*us.user_user_sessions = *value.S.(*uuid.UUID)
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSession.
// This includes values selected through modifiers, order, etc.
func (us *UserSession) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserSession entity.
func (us *UserSession) QueryUser() *UserQuery {
	return NewUserSessionClient(us.config).QueryUser(us)
}

// Update returns a builder for updating this UserSession.
// Note that you need to call UserSession.Unwrap() before calling this method if this UserSession
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSession) Update() *UserSessionUpdateOne {
	return NewUserSessionClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSession entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSession) Unwrap() *UserSession {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserSession is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSession) String() string {
	var builder strings.Builder
	builder.WriteString("UserSession(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("token=")
	builder.WriteString(us.Token)
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(us.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(us.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(us.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// UserSessions is a parsable slice of UserSession.
type UserSessions []*UserSession
