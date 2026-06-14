package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
	"log"
	"net/http"
	"net/url"
	"encoding/json"
	"crypto/tls"
	"time"
)

const (
	var Step string
	var Proxy string
)

type DataIP struct {
	IP string `json:"query"`
	Country string `json:"country"`
}

func ProxyCheck(proxy) (bool, string, string) {
	proxyurl, err := url.Parse(proxy)
	if err != nil {
		fmt.Println(err)
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: http.ProxyURL(proxyurl),
	}
	client := &http.Client{
		Transport: transport,
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest("GET", "http://ip-api.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, "", ""
		fmt.Println(err)
	}
	defer resp.Body.Close
	var data DataIP
	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return true, data.IP, data.Country
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	tg := os.Getenv("TG_TOKEN")
	bot, err := tgbotapi.NewBotAPI(tg)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {}
		
