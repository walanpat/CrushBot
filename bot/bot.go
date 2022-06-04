package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/config"
	"regexp"
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
	if strings.Contains(m.Content, "!roll") {
		//var message = m.Content
		re := regexp.MustCompile(`(?:!roll)\s(\d+)?d(\d+)([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?([\+\-]\d+)?`)
		variablesArr := re.FindAllStringSubmatch(m.Content, -1)
		//variablesArr[0]

		fmt.Println(variablesArr)

		//Here's the thought process Alan.  Write it such that we take in
		//xdzz+a+b+c+d.... STOP at dzz^2 (the d is the stopping letter, triggers a stop)
		//if, another dzz roll is detected, trim previous one by 1
		//(based on the fact I'm running into this silly issue)

		//Regex expression that takes everything up until we encounter a die indicator.
		//trim the last digit that's incorrectly added

		//we will run into matches based on that.
		//then, run a for loop for the number of matches,
		//Convert that into dice rolling baby.

		//Current REGEX Expression in the works
		//(\d+)*d(\d+)([\+\-]\d[^d]*)*
		//It returns all additional +/- constants being added to a dice roll,
		//it DOES NOT fix our 1d20+3d20 issue.

		//for i := 0; i < len(variablesArr); i++ {
		//	for j := 0; j < len(variablesArr); j++ {
		//		_, _ = s.ChannelMessageSend(m.ChannelID, variablesArr[i][j])
		//	}
		//}

		//_, _ = s.ChannelMessageSend(m.ChannelID, re.FindAllStringSubmatch(m.Content, -1)[0][3])

		//fmt.Println("Hit")

		//var dIndex = strings.Index(m.Content, "d")
		//var response string
		//var amountOfRolls, amountError = strconv.Atoi(m.Content[6:dIndex])
		//var diceRolled int
		//var diceError error
		//
		//if len(m.Content) == 10 {
		//	diceRolled, diceError = strconv.Atoi(m.Content[dIndex+1 : dIndex+3])
		//}
		//if len(m.Content) == 11 {
		//	diceRolled, diceError = strconv.Atoi(m.Content[dIndex+1 : dIndex+4])
		//}
		//if len(m.Content) == 12 {
		//	diceRolled, diceError = strconv.Atoi(m.Content[dIndex+1 : dIndex+5])
		//}
		//
		////fmt.Println("length of !roll 1d20 " + strconv.Itoa(len(m.Content)))
		////fmt.Println("amount of rolls " + strconv.Itoa(amountOfRolls))
		////fmt.Println("dice rolled " + strconv.Itoa(diceRolled))
		//
		//if amountError == nil && diceError == nil && diceRolled <= 100 && len(m.Content) < 13 {
		//	response += m.Author.Username + " has rolled.....\n"
		//	for i := 0; i < amountOfRolls; i++ {
		//		response += "Roll " + strconv.Itoa(i+1) + " Value: " + strconv.Itoa(rand.Intn(diceRolled-1)+1) + "\n"
		//	}
		//	_, _ = s.ChannelMessageSend(m.ChannelID, response)
		//} else if diceRolled > 100 {
		//	_, _ = s.ChannelMessageSend(m.ChannelID, "I don't own anything higher than a d100, get your own dice.")
		//
		//} else {
		//	_, _ = s.ChannelMessageSend(m.ChannelID, "There was an error in your roll request.")
		//
		//}

	}
}
