package telegram

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string, webhook string) (error error) {

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = tgbotapi.NewWebhook(webhook + bot.Token)
	if err != nil {
		return err
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		return err
	}

	log.Printf("bot.GetWebhookInfo(): %v\n", info)

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServe("0.0.0.0:8091", nil)

	telegramFSM := newTelegramFSM()
	event := SayHello
	context := PhoneContext{}
	for update := range updates {
		if update.Message != nil { // If we got a message
			if event == GetPhone {
				if update.Message.Contact != nil {
					context = PhoneContext{number: update.Message.Contact.PhoneNumber}
				}
			}
			if event == SayHello {
				telegramFSM.SendEvent(event, nil)
			} else {
				telegramFSM.SendEvent(event, context)
			}

			log.Printf("[%s] %s", update.Message.From.UserName, context.Answer())
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, context.Answer())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}

	return nil
}
