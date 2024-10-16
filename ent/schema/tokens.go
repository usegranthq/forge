package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Tokens holds the schema definition for the Tokens entity.
type Token struct {
	ent.Schema
}

// Fields of the Tokens.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name").NotEmpty(),
		field.String("token").NotEmpty().Sensitive(),
		field.Time("expires_at").Optional(),
		field.Time("last_used_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Tokens.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Unique(),
	}
}
