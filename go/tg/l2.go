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
	"strconv"
)

var userStates = make(map[int64]string)
var Proxy = make(map[int64]string)

type DataIP struct {
	IP string `json:"query"`
	Country string `json:"country"`
	ISP string `json:"isp"`
}

func ProxyCheck(proxy string) (bool, string, string, string) {
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
	req, err := http.NewRequest("GET", "http://ip-api.com/json", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, "Нет", "Нет", "Нет"
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var data DataIP
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	return true, data.IP, data.Country, data.ISP
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
	for update := range updates {
		if update.Message == nil { continue }
		userID := update.Message.From.ID
		if update.Message.IsCommand() {
       			if update.Message.Command() == "start" {
            		userStates[userID] = "wating"
            		bot.Send(tgbotapi.NewMessage(userID, "Введите прокси"))
        	}
        	continue
	}
		switch userStates[userID] {
			case "wating":
				prox := update.Message.Text
				Proxy[userID] = prox
				userStates[userID] = ""
				status, ip, country, isp := ProxyCheck(Proxy[userID])
				stat := strconv.FormatBool(status)
				msg := "Работа прокси: " + stat + "\nIP: " + ip + "\nСтрана: " + country + "\nПровайдер: " + isp
				bot.Send(tgbotapi.NewMessage(userID, msg))
			}
		}
}

