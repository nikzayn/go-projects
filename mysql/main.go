package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("GO Mysql")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/sales")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users VALUES('NIKHIL')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("You have successfully insert the query in the table")
}
