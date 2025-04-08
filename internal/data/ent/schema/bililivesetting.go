package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// BiliLiveSetting holds the schema definition for the BiliLiveSetting entity.
type BiliLiveSetting struct {
	ent.Schema
}

func (BiliLiveSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "bili_live_setting"},
	}
}

// Fields of the BiliLiveSetting.
func (BiliLiveSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("room_id"),
		field.Int64("live_state"),
		field.Time("live_start_time"),
		field.Time("live_end_time"),
		field.Time("create_time"),
		field.Time("update_time"),
	}
}

// Edges of the BiliLiveSetting.
func (BiliLiveSetting) Edges() []ent.Edge {
	return nil
}
