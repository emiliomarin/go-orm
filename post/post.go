package post

import (
	"time"

	"github.com/gofrs/uuid"
)

const (
	PostTable    = "go_orm.post"
	CommentTable = "go_orm.comment"
)

type Post struct {
	ID        uuid.UUID
	CreatedAt time.Time `db:"created_at"`
	Author    string
	Content   string
	Comments  []Comment `gorm:"->;foreignKey:PostID;references:ID"`
}

type Comment struct {
	ID        uuid.UUID
	CreatedAt time.Time `db:"created_at"`
	Author    string
	Content   string
	PostID    uuid.UUID `db:"post_id"`
}

// This is a ad-hoc table for joins needed for
// sqlx to populate the comment data
type PostComment struct {
	Post
	Comment
}
