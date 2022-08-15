package business

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/bot/business/query/builder"
	response "goland-discord-bot/bot/services/responses"

	"goland-discord-bot/bot/services"
	"math"
	"strconv"
	"strings"
)

var RulingUri string
var SetCodeUri string
var Price response.PriceObj

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

		//_, err = s.ChannelFileSend(channelId, data.Name+".png", res)
		EmbeddedCardSending(&data, channelId, s)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetRuling(channelId string, s *discordgo.Session) error {
	//Checking Input
	if RulingUri == "No Rulings Found" || RulingUri == "false" {
		_, _ = s.ChannelMessageSend(channelId, "No Rulings Found")
		return errors.New("no rulings found")
	}
	//Rules Service Request
	data, err := services.GetCardRulingService(RulingUri)
	if err != nil {
		_, _ = s.ChannelMessageSend(channelId, "Error in Retrieving Rules")
		return err
	}
	if data.Object == "error" || len(data.Data) == 0 {
		_, err = s.ChannelMessageSend(channelId, "```ansi\nNo Rulings Found```")
		return errors.New(data.Details)
	}
	//Print/Send out our rules
	for i := 0; i < len(data.Data); i++ {
		_, err = s.ChannelMessageSend(channelId, "```ansi\n"+strconv.Itoa(i+1)+". "+data.Data[i].Comment+"\n```")
	}
	return nil
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

	EmbeddedCardQuerySending(&data, channelId, s)
	//ExtendedMessageSending(&data, channelId, s)
	return
	//Handle our query
	//message := ""
	//for i := 0; i < len(data.Data); i++ {
	//	coloridentityprint := ""
	//
	//	for j := 0; j < len(data.Data[i].ColorIdentity); j++ {
	//		coloridentityprint += data.Data[i].ColorIdentity[j]
	//	}
	//	message += data.Data[i].Name + " " + strconv.Itoa(int(data.Data[i].Cmc)) + " " + coloridentityprint + " " + data.Data[i].ImageUris.Png + "\n"
	//}
	//_, _ = s.ChannelMessageSend(channelId, message)

}

func ExtendedMessageSending(data *response.QueryResponse, channelId string, s *discordgo.Session) {
	message := ""
	//timer1 := time.NewTimer(50 * time.Millisecond)
	for i := 0; i < len(data.Data); i++ {
		//fmt.Println(data.Data[i].Name)
		//<-timer1.C
		//
		//res, err := http.Get(data.Data[i].ImageUris.Png)
		//if err != nil {
		//	_, err = s.ChannelMessageSend(channelId, "Crushcan'tGETthatcardimage:(")
		//	fmt.Println(err)
		//	return
		//}
		//_, err = s.ChannelFileSend(channelId, data.Data[i].Name+".png", res.Body)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//timer1.Reset(100 * time.Millisecond)
		coloridentityprint := ""

		for j := 0; j < len(data.Data[i].ColorIdentity); j++ {
			coloridentityprint += data.Data[i].ColorIdentity[j]
		}
		message += data.Data[i].Name + " " + strconv.Itoa(int(data.Data[i].Cmc)) + " " + coloridentityprint + " " + "\n"
	}

	if len(message) > 2000 {
		iterationsNeeded := int(math.Ceil(float64(len(message)) / 2000))
		fmt.Println(len("```ansi\n" + message[0*2000:(0+1)*2000-11] + "```"))
		for i := 0; i < iterationsNeeded; i++ {
			if i+1 != iterationsNeeded {
				if i == 0 {
					_, err := s.ChannelMessageSend(channelId, "```ansi\n"+message[i*2000:(i+1)*2000-11]+"```")
					if err != nil {
						fmt.Println("Check1")
						fmt.Println(err)
					}
				} else {
					_, err := s.ChannelMessageSend(channelId, "```ansi\n"+message[i*2000-11:(i+1)*2000-11]+"```")
					if err != nil {
						fmt.Println("Check2")

						fmt.Println(err)
					}
				}
			} else {
				var _, err = s.ChannelMessageSend(channelId, "```ansi\n"+message[(i*2000)-11:]+"```")
				if err != nil {
					fmt.Println("Check3")

					fmt.Println(err)
				}
			}
		}
	} else {
		_, _ = s.ChannelMessageSend(channelId, "```ansi\n"+message+"```")

	}
	//_, _ = s.ChannelMessageSend(channelId, message+"```")

}

func EmbeddedCardQuerySending(data *response.QueryResponse, channelId string, s *discordgo.Session) {
	var temp []discordgo.MessageEmbed
	var x []*discordgo.MessageEmbed
	for i := 0; i < data.TotalCards; i++ {
		if len(data.Data[i].CardFaces) > 0 {

			image := discordgo.MessageEmbedImage{
				URL:      data.Data[i].CardFaces[0].ImageUris.Png,
				ProxyURL: "",
				Width:    10,
				Height:   20,
			}
			arrElement := discordgo.MessageEmbed{
				URL:         data.Data[i].ScryfallUri,
				Type:        "",
				Title:       data.Data[i].CardFaces[0].Name,
				Description: "",
				Timestamp:   "",
				Color:       0,
				Footer:      nil,
				Image:       &image,
				Thumbnail:   nil,
				Video:       nil,
				Provider:    nil,
				Author:      nil,
				Fields:      nil,
			}
			temp = append(temp, arrElement)
			x = append(x, &temp[len(temp)-1])

			image2 := discordgo.MessageEmbedImage{
				URL:      data.Data[i].CardFaces[1].ImageUris.Png,
				ProxyURL: "",
				Width:    10,
				Height:   20,
			}
			arrElement2 := discordgo.MessageEmbed{
				URL:         data.Data[i].ScryfallUri,
				Type:        "",
				Title:       data.Data[i].CardFaces[1].Name,
				Description: "",
				Timestamp:   "",
				Color:       0,
				Footer:      nil,
				Image:       &image2,
				Thumbnail:   nil,
				Video:       nil,
				Provider:    nil,
				Author:      nil,
				Fields:      nil,
			}
			temp = append(temp, arrElement2)
			x = append(x, &temp[len(temp)-1])
		} else {
			image := discordgo.MessageEmbedImage{
				URL:      data.Data[i].ImageUris.Png,
				ProxyURL: "",
				Width:    10,
				Height:   20,
			}
			arrElement := discordgo.MessageEmbed{
				URL:         data.Data[i].ScryfallUri,
				Type:        "",
				Title:       data.Data[i].Name,
				Description: "",
				Timestamp:   "",
				Color:       0,
				Footer:      nil,
				Image:       &image,
				Thumbnail:   nil,
				Video:       nil,
				Provider:    nil,
				Author:      nil,
				Fields:      nil,
			}
			temp = append(temp, arrElement)
			x = append(x, &temp[len(temp)-1])
		}
		if len(x) == 10 {
			_, err := s.ChannelMessageSendEmbeds(channelId, x)
			if err != nil {
				fmt.Printf("error sending embeds %q", err)
			}
			x = []*discordgo.MessageEmbed{}
		}
	}
	fmt.Println(x)
	_, err := s.ChannelMessageSendEmbeds(channelId, x)
	if err != nil {
		fmt.Printf("error sending embeds %q", err)
	}

}

func EmbeddedCardSending(data *response.CardResponse, channelId string, s *discordgo.Session) {
	var temp discordgo.MessageEmbed
	var x *discordgo.MessageEmbed

	image := discordgo.MessageEmbedImage{
		URL:      data.ImageUris.Png,
		ProxyURL: "",
		Width:    10,
		Height:   20,
	}
	arrElement := discordgo.MessageEmbed{
		URL:         data.ScryfallUri,
		Type:        "",
		Title:       data.Name,
		Description: "",
		Timestamp:   "",
		Color:       0,
		Footer:      nil,
		Image:       &image,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields:      nil,
	}

	temp = arrElement
	x = &temp

	_, err := s.ChannelMessageSendEmbed(channelId, x)
	if err != nil {
		fmt.Printf("error sending embed %q", err)
	}
}
