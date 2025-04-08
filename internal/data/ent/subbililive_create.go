// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/subbililive"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SubBiliLiveCreate is the builder for creating a SubBiliLive entity.
type SubBiliLiveCreate struct {
	config
	mutation *SubBiliLiveMutation
	hooks    []Hook
}

// SetRoomID sets the "room_id" field.
func (sblc *SubBiliLiveCreate) SetRoomID(i int64) *SubBiliLiveCreate {
	sblc.mutation.SetRoomID(i)
	return sblc
}

// SetLiveState sets the "live_state" field.
func (sblc *SubBiliLiveCreate) SetLiveState(i int64) *SubBiliLiveCreate {
	sblc.mutation.SetLiveState(i)
	return sblc
}

// SetLiveStartTime sets the "live_start_time" field.
func (sblc *SubBiliLiveCreate) SetLiveStartTime(t time.Time) *SubBiliLiveCreate {
	sblc.mutation.SetLiveStartTime(t)
	return sblc
}

// SetLiveEndTime sets the "live_end_time" field.
func (sblc *SubBiliLiveCreate) SetLiveEndTime(t time.Time) *SubBiliLiveCreate {
	sblc.mutation.SetLiveEndTime(t)
	return sblc
}

// SetCreateTime sets the "create_time" field.
func (sblc *SubBiliLiveCreate) SetCreateTime(t time.Time) *SubBiliLiveCreate {
	sblc.mutation.SetCreateTime(t)
	return sblc
}

// SetUpdateTime sets the "update_time" field.
func (sblc *SubBiliLiveCreate) SetUpdateTime(t time.Time) *SubBiliLiveCreate {
	sblc.mutation.SetUpdateTime(t)
	return sblc
}

// SetID sets the "id" field.
func (sblc *SubBiliLiveCreate) SetID(i int64) *SubBiliLiveCreate {
	sblc.mutation.SetID(i)
	return sblc
}

// Mutation returns the SubBiliLiveMutation object of the builder.
func (sblc *SubBiliLiveCreate) Mutation() *SubBiliLiveMutation {
	return sblc.mutation
}

// Save creates the SubBiliLive in the database.
func (sblc *SubBiliLiveCreate) Save(ctx context.Context) (*SubBiliLive, error) {
	return withHooks(ctx, sblc.sqlSave, sblc.mutation, sblc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sblc *SubBiliLiveCreate) SaveX(ctx context.Context) *SubBiliLive {
	v, err := sblc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sblc *SubBiliLiveCreate) Exec(ctx context.Context) error {
	_, err := sblc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sblc *SubBiliLiveCreate) ExecX(ctx context.Context) {
	if err := sblc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sblc *SubBiliLiveCreate) check() error {
	if _, ok := sblc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room_id", err: errors.New(`ent: missing required field "SubBiliLive.room_id"`)}
	}
	if _, ok := sblc.mutation.LiveState(); !ok {
		return &ValidationError{Name: "live_state", err: errors.New(`ent: missing required field "SubBiliLive.live_state"`)}
	}
	if _, ok := sblc.mutation.LiveStartTime(); !ok {
		return &ValidationError{Name: "live_start_time", err: errors.New(`ent: missing required field "SubBiliLive.live_start_time"`)}
	}
	if _, ok := sblc.mutation.LiveEndTime(); !ok {
		return &ValidationError{Name: "live_end_time", err: errors.New(`ent: missing required field "SubBiliLive.live_end_time"`)}
	}
	if _, ok := sblc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "SubBiliLive.create_time"`)}
	}
	if _, ok := sblc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "SubBiliLive.update_time"`)}
	}
	return nil
}

func (sblc *SubBiliLiveCreate) sqlSave(ctx context.Context) (*SubBiliLive, error) {
	if err := sblc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sblc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sblc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sblc.mutation.id = &_node.ID
	sblc.mutation.done = true
	return _node, nil
}

func (sblc *SubBiliLiveCreate) createSpec() (*SubBiliLive, *sqlgraph.CreateSpec) {
	var (
		_node = &SubBiliLive{config: sblc.config}
		_spec = sqlgraph.NewCreateSpec(subbililive.Table, sqlgraph.NewFieldSpec(subbililive.FieldID, field.TypeInt64))
	)
	if id, ok := sblc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sblc.mutation.RoomID(); ok {
		_spec.SetField(subbililive.FieldRoomID, field.TypeInt64, value)
		_node.RoomID = value
	}
	if value, ok := sblc.mutation.LiveState(); ok {
		_spec.SetField(subbililive.FieldLiveState, field.TypeInt64, value)
		_node.LiveState = value
	}
	if value, ok := sblc.mutation.LiveStartTime(); ok {
		_spec.SetField(subbililive.FieldLiveStartTime, field.TypeTime, value)
		_node.LiveStartTime = value
	}
	if value, ok := sblc.mutation.LiveEndTime(); ok {
		_spec.SetField(subbililive.FieldLiveEndTime, field.TypeTime, value)
		_node.LiveEndTime = value
	}
	if value, ok := sblc.mutation.CreateTime(); ok {
		_spec.SetField(subbililive.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sblc.mutation.UpdateTime(); ok {
		_spec.SetField(subbililive.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// SubBiliLiveCreateBulk is the builder for creating many SubBiliLive entities in bulk.
type SubBiliLiveCreateBulk struct {
	config
	err      error
	builders []*SubBiliLiveCreate
}

// Save creates the SubBiliLive entities in the database.
func (sblcb *SubBiliLiveCreateBulk) Save(ctx context.Context) ([]*SubBiliLive, error) {
	if sblcb.err != nil {
		return nil, sblcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sblcb.builders))
	nodes := make([]*SubBiliLive, len(sblcb.builders))
	mutators := make([]Mutator, len(sblcb.builders))
	for i := range sblcb.builders {
		func(i int, root context.Context) {
			builder := sblcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubBiliLiveMutation)
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
					_, err = mutators[i+1].Mutate(root, sblcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sblcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sblcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sblcb *SubBiliLiveCreateBulk) SaveX(ctx context.Context) []*SubBiliLive {
	v, err := sblcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sblcb *SubBiliLiveCreateBulk) Exec(ctx context.Context) error {
	_, err := sblcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sblcb *SubBiliLiveCreateBulk) ExecX(ctx context.Context) {
	if err := sblcb.Exec(ctx); err != nil {
		panic(err)
	}
}
