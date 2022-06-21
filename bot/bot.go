package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/config"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var Id string

//Not sure if this variable/nomenclature will be needed later.  Add to cleanup list.
//var goBot *discordgo.Session
var cachedCardMessageId = ""
var cachedCardSet = ""
var cachedCardRuling = ""

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
	Id = u.ID

	// Adding handler function to handle our messages using AddHandler from discordgo package. We will declare messageHandler function later.
	goBot.AddHandler(messageHandler)
	goBot.AddHandler(reactionHandler)
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
	if m.Author.ID == Id && len(m.Reactions) == 0 && m.Content == "" {
		_ = s.MessageReactionAdd(m.ChannelID, m.Message.ID, "\U0001F4DA")
		_ = s.MessageReactionAdd(m.ChannelID, m.Message.ID, "\U0001F4C5")
	}
	//Bot musn't reply to it's own messages , to confirm it we perform this check.
	if m.Author.ID == Id {
		return
	}
	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}
	if strings.Contains(m.Content, "!roll") {
		//Initializing our "base" regex expression
		re := regexp.MustCompile(`([+\-]?\d+)*d(\d+)([+\-]?\d*[^\dd][^d]+)*`)
		variablesArr := re.FindAllStringSubmatch(m.Content, -1)
		message := "```ansi\n \u001B[0m"
		sumTotal := 0
		for i := 0; i < len(variablesArr); i++ {
			total := 0
			if len(variablesArr) == 1 {
				//This IF statement only occurs if there is a single "roll" (x^n)d(y^n)+(z^n) occurring
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

				//Basic math, modifier regexp
				basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
				message += m.Author.Username + " LETS ROLL\n\n"
				message += "\u001B[3m" + strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "\u001B[0m	"

				//Actual for loop for the "dice rolls"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					rollValueInt := rand.Intn(diceToBeRolled) + 1
					rollValueStr := strconv.Itoa(rollValueInt)
					total += rollValueInt
					if rollValueInt == 1 {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[31m" + rollValueStr + "\u001B[0m]	"
					} else if rollValueInt == diceToBeRolled {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[32m" + rollValueStr + "\u001B[0m]	"
					} else {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[37m" + rollValueStr + "\u001B[0m]	"
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
					message += "\n\n	Modifiers: [\u001B[37m" + strconv.Itoa(arithmeticResult) + "\u001B[0m]\n"
				}
				message += "\n\n	Total Roll: " + "[\u001B[37m" + strconv.Itoa(total) + "\u001B[0m]\n\n"

			} else if i == 0 && len(variablesArr) > 1 {
				//this code block occurs only if (x^n)d(y^n)+(z^n) occurs more than once.
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

				//used to find any modifiers added or subtracted to various dice roll
				basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)

				message += m.Author.Username + " LETS ROLL\n\n"
				message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

				//Actual for loop for the "dice rolls"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					rollValueInt := rand.Intn(diceToBeRolled) + 1
					rollValueStr := strconv.Itoa(rollValueInt)
					total += rollValueInt
					sumTotal += rollValueInt
					if rollValueInt == 1 {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[31m" + rollValueStr + "\u001B[0m]	"
					} else if rollValueInt == diceToBeRolled {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[32m" + rollValueStr + "\u001B[0m]	"
					} else {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[37m" + rollValueStr + "\u001B[0m]	"
					}
				}
				arithmeticResult := 0
				//Adds any and all modifiers, makes sure we don't add the next dice roll
				for j := 0; j < len(arithmetic)-1; j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
				if arithmeticResult != 0 {
					total += arithmeticResult
					message += "\n\n	Modifiers: [\u001B[37m" + strconv.Itoa(arithmeticResult) + "\u001B[0m]\n"
					sumTotal += arithmeticResult
				}
				message += "\n	Total Roll: " + "[\u001B[37m" + strconv.Itoa(total) + "\u001B[0m]\n\n"
			} else {
				//This code block is for ALL code after the original (x^n)d(y^n)+(z^n) amount of dice/things rolled.
				basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
				numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])

				if i != 0 {
					initialArray := basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)
					numbOfRoll, _ = strconv.Atoi(initialArray[len(initialArray)-1][0])
				}

				//this has to be initialized weirdly.  So.  Here's what we do
				diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
				arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
				//message now tells us what we're rolling
				message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

				//Actual for loop for the "dice rolls"
				for j := 0; j < numbOfRoll; j++ {
					//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
					rollValueInt := rand.Intn(diceToBeRolled) + 1
					rollValueStr := strconv.Itoa(rollValueInt)
					total += rollValueInt
					sumTotal += rollValueInt
					if rollValueInt == 1 {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[31m" + rollValueStr + "\u001B[0m]	"
					} else if rollValueInt == diceToBeRolled {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[32m" + rollValueStr + "\u001B[0m]	"
					} else {
						message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[37m" + rollValueStr + "\u001B[0m]	"
					}
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
					message += "\n\n	Modifiers: [\u001B[37m" + strconv.Itoa(arithmeticResult) + "\u001B[0m]\n"
				}
				message += "\n	Total Roll: " + "[\u001B[37m" + strconv.Itoa(total) + "\u001B[0m]\n\n"
				if i+1 == len(variablesArr) {
					message += "\n Sum Total of All Values: \u001B[37m" + strconv.Itoa(sumTotal) + "\u001B[0m"

				}
			}
		}
		message += "\n```"
		_, _ = s.ChannelMessageSend(m.ChannelID, message)
	}
	///test code
	if strings.Contains(m.Content, "ff") {
		//message := "```ansi\n"
		//message += "\u001B[4m bold \u001B[0m   "
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

		//ansi reading
		//```ansi
		//\u001b[{a};{b};{c}m
		//```
		//a is formatting
		//b is background
		//c is text color
		//30: Gray
		//31: Red
		//32: Green
		//33: Yellow
		//34: Blue
		//35: Pink
		//36: Cyan
		//37: White
		//message += "\u001B[0;33m\u001B[0m.\u001B[31m(\u001B[36m[\u001B[34m\\w\u001B[0m.@\u001B[36m]\u001B[34m+\u001B[31m)\u001B[0m!\u001B[34m?\u001B[36m[\u001B[0m+-\u001B[36m]\u001B[34m?\u001B[0m\n\n"
		//message += "\u001B[31m("
		//
		//message += "\n```"
		//
		//_, _ = s.ChannelMessageSend(m.ChannelID, message)
	}
	//Mtg Code
	if strings.Contains(m.Content, "!c") {
		cardName := m.Content[3:len(m.Content)]
		cachedCardRuling = getCard(cardName, m.ChannelID, s)
	}
	if strings.Contains(m.Content, "!rules") {
		if len(cachedCardRuling) > 1 {
			_, _ = s.ChannelMessageSend(m.ChannelID, cachedCardRuling)
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, "No card Selected :( ")

		}

	}

}

func reactionHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	//_, _ = s.ChannelMessageSend(m.ChannelID, m.MessageID)
	decode, length := utf8.DecodeRuneInString(m.Emoji.Name)

	if decode == 128218 && length == 4 && cachedCardRuling != "" && m.MessageReaction.UserID != Id {
		//fmt.Println(len(cachedCardRuling))
		//fmt.Println(cachedCardRuling)
		if len(cachedCardRuling) > 2000 {
			fmt.Println(len(cachedCardRuling))
			fmt.Println(float64(len(cachedCardRuling)) / 2000)
			fmt.Println(int(math.Ceil(float64(len(cachedCardRuling)) / 2000)))
			iterationsNeeded := int(math.Ceil(float64(len(cachedCardRuling)) / 2000))
			for i := 0; i < iterationsNeeded; i++ {
				fmt.Println(i)
				if i+1 != iterationsNeeded {
					fmt.Println(i * 2000)
					fmt.Println((i+1)*2000 - 1)
					if i == 0 {
						_, err := s.ChannelMessageSend(m.ChannelID, cachedCardRuling[i*2000:(i+1)*2000-1]+"```")
						if err != nil {
							fmt.Println(err)
						}
					} else {
						_, err := s.ChannelMessageSend(m.ChannelID, "```ansi\n"+cachedCardRuling[i*2000:(i+1)*2000-1]+"```")
						if err != nil {
							fmt.Println(err)
						}
					}
				} else {
					fmt.Println(i * 2000)
					fmt.Println((i+1)*2000 - 1)
					fmt.Println("final test ")
					var _, err = s.ChannelMessageSend(m.ChannelID, cachedCardRuling[i*2000:len(cachedCardRuling)]+"```")
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			//	if len(cachedCardRuling) < 3999 {
			//		cardRuling1 := cachedCardRuling[0:len(cachedCardRuling)/2] + "```"
			//		cardRuling2 := "```ansi\n" + cachedCardRuling[len(cardRuling1)+1:]
			//		_, err := s.ChannelMessageSend(m.ChannelID, cardRuling1)
			//		if err != nil {
			//			fmt.Println(err)
			//		}
			//		_, err = s.ChannelMessageSend(m.ChannelID, cardRuling2)
			//		if err != nil {
			//			fmt.Println(err)
			//		}
			//	}
			//} else {
			//	_, err := s.ChannelMessageSend(m.ChannelID, cachedCardRuling)
			//	if err != nil {
			//		fmt.Println(err)
			//	}
			//}

		}
	}
	//if m.MessageReaction.Emoji.ID == "\U0001F4C5" {
	//	fmt.Print("emote2")
	//}

}
