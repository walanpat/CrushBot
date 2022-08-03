package mtg

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/bot/mtg/query/builder"
	"goland-discord-bot/bot/mtg/services"
	"strconv"
	"strings"
)

type RulingData struct {
	Object        string `json:"object"`
	OracleId      string `json:"oracle_id"`
	Source        string `json:"source"`
	PublishedDate string `json:"published_at"`
	Comment       string `json:"comment"`
}
type RulingResponse struct {
	Object  string       `json:"object"`
	HasMore bool         `json:"has_more"`
	Data    []RulingData `json:"data"`
	Source  string       `json:"source"`
	Details string       `json:"details"`
}

var RulingUri string
var SetCodeUri string
var Price services.PriceObj

func GetCard(cardName string, channelId string, s *discordgo.Session) {
	//Get card Service Request
	data, err := services.GetCardService(cardName)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Error in Card Retrieval Service")

	}
	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelId, data.Details)
		return
	}

	//Get Card Image Service
	res, err := services.GetCardImageService(data.ImageUris.Png)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush can't GET that card image :(")
		fmt.Println(err)
		return
	}
	//Handling the rest of our attributes (Rulings, Prices, Sets)
	if res != nil {
		if len(data.RulingsUri) > 1 {
			RulingUri = data.RulingsUri
		} else {
			RulingUri = "No Rulings Found"
		}
		if data.Name == "Island" || data.Name == "Plains" || data.Name == "Mountain" || data.Name == "Forest" || data.Name == "Swamp" {
			SetCodeUri = "Basic Lands are Printed in Every Set"
		} else if len(data.SetUri) > 0 {
			SetCodeUri = data.PrintsSearchUri
		}
		if len(data.PurchaseUris.Tcgplayer) > 0 {
			Price = data.Prices
		}
		_, err = s.ChannelFileSend(channelId, data.Name+".png", res)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetRuling(channelId string, s *discordgo.Session) {
	//Checking Input
	if RulingUri == "No Rulings Found" || RulingUri == "false" {
		_, _ = s.ChannelMessageSend(channelId, "No Rulings Found")
		return
	}
	//Rules Service Request
	data, err := services.GetCardRulingService(RulingUri)
	if err != nil {
		_, _ = s.ChannelMessageSend(channelId, "Error in Retrieving Rules")
		return
	}
	if data.Object == "error" || len(data.Data) == 0 {
		_, err = s.ChannelMessageSend(channelId, "```ansi\nNo Rulings Found```")
	}
	//Print/Send out our rules
	for i := 0; i < len(data.Data); i++ {
		_, err = s.ChannelMessageSend(channelId, "```ansi\n"+strconv.Itoa(i+1)+". "+data.Data[i].Comment+"\n```")
	}
}

func GetSets(channelId string, s *discordgo.Session) {
	if SetCodeUri == "Basic Lands are Printed in Every Set" {
		_, _ = s.ChannelMessageSend(channelId, "```ansi\n"+SetCodeUri+"```")
		return
	}
	if SetCodeUri == "No Sets Found" {
		_, _ = s.ChannelMessageSend(channelId, SetCodeUri)
		return
	}

	//Get Sets Service
	data, err := services.GetSetsService(SetCodeUri)
	if err != nil {
		return
	}

	//Formatting/Send Rules
	x := "```ansi\nSets this card has been printed in: "
	if data.HasMore {
	}
	for i := 0; i < len(data.Data); i++ {
		if strings.Contains(x, "\n   "+data.Data[i].SetName) {
			if strings.Contains(x, "\n   "+data.Data[i].SetName+" Promos") && data.Data[i].SetName+" Promos" == data.Data[i-1].SetName {
				x += "\n   " + data.Data[i].SetName
			} else {
				continue
			}
		} else {
			x += "\n   " + data.Data[i].SetName
		}
	}
	x += "\n```"
	_, err = s.ChannelMessageSend(channelId, x)

}

func GetPrice(channelId string, s *discordgo.Session) {
	_, _ = s.ChannelMessageSend(channelId, "```ansi\nScryfall Avg Price: $"+Price.Usd+"```")
}

func GetQuery(userQuery string, channelId string, s *discordgo.Session) {
	//Build our Query
	getUri, err := builder.MtgQueryBuilder(userQuery)
	if err != nil {
		return
	}
	fmt.Println(getUri)

	//Send our query
	data, err := services.GetQueryService(getUri)
	if err != nil {
		_, _ = s.ChannelMessageSend(channelId, "Error with query service return")
		return
	}

	//Handle our query
	message := ""
	for i := 0; i < len(data.Data); i++ {
		coloridentityprint := ""

		for j := 0; j < len(data.Data[i].ColorIdentity); j++ {
			coloridentityprint += data.Data[i].ColorIdentity[j]
		}
		message += data.Data[i].Name + " " + strconv.Itoa(int(data.Data[i].Cmc)) + " " + coloridentityprint + " " + data.Data[i].ImageUris.Png + "\n"
	}
	_, _ = s.ChannelMessageSend(channelId, message)

}
