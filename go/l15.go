package main

import (
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"
)
type MyIpData struct{
	YourIP string `json:"ip"`
	City string `json:"city"`
	Country string `json:"country"`
	ISP string  `json:"org"`
}
type UA struct{
	UserAgent string `json:"user-agent"`
}
func main() {
	fmt.Println("Введите какой хотите user-agent")
	var polz string
	fmt.Scanln(&polz)
	client := &http.Client{
		Timeout: 10 * time.Second,
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
	req2, err := http.NewRequest("GET", "https://ipinfo.io", nil)
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