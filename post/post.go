package post

import (
	"time"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Author    string
	Content   string
	Comments  []Comment `gorm:"foreignkey:PostID"`
}

type Comment struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Author    string
	Content   string
	PostID    uuid.UUID
}
