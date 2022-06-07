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

	//The key to understanding this is understanding how GO uses regexp.
	//If you can understand how it returns regex using the .FindAllStringSubmatch
	//Then you can understand this code

	//To explain it as basic as I can,
	//We are reading a input, (example 1d20+2+6d6)
	//We have 3 if statements here,
	//1.If  there's just a single 1d20 or 3d20 or 900d2 dice roll and modifiers is the first IF block
	//If there's more than 1d20+9+4-2 specifically it looks more like 6d6+1d20+9d10
	//then, we have to use the next 2 if statements
	//2.The next 2 "if" statements will read and calculate and NOT interact with the ending number of a
	//The issue we run into is that our regexp will read 1d20+6d6 and return
	//[[1d20+6 1 20 +6][d6 6]]
	//This causes a problem
	//I had to write more complicated logic to get around this issue and watch our edge cases so that
	//we don't do anything insane like adding a "# of rolls" variable to our basic addition modifiers

	if strings.Contains(m.Content, "ff") {
		message := "```ansi\n"
		////message += "**bold** "
		////message += " *italics* "
		////message += "__underlined text__ "
		////message += "`Highlighted text` "
		////message += "\n > quote text \n"
		////message += "~~strikethrough text ~~ \n"
		//message += " ```diff\n- Discord red text\n```"
		//message += "```css\n[Discord orange-red text]\n```"
		//message += "```fix\nDiscord yellow text\n```"
		//message += "```apache\nDiscord_dark_yellow_text\n```"
		//message += "```css\n.Discord_blue_text\n```"
		//message += "```ini\n[Discord blue text]\n```"
		//message += "```diff\n+ Discord light green text\n```"
		//message += "```yaml\nCyan text in Discord\n```"
		//message += "\n[](Red text in Discord)\n```"

		//message += "\n- Red text in Discord\n+ Light green text in Discord\n```"

		//This is the key to what we want to do.  CHECK THIS SHIT KING

		message += "\u001B[0;33m\u001B[0m.\u001B[31m(\u001B[36m[\u001B[34m\\w\u001B[0m.@\u001B[36m]\u001B[34m+\u001B[31m)\u001B[0m!\u001B[34m?\u001B[36m[\u001B[0m+-\u001B[36m]\u001B[34m?\u001B[0m\n\n"
		message += "\u001B[31m("

		message += "\n```"

		_, _ = s.ChannelMessageSend(m.ChannelID, message)
	}

	if strings.Contains(m.Content, "!roll") {
		//Initializing our "base" regex expression
		re := regexp.MustCompile(`([\+\-]?\d+)*d(\d+)([\+\-]?\d*[^\dd][^d]+)*`)
		variablesArr := re.FindAllStringSubmatch(m.Content, -1)
		message := "```\n"
		//message := ""
		sumTotal := 0
		for i := 0; i < len(variablesArr); i++ {
			total := 0
			if len(variablesArr) == 1 {
				//This IF statement only occurs if there is a single "roll" (x^n)d(y^n)+(z^n) occurring
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

				//Basic math, modifier regexp
				basicMathRE := regexp.MustCompile(`([\+\-]?\d*)`)
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
				message += m.Author.Username + " LETS ROLL\n"
				message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "		"

				//Actual for loop for the "dice rolls"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					rollValueInt := rand.Intn(diceToBeRolled) + 1
					rollValueStr := strconv.Itoa(rollValueInt)
					total += rollValueInt
					if rollValueInt == 1 {
						message += "Roll " + strconv.Itoa(j+1) + ": [" + "\u001B[31m" + rollValueStr + "\u001B[0;33m\u001B[0m]		"

					} else if rollValueInt == diceToBeRolled {
						message += "Roll " + strconv.Itoa(j+1) + ": [__**`" + rollValueStr + "`**__]		"

					} else {
						message += "Roll " + strconv.Itoa(j+1) + ": [`" + rollValueStr + "`]		"

					}
				}
				arithmeticResult := 0
				//Adds any and all modifiers
				for j := 0; j < len(arithmetic); j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
				if arithmeticResult != 0 {
					total += arithmeticResult
					message += "Modifiers: [" + strconv.Itoa(arithmeticResult) + "]"
				}
				message += "\nTotal Roll: " + "[" + strconv.Itoa(total) + "]\n"

			} else if i == 0 && len(variablesArr) > 1 {
				//this code block occurs only if (x^n)d(y^n)+(z^n) occurs more than once.
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

				//used to find any modifiers added or subtracted to various dice roll
				basicMathRE := regexp.MustCompile(`([\+\-]?\d*)`)
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)

				message += m.Author.Username + " LETS ROLL\n"
				message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "		"

				//Actual for loop for the "dice rolls"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					rollValueInt := rand.Intn(diceToBeRolled) + 1
					rollValueStr := strconv.Itoa(rollValueInt)
					total += rollValueInt
					sumTotal += rollValueInt
					message += "Roll " + strconv.Itoa(j+1) + ": [" + rollValueStr + "]		"
				}
				arithmeticResult := 0
				//Adds any and all modifiers, makes sure we don't add the next dice roll
				for j := 0; j < len(arithmetic)-1; j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x

				}
				if arithmeticResult != 0 {
					total += arithmeticResult
					message += "Modifiers: [" + strconv.Itoa(arithmeticResult) + "]"
					sumTotal += arithmeticResult

				}
				message += "\n Total Roll: " + "[" + strconv.Itoa(total) + "]\n"

			} else {
				//This code block is for ALL code after the original (x^n)d(y^n)+(z^n) amount of dice/things rolled.
				basicMathRE := regexp.MustCompile(`([\+\-]?\d*)`)
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])

				if i != 0 {
					initialArray := basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)
					numbOfRoll, _ = strconv.Atoi(initialArray[len(initialArray)-1][0])
				}
				//this has to be initialized weirdly.  So.  Here's what we do
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
				//message now tells us what we're rolling
				message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "		"

				//Actual for loop for the "dice rolls"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					rollValueInt := rand.Intn(diceToBeRolled) + 1
					rollValueStr := strconv.Itoa(rollValueInt)
					total += rollValueInt
					sumTotal += rollValueInt
					message += "Roll " + strconv.Itoa(j+1) + ": [" + rollValueStr + "]		"
				}
				arithmeticResult := 0
				//If, we are NOT at the last modifier/read value
				if i+1 != len(variablesArr) {
					for j := 0; j < len(arithmetic)-1; j++ {
						x, _ := strconv.Atoi(arithmetic[j][0])
						arithmeticResult += x
					}
				} else {
					//If it is the last value, then add it if there's any modifiers/arithmetic values.
					for j := 0; j < len(arithmetic); j++ {
						x, _ := strconv.Atoi(arithmetic[j][0])
						arithmeticResult += x
					}
				}
				if arithmeticResult != 0 {
					total += arithmeticResult
					sumTotal += arithmeticResult
					message += "Modifiers: [" + strconv.Itoa(arithmeticResult) + "]"
				}
				message += "\n  Total Roll: " + "[" + strconv.Itoa(total) + "]\n"
				if i+1 == len(variablesArr) {
					message += "\n Sum Total of All Values: " + strconv.Itoa(sumTotal)

				}
			}
		}
		message += "\n```"

		_, _ = s.ChannelMessageSend(m.ChannelID, message)

	}
}
