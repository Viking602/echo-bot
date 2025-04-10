// Code generated by ent, DO NOT EDIT.

package sub

import (
	"echo/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldID, id))
}

// SubType applies equality check predicate on the "sub_type" field. It's identical to SubTypeEQ.
func SubType(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldSubType, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldGroupID, v))
}

// SubID applies equality check predicate on the "sub_id" field. It's identical to SubIDEQ.
func SubID(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldSubID, v))
}

// BotID applies equality check predicate on the "bot_id" field. It's identical to BotIDEQ.
func BotID(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldBotID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldStatus, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldUpdateTime, v))
}

// SubTypeEQ applies the EQ predicate on the "sub_type" field.
func SubTypeEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldSubType, v))
}

// SubTypeNEQ applies the NEQ predicate on the "sub_type" field.
func SubTypeNEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldSubType, v))
}

// SubTypeIn applies the In predicate on the "sub_type" field.
func SubTypeIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldSubType, vs...))
}

// SubTypeNotIn applies the NotIn predicate on the "sub_type" field.
func SubTypeNotIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldSubType, vs...))
}

// SubTypeGT applies the GT predicate on the "sub_type" field.
func SubTypeGT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldSubType, v))
}

// SubTypeGTE applies the GTE predicate on the "sub_type" field.
func SubTypeGTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldSubType, v))
}

// SubTypeLT applies the LT predicate on the "sub_type" field.
func SubTypeLT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldSubType, v))
}

// SubTypeLTE applies the LTE predicate on the "sub_type" field.
func SubTypeLTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldSubType, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldGroupID, v))
}

// SubIDEQ applies the EQ predicate on the "sub_id" field.
func SubIDEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldSubID, v))
}

// SubIDNEQ applies the NEQ predicate on the "sub_id" field.
func SubIDNEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldSubID, v))
}

// SubIDIn applies the In predicate on the "sub_id" field.
func SubIDIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldSubID, vs...))
}

// SubIDNotIn applies the NotIn predicate on the "sub_id" field.
func SubIDNotIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldSubID, vs...))
}

// SubIDGT applies the GT predicate on the "sub_id" field.
func SubIDGT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldSubID, v))
}

// SubIDGTE applies the GTE predicate on the "sub_id" field.
func SubIDGTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldSubID, v))
}

// SubIDLT applies the LT predicate on the "sub_id" field.
func SubIDLT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldSubID, v))
}

// SubIDLTE applies the LTE predicate on the "sub_id" field.
func SubIDLTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldSubID, v))
}

// BotIDEQ applies the EQ predicate on the "bot_id" field.
func BotIDEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldBotID, v))
}

// BotIDNEQ applies the NEQ predicate on the "bot_id" field.
func BotIDNEQ(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldBotID, v))
}

// BotIDIn applies the In predicate on the "bot_id" field.
func BotIDIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldBotID, vs...))
}

// BotIDNotIn applies the NotIn predicate on the "bot_id" field.
func BotIDNotIn(vs ...int64) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldBotID, vs...))
}

// BotIDGT applies the GT predicate on the "bot_id" field.
func BotIDGT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldBotID, v))
}

// BotIDGTE applies the GTE predicate on the "bot_id" field.
func BotIDGTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldBotID, v))
}

// BotIDLT applies the LT predicate on the "bot_id" field.
func BotIDLT(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldBotID, v))
}

// BotIDLTE applies the LTE predicate on the "bot_id" field.
func BotIDLTE(v int64) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldBotID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldStatus, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Sub {
	return predicate.Sub(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Sub {
	return predicate.Sub(sql.FieldNotNull(FieldCreateTime))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Sub {
	return predicate.Sub(sql.FieldLTE(FieldUpdateTime, v))
}

// UpdateTimeIsNil applies the IsNil predicate on the "update_time" field.
func UpdateTimeIsNil() predicate.Sub {
	return predicate.Sub(sql.FieldIsNull(FieldUpdateTime))
}

// UpdateTimeNotNil applies the NotNil predicate on the "update_time" field.
func UpdateTimeNotNil() predicate.Sub {
	return predicate.Sub(sql.FieldNotNull(FieldUpdateTime))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sub) predicate.Sub {
	return predicate.Sub(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sub) predicate.Sub {
	return predicate.Sub(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sub) predicate.Sub {
	return predicate.Sub(sql.NotPredicates(p))
}
