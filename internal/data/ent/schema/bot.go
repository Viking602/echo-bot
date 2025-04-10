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
		field.Int64("bot_id"),
		field.String("bot_name"),
		field.Int64("self_id"),
		field.String("access_token"),
		field.Int("status").Default(1),
		field.Time("last_online_time"),
		field.String("last_online_ip"),
		field.Time("create_time").Optional(),
		field.Time("update_time").Optional(),
	}
}

// Edges of the Bot.
func (Bot) Edges() []ent.Edge {
	return nil
}
