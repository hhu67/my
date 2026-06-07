package main

import (
	"fmt"
	"encoding/json"
	"log"
	"os"
)

type Json struct{
	Name string `json:"name"`
	Version int `json:"ver"`
	Active bool `json:"act"`
}

func main() {
	jsonn, err := os.ReadFile("l13.json")
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}
	var conf Json
	err = json.Unmarshal(jsonn, &conf)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}
	if conf.Active {
		fmt.Println("Статус:", conf.Active)
	}
}
