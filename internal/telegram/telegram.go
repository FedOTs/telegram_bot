package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string, webhook string) (bot *tgbotapi.BotAPI, error error) {

	bot_, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	_, err = tgbotapi.NewWebhook(webhook)
	if err != nil {
		return nil, err
	}

	return bot_, nil
}
