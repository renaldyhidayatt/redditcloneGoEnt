// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renaldyhidayatt/redditgoent/ent/post"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
	"github.com/renaldyhidayatt/redditgoent/ent/vote"
)

// VoteCreate is the builder for creating a Vote entity.
type VoteCreate struct {
	config
	mutation *VoteMutation
	hooks    []Hook
}

// SetVote sets the "vote" field.
func (vc *VoteCreate) SetVote(v vote.Vote) *VoteCreate {
	vc.mutation.SetVote(v)
	return vc
}

// AddPostIDs adds the "post" edge to the Post entity by IDs.
func (vc *VoteCreate) AddPostIDs(ids ...int) *VoteCreate {
	vc.mutation.AddPostIDs(ids...)
	return vc
}

// AddPost adds the "post" edges to the Post entity.
func (vc *VoteCreate) AddPost(p ...*Post) *VoteCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vc.AddPostIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (vc *VoteCreate) AddUserIDs(ids ...int) *VoteCreate {
	vc.mutation.AddUserIDs(ids...)
	return vc
}

// AddUser adds the "user" edges to the User entity.
func (vc *VoteCreate) AddUser(u ...*User) *VoteCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return vc.AddUserIDs(ids...)
}

// Mutation returns the VoteMutation object of the builder.
func (vc *VoteCreate) Mutation() *VoteMutation {
	return vc.mutation
}

// Save creates the Vote in the database.
func (vc *VoteCreate) Save(ctx context.Context) (*Vote, error) {
	return withHooks[*Vote, VoteMutation](ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VoteCreate) SaveX(ctx context.Context) *Vote {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VoteCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VoteCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VoteCreate) check() error {
	if _, ok := vc.mutation.Vote(); !ok {
		return &ValidationError{Name: "vote", err: errors.New(`ent: missing required field "Vote.vote"`)}
	}
	if v, ok := vc.mutation.Vote(); ok {
		if err := vote.VoteValidator(v); err != nil {
			return &ValidationError{Name: "vote", err: fmt.Errorf(`ent: validator failed for field "Vote.vote": %w`, err)}
		}
	}
	if len(vc.mutation.PostIDs()) == 0 {
		return &ValidationError{Name: "post", err: errors.New(`ent: missing required edge "Vote.post"`)}
	}
	if len(vc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Vote.user"`)}
	}
	return nil
}

func (vc *VoteCreate) sqlSave(ctx context.Context) (*Vote, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VoteCreate) createSpec() (*Vote, *sqlgraph.CreateSpec) {
	var (
		_node = &Vote{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vote.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: vote.FieldID,
			},
		}
	)
	if value, ok := vc.mutation.Vote(); ok {
		_spec.SetField(vote.FieldVote, field.TypeEnum, value)
		_node.Vote = value
	}
	if nodes := vc.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vote.PostTable,
			Columns: []string{vote.PostColumn},
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
	if nodes := vc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   vote.UserTable,
			Columns: []string{vote.UserColumn},
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

// VoteCreateBulk is the builder for creating many Vote entities in bulk.
type VoteCreateBulk struct {
	config
	builders []*VoteCreate
}

// Save creates the Vote entities in the database.
func (vcb *VoteCreateBulk) Save(ctx context.Context) ([]*Vote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vote, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VoteMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VoteCreateBulk) SaveX(ctx context.Context) []*Vote {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VoteCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VoteCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}