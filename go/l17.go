package main

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"net/http"
	"time"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"fmt"
)

type OpenWeather struct{
	City string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func WeatherReq() (string, float64) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	api := os.Getenv("openweathermap")
	city := os.Getenv("city_openweather")
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	link := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=" + api
	url := strings.ReplaceAll(link, " ", "%20")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	defer resp.Body.Close()
	var WeatherData OpenWeather
	err = json.NewDecoder(resp.Body).Decode(&WeatherData)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return WeatherData.City, WeatherData.Main.Temp
}

func main() {
	err := godotenv.Load()
	if err != nil { log.Fatalf("%v\n", err) }
	broker := os.Getenv("broker_mqtt")
	ClientID := os.Getenv("clientid_mqtt")
	topic := os.Getenv("topic_mqtt")
	username := os.Getenv("login_mqtt")
	password := os.Getenv("password_mqtt")

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(ClientID)
	opts.SetConnectTimeout(5 * time.Second)
	opts.SetUsername(username)
        opts.SetPassword(password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil { log.Fatalf("%v\n", token.Error()) }
	defer client.Disconnect(250)
	fmt.Println("Connect to a broker")

	city, temp := WeatherReq()
	payload := fmt.Sprintf("City: %s, Temperature: %.2f°C", city, temp)

	token := client.Publish(topic, 1, false, payload)
	ticket := time.NewTicker(2 * time.Second)
	defer ticket.Stop()
	token.Wait()
	if err = token.Error(); err != nil {
		log.Fatalf("%v\n", token.Error)
	} else {
		fmt.Println("Sent in", topic)
	}
}
