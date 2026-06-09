package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"github.com/joho/godotenv"
	"os"
	"net/url"
	"crypto/tls"
)
type MyIpData struct{
	YourIP string `json:"query"`
	City string `json:"city"`
	Country string `json:"country"`
	ISP string  `json:"isp"`
}
type UA struct{
	UserAgent string `json:"user-agent"`
}

func proxy() (bool, string) {
	err := godotenv.Load()
	if err != nil {
		return false, ""
	}
	proxy := os.Getenv("proxy")
	if proxy == "" {
		return false, ""
	}
	return true, proxy
}

func main() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	er, proxis := proxy()
	if er {
		proxyURL, err := url.Parse(proxis)
		if err != nil {
			return
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	fmt.Println("Введите какой хотите user-agent")
	var polz string
	fmt.Scanln(&polz)
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest("GET", "https://httpbin.org/user-agent", nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	req.Header.Set("User-Agent", polz)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	req2, err := http.NewRequest("GET", "http://ip-api.com/json", nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	resp2, err := client.Do(req2)
	defer resp.Body.Close()
	var file UA
	err = json.NewDecoder(resp.Body).Decode(&file)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	var data MyIpData
	err = json.NewDecoder(resp2.Body).Decode(&data)
	fmt.Printf("Ваш IP: %s\nВаш город по IP: %s\nВаша страна по IP: %s\nВаш провайдер по IP: %s\nВаш User-Agent: %s\n", data.YourIP, data.City, data.Country, data.ISP, file.UserAgent)
}