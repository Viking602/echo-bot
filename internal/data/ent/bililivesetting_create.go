// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/bililivesetting"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BiliLiveSettingCreate is the builder for creating a BiliLiveSetting entity.
type BiliLiveSettingCreate struct {
	config
	mutation *BiliLiveSettingMutation
	hooks    []Hook
}

// SetRoomID sets the "room_id" field.
func (blsc *BiliLiveSettingCreate) SetRoomID(i int64) *BiliLiveSettingCreate {
	blsc.mutation.SetRoomID(i)
	return blsc
}

// SetLiveState sets the "live_state" field.
func (blsc *BiliLiveSettingCreate) SetLiveState(i int64) *BiliLiveSettingCreate {
	blsc.mutation.SetLiveState(i)
	return blsc
}

// SetLiveStartTime sets the "live_start_time" field.
func (blsc *BiliLiveSettingCreate) SetLiveStartTime(t time.Time) *BiliLiveSettingCreate {
	blsc.mutation.SetLiveStartTime(t)
	return blsc
}

// SetLiveEndTime sets the "live_end_time" field.
func (blsc *BiliLiveSettingCreate) SetLiveEndTime(t time.Time) *BiliLiveSettingCreate {
	blsc.mutation.SetLiveEndTime(t)
	return blsc
}

// SetCreateTime sets the "create_time" field.
func (blsc *BiliLiveSettingCreate) SetCreateTime(t time.Time) *BiliLiveSettingCreate {
	blsc.mutation.SetCreateTime(t)
	return blsc
}

// SetUpdateTime sets the "update_time" field.
func (blsc *BiliLiveSettingCreate) SetUpdateTime(t time.Time) *BiliLiveSettingCreate {
	blsc.mutation.SetUpdateTime(t)
	return blsc
}

// SetID sets the "id" field.
func (blsc *BiliLiveSettingCreate) SetID(i int64) *BiliLiveSettingCreate {
	blsc.mutation.SetID(i)
	return blsc
}

// Mutation returns the BiliLiveSettingMutation object of the builder.
func (blsc *BiliLiveSettingCreate) Mutation() *BiliLiveSettingMutation {
	return blsc.mutation
}

// Save creates the BiliLiveSetting in the database.
func (blsc *BiliLiveSettingCreate) Save(ctx context.Context) (*BiliLiveSetting, error) {
	return withHooks(ctx, blsc.sqlSave, blsc.mutation, blsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (blsc *BiliLiveSettingCreate) SaveX(ctx context.Context) *BiliLiveSetting {
	v, err := blsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (blsc *BiliLiveSettingCreate) Exec(ctx context.Context) error {
	_, err := blsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blsc *BiliLiveSettingCreate) ExecX(ctx context.Context) {
	if err := blsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (blsc *BiliLiveSettingCreate) check() error {
	if _, ok := blsc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room_id", err: errors.New(`ent: missing required field "BiliLiveSetting.room_id"`)}
	}
	if _, ok := blsc.mutation.LiveState(); !ok {
		return &ValidationError{Name: "live_state", err: errors.New(`ent: missing required field "BiliLiveSetting.live_state"`)}
	}
	if _, ok := blsc.mutation.LiveStartTime(); !ok {
		return &ValidationError{Name: "live_start_time", err: errors.New(`ent: missing required field "BiliLiveSetting.live_start_time"`)}
	}
	if _, ok := blsc.mutation.LiveEndTime(); !ok {
		return &ValidationError{Name: "live_end_time", err: errors.New(`ent: missing required field "BiliLiveSetting.live_end_time"`)}
	}
	if _, ok := blsc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "BiliLiveSetting.create_time"`)}
	}
	if _, ok := blsc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "BiliLiveSetting.update_time"`)}
	}
	return nil
}

func (blsc *BiliLiveSettingCreate) sqlSave(ctx context.Context) (*BiliLiveSetting, error) {
	if err := blsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := blsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, blsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	blsc.mutation.id = &_node.ID
	blsc.mutation.done = true
	return _node, nil
}

func (blsc *BiliLiveSettingCreate) createSpec() (*BiliLiveSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &BiliLiveSetting{config: blsc.config}
		_spec = sqlgraph.NewCreateSpec(bililivesetting.Table, sqlgraph.NewFieldSpec(bililivesetting.FieldID, field.TypeInt64))
	)
	if id, ok := blsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := blsc.mutation.RoomID(); ok {
		_spec.SetField(bililivesetting.FieldRoomID, field.TypeInt64, value)
		_node.RoomID = value
	}
	if value, ok := blsc.mutation.LiveState(); ok {
		_spec.SetField(bililivesetting.FieldLiveState, field.TypeInt64, value)
		_node.LiveState = value
	}
	if value, ok := blsc.mutation.LiveStartTime(); ok {
		_spec.SetField(bililivesetting.FieldLiveStartTime, field.TypeTime, value)
		_node.LiveStartTime = value
	}
	if value, ok := blsc.mutation.LiveEndTime(); ok {
		_spec.SetField(bililivesetting.FieldLiveEndTime, field.TypeTime, value)
		_node.LiveEndTime = value
	}
	if value, ok := blsc.mutation.CreateTime(); ok {
		_spec.SetField(bililivesetting.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := blsc.mutation.UpdateTime(); ok {
		_spec.SetField(bililivesetting.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// BiliLiveSettingCreateBulk is the builder for creating many BiliLiveSetting entities in bulk.
type BiliLiveSettingCreateBulk struct {
	config
	err      error
	builders []*BiliLiveSettingCreate
}

// Save creates the BiliLiveSetting entities in the database.
func (blscb *BiliLiveSettingCreateBulk) Save(ctx context.Context) ([]*BiliLiveSetting, error) {
	if blscb.err != nil {
		return nil, blscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(blscb.builders))
	nodes := make([]*BiliLiveSetting, len(blscb.builders))
	mutators := make([]Mutator, len(blscb.builders))
	for i := range blscb.builders {
		func(i int, root context.Context) {
			builder := blscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BiliLiveSettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, blscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, blscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, blscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (blscb *BiliLiveSettingCreateBulk) SaveX(ctx context.Context) []*BiliLiveSetting {
	v, err := blscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (blscb *BiliLiveSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := blscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (blscb *BiliLiveSettingCreateBulk) ExecX(ctx context.Context) {
	if err := blscb.Exec(ctx); err != nil {
		panic(err)
	}
}
