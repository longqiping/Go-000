package main

import (
	"Go-000/Week02/web/data"
	"database/sql"
	"errors"
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

/*
type AppError struct {
	Message string
	Err     error
}
*/
func main() {
	db, err := data.Posts(40)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("This is a sql.errNoRows")
	} else {
		fmt.Println("This is not a sql.errNoRows")
	}

	fmt.Println(db, err)

}
