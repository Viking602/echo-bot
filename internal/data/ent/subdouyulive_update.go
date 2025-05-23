// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/predicate"
	"echo/internal/data/ent/subdouyulive"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubDouyuLiveUpdate is the builder for updating SubDouyuLive entities.
type SubDouyuLiveUpdate struct {
	config
	hooks    []Hook
	mutation *SubDouyuLiveMutation
}

// Where appends a list predicates to the SubDouyuLiveUpdate builder.
func (sdlu *SubDouyuLiveUpdate) Where(ps ...predicate.SubDouyuLive) *SubDouyuLiveUpdate {
	sdlu.mutation.Where(ps...)
	return sdlu
}

// SetRoomID sets the "room_id" field.
func (sdlu *SubDouyuLiveUpdate) SetRoomID(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.ResetRoomID()
	sdlu.mutation.SetRoomID(i)
	return sdlu
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (sdlu *SubDouyuLiveUpdate) SetNillableRoomID(i *int64) *SubDouyuLiveUpdate {
	if i != nil {
		sdlu.SetRoomID(*i)
	}
	return sdlu
}

// AddRoomID adds i to the "room_id" field.
func (sdlu *SubDouyuLiveUpdate) AddRoomID(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.AddRoomID(i)
	return sdlu
}

// SetLiveState sets the "live_state" field.
func (sdlu *SubDouyuLiveUpdate) SetLiveState(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.ResetLiveState()
	sdlu.mutation.SetLiveState(i)
	return sdlu
}

// SetNillableLiveState sets the "live_state" field if the given value is not nil.
func (sdlu *SubDouyuLiveUpdate) SetNillableLiveState(i *int64) *SubDouyuLiveUpdate {
	if i != nil {
		sdlu.SetLiveState(*i)
	}
	return sdlu
}

// AddLiveState adds i to the "live_state" field.
func (sdlu *SubDouyuLiveUpdate) AddLiveState(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.AddLiveState(i)
	return sdlu
}

// SetLiveStartTime sets the "live_start_time" field.
func (sdlu *SubDouyuLiveUpdate) SetLiveStartTime(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.ResetLiveStartTime()
	sdlu.mutation.SetLiveStartTime(i)
	return sdlu
}

// SetNillableLiveStartTime sets the "live_start_time" field if the given value is not nil.
func (sdlu *SubDouyuLiveUpdate) SetNillableLiveStartTime(i *int64) *SubDouyuLiveUpdate {
	if i != nil {
		sdlu.SetLiveStartTime(*i)
	}
	return sdlu
}

// AddLiveStartTime adds i to the "live_start_time" field.
func (sdlu *SubDouyuLiveUpdate) AddLiveStartTime(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.AddLiveStartTime(i)
	return sdlu
}

// SetLiveEndTime sets the "live_end_time" field.
func (sdlu *SubDouyuLiveUpdate) SetLiveEndTime(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.ResetLiveEndTime()
	sdlu.mutation.SetLiveEndTime(i)
	return sdlu
}

// SetNillableLiveEndTime sets the "live_end_time" field if the given value is not nil.
func (sdlu *SubDouyuLiveUpdate) SetNillableLiveEndTime(i *int64) *SubDouyuLiveUpdate {
	if i != nil {
		sdlu.SetLiveEndTime(*i)
	}
	return sdlu
}

// AddLiveEndTime adds i to the "live_end_time" field.
func (sdlu *SubDouyuLiveUpdate) AddLiveEndTime(i int64) *SubDouyuLiveUpdate {
	sdlu.mutation.AddLiveEndTime(i)
	return sdlu
}

// SetCreateTime sets the "create_time" field.
func (sdlu *SubDouyuLiveUpdate) SetCreateTime(t time.Time) *SubDouyuLiveUpdate {
	sdlu.mutation.SetCreateTime(t)
	return sdlu
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sdlu *SubDouyuLiveUpdate) SetNillableCreateTime(t *time.Time) *SubDouyuLiveUpdate {
	if t != nil {
		sdlu.SetCreateTime(*t)
	}
	return sdlu
}

// ClearCreateTime clears the value of the "create_time" field.
func (sdlu *SubDouyuLiveUpdate) ClearCreateTime() *SubDouyuLiveUpdate {
	sdlu.mutation.ClearCreateTime()
	return sdlu
}

// SetUpdateTime sets the "update_time" field.
func (sdlu *SubDouyuLiveUpdate) SetUpdateTime(t time.Time) *SubDouyuLiveUpdate {
	sdlu.mutation.SetUpdateTime(t)
	return sdlu
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sdlu *SubDouyuLiveUpdate) SetNillableUpdateTime(t *time.Time) *SubDouyuLiveUpdate {
	if t != nil {
		sdlu.SetUpdateTime(*t)
	}
	return sdlu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (sdlu *SubDouyuLiveUpdate) ClearUpdateTime() *SubDouyuLiveUpdate {
	sdlu.mutation.ClearUpdateTime()
	return sdlu
}

// Mutation returns the SubDouyuLiveMutation object of the builder.
func (sdlu *SubDouyuLiveUpdate) Mutation() *SubDouyuLiveMutation {
	return sdlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sdlu *SubDouyuLiveUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sdlu.sqlSave, sdlu.mutation, sdlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sdlu *SubDouyuLiveUpdate) SaveX(ctx context.Context) int {
	affected, err := sdlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sdlu *SubDouyuLiveUpdate) Exec(ctx context.Context) error {
	_, err := sdlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdlu *SubDouyuLiveUpdate) ExecX(ctx context.Context) {
	if err := sdlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sdlu *SubDouyuLiveUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(subdouyulive.Table, subdouyulive.Columns, sqlgraph.NewFieldSpec(subdouyulive.FieldID, field.TypeInt64))
	if ps := sdlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdlu.mutation.RoomID(); ok {
		_spec.SetField(subdouyulive.FieldRoomID, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.AddedRoomID(); ok {
		_spec.AddField(subdouyulive.FieldRoomID, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.LiveState(); ok {
		_spec.SetField(subdouyulive.FieldLiveState, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.AddedLiveState(); ok {
		_spec.AddField(subdouyulive.FieldLiveState, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.LiveStartTime(); ok {
		_spec.SetField(subdouyulive.FieldLiveStartTime, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.AddedLiveStartTime(); ok {
		_spec.AddField(subdouyulive.FieldLiveStartTime, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.LiveEndTime(); ok {
		_spec.SetField(subdouyulive.FieldLiveEndTime, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.AddedLiveEndTime(); ok {
		_spec.AddField(subdouyulive.FieldLiveEndTime, field.TypeInt64, value)
	}
	if value, ok := sdlu.mutation.CreateTime(); ok {
		_spec.SetField(subdouyulive.FieldCreateTime, field.TypeTime, value)
	}
	if sdlu.mutation.CreateTimeCleared() {
		_spec.ClearField(subdouyulive.FieldCreateTime, field.TypeTime)
	}
	if value, ok := sdlu.mutation.UpdateTime(); ok {
		_spec.SetField(subdouyulive.FieldUpdateTime, field.TypeTime, value)
	}
	if sdlu.mutation.UpdateTimeCleared() {
		_spec.ClearField(subdouyulive.FieldUpdateTime, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sdlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subdouyulive.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sdlu.mutation.done = true
	return n, nil
}

// SubDouyuLiveUpdateOne is the builder for updating a single SubDouyuLive entity.
type SubDouyuLiveUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubDouyuLiveMutation
}

// SetRoomID sets the "room_id" field.
func (sdluo *SubDouyuLiveUpdateOne) SetRoomID(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.ResetRoomID()
	sdluo.mutation.SetRoomID(i)
	return sdluo
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (sdluo *SubDouyuLiveUpdateOne) SetNillableRoomID(i *int64) *SubDouyuLiveUpdateOne {
	if i != nil {
		sdluo.SetRoomID(*i)
	}
	return sdluo
}

// AddRoomID adds i to the "room_id" field.
func (sdluo *SubDouyuLiveUpdateOne) AddRoomID(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.AddRoomID(i)
	return sdluo
}

// SetLiveState sets the "live_state" field.
func (sdluo *SubDouyuLiveUpdateOne) SetLiveState(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.ResetLiveState()
	sdluo.mutation.SetLiveState(i)
	return sdluo
}

// SetNillableLiveState sets the "live_state" field if the given value is not nil.
func (sdluo *SubDouyuLiveUpdateOne) SetNillableLiveState(i *int64) *SubDouyuLiveUpdateOne {
	if i != nil {
		sdluo.SetLiveState(*i)
	}
	return sdluo
}

// AddLiveState adds i to the "live_state" field.
func (sdluo *SubDouyuLiveUpdateOne) AddLiveState(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.AddLiveState(i)
	return sdluo
}

// SetLiveStartTime sets the "live_start_time" field.
func (sdluo *SubDouyuLiveUpdateOne) SetLiveStartTime(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.ResetLiveStartTime()
	sdluo.mutation.SetLiveStartTime(i)
	return sdluo
}

// SetNillableLiveStartTime sets the "live_start_time" field if the given value is not nil.
func (sdluo *SubDouyuLiveUpdateOne) SetNillableLiveStartTime(i *int64) *SubDouyuLiveUpdateOne {
	if i != nil {
		sdluo.SetLiveStartTime(*i)
	}
	return sdluo
}

// AddLiveStartTime adds i to the "live_start_time" field.
func (sdluo *SubDouyuLiveUpdateOne) AddLiveStartTime(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.AddLiveStartTime(i)
	return sdluo
}

// SetLiveEndTime sets the "live_end_time" field.
func (sdluo *SubDouyuLiveUpdateOne) SetLiveEndTime(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.ResetLiveEndTime()
	sdluo.mutation.SetLiveEndTime(i)
	return sdluo
}

// SetNillableLiveEndTime sets the "live_end_time" field if the given value is not nil.
func (sdluo *SubDouyuLiveUpdateOne) SetNillableLiveEndTime(i *int64) *SubDouyuLiveUpdateOne {
	if i != nil {
		sdluo.SetLiveEndTime(*i)
	}
	return sdluo
}

// AddLiveEndTime adds i to the "live_end_time" field.
func (sdluo *SubDouyuLiveUpdateOne) AddLiveEndTime(i int64) *SubDouyuLiveUpdateOne {
	sdluo.mutation.AddLiveEndTime(i)
	return sdluo
}

// SetCreateTime sets the "create_time" field.
func (sdluo *SubDouyuLiveUpdateOne) SetCreateTime(t time.Time) *SubDouyuLiveUpdateOne {
	sdluo.mutation.SetCreateTime(t)
	return sdluo
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sdluo *SubDouyuLiveUpdateOne) SetNillableCreateTime(t *time.Time) *SubDouyuLiveUpdateOne {
	if t != nil {
		sdluo.SetCreateTime(*t)
	}
	return sdluo
}

// ClearCreateTime clears the value of the "create_time" field.
func (sdluo *SubDouyuLiveUpdateOne) ClearCreateTime() *SubDouyuLiveUpdateOne {
	sdluo.mutation.ClearCreateTime()
	return sdluo
}

// SetUpdateTime sets the "update_time" field.
func (sdluo *SubDouyuLiveUpdateOne) SetUpdateTime(t time.Time) *SubDouyuLiveUpdateOne {
	sdluo.mutation.SetUpdateTime(t)
	return sdluo
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sdluo *SubDouyuLiveUpdateOne) SetNillableUpdateTime(t *time.Time) *SubDouyuLiveUpdateOne {
	if t != nil {
		sdluo.SetUpdateTime(*t)
	}
	return sdluo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (sdluo *SubDouyuLiveUpdateOne) ClearUpdateTime() *SubDouyuLiveUpdateOne {
	sdluo.mutation.ClearUpdateTime()
	return sdluo
}

// Mutation returns the SubDouyuLiveMutation object of the builder.
func (sdluo *SubDouyuLiveUpdateOne) Mutation() *SubDouyuLiveMutation {
	return sdluo.mutation
}

// Where appends a list predicates to the SubDouyuLiveUpdate builder.
func (sdluo *SubDouyuLiveUpdateOne) Where(ps ...predicate.SubDouyuLive) *SubDouyuLiveUpdateOne {
	sdluo.mutation.Where(ps...)
	return sdluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sdluo *SubDouyuLiveUpdateOne) Select(field string, fields ...string) *SubDouyuLiveUpdateOne {
	sdluo.fields = append([]string{field}, fields...)
	return sdluo
}

// Save executes the query and returns the updated SubDouyuLive entity.
func (sdluo *SubDouyuLiveUpdateOne) Save(ctx context.Context) (*SubDouyuLive, error) {
	return withHooks(ctx, sdluo.sqlSave, sdluo.mutation, sdluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sdluo *SubDouyuLiveUpdateOne) SaveX(ctx context.Context) *SubDouyuLive {
	node, err := sdluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sdluo *SubDouyuLiveUpdateOne) Exec(ctx context.Context) error {
	_, err := sdluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdluo *SubDouyuLiveUpdateOne) ExecX(ctx context.Context) {
	if err := sdluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sdluo *SubDouyuLiveUpdateOne) sqlSave(ctx context.Context) (_node *SubDouyuLive, err error) {
	_spec := sqlgraph.NewUpdateSpec(subdouyulive.Table, subdouyulive.Columns, sqlgraph.NewFieldSpec(subdouyulive.FieldID, field.TypeInt64))
	id, ok := sdluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubDouyuLive.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sdluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subdouyulive.FieldID)
		for _, f := range fields {
			if !subdouyulive.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subdouyulive.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sdluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sdluo.mutation.RoomID(); ok {
		_spec.SetField(subdouyulive.FieldRoomID, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.AddedRoomID(); ok {
		_spec.AddField(subdouyulive.FieldRoomID, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.LiveState(); ok {
		_spec.SetField(subdouyulive.FieldLiveState, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.AddedLiveState(); ok {
		_spec.AddField(subdouyulive.FieldLiveState, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.LiveStartTime(); ok {
		_spec.SetField(subdouyulive.FieldLiveStartTime, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.AddedLiveStartTime(); ok {
		_spec.AddField(subdouyulive.FieldLiveStartTime, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.LiveEndTime(); ok {
		_spec.SetField(subdouyulive.FieldLiveEndTime, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.AddedLiveEndTime(); ok {
		_spec.AddField(subdouyulive.FieldLiveEndTime, field.TypeInt64, value)
	}
	if value, ok := sdluo.mutation.CreateTime(); ok {
		_spec.SetField(subdouyulive.FieldCreateTime, field.TypeTime, value)
	}
	if sdluo.mutation.CreateTimeCleared() {
		_spec.ClearField(subdouyulive.FieldCreateTime, field.TypeTime)
	}
	if value, ok := sdluo.mutation.UpdateTime(); ok {
		_spec.SetField(subdouyulive.FieldUpdateTime, field.TypeTime, value)
	}
	if sdluo.mutation.UpdateTimeCleared() {
		_spec.ClearField(subdouyulive.FieldUpdateTime, field.TypeTime)
	}
	_node = &SubDouyuLive{config: sdluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sdluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subdouyulive.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sdluo.mutation.done = true
	return _node, nil
}
