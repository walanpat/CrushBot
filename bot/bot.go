package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/config"
	"math/rand"
	"strconv"
	"strings"
)

var BotId string
var goBot *discordgo.Session

func Start() {

	//creating new bot session
	goBot, err := discordgo.New("Bot " + config.Token)

	//Handling error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Making our bot a user using User function .
	u, err := goBot.User("@me")
	//Handlinf error
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Storing our id from u to BotId .
	BotId = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	//Error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//If every thing works fine we will be printing this.
	fmt.Println("Bot is running !")
}

//Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == BotId {
		return
	}
	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

	//Here is our code specifically for responding to a roll request
	if strings.Contains(m.Content, "!roll ") {
		var dIndex = strings.Index(m.Content, "d")
		var response string
		var amountOfRolls, amountError = strconv.Atoi(m.Content[6:dIndex])
		var diceRolled int
		var diceError error

		if len(m.Content) == 10 {
			diceRolled, diceError = strconv.Atoi(m.Content[dIndex+1 : dIndex+3])
		}
		if len(m.Content) == 11 {
			diceRolled, diceError = strconv.Atoi(m.Content[dIndex+1 : dIndex+4])
		}
		if len(m.Content) == 12 {
			diceRolled, diceError = strconv.Atoi(m.Content[dIndex+1 : dIndex+5])
		}

		//fmt.Println("length of !roll 1d20 " + strconv.Itoa(len(m.Content)))
		//fmt.Println("amount of rolls " + strconv.Itoa(amountOfRolls))
		//fmt.Println("dice rolled " + strconv.Itoa(diceRolled))

		if amountError == nil && diceError == nil && diceRolled <= 100 {
			response += m.Author.Username + " has rolled.....\n"
			for i := 0; i < amountOfRolls; i++ {
				response += "Roll " + strconv.Itoa(i+1) + " Value: " + strconv.Itoa(rand.Intn(diceRolled-1)+1) + "\n"
			}
			_, _ = s.ChannelMessageSend(m.ChannelID, response)
		} else if diceRolled > 100 {
			_, _ = s.ChannelMessageSend(m.ChannelID, "I don't own anything higher than a d100, get your own dice.")

		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "There was an error in your roll request.")

		}

	}
}
