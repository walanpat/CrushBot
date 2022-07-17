package main

import (
	"fmt"
	"goland-discord-bot/bot"
	"goland-discord-bot/config"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
