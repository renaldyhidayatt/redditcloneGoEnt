// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/renaldyhidayatt/redditgoent/ent/post"
	"github.com/renaldyhidayatt/redditgoent/ent/predicate"
	"github.com/renaldyhidayatt/redditgoent/ent/subreddit"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
)

// PostUpdate is the builder for updating Post entities.
type PostUpdate struct {
	config
	hooks    []Hook
	mutation *PostMutation
}

// Where appends a list predicates to the PostUpdate builder.
func (pu *PostUpdate) Where(ps ...predicate.Post) *PostUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPostname sets the "postname" field.
func (pu *PostUpdate) SetPostname(s string) *PostUpdate {
	pu.mutation.SetPostname(s)
	return pu
}

// SetURL sets the "url" field.
func (pu *PostUpdate) SetURL(s string) *PostUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PostUpdate) SetDescription(s string) *PostUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetVoteCount sets the "voteCount" field.
func (pu *PostUpdate) SetVoteCount(i int) *PostUpdate {
	pu.mutation.ResetVoteCount()
	pu.mutation.SetVoteCount(i)
	return pu
}

// SetNillableVoteCount sets the "voteCount" field if the given value is not nil.
func (pu *PostUpdate) SetNillableVoteCount(i *int) *PostUpdate {
	if i != nil {
		pu.SetVoteCount(*i)
	}
	return pu
}

// AddVoteCount adds i to the "voteCount" field.
func (pu *PostUpdate) AddVoteCount(i int) *PostUpdate {
	pu.mutation.AddVoteCount(i)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PostUpdate) SetCreatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PostUpdate) SetNillableCreatedAt(t *time.Time) *PostUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PostUpdate) SetUpdatedAt(t time.Time) *PostUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PostUpdate) ClearUpdatedAt() *PostUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (pu *PostUpdate) AddUserIDs(ids ...int) *PostUpdate {
	pu.mutation.AddUserIDs(ids...)
	return pu
}

// AddUser adds the "user" edges to the User entity.
func (pu *PostUpdate) AddUser(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddUserIDs(ids...)
}

// AddSubredditIDs adds the "subreddit" edge to the Subreddit entity by IDs.
func (pu *PostUpdate) AddSubredditIDs(ids ...int) *PostUpdate {
	pu.mutation.AddSubredditIDs(ids...)
	return pu
}

// AddSubreddit adds the "subreddit" edges to the Subreddit entity.
func (pu *PostUpdate) AddSubreddit(s ...*Subreddit) *PostUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSubredditIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pu *PostUpdate) Mutation() *PostMutation {
	return pu.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (pu *PostUpdate) ClearUser() *PostUpdate {
	pu.mutation.ClearUser()
	return pu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (pu *PostUpdate) RemoveUserIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveUserIDs(ids...)
	return pu
}

// RemoveUser removes "user" edges to User entities.
func (pu *PostUpdate) RemoveUser(u ...*User) *PostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveUserIDs(ids...)
}

// ClearSubreddit clears all "subreddit" edges to the Subreddit entity.
func (pu *PostUpdate) ClearSubreddit() *PostUpdate {
	pu.mutation.ClearSubreddit()
	return pu
}

// RemoveSubredditIDs removes the "subreddit" edge to Subreddit entities by IDs.
func (pu *PostUpdate) RemoveSubredditIDs(ids ...int) *PostUpdate {
	pu.mutation.RemoveSubredditIDs(ids...)
	return pu
}

// RemoveSubreddit removes "subreddit" edges to Subreddit entities.
func (pu *PostUpdate) RemoveSubreddit(s ...*Subreddit) *PostUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSubredditIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PostUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks[int, PostMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PostUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PostUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PostUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PostUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok && !pu.mutation.UpdatedAtCleared() {
		v := post.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Postname(); ok {
		_spec.SetField(post.FieldPostname, field.TypeString, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(post.FieldURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.VoteCount(); ok {
		_spec.SetField(post.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedVoteCount(); ok {
		_spec.AddField(post.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.ClearField(post.FieldUpdatedAt, field.TypeTime)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.UserTable,
			Columns: post.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedUserIDs(); len(nodes) > 0 && !pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.UserTable,
			Columns: post.UserPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.UserTable,
			Columns: post.UserPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SubredditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.SubredditTable,
			Columns: post.SubredditPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subreddit.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSubredditIDs(); len(nodes) > 0 && !pu.mutation.SubredditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.SubredditTable,
			Columns: post.SubredditPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subreddit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SubredditIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.SubredditTable,
			Columns: post.SubredditPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subreddit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PostUpdateOne is the builder for updating a single Post entity.
type PostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PostMutation
}

// SetPostname sets the "postname" field.
func (puo *PostUpdateOne) SetPostname(s string) *PostUpdateOne {
	puo.mutation.SetPostname(s)
	return puo
}

// SetURL sets the "url" field.
func (puo *PostUpdateOne) SetURL(s string) *PostUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PostUpdateOne) SetDescription(s string) *PostUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetVoteCount sets the "voteCount" field.
func (puo *PostUpdateOne) SetVoteCount(i int) *PostUpdateOne {
	puo.mutation.ResetVoteCount()
	puo.mutation.SetVoteCount(i)
	return puo
}

// SetNillableVoteCount sets the "voteCount" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableVoteCount(i *int) *PostUpdateOne {
	if i != nil {
		puo.SetVoteCount(*i)
	}
	return puo
}

// AddVoteCount adds i to the "voteCount" field.
func (puo *PostUpdateOne) AddVoteCount(i int) *PostUpdateOne {
	puo.mutation.AddVoteCount(i)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PostUpdateOne) SetCreatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PostUpdateOne) SetNillableCreatedAt(t *time.Time) *PostUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PostUpdateOne) SetUpdatedAt(t time.Time) *PostUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PostUpdateOne) ClearUpdatedAt() *PostUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (puo *PostUpdateOne) AddUserIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddUserIDs(ids...)
	return puo
}

// AddUser adds the "user" edges to the User entity.
func (puo *PostUpdateOne) AddUser(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddUserIDs(ids...)
}

// AddSubredditIDs adds the "subreddit" edge to the Subreddit entity by IDs.
func (puo *PostUpdateOne) AddSubredditIDs(ids ...int) *PostUpdateOne {
	puo.mutation.AddSubredditIDs(ids...)
	return puo
}

// AddSubreddit adds the "subreddit" edges to the Subreddit entity.
func (puo *PostUpdateOne) AddSubreddit(s ...*Subreddit) *PostUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSubredditIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (puo *PostUpdateOne) Mutation() *PostMutation {
	return puo.mutation
}

// ClearUser clears all "user" edges to the User entity.
func (puo *PostUpdateOne) ClearUser() *PostUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (puo *PostUpdateOne) RemoveUserIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveUserIDs(ids...)
	return puo
}

// RemoveUser removes "user" edges to User entities.
func (puo *PostUpdateOne) RemoveUser(u ...*User) *PostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveUserIDs(ids...)
}

// ClearSubreddit clears all "subreddit" edges to the Subreddit entity.
func (puo *PostUpdateOne) ClearSubreddit() *PostUpdateOne {
	puo.mutation.ClearSubreddit()
	return puo
}

// RemoveSubredditIDs removes the "subreddit" edge to Subreddit entities by IDs.
func (puo *PostUpdateOne) RemoveSubredditIDs(ids ...int) *PostUpdateOne {
	puo.mutation.RemoveSubredditIDs(ids...)
	return puo
}

// RemoveSubreddit removes "subreddit" edges to Subreddit entities.
func (puo *PostUpdateOne) RemoveSubreddit(s ...*Subreddit) *PostUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSubredditIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PostUpdateOne) Select(field string, fields ...string) *PostUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Post entity.
func (puo *PostUpdateOne) Save(ctx context.Context) (*Post, error) {
	puo.defaults()
	return withHooks[*Post, PostMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PostUpdateOne) SaveX(ctx context.Context) *Post {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PostUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PostUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PostUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok && !puo.mutation.UpdatedAtCleared() {
		v := post.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PostUpdateOne) sqlSave(ctx context.Context) (_node *Post, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   post.Table,
			Columns: post.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: post.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Post.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, post.FieldID)
		for _, f := range fields {
			if !post.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != post.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Postname(); ok {
		_spec.SetField(post.FieldPostname, field.TypeString, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(post.FieldURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.VoteCount(); ok {
		_spec.SetField(post.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedVoteCount(); ok {
		_spec.AddField(post.FieldVoteCount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.ClearField(post.FieldUpdatedAt, field.TypeTime)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.UserTable,
			Columns: post.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedUserIDs(); len(nodes) > 0 && !puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.UserTable,
			Columns: post.UserPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.UserTable,
			Columns: post.UserPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SubredditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.SubredditTable,
			Columns: post.SubredditPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subreddit.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSubredditIDs(); len(nodes) > 0 && !puo.mutation.SubredditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.SubredditTable,
			Columns: post.SubredditPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subreddit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SubredditIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   post.SubredditTable,
			Columns: post.SubredditPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: subreddit.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Post{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{post.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
