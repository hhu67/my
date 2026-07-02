package main

import (
	"fmt"
	//"net/http"
	//"io"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	fmt.Printf("Your name? ")
	var username string
	fmt.Scan(&username)
	userTime := fmt.Sprintf("%v", time.Now())
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS users (id INTEGER NOT NULL PRIMARY KEY, name TEXT, time TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("INSERT INTO users(name, time) VALUES(?, ?)", username, userTime)
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT id, name, time FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var datatime string
		rows.Scan(&id, &name, &datatime)
		fmt.Printf("Name %s, ID %d, Time added %s\n", name, id, datatime)
	}
}
