package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserVerification holds the schema definition for the UserVerification entity.
type UserVerification struct {
	ent.Schema
}

// Fields of the UserVerification.
func (UserVerification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("attempt_id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("code").NotEmpty(),
		field.Int("attempts").Default(0),
		field.Time("expires_at"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserVerification.
func (UserVerification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_verifications").
			Unique().
			Required(),
	}
}
