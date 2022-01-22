package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished executing
	defer db.Close()

	fmt.Println("Database opened")

}
