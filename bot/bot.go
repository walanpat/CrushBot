package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/bot/business"
	"goland-discord-bot/bot/business/dicerolling"
	"goland-discord-bot/config"
	"strings"
	"time"
	"unicode/utf8"
)

var Id string

//Not sure if this variable/nomenclature will be needed later.  Add to clean up list.
//var goBot *discordgo.Session

var mtgSetMessageFlag = false
var mtgRulesMessageFlag = false
var mtgPriceMessageFlag = false

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
	//Handling error
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

// Definition of messageHandler function it takes two arguments first one is discordgo.Session which is s , second one is discordgo.MessageCreate which is m.
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Allows us a "time buffer" so that we don't react too fast to our own card search
	if m.Author.ID == Id && len(m.Reactions) == 0 && m.Content == "" {
		cachedCardTimer := time.NewTimer(5 * time.Millisecond)
		<-cachedCardTimer.C
		_ = s.MessageReactionAdd(m.ChannelID, m.Message.ID, "\U0001F4DA")
		cachedCardTimer.Reset(5 * time.Millisecond)
		<-cachedCardTimer.C
		_ = s.MessageReactionAdd(m.ChannelID, m.Message.ID, "\U0001F4C5")
		cachedCardTimer.Reset(5 * time.Millisecond)
		<-cachedCardTimer.C
		_ = s.MessageReactionAdd(m.ChannelID, m.Message.ID, "\U0001F4B5")

	}

	//Bot mustn't reply to its own messages , to confirm it we perform this check.
	if m.Author.ID == Id {
		return
	}

	//If we message ping to our bot in our discord it will return us pong .
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

	//Dice Rolling Code
	if strings.Contains(m.Content, "!roll") {
		message, err := dicerolling.DiceRollGeneric(m)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, message)
		}
	}

	//Rolls 6 different 5e stats, drops the lowest.
	if strings.Contains(m.Content, "!stats") {
		message, err := dicerolling.FiveEStats()
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, message)
	}

	//Mtg card request Code
	if strings.Contains(m.Content, "!c") {
		if m.Content[0:3] != "!c " {
			return
		}

		cardName := strings.ReplaceAll(m.Content[3:len(m.Content)], " ", "+")
		business.GetCard(cardName, m.ChannelID, s)

		mtgRulesMessageFlag = false
		mtgSetMessageFlag = false
		mtgPriceMessageFlag = false
	}

	if strings.Contains(m.Content, "!initiative ") {
		//!Initiative name, +4, name, +2, name, +4
		//Take map[string][int]
		if len(m.Content) > 12 {
			_, _ = s.ChannelMessageSend(m.ChannelID, dicerolling.InitiativeRoller(m.Content))
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, dicerolling.InitiativeRoller("Initiative roll needs 2 things.  a play"))

		}
	}
	if strings.Contains(m.Content, "[") || strings.Contains(m.Content, "]") {
		cardName := m.Content[strings.IndexRune(m.Content, '[')+1 : strings.IndexRune(m.Content, ']')]
		cardName = strings.ReplaceAll(cardName, " ", "+")

		business.GetCard(cardName, m.ChannelID, s)
		mtgRulesMessageFlag = false
		mtgSetMessageFlag = false
		mtgPriceMessageFlag = false
	}

	//Mtg card query request code
	if strings.Contains(m.Content, "!q") && m.Author.ID != Id {
		if len(m.Content) > 4 {
			business.GetQuery(m.Content, m.ChannelID, s)
		}
	}

	//Encode testing code
	if strings.Contains(m.Content, "!encode") {
		y := discordgo.MessageEmbed{
			URL:         "https://www.youtube.com/",
			Type:        "Youtube",
			Title:       "title",
			Description: "Youtube Embed description",
			Timestamp:   "",
			Color:       0,
			Footer:      nil,
			Image:       nil,
			Thumbnail:   nil,
			Video:       nil,
			Provider:    nil,
			Author:      nil,
			Fields:      nil,
		}
		z := discordgo.MessageEmbed{
			URL:         "https://www.google.com/",
			Type:        "Google",
			Title:       "title",
			Description: "Google Embed description",
			Timestamp:   "",
			Color:       0,
			Footer:      nil,
			Image:       nil,
			Thumbnail:   nil,
			Video:       nil,
			Provider:    nil,
			Author:      nil,
			Fields:      nil,
		}
		var temp [2]discordgo.MessageEmbed
		temp[0] = y
		temp[1] = z

		x := []*discordgo.MessageEmbed{&temp[0], &temp[1]}

		_, _ = s.ChannelMessageSendEmbeds(m.ChannelID, x)

	}

	//if m.Content[0:6] == "!play " {
	//
	//}

}

func reactionHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	//Deconstructs emojis into a 6 digit int
	decode, length := utf8.DecodeRuneInString(m.Emoji.Name)
	//Code for getting the ruling
	if decode == 128218 && length == 4 && m.MessageReaction.UserID != Id && mtgRulesMessageFlag == false {
		err := business.GetRuling(m.ChannelID, s)
		if err != nil {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Error getting rulings.")
			return
		} else {
			mtgRulesMessageFlag = true
		}
	}
	//Code for getting sets
	if decode == 128197 && length == 4 && m.MessageReaction.UserID != Id && mtgSetMessageFlag == false {
		business.GetSets(m.ChannelID, s)
		mtgSetMessageFlag = true
	}
	//Code for getting price
	if decode == 128181 && length == 4 && m.MessageReaction.UserID != Id && mtgPriceMessageFlag == false {
		business.GetPrice(m.ChannelID, s)
		mtgPriceMessageFlag = true
	}
}
