// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/renaldyhidayatt/redditgoent/ent/subreddit"
)

// Subreddit is the model entity for the Subreddit schema.
type Subreddit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Comunity name is required
	Name string `json:"name,omitempty"`
	// Description is required
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubredditQuery when eager-loading is set.
	Edges SubredditEdges `json:"edges"`
}

// SubredditEdges holds the relations/edges for other nodes in the graph.
type SubredditEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e SubredditEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e SubredditEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subreddit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subreddit.FieldID:
			values[i] = new(sql.NullInt64)
		case subreddit.FieldName, subreddit.FieldDescription:
			values[i] = new(sql.NullString)
		case subreddit.FieldCreatedAt, subreddit.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subreddit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subreddit fields.
func (s *Subreddit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subreddit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subreddit.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case subreddit.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case subreddit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case subreddit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the Subreddit entity.
func (s *Subreddit) QueryPosts() *PostQuery {
	return NewSubredditClient(s.config).QueryPosts(s)
}

// QueryUser queries the "user" edge of the Subreddit entity.
func (s *Subreddit) QueryUser() *UserQuery {
	return NewSubredditClient(s.config).QueryUser(s)
}

// Update returns a builder for updating this Subreddit.
// Note that you need to call Subreddit.Unwrap() before calling this method if this Subreddit
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subreddit) Update() *SubredditUpdateOne {
	return NewSubredditClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Subreddit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subreddit) Unwrap() *Subreddit {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subreddit is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subreddit) String() string {
	var builder strings.Builder
	builder.WriteString("Subreddit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(s.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Subreddits is a parsable slice of Subreddit.
type Subreddits []*Subreddit

func (s Subreddits) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}