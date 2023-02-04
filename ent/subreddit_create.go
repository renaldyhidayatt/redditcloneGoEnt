// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renaldyhidayatt/redditgoent/ent/post"
	"github.com/renaldyhidayatt/redditgoent/ent/subreddit"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
)

// SubredditCreate is the builder for creating a Subreddit entity.
type SubredditCreate struct {
	config
	mutation *SubredditMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SubredditCreate) SetName(s string) *SubredditCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SubredditCreate) SetDescription(s string) *SubredditCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubredditCreate) SetCreatedAt(t time.Time) *SubredditCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubredditCreate) SetNillableCreatedAt(t *time.Time) *SubredditCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubredditCreate) SetUpdatedAt(t time.Time) *SubredditCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubredditCreate) SetNillableUpdatedAt(t *time.Time) *SubredditCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (sc *SubredditCreate) AddPostIDs(ids ...int) *SubredditCreate {
	sc.mutation.AddPostIDs(ids...)
	return sc
}

// AddPosts adds the "posts" edges to the Post entity.
func (sc *SubredditCreate) AddPosts(p ...*Post) *SubredditCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddPostIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (sc *SubredditCreate) AddUserIDs(ids ...int) *SubredditCreate {
	sc.mutation.AddUserIDs(ids...)
	return sc
}

// AddUser adds the "user" edges to the User entity.
func (sc *SubredditCreate) AddUser(u ...*User) *SubredditCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return sc.AddUserIDs(ids...)
}

// Mutation returns the SubredditMutation object of the builder.
func (sc *SubredditCreate) Mutation() *SubredditMutation {
	return sc.mutation
}

// Save creates the Subreddit in the database.
func (sc *SubredditCreate) Save(ctx context.Context) (*Subreddit, error) {
	sc.defaults()
	return withHooks[*Subreddit, SubredditMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubredditCreate) SaveX(ctx context.Context) *Subreddit {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubredditCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubredditCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubredditCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := subreddit.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubredditCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Subreddit.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := subreddit.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Subreddit.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Subreddit.description"`)}
	}
	if v, ok := sc.mutation.Description(); ok {
		if err := subreddit.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Subreddit.description": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Subreddit.created_at"`)}
	}
	if len(sc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Subreddit.user"`)}
	}
	return nil
}

func (sc *SubredditCreate) sqlSave(ctx context.Context) (*Subreddit, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubredditCreate) createSpec() (*Subreddit, *sqlgraph.CreateSpec) {
	var (
		_node = &Subreddit{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subreddit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subreddit.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(subreddit.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(subreddit.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subreddit.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(subreddit.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := sc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subreddit.PostsTable,
			Columns: subreddit.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   subreddit.UserTable,
			Columns: []string{subreddit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubredditCreateBulk is the builder for creating many Subreddit entities in bulk.
type SubredditCreateBulk struct {
	config
	builders []*SubredditCreate
}

// Save creates the Subreddit entities in the database.
func (scb *SubredditCreateBulk) Save(ctx context.Context) ([]*Subreddit, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subreddit, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubredditMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubredditCreateBulk) SaveX(ctx context.Context) []*Subreddit {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubredditCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubredditCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}