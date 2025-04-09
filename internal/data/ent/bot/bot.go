// Code generated by ent, DO NOT EDIT.

package bot

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the bot type in the database.
	Label = "bot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBotID holds the string denoting the bot_id field in the database.
	FieldBotID = "bot_id"
	// FieldBotName holds the string denoting the bot_name field in the database.
	FieldBotName = "bot_name"
	// FieldSelfID holds the string denoting the self_id field in the database.
	FieldSelfID = "self_id"
	// FieldAccessToken holds the string denoting the access_token field in the database.
	FieldAccessToken = "access_token"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastOnlineTime holds the string denoting the last_online_time field in the database.
	FieldLastOnlineTime = "last_online_time"
	// FieldLastOnlineIP holds the string denoting the last_online_ip field in the database.
	FieldLastOnlineIP = "last_online_ip"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// Table holds the table name of the bot in the database.
	Table = "bots"
)

// Columns holds all SQL columns for bot fields.
var Columns = []string{
	FieldID,
	FieldBotID,
	FieldBotName,
	FieldSelfID,
	FieldAccessToken,
	FieldStatus,
	FieldLastOnlineTime,
	FieldLastOnlineIP,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int
)

// OrderOption defines the ordering options for the Bot queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBotID orders the results by the bot_id field.
func ByBotID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBotID, opts...).ToFunc()
}

// ByBotName orders the results by the bot_name field.
func ByBotName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBotName, opts...).ToFunc()
}

// BySelfID orders the results by the self_id field.
func BySelfID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSelfID, opts...).ToFunc()
}

// ByAccessToken orders the results by the access_token field.
func ByAccessToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessToken, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastOnlineTime orders the results by the last_online_time field.
func ByLastOnlineTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastOnlineTime, opts...).ToFunc()
}

// ByLastOnlineIP orders the results by the last_online_ip field.
func ByLastOnlineIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastOnlineIP, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}
