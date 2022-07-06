package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type imageUris struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}
type legalitiesObj struct {
	Standard        string `json:"standard"`
	Future          string `json:"future"`
	Historic        string `json:"historic"`
	Gladiator       string `json:"gladiator"`
	Pioneer         string `json:"pioneer"`
	Explorer        string `json:"explorer"`
	Modern          string `json:"modern"`
	Legacy          string `json:"legacy"`
	Pauper          string `json:"pauper"`
	Vintage         string `json:"vintage"`
	Penny           string `json:"penny"`
	Commander       string `json:"commander"`
	Brawl           string `json:"brawl"`
	HistoricBrawl   string `json:"historicbrawl"`
	Alchemy         string `json:"alchemy"`
	PauperCommander string `json:"paupercommander"`
	Duel            string `json:"duel"`
	OldCchool       string `json:"oldschool"`
	PreModern       string `json:"premodern"`
}
type priceObj struct {
	Usd       string `json:"usd"`
	UsdFoil   string `json:"usd_foil"`
	UsdEtched string `json:"usd_etched"`
	Eur       string `json:"eur"`
	EurFoil   string `json:"eur_foil"`
	Tix       string `json:"tix"`
}
type relatedUrisObj struct {
	Gatherer                  string `json:"gatherer"`
	TcgplayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TcgplayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	Edhrec                    string `json:"edhrec"`
}
type purchaseUrisObj struct {
	Tcgplayer   string `json:"tcgplayer"`
	Cardmarket  string `json:"cardmarket"`
	Cardhoarder string `json:"cardhoarder"`
}
type CardResponse struct {
	Object          string          `json:"object"`
	Id              string          `json:"id"`
	OracleId        string          `json:"oracle_id"`
	MultiverseIds   []int           `json:"multiverse_ids"`
	TcgplayerId     int             `json:"tcgplayer_id"`
	Name            string          `json:"name"`
	Lang            string          `json:"lang"`
	ReleasedAt      string          `json:"released_at"`
	Uri             string          `json:"uri"`
	ScryfallUri     string          `json:"scryfall_uri"`
	Layout          string          `json:"layout"`
	HighresImage    bool            `json:"highres_image"`
	ImageStatus     string          `json:"image_status"`
	ImageUris       imageUris       `json:"image_uris"`
	ManaCost        string          `json:"mana_cost"`
	Cmc             float32         `json:"cmc"`
	TypeLine        string          `json:"type_line"`
	OracleText      string          `json:"oracle_text"`
	Colors          []string        `json:"colors"`
	ColorIdentity   []string        `json:"color_identity"`
	Keywords        []string        `json:"keywords"`
	Legalities      legalitiesObj   `json:"legalities"`
	Games           []string        `json:"games"`
	Reserved        bool            `json:"reserved"`
	Foil            bool            `json:"foil"`
	Finishes        []string        `json:"finishes"`
	Oversized       bool            `json:"oversized"`
	Promo           bool            `json:"promo"`
	Reprint         bool            `json:"reprint"`
	Variation       bool            `json:"variation"`
	SetId           string          `json:"set_id"`
	Set             string          `json:"set"`
	SetName         string          `json:"set_name"`
	SetType         string          `json:"set_type"`
	SetUri          string          `json:"set_uri"`
	SetSearchUri    string          `json:"set_search_uri"`
	ScryfallSetUri  string          `json:"scryfall_set_uri"`
	CollectorNumber string          `json:"collector_number"`
	Digital         bool            `json:"digital"`
	Rarity          string          `json:"rarity"`
	CardBackId      string          `json:"card_back_id"`
	Artist          string          `json:"artist"`
	ArtistIds       []string        `json:"artist_ids"`
	IllustrationId  string          `json:"illustration_id"`
	BorderColor     string          `json:"border_color"`
	Frame           string          `json:"frame"`
	SecurityStamp   string          `json:"security_stamp"`
	FullArt         bool            `json:"full_art"`
	Textless        bool            `json:"textless"`
	Booster         bool            `json:"booster"`
	StorySpotlight  bool            `json:"story_spotlight"`
	EdhrecRank      int             `json:"edhrec_rank"`
	PennyRank       int             `json:"penny_rank"`
	Prices          priceObj        `json:"prices"`
	RelatedUris     relatedUrisObj  `json:"related_uris"`
	PurchaseUris    purchaseUrisObj `json:"purchase_uris"`
	Details         string          `json:"details"`
	RulingsUri      string          `json:"rulings_uri"`
	PrintsSearchUri string          `json:"prints_search_uri"`
}
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
type SetListResponse struct {
	Object      string         `json:"object"`
	TotalCards  int            `json:"total_cards"`
	HasMore     bool           `json:"has_more"`
	Data        []CardResponse `json:"data"`
	Id          string         `json:"id"`
	Code        string         `json:"code"`
	MtgoCode    string         `json:"mtg_code"`
	ArenaCode   string         `json:"arena_code"`
	TcgplayerId string         `json:"tcgplayer_id"`
	Name        string         `json:"name"`
	Uri         string         `json:"uri"`
	ScryfallUri string         `json:"scryfall_uri"`
	SearchUri   string         `json:"search_uri"`
	ReleasedAt  string         `json:"released_at"`
	SetType     string         `json:"set_type"`
	CardCount   string         `json:"card_count"`
	PrintedSize int            `json:"printed_size"`
	Digital     bool           `json:"digital"`
	NonfoilOnly bool           `json:"nonfoil_only"`
	FoilOnly    bool           `json:"foil_only"`
	IconSvgUri  string         `json:"icon_svg_uri"`
}

var RulingUri string
var SetCodeUri string
var Price priceObj

func getCard(cardName string, channelId string, s *discordgo.Session) {
	res, err := http.Get("https://api.scryfall.com/cards/named?fuzzy=" + cardName)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush tried. API said no  :(")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Card database said no to that :(")
	}

	var data CardResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelId, data.Details)
	}

	res, err = http.Get(data.ImageUris.Png)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush can't GET that card image :(")
		fmt.Println(err)
		return
	}
	if res.StatusCode == 200 {
		if len(data.RulingsUri) > 1 {
			RulingUri = data.RulingsUri
		} else {
			RulingUri = "No Rulings Found"
		}
		if len(data.SetUri) > 0 {
			SetCodeUri = data.PrintsSearchUri
		}
		if len(data.PurchaseUris.Tcgplayer) > 0 {
			Price = data.Prices
		}
		_, err = s.ChannelFileSend(channelId, data.Name+".png", res.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
	//fmt.Println(data)
}

func getRuling(channelId string, s *discordgo.Session) {
	if RulingUri == "No Rulings Found" || RulingUri == "false" {
		_, _ = s.ChannelMessageSend(channelId, "No Rulings Found")
		return
	}
	res, err := http.Get(RulingUri)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush tried. API said no  :(")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Card database said no to that :(")
	}
	var data RulingResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	if data.Object == "error" || len(data.Data) == 0 {
		_, err = s.ChannelMessageSend(channelId, "```ansi\nNo Rulings Found```")
	}
	for i := 0; i < len(data.Data); i++ {
		_, err = s.ChannelMessageSend(channelId, "```ansi\n"+strconv.Itoa(i+1)+". "+data.Data[i].Comment+"\n```")
	}
}

func getSets(channelId string, s *discordgo.Session) {
	if SetCodeUri == "No Sets Found" {
		_, _ = s.ChannelMessageSend(channelId, SetCodeUri)
		return
	}
	res, err := http.Get(SetCodeUri)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush tried. API said no  :(")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Error Unmarshalling Json :(")
	}
	var data SetListResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		_, err = s.ChannelMessageSend(channelId, "Can't unmarshal JSON")
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelId, "Error in Json Object")
	}
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
	_, err = s.ChannelMessageSend(channelId, x)

}

func getPrice(channelId string, s *discordgo.Session) {
	_, _ = s.ChannelMessageSend(channelId, "```ansi\n Scryfall Avg Price: $"+Price.Usd+"```")
}
