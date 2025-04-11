package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SubDouyuLive holds the schema definition for the Sub_douyu_live entity.
type SubDouyuLive struct {
	ent.Schema
}

func (SubDouyuLive) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sub_douyu_live"},
	}
}

// Fields of the SubDouyuLive.
func (SubDouyuLive) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("room_id"),
		field.Int64("live_state"),
		field.Int64("live_start_time"),
		field.Int64("live_end_time"),
		field.Time("create_time").Optional(),
		field.Time("update_time").Optional(),
	}
}

// Edges of the SubDouyuLive.
func (SubDouyuLive) Edges() []ent.Edge {
	return nil
}
