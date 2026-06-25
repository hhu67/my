package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"bytes"
	"io"
)

type UserStruct struct{
	Age int `json:"age"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := UserStruct{
		Age: 66,
		Name: "Sergey",
		Email: "Likhodeev.60@mail.ru",
	}
	JsonData, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:7777/user", bytes.NewBuffer(JsonData))
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status code %d\n", resp.StatusCode)
	fmt.Println(string(body))
}