package sql

import (
	"log"

	"time"

	"github.com/emiliomarin/go-orm/post"
	"github.com/gofrs/uuid"

	"database/sql"

	_ "github.com/lib/pq"
)

// Run will insert a new object, get it and then
// find objects with a specific filter
func Run() {
	db := setupConnection()
	defer db.Close()

	p, _ := insertData(db)
	getData(db, p.ID)
}

func insertData(db *sql.DB) (*post.Post, *post.Comment) {
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

func getData(db *sql.DB, postID uuid.UUID) {
	p := post.Post{}
	p.Comments = make([]post.Comment, 1)
	rows, err := db.Query("SELECT p.id, p.created_at, p.author, p.content, c.id, c.author, c.content, c.created_at, c.post_id FROM go_orm.post p JOIN go_orm.comment c ON c.post_id=p.id WHERE p.id=$1", postID)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err = rows.Scan(&p.ID, &p.CreatedAt, &p.Author, &p.Content, &p.Comments[0].ID, &p.Comments[0].Author, &p.Comments[0].Content, &p.Comments[0].CreatedAt, &p.Comments[0].PostID)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func setupConnection() *sql.DB {
	// Opening a driver typically will not attempt to connect to the database.
	db, err := sql.Open("postgres", "user=arexdb_dev password=arexdb_dev dbname=go_orm sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
