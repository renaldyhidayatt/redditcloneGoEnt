// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "postname", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "vote_count", Type: field.TypeInt, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "comment_post", Type: field.TypeInt, Nullable: true},
		{Name: "vote_post", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_comments_post",
				Columns:    []*schema.Column{PostsColumns[7]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_votes_post",
				Columns:    []*schema.Column{PostsColumns[8]},
				RefColumns: []*schema.Column{VotesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubredditsColumns holds the columns for the "subreddits" table.
	SubredditsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// SubredditsTable holds the schema information for the "subreddits" table.
	SubredditsTable = &schema.Table{
		Name:       "subreddits",
		Columns:    SubredditsColumns,
		PrimaryKey: []*schema.Column{SubredditsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "enabled", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "comment_user", Type: field.TypeInt, Nullable: true},
		{Name: "subreddit_user", Type: field.TypeInt, Nullable: true},
		{Name: "verification_token_user", Type: field.TypeInt, Nullable: true},
		{Name: "vote_user", Type: field.TypeInt, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_comments_user",
				Columns:    []*schema.Column{UsersColumns[7]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_subreddits_user",
				Columns:    []*schema.Column{UsersColumns[8]},
				RefColumns: []*schema.Column{SubredditsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_verification_tokens_user",
				Columns:    []*schema.Column{UsersColumns[9]},
				RefColumns: []*schema.Column{VerificationTokensColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_votes_user",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{VotesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// VerificationTokensColumns holds the columns for the "verification_tokens" table.
	VerificationTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// VerificationTokensTable holds the schema information for the "verification_tokens" table.
	VerificationTokensTable = &schema.Table{
		Name:       "verification_tokens",
		Columns:    VerificationTokensColumns,
		PrimaryKey: []*schema.Column{VerificationTokensColumns[0]},
	}
	// VotesColumns holds the columns for the "votes" table.
	VotesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "vote", Type: field.TypeEnum, Enums: []string{"upvote", "downvote"}},
	}
	// VotesTable holds the schema information for the "votes" table.
	VotesTable = &schema.Table{
		Name:       "votes",
		Columns:    VotesColumns,
		PrimaryKey: []*schema.Column{VotesColumns[0]},
	}
	// PostUserColumns holds the columns for the "post_user" table.
	PostUserColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// PostUserTable holds the schema information for the "post_user" table.
	PostUserTable = &schema.Table{
		Name:       "post_user",
		Columns:    PostUserColumns,
		PrimaryKey: []*schema.Column{PostUserColumns[0], PostUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_user_post_id",
				Columns:    []*schema.Column{PostUserColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_user_user_id",
				Columns:    []*schema.Column{PostUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostSubredditColumns holds the columns for the "post_subreddit" table.
	PostSubredditColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "subreddit_id", Type: field.TypeInt},
	}
	// PostSubredditTable holds the schema information for the "post_subreddit" table.
	PostSubredditTable = &schema.Table{
		Name:       "post_subreddit",
		Columns:    PostSubredditColumns,
		PrimaryKey: []*schema.Column{PostSubredditColumns[0], PostSubredditColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_subreddit_post_id",
				Columns:    []*schema.Column{PostSubredditColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_subreddit_subreddit_id",
				Columns:    []*schema.Column{PostSubredditColumns[1]},
				RefColumns: []*schema.Column{SubredditsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		PostsTable,
		SubredditsTable,
		UsersTable,
		VerificationTokensTable,
		VotesTable,
		PostUserTable,
		PostSubredditTable,
	}
)

func init() {
	PostsTable.ForeignKeys[0].RefTable = CommentsTable
	PostsTable.ForeignKeys[1].RefTable = VotesTable
	UsersTable.ForeignKeys[0].RefTable = CommentsTable
	UsersTable.ForeignKeys[1].RefTable = SubredditsTable
	UsersTable.ForeignKeys[2].RefTable = VerificationTokensTable
	UsersTable.ForeignKeys[3].RefTable = VotesTable
	PostUserTable.ForeignKeys[0].RefTable = PostsTable
	PostUserTable.ForeignKeys[1].RefTable = UsersTable
	PostSubredditTable.ForeignKeys[0].RefTable = PostsTable
	PostSubredditTable.ForeignKeys[1].RefTable = SubredditsTable
}
