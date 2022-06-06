package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/config"
	"math/rand"
	"regexp"
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
	if strings.Contains(m.Content, "!roll") {
		re := regexp.MustCompile(`([\+\-]?\d+)*d(\d+)([\+\-]?\d*[^\dd][^d]+)*`)
		variablesArr := re.FindAllStringSubmatch(m.Content, -1)

		fmt.Println(variablesArr)
		message := ""

		for i := 0; i < len(variablesArr); i++ {
			//Alright, our baseline tomfoolery is up to speed.
			//You can roll 1 dice with any number of modifiers (positive or negative) here
			if len(variablesArr) == 1 {
				fmt.Println("TESTING BASE LINE")
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
				basicMathRE := regexp.MustCompile(`([\+\-]?\d*)`)
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
				fmt.Println(arithmetic)
				message += m.Author.Username + " LETS ROLL\n"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					message += "Roll " + strconv.Itoa(j+1) + ": " + strconv.Itoa(rand.Intn(diceToBeRolled+1)) + "\n"
				}
				arithmeticResult := 0

				for j := 0; j < len(arithmetic); j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
				if arithmeticResult != 0 {
					message += "Modifiers: " + strconv.Itoa(arithmeticResult)
				}

			} else if i == 0 && len(variablesArr) > 1 {
				fmt.Println("hit 0")
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

				basicMathRE := regexp.MustCompile(`([\+\-]?\d*)*`)
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
				//fmt.Println(arithmetic)
				message += m.Author.Username + " LETS ROLL\n"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					message += "Roll " + strconv.Itoa(j+1) + ": " + strconv.Itoa(rand.Intn(diceToBeRolled+1)) + "\n"
				}
				arithmeticResult := 0
				for j := 0; j < len(arithmetic)-1; j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
				if arithmeticResult != 0 {
					message += "Modifiers: " + strconv.Itoa(arithmeticResult)
				}

			} else {
				fmt.Println("HIT HIT HIT")
				//Code for if there are MULTIPLE types of dice rolls in a single situation.
				basicMathRE := regexp.MustCompile(`([\+\-]?\d*)*`)

				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				//Make sure this pulls the correct sign if needeD????????????? not sure if necessary
				if i != 0 {
					numbOfRoll, _ = strconv.Atoi(basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)[0][len(basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)[0])-1])
				}
				//this has to be initialized weirdly.  So.  Here's what we do
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][2], -1)
				message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "\n"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					message += "Roll " + strconv.Itoa(j+1) + ": " + strconv.Itoa(rand.Intn(diceToBeRolled+1)) + "\n"
				}
				arithmeticResult := 0
				if i+1 != len(variablesArr) {
					for j := 1; j < len(arithmetic)-1; j++ {
						arithmeticResult, _ = strconv.Atoi(arithmetic[0][j])
					}
				} else {
					for j := 1; j < len(arithmetic); j++ {
						arithmeticResult, _ = strconv.Atoi(arithmetic[0][j])
					}
				}
				if arithmeticResult != 0 {
					message += "Modifiers: " + strconv.Itoa(arithmeticResult)
				}
			}

		}
		_, _ = s.ChannelMessageSend(m.ChannelID, message)

	}
}
