package main

import (
	_const "http.exchange/src/const"
	tg_bot "http.exchange/src/telegram-bot"
	ut "http.exchange/src/utils"
	"log"
)

func main() {
	run()
}
func run() {
	err := tg_bot.CreateBot(_const.Token)
	if !ut.IsErrNil(err) {
		log.Fatalln(err)
	}
	tg_bot.StartLongPolling()
}
