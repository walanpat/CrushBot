package business

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"

	"goland-discord-bot/bot/business/query/builder"
	"goland-discord-bot/bot/services"
	response "goland-discord-bot/bot/services/responses"
)

var RulingsURI string
var SetCodeURI string
var Price response.PriceObj

const (
	notFoundError = "Scryfall could not find what you searched for."
)

func GetCard(cardName string, channelID string, s *discordgo.Session) {
	//Get card Service Request
	data, err := services.GetCardService(cardName)
	if err != nil {
		_, err := s.ChannelMessageSend(channelID, "Error in Card Retrieval Service")
		if err != nil {
			fmt.Printf("\nError sending message: %v\n", err)
		}
	}

	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelID, data.Details)
		if err != nil {
			_, err := s.ChannelMessageSend(channelID, "Error in Card Retrieval Service")
			if err != nil {
				fmt.Printf("\nError sending message: %v\n", err)
			}
		}
	}

	//Get Card Image Service
	res, err := services.GetCardImageService(data.ImageURIs.Png)
	if err != nil {
		_, err = s.ChannelMessageSend(channelID, "Crush can't GET that card image :(")
		if err != nil {
			fmt.Printf("\nError sending message: %v\n", err)
		}
		fmt.Println(err)
		return
	}
	//Handling the rest of our attributes (Rulings, Prices, Sets)
	if res != nil {
		if len(data.RulingsURI) > 1 {
			RulingsURI = data.RulingsURI
		} else {
			RulingsURI = "No Rulings Found"
		}
		if data.Name == "Island" || data.Name == "Plains" || data.Name == "Mountain" || data.Name == "Forest" || data.Name == "Swamp" {
			SetCodeURI = "Basic Lands are Printed in Every Set"
		} else if len(data.SetURI) > 0 {
			SetCodeURI = data.PrintsSearchURI
		}
		if len(data.PurchaseURIs.Tcgplayer) > 0 {
			Price = data.Prices
		}

		//_, err = s.ChannelFileSend(channelID, data.Name+".png", res)
		EmbeddedCardSending(&data, channelID, s)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetRuling(channelID string, s *discordgo.Session) error {
	//Checking Input
	if RulingsURI == "No Rulings Found" || RulingsURI == "false" {
		_, err := s.ChannelMessageSend(channelID, "No Rulings Found")
		if err != nil {
			fmt.Printf("\nError sending message: %v\n", err)
		}
		return errors.New("no rulings found")
	}
	//Rules Service Request
	data, err := services.GetCardRulingService(RulingsURI)
	if err != nil {
		_, _ = s.ChannelMessageSend(channelID, "Error in Retrieving Rules")
		return err
	}
	if data.Object == "error" || len(data.Data) == 0 {
		_, err = s.ChannelMessageSend(channelID, "```ansi\nNo Rulings Found```")
		return errors.New(data.Details)
	}
	//Print/Send out our rules
	for i := 0; i < len(data.Data); i++ {
		_, err = s.ChannelMessageSend(channelID, "```ansi\n"+strconv.Itoa(i+1)+". "+data.Data[i].Comment+"\n```")
	}
	return nil
}

func GetSets(channelID string, s *discordgo.Session) {
	if SetCodeURI == "Basic Lands are Printed in Every Set" {
		_, _ = s.ChannelMessageSend(channelID, "```ansi\n"+SetCodeURI+"```")
		return
	}
	if SetCodeURI == "No Sets Found" {
		_, _ = s.ChannelMessageSend(channelID, SetCodeURI)
		return
	}

	//Get Sets Service
	data, err := services.GetSetsService(SetCodeURI)
	if err != nil {
		return
	}

	//Formatting/Send Rules
	x := "```ansi\nSets this card has been printed in: "

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
	_, err = s.ChannelMessageSend(channelID, x)

}

func GetPrice(channelID string, s *discordgo.Session) {
	_, _ = s.ChannelMessageSend(channelID, "```ansi\nScryfall Avg Price: $"+Price.Usd+"```")
}

func GetQuery(userQuery string, channelID string, s *discordgo.Session) {
	//Build our Query
	getUri, err := builder.MtgQueryBuilder(userQuery)
	if err != nil {
		_, _ = s.ChannelMessageSend(channelID, "```ansi\n"+err.Error()+"```")
		return
	}
	fmt.Println(err)
	fmt.Println(getUri)

	//Send our query
	data, err := services.GetQueryService(getUri)
	if err != nil {
		if err.Error() == "scryfall returned an error object, either nothing was found or there is a bad input" {
			_, _ = s.ChannelMessageSend(channelID, notFoundError)
			return
		}
		_, _ = s.ChannelMessageSend(channelID, err.Error())
		return
	}
	if data.TotalCards > 30 {
		arrElement := discordgo.MessageEmbed{
			URL:         "https://scryfall.com/search?q=" + getUri[40:],
			Type:        "",
			Title:       "That's a lot of cards\n Here's the scryfall link instead:",
			Description: "",
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
		temp := arrElement
		x := &temp

		_, err := s.ChannelMessageSendEmbed(channelID, x)
		if err != nil {
			fmt.Printf("error sending embed %q", err)
		}

		return
	} else {
		EmbeddedCardQuerySending(&data, channelID, s)
		return
	}

}

func ExtendedMessageSending(data *response.QueryResponse, channelID string, s *discordgo.Session) {
	message := ""
	//timer1 := time.NewTimer(50 * time.Millisecond)
	for i := 0; i < len(data.Data); i++ {
		//fmt.Println(data.Data[i].Name)
		//<-timer1.C
		//
		//res, err := http.Get(data.Data[i].ImageURIs.Png)
		//if err != nil {
		//	_, err = s.ChannelMessageSend(channelID, "Crushcan'tGETthatcardimage:(")
		//	fmt.Println(err)
		//	return
		//}
		//_, err = s.ChannelFileSend(channelID, data.Data[i].Name+".png", res.Body)
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
					_, err := s.ChannelMessageSend(channelID, "```ansi\n"+message[i*2000:(i+1)*2000-11]+"```")
					if err != nil {
						fmt.Println("Check1")
						fmt.Println(err)
					}
				} else {
					_, err := s.ChannelMessageSend(channelID, "```ansi\n"+message[i*2000-11:(i+1)*2000-11]+"```")
					if err != nil {
						fmt.Println("Check2")

						fmt.Println(err)
					}
				}
			} else {
				var _, err = s.ChannelMessageSend(channelID, "```ansi\n"+message[(i*2000)-11:]+"```")
				if err != nil {
					fmt.Println("Check3")

					fmt.Println(err)
				}
			}
		}
	} else {
		_, _ = s.ChannelMessageSend(channelID, "```ansi\n"+message+"```")

	}
	//_, _ = s.ChannelMessageSend(channelID, message+"```")

}

func EmbeddedCardQuerySending(data *response.QueryResponse, channelID string, s *discordgo.Session) {
	var temp []discordgo.MessageEmbed
	var x []*discordgo.MessageEmbed
	for i := 0; i < data.TotalCards; i++ {
		if len(data.Data[i].CardFaces) > 0 {

			image := discordgo.MessageEmbedImage{
				URL:      data.Data[i].CardFaces[0].ImageURIs.Png,
				ProxyURL: "",
				Width:    10,
				Height:   20,
			}
			arrElement := discordgo.MessageEmbed{
				URL:         data.Data[i].ScryfallURI,
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
				URL:      data.Data[i].CardFaces[1].ImageURIs.Png,
				ProxyURL: "",
				Width:    10,
				Height:   20,
			}
			arrElement2 := discordgo.MessageEmbed{
				URL:         data.Data[i].ScryfallURI,
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
				URL:      data.Data[i].ImageURIs.Png,
				ProxyURL: "",
				Width:    10,
				Height:   20,
			}
			arrElement := discordgo.MessageEmbed{
				URL:         data.Data[i].ScryfallURI,
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
			_, err := s.ChannelMessageSendEmbeds(channelID, x)
			if err != nil {
				fmt.Printf("error sending embeds %q", err)
			}
			x = []*discordgo.MessageEmbed{}
		}
	}
	fmt.Println(x)
	_, err := s.ChannelMessageSendEmbeds(channelID, x)
	if err != nil {
		fmt.Printf("error sending embeds %q", err)
	}

}

func EmbeddedCardSending(data *response.CardResponse, channelID string, s *discordgo.Session) {
	var temp discordgo.MessageEmbed
	var x *discordgo.MessageEmbed

	image := discordgo.MessageEmbedImage{
		URL:      data.ImageURIs.Png,
		ProxyURL: "",
		Width:    10,
		Height:   20,
	}
	arrElement := discordgo.MessageEmbed{
		URL:         data.ScryfallURI,
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

	_, err := s.ChannelMessageSendEmbed(channelID, x)
	if err != nil {
		fmt.Printf("error sending embed %q", err)
	}
}
