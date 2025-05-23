// Code generated by ent, DO NOT EDIT.

package bot

import (
	"echo/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldID, id))
}

// BotID applies equality check predicate on the "bot_id" field. It's identical to BotIDEQ.
func BotID(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldBotID, v))
}

// BotName applies equality check predicate on the "bot_name" field. It's identical to BotNameEQ.
func BotName(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldBotName, v))
}

// SelfID applies equality check predicate on the "self_id" field. It's identical to SelfIDEQ.
func SelfID(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldSelfID, v))
}

// AccessToken applies equality check predicate on the "access_token" field. It's identical to AccessTokenEQ.
func AccessToken(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldAccessToken, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldStatus, v))
}

// LastOnlineTime applies equality check predicate on the "last_online_time" field. It's identical to LastOnlineTimeEQ.
func LastOnlineTime(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldLastOnlineTime, v))
}

// LastOnlineIP applies equality check predicate on the "last_online_ip" field. It's identical to LastOnlineIPEQ.
func LastOnlineIP(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldLastOnlineIP, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldUpdateTime, v))
}

// BotIDEQ applies the EQ predicate on the "bot_id" field.
func BotIDEQ(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldBotID, v))
}

// BotIDNEQ applies the NEQ predicate on the "bot_id" field.
func BotIDNEQ(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldBotID, v))
}

// BotIDIn applies the In predicate on the "bot_id" field.
func BotIDIn(vs ...int64) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldBotID, vs...))
}

// BotIDNotIn applies the NotIn predicate on the "bot_id" field.
func BotIDNotIn(vs ...int64) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldBotID, vs...))
}

// BotIDGT applies the GT predicate on the "bot_id" field.
func BotIDGT(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldBotID, v))
}

// BotIDGTE applies the GTE predicate on the "bot_id" field.
func BotIDGTE(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldBotID, v))
}

// BotIDLT applies the LT predicate on the "bot_id" field.
func BotIDLT(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldBotID, v))
}

// BotIDLTE applies the LTE predicate on the "bot_id" field.
func BotIDLTE(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldBotID, v))
}

// BotNameEQ applies the EQ predicate on the "bot_name" field.
func BotNameEQ(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldBotName, v))
}

// BotNameNEQ applies the NEQ predicate on the "bot_name" field.
func BotNameNEQ(v string) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldBotName, v))
}

// BotNameIn applies the In predicate on the "bot_name" field.
func BotNameIn(vs ...string) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldBotName, vs...))
}

// BotNameNotIn applies the NotIn predicate on the "bot_name" field.
func BotNameNotIn(vs ...string) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldBotName, vs...))
}

// BotNameGT applies the GT predicate on the "bot_name" field.
func BotNameGT(v string) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldBotName, v))
}

// BotNameGTE applies the GTE predicate on the "bot_name" field.
func BotNameGTE(v string) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldBotName, v))
}

// BotNameLT applies the LT predicate on the "bot_name" field.
func BotNameLT(v string) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldBotName, v))
}

// BotNameLTE applies the LTE predicate on the "bot_name" field.
func BotNameLTE(v string) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldBotName, v))
}

// BotNameContains applies the Contains predicate on the "bot_name" field.
func BotNameContains(v string) predicate.Bot {
	return predicate.Bot(sql.FieldContains(FieldBotName, v))
}

// BotNameHasPrefix applies the HasPrefix predicate on the "bot_name" field.
func BotNameHasPrefix(v string) predicate.Bot {
	return predicate.Bot(sql.FieldHasPrefix(FieldBotName, v))
}

// BotNameHasSuffix applies the HasSuffix predicate on the "bot_name" field.
func BotNameHasSuffix(v string) predicate.Bot {
	return predicate.Bot(sql.FieldHasSuffix(FieldBotName, v))
}

// BotNameEqualFold applies the EqualFold predicate on the "bot_name" field.
func BotNameEqualFold(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEqualFold(FieldBotName, v))
}

// BotNameContainsFold applies the ContainsFold predicate on the "bot_name" field.
func BotNameContainsFold(v string) predicate.Bot {
	return predicate.Bot(sql.FieldContainsFold(FieldBotName, v))
}

// SelfIDEQ applies the EQ predicate on the "self_id" field.
func SelfIDEQ(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldSelfID, v))
}

// SelfIDNEQ applies the NEQ predicate on the "self_id" field.
func SelfIDNEQ(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldSelfID, v))
}

// SelfIDIn applies the In predicate on the "self_id" field.
func SelfIDIn(vs ...int64) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldSelfID, vs...))
}

// SelfIDNotIn applies the NotIn predicate on the "self_id" field.
func SelfIDNotIn(vs ...int64) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldSelfID, vs...))
}

// SelfIDGT applies the GT predicate on the "self_id" field.
func SelfIDGT(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldSelfID, v))
}

// SelfIDGTE applies the GTE predicate on the "self_id" field.
func SelfIDGTE(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldSelfID, v))
}

// SelfIDLT applies the LT predicate on the "self_id" field.
func SelfIDLT(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldSelfID, v))
}

// SelfIDLTE applies the LTE predicate on the "self_id" field.
func SelfIDLTE(v int64) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldSelfID, v))
}

// AccessTokenEQ applies the EQ predicate on the "access_token" field.
func AccessTokenEQ(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldAccessToken, v))
}

// AccessTokenNEQ applies the NEQ predicate on the "access_token" field.
func AccessTokenNEQ(v string) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldAccessToken, v))
}

// AccessTokenIn applies the In predicate on the "access_token" field.
func AccessTokenIn(vs ...string) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldAccessToken, vs...))
}

// AccessTokenNotIn applies the NotIn predicate on the "access_token" field.
func AccessTokenNotIn(vs ...string) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldAccessToken, vs...))
}

// AccessTokenGT applies the GT predicate on the "access_token" field.
func AccessTokenGT(v string) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldAccessToken, v))
}

// AccessTokenGTE applies the GTE predicate on the "access_token" field.
func AccessTokenGTE(v string) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldAccessToken, v))
}

// AccessTokenLT applies the LT predicate on the "access_token" field.
func AccessTokenLT(v string) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldAccessToken, v))
}

// AccessTokenLTE applies the LTE predicate on the "access_token" field.
func AccessTokenLTE(v string) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldAccessToken, v))
}

// AccessTokenContains applies the Contains predicate on the "access_token" field.
func AccessTokenContains(v string) predicate.Bot {
	return predicate.Bot(sql.FieldContains(FieldAccessToken, v))
}

// AccessTokenHasPrefix applies the HasPrefix predicate on the "access_token" field.
func AccessTokenHasPrefix(v string) predicate.Bot {
	return predicate.Bot(sql.FieldHasPrefix(FieldAccessToken, v))
}

// AccessTokenHasSuffix applies the HasSuffix predicate on the "access_token" field.
func AccessTokenHasSuffix(v string) predicate.Bot {
	return predicate.Bot(sql.FieldHasSuffix(FieldAccessToken, v))
}

// AccessTokenEqualFold applies the EqualFold predicate on the "access_token" field.
func AccessTokenEqualFold(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEqualFold(FieldAccessToken, v))
}

// AccessTokenContainsFold applies the ContainsFold predicate on the "access_token" field.
func AccessTokenContainsFold(v string) predicate.Bot {
	return predicate.Bot(sql.FieldContainsFold(FieldAccessToken, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldStatus, v))
}

// LastOnlineTimeEQ applies the EQ predicate on the "last_online_time" field.
func LastOnlineTimeEQ(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldLastOnlineTime, v))
}

// LastOnlineTimeNEQ applies the NEQ predicate on the "last_online_time" field.
func LastOnlineTimeNEQ(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldLastOnlineTime, v))
}

// LastOnlineTimeIn applies the In predicate on the "last_online_time" field.
func LastOnlineTimeIn(vs ...time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldLastOnlineTime, vs...))
}

// LastOnlineTimeNotIn applies the NotIn predicate on the "last_online_time" field.
func LastOnlineTimeNotIn(vs ...time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldLastOnlineTime, vs...))
}

// LastOnlineTimeGT applies the GT predicate on the "last_online_time" field.
func LastOnlineTimeGT(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldLastOnlineTime, v))
}

// LastOnlineTimeGTE applies the GTE predicate on the "last_online_time" field.
func LastOnlineTimeGTE(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldLastOnlineTime, v))
}

// LastOnlineTimeLT applies the LT predicate on the "last_online_time" field.
func LastOnlineTimeLT(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldLastOnlineTime, v))
}

// LastOnlineTimeLTE applies the LTE predicate on the "last_online_time" field.
func LastOnlineTimeLTE(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldLastOnlineTime, v))
}

// LastOnlineIPEQ applies the EQ predicate on the "last_online_ip" field.
func LastOnlineIPEQ(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldLastOnlineIP, v))
}

// LastOnlineIPNEQ applies the NEQ predicate on the "last_online_ip" field.
func LastOnlineIPNEQ(v string) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldLastOnlineIP, v))
}

// LastOnlineIPIn applies the In predicate on the "last_online_ip" field.
func LastOnlineIPIn(vs ...string) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldLastOnlineIP, vs...))
}

// LastOnlineIPNotIn applies the NotIn predicate on the "last_online_ip" field.
func LastOnlineIPNotIn(vs ...string) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldLastOnlineIP, vs...))
}

// LastOnlineIPGT applies the GT predicate on the "last_online_ip" field.
func LastOnlineIPGT(v string) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldLastOnlineIP, v))
}

// LastOnlineIPGTE applies the GTE predicate on the "last_online_ip" field.
func LastOnlineIPGTE(v string) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldLastOnlineIP, v))
}

// LastOnlineIPLT applies the LT predicate on the "last_online_ip" field.
func LastOnlineIPLT(v string) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldLastOnlineIP, v))
}

// LastOnlineIPLTE applies the LTE predicate on the "last_online_ip" field.
func LastOnlineIPLTE(v string) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldLastOnlineIP, v))
}

// LastOnlineIPContains applies the Contains predicate on the "last_online_ip" field.
func LastOnlineIPContains(v string) predicate.Bot {
	return predicate.Bot(sql.FieldContains(FieldLastOnlineIP, v))
}

// LastOnlineIPHasPrefix applies the HasPrefix predicate on the "last_online_ip" field.
func LastOnlineIPHasPrefix(v string) predicate.Bot {
	return predicate.Bot(sql.FieldHasPrefix(FieldLastOnlineIP, v))
}

// LastOnlineIPHasSuffix applies the HasSuffix predicate on the "last_online_ip" field.
func LastOnlineIPHasSuffix(v string) predicate.Bot {
	return predicate.Bot(sql.FieldHasSuffix(FieldLastOnlineIP, v))
}

// LastOnlineIPEqualFold applies the EqualFold predicate on the "last_online_ip" field.
func LastOnlineIPEqualFold(v string) predicate.Bot {
	return predicate.Bot(sql.FieldEqualFold(FieldLastOnlineIP, v))
}

// LastOnlineIPContainsFold applies the ContainsFold predicate on the "last_online_ip" field.
func LastOnlineIPContainsFold(v string) predicate.Bot {
	return predicate.Bot(sql.FieldContainsFold(FieldLastOnlineIP, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Bot {
	return predicate.Bot(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Bot {
	return predicate.Bot(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Bot {
	return predicate.Bot(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Bot {
	return predicate.Bot(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Bot {
	return predicate.Bot(sql.FieldNotNull(FieldUpdateTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Bot) predicate.Bot {
	return predicate.Bot(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Bot) predicate.Bot {
	return predicate.Bot(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Bot) predicate.Bot {
	return predicate.Bot(sql.NotPredicates(p))
}
