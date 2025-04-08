// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BiliLiveSettingColumns holds the columns for the "bili_live_setting" table.
	BiliLiveSettingColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "room_id", Type: field.TypeInt64},
		{Name: "live_state", Type: field.TypeInt64},
		{Name: "live_start_time", Type: field.TypeTime},
		{Name: "live_end_time", Type: field.TypeTime},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// BiliLiveSettingTable holds the schema information for the "bili_live_setting" table.
	BiliLiveSettingTable = &schema.Table{
		Name:       "bili_live_setting",
		Columns:    BiliLiveSettingColumns,
		PrimaryKey: []*schema.Column{BiliLiveSettingColumns[0]},
	}
	// BotsColumns holds the columns for the "bots" table.
	BotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "bot_name", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt, Default: 1},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// BotsTable holds the schema information for the "bots" table.
	BotsTable = &schema.Table{
		Name:       "bots",
		Columns:    BotsColumns,
		PrimaryKey: []*schema.Column{BotsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BiliLiveSettingTable,
		BotsTable,
	}
)

func init() {
	BiliLiveSettingTable.Annotation = &entsql.Annotation{
		Table: "bili_live_setting",
	}
}
