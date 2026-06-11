package main

import (
	"fmt"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"os"
)

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
	fmt.Printf("Работает как %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	upd := bot.GetUpdatesChan(u)
	for update := range upd {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сергей")
			bot.Send(msg)
		}
	}
}
