package schema

import "entgo.io/ent"

// Files holds the schema definition for the Files entity.
type Files struct {
	ent.Schema
}

// Fields of the Files.
func (Files) Fields() []ent.Field {
	return nil
}

// Edges of the Files.
func (Files) Edges() []ent.Edge {
	return nil
}
