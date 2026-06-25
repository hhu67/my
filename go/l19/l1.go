package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type User struct{
	Age int `json:"age"`
	Name string `json:"name"`
	Email string `json:"email"`
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
	
	setChan <- data
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received"))
}

func WatchChanell() {
	for u := range setChan {
		JsonData, err := json.Marshal(u)
		if err != nil { continue }
		fmt.Println(string(JsonData))
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
