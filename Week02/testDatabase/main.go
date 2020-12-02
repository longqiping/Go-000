package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=10.3.134.110 user=greek dbname=greek password=Greek#007 sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	var id int = 100
	var content string
	var author string
	err := Db.QueryRow("SELECT content, author FROM posts where id=$1", id).Scan(&content, &author)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println(err)
		fmt.Println(errors.Is(err, sql.ErrNoRows))
	case err != nil:
		fmt.Println("query error:", err)
		fmt.Println(err)
	default:
		fmt.Printf("content is %v, author is %v\n", content, author)
		fmt.Println(err)
	}

}
