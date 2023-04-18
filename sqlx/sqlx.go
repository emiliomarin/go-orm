package sqlx

import (
	"log"

	"time"

	"github.com/emiliomarin/go-orm/post"
	"github.com/gofrs/uuid"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Run will insert a new object, get it and then
// find objects with a specific filter
func Run() {
	db := setupConnection()
	defer func() {
		db.DB.Close()
	}()

	p, _ := insertData(db)
	getData(db, p.ID)
}

func insertData(db *sqlx.DB) (*post.Post, *post.Comment) {
	p := &post.Post{
		ID:        uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now(),
		Author:    "some-author",
		Content:   "some-content",
	}

	_, err := db.Exec("INSERT INTO go_orm.post (id, created_at, author, content) VALUES ($1, $2, $3, $4)", p.ID, p.CreatedAt, p.Author, p.Content)
	if err != nil {
		log.Fatal(err)
	}

	c := &post.Comment{
		ID:        uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now(),
		Author:    "some-author",
		Content:   "some-content",
		PostID:    p.ID,
	}
	_, err = db.Exec("INSERT INTO go_orm.comment (id, created_at, author, content, post_id) VALUES ($1, $2, $3, $4, $5)", c.ID, c.CreatedAt, c.Author, c.Content, c.PostID)
	if err != nil {
		log.Fatal(err)
	}
	return p, c
}

func getData(db *sqlx.DB, postID uuid.UUID) {
	// It is also possible to scan the rows like in sql and pgx and populate the struct ourselves
	p := []post.PostComment{}
	err := db.Select(&p, "SELECT * FROM go_orm.post p JOIN go_orm.comment c ON c.post_id=p.id WHERE p.id=$1", postID)
	if err != nil {
		log.Fatal(err)
	}
}

func setupConnection() *sqlx.DB {
	// Setup connection
	db, err := sqlx.Connect("postgres", "user=arexdb_dev password=arexdb_dev dbname=go_orm sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
