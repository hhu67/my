package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

type OpenWeather struct {
	City string `json:"name"`
	Main struct {
		Temp     float64 `json:"temp"`
		FeelLike float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		MainWeather string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

func WeatherReq() (string, float64, string, float64, float64, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("upload your .env file %v\n", err)
	}
	api := os.Getenv("openweathermap")
	if api == "" {
		log.Fatalf("add your openweathermap api in .env file\n")
	}
	city := os.Getenv("city_openweather")
	if city == "" {
		log.Fatalf("add your city for openweathermap in .env file\n")
	}
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	lin := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=metric&appid=" + api
	url := strings.ReplaceAll(lin, " ", "%20")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("openweathermap return %d\n", resp.StatusCode)
	}
	defer resp.Body.Close()
	var WeatherData OpenWeather
	err = json.NewDecoder(resp.Body).Decode(&WeatherData)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	return WeatherData.City, WeatherData.Main.Temp, WeatherData.Weather[0].MainWeather, WeatherData.Wind.Speed, WeatherData.Main.FeelLike, WeatherData.Weather[0].Description
}

func Send(client mqtt.Client, topic string) {
	city, temp, main, speed, feel, description := WeatherReq()
	payload := fmt.Sprintf("City: %s, Temperature: %.2f°C, Feel like %.2f°C\nWeather: %s, %s, Wind %.2fm/s", city, temp, feel, main, description, speed)

	token := client.Publish(topic, 1, false, payload)
	token.Wait()
	if err := token.Error(); err != nil {
		log.Fatalf("%v\n", token.Error())
	} else {
		fmt.Println("Sent in", topic)
	}
}

func main() {
	opts := mqtt.NewClientOptions()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	broker := os.Getenv("broker_mqtt")
	if broker == "" {
		log.Fatalf("add yout broker in .env file\n")
	}
	ClientID := os.Getenv("clientid_mqtt")
	if ClientID == "" {
		log.Fatalf("add your client id in .env file")
	}
	topic := os.Getenv("topic_mqtt")
	if topic == "" {
		log.Fatalf("add your topic in .env file")
	}
	username := os.Getenv("login_mqtt")
	if username != "" {
		opts.SetUsername(username)
	}
	password := os.Getenv("password_mqtt")
	if password != "" {
		opts.SetPassword(password)
	}
	TimeHour := os.Getenv("TimeHourSendMQTT")

	opts.AddBroker(broker)
	opts.SetClientID(ClientID)
	opts.SetConnectTimeout(5 * time.Second)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("%v\n", token.Error())
	}
	defer client.Disconnect(250)

	fmt.Println("Connect to a broker")
	var ticker *time.Ticker
	if TimeHour != "" {
		ReallyTimeHour, err := strconv.Atoi(TimeHour)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		ticker = time.NewTicker(time.Duration(ReallyTimeHour) * time.Hour)
	} else {
		ticker = time.NewTicker(2 * time.Hour)
	}
	defer ticker.Stop()
	Send(client, topic)

	for range ticker.C {
		Send(client, topic)
	}
}
