package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"telegram_bot/internal/config"
	"telegram_bot/internal/telegram"
)

func main() {

	configString := "configs" + string(os.PathSeparator) + "config.json"

	conf, err := config.LoadConfiguration(configString)

	if err != nil {
		log.Printf("Error load configuration %v", err)
	}

	bot, err := telegram.StartBot(conf.Telegram.Token, conf.Telegram.WebHook)

	if err != nil {
		log.Printf("Error telegram bot %v", err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	http.ListenAndServe("0.0.0.0:8091", nil)

	for update := range updates {
		fmt.Println(update)
	}

	fmt.Println(conf, configString)
}
