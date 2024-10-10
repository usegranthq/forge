package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("email").NotEmpty().Unique(),
		field.Time("last_login").Optional(),
		field.Time("verified_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_sessions", UserSession.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("projects", Project.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("user_verifications", UserVerification.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
