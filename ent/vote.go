// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/renaldyhidayatt/redditgoent/ent/vote"
)

// Vote is the model entity for the Vote schema.
type Vote struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pilih upvote atau downvote
	Vote vote.Vote `json:"vote,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VoteQuery when eager-loading is set.
	Edges VoteEdges `json:"edges"`
}

// VoteEdges holds the relations/edges for other nodes in the graph.
type VoteEdges struct {
	// Post holds the value of the post edge.
	Post []*Post `json:"post,omitempty"`
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading.
func (e VoteEdges) PostOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e VoteEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			values[i] = new(sql.NullInt64)
		case vote.FieldVote:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Vote", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vote fields.
func (v *Vote) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vote.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case vote.FieldVote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vote", values[i])
			} else if value.Valid {
				v.Vote = vote.Vote(value.String)
			}
		}
	}
	return nil
}

// QueryPost queries the "post" edge of the Vote entity.
func (v *Vote) QueryPost() *PostQuery {
	return NewVoteClient(v.config).QueryPost(v)
}

// QueryUser queries the "user" edge of the Vote entity.
func (v *Vote) QueryUser() *UserQuery {
	return NewVoteClient(v.config).QueryUser(v)
}

// Update returns a builder for updating this Vote.
// Note that you need to call Vote.Unwrap() before calling this method if this Vote
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vote) Update() *VoteUpdateOne {
	return NewVoteClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vote) Unwrap() *Vote {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Vote is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vote) String() string {
	var builder strings.Builder
	builder.WriteString("Vote(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("vote=")
	builder.WriteString(fmt.Sprintf("%v", v.Vote))
	builder.WriteByte(')')
	return builder.String()
}

// Votes is a parsable slice of Vote.
type Votes []*Vote

func (v Votes) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}
