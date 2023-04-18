package main

import (
	"fmt"
	"log"
	"time"

	"github.com/emiliomarin/go-orm/gorm"
	"github.com/emiliomarin/go-orm/post"

	"github.com/gofrs/uuid"
	"gorm.io/driver/postgres"
	gormio "gorm.io/gorm"
)

func main() {
	db := setupDB()
	defer func() {
		db.Table("go_orm.comment").Delete("", "id IS NOT NULL")
		db.Table("go_orm.post").Delete("", "id IS NOT NULL")
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	gorm.Run()
}

// Setup the DB with some intiial data so it's not fully empty
func setupDB() *gormio.DB {
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
	db, err := gormio.Open(postgres.Open(conn), &gormio.Config{})
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000; i++ {
		p := post.Post{
			ID:        uuid.Must(uuid.NewV4()),
			CreatedAt: time.Now(),
			Author:    fmt.Sprintf("Author-%d", i),
			Content:   "some content",
		}

		if err := db.Table("go_orm.post").Create(&p).Error; err != nil {
			log.Fatal(err)
		}

		comment := post.Comment{
			ID:        uuid.Must(uuid.NewV4()),
			CreatedAt: time.Now(),
			Author:    fmt.Sprintf("Author-%d", i),
			Content:   "some comment",
			PostID:    p.ID,
		}
		comment2 := post.Comment{
			ID:        uuid.Must(uuid.NewV4()),
			CreatedAt: time.Now(),
			Author:    fmt.Sprintf("Author-%d", i),
			Content:   "some comment",
			PostID:    p.ID,
		}

		if err := db.Table("go_orm.comment").Create(&comment).Error; err != nil {
			log.Fatal(err)
		}
		if err := db.Table("go_orm.comment").Create(&comment2).Error; err != nil {
			log.Fatal(err)
		}
	}
	return db
}
