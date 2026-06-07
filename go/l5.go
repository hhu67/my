package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
		return
	}
	api := os.Getenv("api_key")
	fmt.Println(api)
}
