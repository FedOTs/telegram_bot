package main

import (
	"log"
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

	telegram.StartBot(conf.Telegram.Token, conf.Telegram.WebHook)

	if err != nil {
		log.Printf("Error telegram bot %v", err)
	}

}
