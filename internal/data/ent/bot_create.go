// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/bot"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BotCreate is the builder for creating a Bot entity.
type BotCreate struct {
	config
	mutation *BotMutation
	hooks    []Hook
}

// SetBotName sets the "bot_name" field.
func (bc *BotCreate) SetBotName(s string) *BotCreate {
	bc.mutation.SetBotName(s)
	return bc
}

// SetStatus sets the "status" field.
func (bc *BotCreate) SetStatus(i int) *BotCreate {
	bc.mutation.SetStatus(i)
	return bc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bc *BotCreate) SetNillableStatus(i *int) *BotCreate {
	if i != nil {
		bc.SetStatus(*i)
	}
	return bc
}

// SetCreateTime sets the "create_time" field.
func (bc *BotCreate) SetCreateTime(t time.Time) *BotCreate {
	bc.mutation.SetCreateTime(t)
	return bc
}

// SetUpdateTime sets the "update_time" field.
func (bc *BotCreate) SetUpdateTime(t time.Time) *BotCreate {
	bc.mutation.SetUpdateTime(t)
	return bc
}

// SetID sets the "id" field.
func (bc *BotCreate) SetID(i int64) *BotCreate {
	bc.mutation.SetID(i)
	return bc
}

// Mutation returns the BotMutation object of the builder.
func (bc *BotCreate) Mutation() *BotMutation {
	return bc.mutation
}

// Save creates the Bot in the database.
func (bc *BotCreate) Save(ctx context.Context) (*Bot, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BotCreate) SaveX(ctx context.Context) *Bot {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BotCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BotCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BotCreate) defaults() {
	if _, ok := bc.mutation.Status(); !ok {
		v := bot.DefaultStatus
		bc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BotCreate) check() error {
	if _, ok := bc.mutation.BotName(); !ok {
		return &ValidationError{Name: "bot_name", err: errors.New(`ent: missing required field "Bot.bot_name"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Bot.status"`)}
	}
	if _, ok := bc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Bot.create_time"`)}
	}
	if _, ok := bc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Bot.update_time"`)}
	}
	return nil
}

func (bc *BotCreate) sqlSave(ctx context.Context) (*Bot, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BotCreate) createSpec() (*Bot, *sqlgraph.CreateSpec) {
	var (
		_node = &Bot{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(bot.Table, sqlgraph.NewFieldSpec(bot.FieldID, field.TypeInt64))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.BotName(); ok {
		_spec.SetField(bot.FieldBotName, field.TypeString, value)
		_node.BotName = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(bot.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.CreateTime(); ok {
		_spec.SetField(bot.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := bc.mutation.UpdateTime(); ok {
		_spec.SetField(bot.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// BotCreateBulk is the builder for creating many Bot entities in bulk.
type BotCreateBulk struct {
	config
	err      error
	builders []*BotCreate
}

// Save creates the Bot entities in the database.
func (bcb *BotCreateBulk) Save(ctx context.Context) ([]*Bot, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bot, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BotMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BotCreateBulk) SaveX(ctx context.Context) []*Bot {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BotCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BotCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
