package main

import (
	"fmt"
	//"net/http"
	//"io"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Printf("Your name? ")
	var username string
	fmt.Scan(&username)
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users(name) VALUES(?)", username)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Printf("Name %s, ID %d\n", name, id)
	}
}
