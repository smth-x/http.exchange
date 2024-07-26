package main

import (
	"gopkg.in/ini.v1"
	tg_bot "http.exchange/src/telegram-bot"
	ut "http.exchange/src/utils"
	"log"
	"os"
)

func main() {
	run()
}
func run() {
	err := tg_bot.CreateBot(getValueFromIniFile("token.ini", "", "token")) //sectionInFile have to be "", case token placed not in sections
	if !ut.IsErrNil(err) {
		log.Fatalln(err)
	}
	tg_bot.StartLongPolling()
}

func getValueFromIniFile(iniFileName, sectionInFile, key string) string {
	file, err := os.Open(iniFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := ini.Load(buffer)
	if err != nil {
		log.Fatal(err)
	}

	section, _ := cfg.GetSection(sectionInFile)
	k, _ := section.GetKey(key)
	return k.String()
}
