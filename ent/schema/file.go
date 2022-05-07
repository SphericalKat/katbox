package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the Files.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("storage_key").NotEmpty().Unique(),
		field.Time("expires_at").Default(func() time.Time {
			return time.Now().Add(24 * 7 * time.Hour)
		}),
	}
}

// Edges of the Files.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("files").Unique().Required(),
	}
}
