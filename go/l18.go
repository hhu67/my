package main

import (
	"fmt"
	// "os"
	// "github.com/joho/godotenv"
	"encoding/json"
	"net/http"
	"time"
	"bytes"
	"log"
)

type Send struct{
	Name string `json:"name"`
	Age int `json:"id"`
}

func main() {
	fmt.Println("Name for body")
	var name string
	fmt.Scanln(&name)
	data := Send{
		Name: name,
		Age: 66,
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	req, err := http.NewRequest("POST", "https://api.restful-api.dev/objects", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer resp.Body.Close()
	fmt.Println("Status code", resp.StatusCode)
}