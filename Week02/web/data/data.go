package data

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=10.3.134.110 user=greek dbname=greek password=Greek#007 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func Posts(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, content, author FROM posts where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
