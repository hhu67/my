package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func add(text string, govno string) {
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil { panic(err) }
	defer db.Close()
	sqlStmt := `CREATE TABLE IF NOT EXISTS l2 (id INTEGER NOT NULL PRIMARY KEY, name TEXT, govno TEXT)`
	_, err = db.Exec(sqlStmt)
	if err != nil {panic(err)}
	_, err = db.Exec("INSERT INTO l2 (name, govno) VALUES(?, ?)", text, govno)
	if err != nil { panic(err) }
}

func view() {
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {panic(err)}
	defer db.Close()
	rows, err := db.Query("SELECT id, name, govno FROM l2")
	if err != nil {panic(err)}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var govno string
		rows.Scan(&id, &name, &govno)
		StringView := fmt.Sprintf("ID %d, Name %s, Thought %s\n", id, name, govno)
		fmt.Printf("%s", StringView)
	}
}

func DeleteByID(id string) {
	db, err := sql.Open("sqlite3", "./base.db")
	if err != nil {panic(err)}
	_, err = db.Exec("DELETE FROM l2 WHERE id = ?", id)
	if err != nil {panic(err)}
	fmt.Printf("delete by id %s\n", id)
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("1) Added in DB\n2) View all DB\n3) Delete by index\n4) Exit\n")
		fmt.Printf("Your choise(1/2/3/4) ")
		var choise int
		fmt.Scan(&choise)
		if choise == 1 {
			fmt.Printf("Your name ")
			scan.Scan()
			name := scan.Text()
			fmt.Printf("enter the thought ")
			scan.Scan()
			thought := strings.TrimSpace(scan.Text())
			add(name, thought)
			fmt.Printf("Added %s and %s\n", name, thought)
		} else if choise == 2 {
			fmt.Println("------------------------")
			view()
			fmt.Println("------------------------")
		} else if choise == 4 {
			return
		} else if choise == 3 {
			fmt.Printf("ID ")
			scan.Scan()
			id := scan.Text()
			DeleteByID(id)
		} else {
			fmt.Printf("Incorrect number\n ")
		}

	}
}
