package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"bytes"
	"io"
	"strings"
)

type UserStruct struct{
	Age int `json:"age"`
	Name string `json:"name"`
	Email string `json:"email"`
	Alive bool `json:"alive"`
}

func polz() (int, string, string, bool) {
	fmt.Printf("Your age? ")
	var age int
	fmt.Scan(&age)
	fmt.Printf("Your name? ")
	var name string
	fmt.Scan(&name)
	fmt.Printf("Your email? ")
	var email string
	fmt.Scan(&email)
	fmt.Printf("Your alive?(y/n) ")
	var live string
	fmt.Scan(&live)
	LowerLive := strings.ToLower(live)
	var alive bool
	if LowerLive == "y" {
		alive = true
	} else if LowerLive == "n" {
		alive = false
	}
	return age, name, email, alive
}

func main() {
	age, name, email, alive := polz()
	user := UserStruct{
		Age: age,
		Name: name,
		Email: email,
		Alive: alive,
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
