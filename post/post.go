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
	CreatedAt time.Time
	Author    string
	Content   string
	Comments  []Comment `gorm:"->;foreignKey:PostID;references:ID"`
}

type Comment struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Author    string
	Content   string
	PostID    uuid.UUID
}
