package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Sub holds the schema definition for the Sub entity.
type Sub struct {
	ent.Schema
}

// Fields of the Sub.
func (Sub) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("sub_type"),
		field.Int64("group_id"),
		field.Int64("sub_id"),
		field.Int64("bot_id"),
		field.Int("status").Default(1),
		field.Time("create_time"),
		field.Time("update_time"),
	}
}

// Edges of the Sub.
func (Sub) Edges() []ent.Edge {
	return nil
}
