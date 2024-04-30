package main

import (
	"fmt"
	"goland-discord-bot/bot"
	"goland-discord-bot/bot/commons"
	"goland-discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = commons.InitializeResponses()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	bot.Start()

	<-make(chan struct{})
}
