package main

import (
	"Go-000/Week02/data"
	"database/sql"
	"errors"
	"fmt"
)

/*
type Post struct {
	Id      int
	Content string
	Author  string
}
*/
/*
type AppError struct {
	Message string
	Err     error
}
*/
func main() {
	db, err := data.Posts(4)
	switch {
	case err == sql.ErrNoRows:
		errors.Is(err, sql.ErrNoRows)
		fmt.Println("这是一个sql.ErrNoRows")
	case err != nil:
		fmt.Println("查询错误:", err)
	default:
		fmt.Println(db, err)
	}

}
