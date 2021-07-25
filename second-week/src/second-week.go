package main

/*
1. 我们在数据库操作的时候，
比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

import (
	"database/sql"
	"fmt"
)

// Das
func Dao() (int, error) {
	status, err := sql.ErrNoRows("select * from table")

	if err != nil {
		return 1, err
	}
	_ = status
}
func main() {
	check, err := Dao()

	if err != nil {
		fmt.Printf("error + %v %d", err)
	}
}
