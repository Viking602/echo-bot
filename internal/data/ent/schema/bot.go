package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Bot holds the schema definition for the Bot entity.
type Bot struct {
	ent.Schema
}

// Fields of the Bot.
func (Bot) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.String("bot_name"),
		field.Int("status").Default(1),
		field.Time("create_time"),
		field.Time("update_time"),
	}
}

// Edges of the Bot.
func (Bot) Edges() []ent.Edge {
	return nil
}
