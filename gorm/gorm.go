package gorm

import (
	"fmt"
	"log"

	"time"

	"github.com/emiliomarin/go-orm/post"
	"github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Run will insert a new object, get it and then
// find objects with a specific filter
func Run() {
	db := setupConnection()
	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	p, _ := insertData(db)
	getData(db, p.ID)
}

func insertData(db *gorm.DB) (*post.Post, *post.Comment) {
	p := &post.Post{
		ID:        uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now(),
		Author:    "some-author",
		Content:   "some-content",
	}
	if err := db.Table(post.PostTable).Create(p).Error; err != nil {
		log.Fatal(err)
	}

	c := &post.Comment{
		ID:        uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now(),
		Author:    "some-author",
		Content:   "some-content",
		PostID:    p.ID,
	}
	if err := db.Table(post.CommentTable).Create(c).Error; err != nil {
		log.Fatal(err)
	}
	return p, c
}

func getData(db *gorm.DB, postID uuid.UUID) {
	p := post.Post{}
	err := db.
		Table(post.PostTable).
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Table(post.CommentTable)
		}).
		Where("id = ?", postID).First(&p).
		Error
	if err != nil {
		log.Fatal(err)
	}

	// We could also run the query raw
	// p = post.Post{}
	// err = db.Debug().Raw("SELECT * from go_orm.post p LEFT JOIN go_orm.comment c ON c.post_id=p.id where p.id = ?", postID).Scan(&p).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func setupConnection() *gorm.DB {
	// Setup connection
	conn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		"localhost",
		"5432",
		"go_orm",
		"arexdb_dev",
		"arexdb_dev",
		"disable",
	)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
