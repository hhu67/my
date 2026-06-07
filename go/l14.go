package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/joho/godotenv"
	"os"
)

type IPR struct{
	IP string `json:"origin"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	proxy := os.Getenv("proxy")
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Fatalln(err)
	}
	tran := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	myclient := &http.Client {
		Transport: tran,
	}
	url, err := myclient.Get("https://httpbin.org/ip")
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer url.Body.Close()
	var ipresp IPR
	err = json.NewDecoder(url.Body).Decode(&ipresp)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Println(ipresp.IP)
}
