// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/renaldyhidayatt/redditgoent/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// Postname applies equality check predicate on the "postname" field. It's identical to PostnameEQ.
func Postname(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPostname, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldURL, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDescription, v))
}

// VoteCount applies equality check predicate on the "voteCount" field. It's identical to VoteCountEQ.
func VoteCount(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldVoteCount, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// PostnameEQ applies the EQ predicate on the "postname" field.
func PostnameEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldPostname, v))
}

// PostnameNEQ applies the NEQ predicate on the "postname" field.
func PostnameNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldPostname, v))
}

// PostnameIn applies the In predicate on the "postname" field.
func PostnameIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldPostname, vs...))
}

// PostnameNotIn applies the NotIn predicate on the "postname" field.
func PostnameNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldPostname, vs...))
}

// PostnameGT applies the GT predicate on the "postname" field.
func PostnameGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldPostname, v))
}

// PostnameGTE applies the GTE predicate on the "postname" field.
func PostnameGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldPostname, v))
}

// PostnameLT applies the LT predicate on the "postname" field.
func PostnameLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldPostname, v))
}

// PostnameLTE applies the LTE predicate on the "postname" field.
func PostnameLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldPostname, v))
}

// PostnameContains applies the Contains predicate on the "postname" field.
func PostnameContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldPostname, v))
}

// PostnameHasPrefix applies the HasPrefix predicate on the "postname" field.
func PostnameHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldPostname, v))
}

// PostnameHasSuffix applies the HasSuffix predicate on the "postname" field.
func PostnameHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldPostname, v))
}

// PostnameEqualFold applies the EqualFold predicate on the "postname" field.
func PostnameEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldPostname, v))
}

// PostnameContainsFold applies the ContainsFold predicate on the "postname" field.
func PostnameContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldPostname, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldURL, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldDescription, v))
}

// VoteCountEQ applies the EQ predicate on the "voteCount" field.
func VoteCountEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldVoteCount, v))
}

// VoteCountNEQ applies the NEQ predicate on the "voteCount" field.
func VoteCountNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldVoteCount, v))
}

// VoteCountIn applies the In predicate on the "voteCount" field.
func VoteCountIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldVoteCount, vs...))
}

// VoteCountNotIn applies the NotIn predicate on the "voteCount" field.
func VoteCountNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldVoteCount, vs...))
}

// VoteCountGT applies the GT predicate on the "voteCount" field.
func VoteCountGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldVoteCount, v))
}

// VoteCountGTE applies the GTE predicate on the "voteCount" field.
func VoteCountGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldVoteCount, v))
}

// VoteCountLT applies the LT predicate on the "voteCount" field.
func VoteCountLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldVoteCount, v))
}

// VoteCountLTE applies the LTE predicate on the "voteCount" field.
func VoteCountLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldVoteCount, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldUpdatedAt))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubreddit applies the HasEdge predicate on the "subreddit" edge.
func HasSubreddit() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubredditTable, SubredditPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubredditWith applies the HasEdge predicate on the "subreddit" edge with a given conditions (other predicates).
func HasSubredditWith(preds ...predicate.Subreddit) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SubredditInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, SubredditTable, SubredditPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		p(s.Not())
	})
}