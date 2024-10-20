package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProjectDomain holds the schema definition for the ProjectDomain entity.
type ProjectDomain struct {
	ent.Schema
}

// Fields of the ProjectDomain.
func (ProjectDomain) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("domain").NotEmpty(),
		field.Bool("verified").Default(false),
		field.String("verified_at").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ProjectDomain.
func (ProjectDomain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("domain").
			Unique().
			Required(),
	}
}
