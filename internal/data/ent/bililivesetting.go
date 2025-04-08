// Code generated by ent, DO NOT EDIT.

package ent

import (
	"echo/internal/data/ent/bililivesetting"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BiliLiveSetting is the model entity for the BiliLiveSetting schema.
type BiliLiveSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// RoomID holds the value of the "room_id" field.
	RoomID int64 `json:"room_id,omitempty"`
	// LiveState holds the value of the "live_state" field.
	LiveState int64 `json:"live_state,omitempty"`
	// LiveStartTime holds the value of the "live_start_time" field.
	LiveStartTime time.Time `json:"live_start_time,omitempty"`
	// LiveEndTime holds the value of the "live_end_time" field.
	LiveEndTime time.Time `json:"live_end_time,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime   time.Time `json:"update_time,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BiliLiveSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bililivesetting.FieldID, bililivesetting.FieldRoomID, bililivesetting.FieldLiveState:
			values[i] = new(sql.NullInt64)
		case bililivesetting.FieldLiveStartTime, bililivesetting.FieldLiveEndTime, bililivesetting.FieldCreateTime, bililivesetting.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BiliLiveSetting fields.
func (bls *BiliLiveSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bililivesetting.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bls.ID = int64(value.Int64)
		case bililivesetting.FieldRoomID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field room_id", values[i])
			} else if value.Valid {
				bls.RoomID = value.Int64
			}
		case bililivesetting.FieldLiveState:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field live_state", values[i])
			} else if value.Valid {
				bls.LiveState = value.Int64
			}
		case bililivesetting.FieldLiveStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field live_start_time", values[i])
			} else if value.Valid {
				bls.LiveStartTime = value.Time
			}
		case bililivesetting.FieldLiveEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field live_end_time", values[i])
			} else if value.Valid {
				bls.LiveEndTime = value.Time
			}
		case bililivesetting.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				bls.CreateTime = value.Time
			}
		case bililivesetting.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				bls.UpdateTime = value.Time
			}
		default:
			bls.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BiliLiveSetting.
// This includes values selected through modifiers, order, etc.
func (bls *BiliLiveSetting) Value(name string) (ent.Value, error) {
	return bls.selectValues.Get(name)
}

// Update returns a builder for updating this BiliLiveSetting.
// Note that you need to call BiliLiveSetting.Unwrap() before calling this method if this BiliLiveSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (bls *BiliLiveSetting) Update() *BiliLiveSettingUpdateOne {
	return NewBiliLiveSettingClient(bls.config).UpdateOne(bls)
}

// Unwrap unwraps the BiliLiveSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bls *BiliLiveSetting) Unwrap() *BiliLiveSetting {
	_tx, ok := bls.config.driver.(*txDriver)
	if !ok {
		panic("ent: BiliLiveSetting is not a transactional entity")
	}
	bls.config.driver = _tx.drv
	return bls
}

// String implements the fmt.Stringer.
func (bls *BiliLiveSetting) String() string {
	var builder strings.Builder
	builder.WriteString("BiliLiveSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bls.ID))
	builder.WriteString("room_id=")
	builder.WriteString(fmt.Sprintf("%v", bls.RoomID))
	builder.WriteString(", ")
	builder.WriteString("live_state=")
	builder.WriteString(fmt.Sprintf("%v", bls.LiveState))
	builder.WriteString(", ")
	builder.WriteString("live_start_time=")
	builder.WriteString(bls.LiveStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("live_end_time=")
	builder.WriteString(bls.LiveEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("create_time=")
	builder.WriteString(bls.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(bls.UpdateTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BiliLiveSettings is a parsable slice of BiliLiveSetting.
type BiliLiveSettings []*BiliLiveSetting
