package main

import (
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

	info, err := bot.GetWebhookInfo()

	if err != nil {
		log.Printf("bot.GetWebhookInfo() err : %v", err)
	}

	log.Printf("bot.GetWebhookInfo(): %v\n", info)

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServe("0.0.0.0:8091", nil)

	for update := range updates {
		//log.Printf("update %v\n", update)
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			//msg.ReplyToMessageID = update.Message.MessageID

			//bot.Send(msg)
		}
	}

}
