package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	PathFile := os.Getenv("html_file_for_l16.go") // Enter the file path for the html_file_for_l16.go variable into the env file
	html, err := os.ReadFile(PathFile)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Fprint(w, string(html))
}

func main() {
	http.HandleFunc("/VSZw5p9vwaCxQgU/efNCQ3Vz3L2P8ZGv3VWp/ybZAUeqSMbPpcDsc7G6e/", handler)
	err := http.ListenAndServe(":23455", nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
