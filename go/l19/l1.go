package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type User struct{
	Age int `json:"age"`
	Name string `json:"name"`
	Email string `json:"email"`
	Alive bool `json:"alive"`
}

var setChan = make(chan User)

func handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not",  http.StatusMethodNotAllowed)
		return
	}

	var data User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w,"no", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	i := 0
	if data.Age > 150 {
		i++
	}
	CheckEmail := strings.ContainsAny(data.Email, "@.")
	if CheckEmail != true {
		i++
	}
	if data.Alive != true {
		i++
	}
	if i != 0 {
		msg := fmt.Sprintf("Error: %d", i)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	
	setChan <- data
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received"))
}

func WatchChanell() {
	for u := range setChan {
		db, err := sql.Open("sqlite3", "./base.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		sqlStmt :=  `CREATE TABLE IF NOT EXISTS l1 (
			id INTEGER NOT NULL PRIMARY KEY,
			age INTEGER DEFAULT 0,
			name TEXT,
			email TEXT,
			alive INTEGER NOT NULL DEFAULT 0 CHECK (alive IN (0, 1))
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			panic(err)
		}
		var Alive int
		if u.Alive {
			Alive = 1
		} else {
			Alive = 0
		}
		_, err = db.Exec("INSERT INTO l1 (age, name, email, alive) VALUES(?, ?, ?, ?)", u.Age, u.Name, u.Email, Alive)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		rows, err := db.Query("SELECT id, age, name, email, alive FROM l1")
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var id int
			var age int
			var name string
			var email string
			var alive int
			var boolAlive bool
			rows.Scan(&id, &age, &name, &email, &alive)
			if alive == 1 {
				boolAlive = true
			}
			if alive == 0 {
				boolAlive = false
			}
			StrAnswer := fmt.Sprintf("ID %d, Age %d, Name %s, Email %s, Alive %t\n", id, age, name, email, boolAlive)
			fmt.Printf("%s", StrAnswer)
		}
	}
}

func main() {
	go WatchChanell()
	http.HandleFunc("/user", handle)
	fmt.Println(":7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		panic(err)
	}
}
