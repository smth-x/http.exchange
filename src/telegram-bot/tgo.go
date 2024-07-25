package telegram_bot

import (
	tg "github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegoutil"
	ut "http.exchange/src/utils"
	"log"
	"strings"
)

var (
	Bot    *tg.Bot
	rateOf string
	rateTo string
)

func CreateBot(token string) error {
	newBot, err := tg.NewBot(token, tg.WithDefaultDebugLogger())
	if !ut.IsErrNil(err) {
		return err
	}
	Bot = newBot
	return nil
}
func StartLongPolling() {
	updates, _ := Bot.UpdatesViaLongPolling(nil)
	defer Bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			err := handle(update)
			log.Printf("handling error %v\n", err)
		}
	}
}
func handle(input tg.Update) error {
	chatID := input.Message.Chat.ID
	message := input.Message.Text
	sendMessage := func(text string, withKeyboard bool) error {
		msg := telegoutil.Message(tg.ChatID{ID: chatID}, text)
		if withKeyboard {
			msg.WithReplyMarkup(KeyBoard)
		}
		_, err := Bot.SendMessage(msg)
		if !ut.IsErrNil(err) {
			return err
		}
		return nil
	}
	log.Printf("[%s - %d] %s", input.Message.From.Username, chatID, message)

	if strings.EqualFold(message, "/start") {
		err := sendMessage("Welcome! Write /rate to find out exchange rate of some currency", false)
		if !ut.IsErrNil(err) {
			return err
		}
	} else if strings.EqualFold(message, "/rate") {
		msg := telegoutil.Message(tg.ChatID{ID: chatID}, "What currency exchange rate are you interested in?").WithReplyMarkup(KeyBoard)
		_, err := Bot.SendMessage(msg)
		if !ut.IsErrNil(err) {
			return err
		}
		return nil
	}

	return nil
}
