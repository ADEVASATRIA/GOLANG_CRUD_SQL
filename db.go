package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err  := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/la_golang")
	// if err != nil {
	// 	panic(err.Error())
	// }

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
	} else {
		fmt.Println("Successfully connected to the database")
	}
	insert(db)
	defer db.Close()
}

func insert(db *sql.DB) {
	insert, err := db.Query("INSERT INTO la_user(nim , nama , email ,no_hp)VALUES(2502012464,'adeva','adevawibawa@gmail.com',081280844596)")
	if err != nil {

		panic(err.Error())
		
		} else {
		
		insert.Close()
		
		}
}
