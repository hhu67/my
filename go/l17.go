package main

import (
	"log"
	"github.com/joho/godotenv"
	"os"
	"strings"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
)

type OpenWeather struct{
	City string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
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
	var WeatherData OpenWeather
	err = json.NewDecoder(resp.Body).Decode(&WeatherData)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Println(WeatherData.City, WeatherData.Main.Temp)
}