package telegram_bot

import (
	tg "github.com/mymmrac/telego"
	ut "http.exchange/src/utils"
	"log"
)

var Bot *tg.Bot

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
			if !ut.IsErrNil(err) {
				log.Println(err)
			}
		}
	}
}
func handle(input tg.Update) error {
	chatID := input.Message.Chat.ID
	message := input.Message.Text
	log.Println(chatID, message)

	return nil
}
