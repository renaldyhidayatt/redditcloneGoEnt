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
	"github.com/renaldyhidayatt/redditgoent/ent/comment"
	"github.com/renaldyhidayatt/redditgoent/ent/post"
	"github.com/renaldyhidayatt/redditgoent/ent/predicate"
	"github.com/renaldyhidayatt/redditgoent/ent/user"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetText sets the "text" field.
func (cu *CommentUpdate) SetText(s string) *CommentUpdate {
	cu.mutation.SetText(s)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(t *time.Time) *CommentUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(t time.Time) *CommentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cu *CommentUpdate) ClearUpdatedAt() *CommentUpdate {
	cu.mutation.ClearUpdatedAt()
	return cu
}

// AddPostIDs adds the "post" edge to the Post entity by IDs.
func (cu *CommentUpdate) AddPostIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddPostIDs(ids...)
	return cu
}

// AddPost adds the "post" edges to the Post entity.
func (cu *CommentUpdate) AddPost(p ...*Post) *CommentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPostIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (cu *CommentUpdate) AddUserIDs(ids ...int) *CommentUpdate {
	cu.mutation.AddUserIDs(ids...)
	return cu
}

// AddUser adds the "user" edges to the User entity.
func (cu *CommentUpdate) AddUser(u ...*User) *CommentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddUserIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// ClearPost clears all "post" edges to the Post entity.
func (cu *CommentUpdate) ClearPost() *CommentUpdate {
	cu.mutation.ClearPost()
	return cu
}

// RemovePostIDs removes the "post" edge to Post entities by IDs.
func (cu *CommentUpdate) RemovePostIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemovePostIDs(ids...)
	return cu
}

// RemovePost removes "post" edges to Post entities.
func (cu *CommentUpdate) RemovePost(p ...*Post) *CommentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePostIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (cu *CommentUpdate) ClearUser() *CommentUpdate {
	cu.mutation.ClearUser()
	return cu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (cu *CommentUpdate) RemoveUserIDs(ids ...int) *CommentUpdate {
	cu.mutation.RemoveUserIDs(ids...)
	return cu
}

// RemoveUser removes "user" edges to User entities.
func (cu *CommentUpdate) RemoveUser(u ...*User) *CommentUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks[int, CommentMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok && !cu.mutation.UpdatedAtCleared() {
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if cu.mutation.UpdatedAtCleared() {
		_spec.ClearField(comment.FieldUpdatedAt, field.TypeTime)
	}
	if cu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPostIDs(); len(nodes) > 0 && !cu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
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
	if nodes := cu.mutation.RemovedUserIDs(); len(nodes) > 0 && !cu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
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
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetText sets the "text" field.
func (cuo *CommentUpdateOne) SetText(s string) *CommentUpdateOne {
	cuo.mutation.SetText(s)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(t *time.Time) *CommentUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(t time.Time) *CommentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cuo *CommentUpdateOne) ClearUpdatedAt() *CommentUpdateOne {
	cuo.mutation.ClearUpdatedAt()
	return cuo
}

// AddPostIDs adds the "post" edge to the Post entity by IDs.
func (cuo *CommentUpdateOne) AddPostIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddPostIDs(ids...)
	return cuo
}

// AddPost adds the "post" edges to the Post entity.
func (cuo *CommentUpdateOne) AddPost(p ...*Post) *CommentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPostIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (cuo *CommentUpdateOne) AddUserIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.AddUserIDs(ids...)
	return cuo
}

// AddUser adds the "user" edges to the User entity.
func (cuo *CommentUpdateOne) AddUser(u ...*User) *CommentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddUserIDs(ids...)
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// ClearPost clears all "post" edges to the Post entity.
func (cuo *CommentUpdateOne) ClearPost() *CommentUpdateOne {
	cuo.mutation.ClearPost()
	return cuo
}

// RemovePostIDs removes the "post" edge to Post entities by IDs.
func (cuo *CommentUpdateOne) RemovePostIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemovePostIDs(ids...)
	return cuo
}

// RemovePost removes "post" edges to Post entities.
func (cuo *CommentUpdateOne) RemovePost(p ...*Post) *CommentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePostIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (cuo *CommentUpdateOne) ClearUser() *CommentUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (cuo *CommentUpdateOne) RemoveUserIDs(ids ...int) *CommentUpdateOne {
	cuo.mutation.RemoveUserIDs(ids...)
	return cuo
}

// RemoveUser removes "user" edges to User entities.
func (cuo *CommentUpdateOne) RemoveUser(u ...*User) *CommentUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	cuo.defaults()
	return withHooks[*Comment, CommentMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok && !cuo.mutation.UpdatedAtCleared() {
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeTime, value)
	}
	if cuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(comment.FieldUpdatedAt, field.TypeTime)
	}
	if cuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPostIDs(); len(nodes) > 0 && !cuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
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
	if nodes := cuo.mutation.RemovedUserIDs(); len(nodes) > 0 && !cuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
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
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   comment.UserTable,
			Columns: []string{comment.UserColumn},
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
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}