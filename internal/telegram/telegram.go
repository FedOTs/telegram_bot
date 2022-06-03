package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string, webhook string) (bot *tgbotapi.BotAPI, error error) {

	bot_, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot_.Debug = true

	log.Printf("Authorized on account %s", bot_.Self.UserName)

	_, err = tgbotapi.NewWebhook(webhook + bot_.Token)
	if err != nil {
		return nil, err
	}

	return bot_, nil
}

func SendMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) (error error) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)

	return nil
}
