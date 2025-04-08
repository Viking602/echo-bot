// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"echo/internal/data/ent/bot"
	"echo/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BotDelete is the builder for deleting a Bot entity.
type BotDelete struct {
	config
	hooks    []Hook
	mutation *BotMutation
}

// Where appends a list predicates to the BotDelete builder.
func (bd *BotDelete) Where(ps ...predicate.Bot) *BotDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BotDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BotDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bot.Table, sqlgraph.NewFieldSpec(bot.FieldID, field.TypeInt64))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BotDeleteOne is the builder for deleting a single Bot entity.
type BotDeleteOne struct {
	bd *BotDelete
}

// Where appends a list predicates to the BotDelete builder.
func (bdo *BotDeleteOne) Where(ps ...predicate.Bot) *BotDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BotDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BotDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}
