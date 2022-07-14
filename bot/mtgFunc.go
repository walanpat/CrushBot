package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"math"
	"net/http"
	"regexp"
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
	NextPage    string         `json:"next_page"`
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
type QueryResponse struct {
	Object     string         `json:"object"`
	TotalCards int            `json:"total_cards"`
	HasMore    bool           `json:"has_more"`
	NextPage   string         `json:"next_page"`
	Data       []CardResponse `json:"data"`
	Details    string         `json:"details"`
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
		return
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
		if data.Name == "Island" || data.Name == "Plains" || data.Name == "Mountain" || data.Name == "Forest" || data.Name == "Swamp" {
			SetCodeUri = "Basic Lands are Printed in Every Set"
		} else if len(data.SetUri) > 0 {
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
	if SetCodeUri == "Basic Lands are Printed in Every Set" {
		_, _ = s.ChannelMessageSend(channelId, "```ansi\n"+SetCodeUri+"```")
		return
	}
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

func getPrice(channelId string, s *discordgo.Session) {
	_, _ = s.ChannelMessageSend(channelId, "```ansi\nScryfall Avg Price: $"+Price.Usd+"```")
}

//Search by

//Color
//cmc
//power
//toughness
//Types
//Subtypes (legendary?)
//keywords
//Art content
//oracletags

//Query Format:
//!q
//color:r||g||r+g,
//cmc:>=3,
//type:instant,
//subtype:goblin+soldier,
//power:4,
//toughness:4,
//text:Enters the battlefield Tapped
//rarity:r,
//art:squirrel,
//function:removal
//is:etb (this is specific to certain shortcults.

//typeRe := regexp.MustCompile(`type:[a-z ]*`)
//colorRe := regexp.MustCompile(`color:[a-z ]*`)
//cmcRe  := regexp.MustCompile(`cmc:[a-z ]*`)
//powerRe :=  regexp.MustCompile(`power:[a-z ]*`)
//toughnessRe  := regexp.MustCompile(`toughness:[a-z ]*`)
//textRe  := regexp.MustCompile(`text:[a-z ]*`)
//rarityRe  := regexp.MustCompile(`rarity:[a-z ]*`)
//artRe  := regexp.MustCompile(`art:[a-z ]*`)
//functionRe  := regexp.MustCompile(`function:[a-z ]*`)
//isRe  := regexp.MustCompile(`is:[a-z ]*`)
func getQuery(userQuery string, channelId string, s *discordgo.Session) {
	//notes
	//Each portion of the query is separated by a +,

	//CMC is 2 cmc%3D{2}
	//CMC is 5 and card color is blue c%3Au+cmc%3D5
	//Cards with one green and blue in their mana costs mana%3A%7BG%7D%7BU%7D
	//Cards with two generic and ATLEAST two white in their mana cost m%3A2WW
	//Cards that cost more than three generic, one white, and one blue mana  m>3WU
	//Cards with one phyrexian red mana in their cost m%3A%7BR%2FP%7D
	//Cards that produce blue and White Mana produces%3Dwu

	//Card color is red and green c%3A{rg}
	//Cards that are atleast white and blue, not red color>%3Duw+-c%3Ared
	//Cards that are instants that you can play with an Esper commander id<%3Desper+t%3Ainstant
	//Cards that are Colorless identity lands id%3Ac+t%3Aland

	//Card Type is merfolk and legendary  t%3Amerfolk + t%3Alegend
	//Card type is golbin and NOT creature t%3Agoblin+-t%3Acreature

	//card text is draw, type of card is creature o%3Adraw+t%3Acreature
	//Card text that is Enters the battlefield tapped o%3A"~+enters+the+battlefield+tapped"

	//Cards with 8 or more power pow>%3D8
	//White creatures that are power heavy pow>tou+c%3Aw+t%3Acreature
	//Planeswalkers with loyalty 3 to start t%3Aplaneswalker+loy%3D3

	//Card is a rare artifact r%3Acommon+t%3Aartifact
	//Cards at rare or higher (rares and mythics) r>%3Dr
	//Cards printed as commons for the first time in Iconic Masters rarity%3Acommon+e%3Aima+new%3Ararity
	//Non-rare printings of cards that have been printed at rare in%3Arare+-rarity%3Arare

	//Cards from War of the Spark    https://scryfall.com/sets/war?as=grid&order=set (note this will need to be edited
	//Cards available inside War of the spark booster boxes e%3Awar+is%3Abooster
	//Cards in Zendikar Block (but using the world wake code) https://scryfall.com/search?q=b%3Awwk
	//Cards that were in both Alpha and Magic 2015 https://scryfall.com/search?q=in%3Alea+in%3Am15
	//Cards that are legendary and have NEVER been printedin a booster set https://scryfall.com/search?q=t%3Alegendary+-in%3Abooster
	//Prerelease promos with a date stamp cards is%3Adatestamped+is%3Aprerelease

	//Cards that have art that contain a squirrel https://scryfall.com/search?q=art%3Asquirrel
	//Cards that cause removal (oracle tag) https://scryfall.com/search?q=function%3Aremoval

	//https://api.scryfall.com/cards/search?q=c%3Awhite+cmc%3D1
	//res, _ := http.Get("https://api.scryfall.com/cards/search?q=" + userQuery)

	typeRe := regexp.MustCompile(`type:[a-z ]*`)
	variablesArr := typeRe.FindStringSubmatch(userQuery)
	cardTypeUri := "t%3A" + variablesArr[0][5:len(variablesArr[0])]
	cardTypeUri = strings.ReplaceAll(cardTypeUri, " ", "+t%3A")
	fmt.Println(cardTypeUri)

	res, err := http.Get("https://api.scryfall.com/cards/search?q=" + cardTypeUri)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush tried. API said no  :(")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Card database said no to that :(")
	}

	var data QueryResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelId, data.Details)
		return
		//_, _ = s.ChannelMessageSend(channelId, "```ansi\n ```")
	}

	message := ""
	for i := 0; i < len(data.Data); i++ {
		message += data.Data[i].Name + " "
	}
	//fmt.Println(data)
	fmt.Println(len(message))
	if len(message) > 2000 {
		iterationsNeeded := int(math.Ceil(float64(len(message)) / 2000))
		fmt.Println(len(message[0*2000 : (0+1)*2000]))
		for i := 0; i < iterationsNeeded; i++ {
			if i+1 != iterationsNeeded {
				if i == 0 {
					_, err := s.ChannelMessageSend(channelId, "```ansi\n"+message[i*2000:(i+1)*2000]+"```")
					if err != nil {
						fmt.Println("Check1")
						fmt.Println(err)
					}
				} else {
					_, err := s.ChannelMessageSend(channelId, "```ansi\n"+message[i*2000:(i+1)*2000]+"```")
					if err != nil {
						fmt.Println("Check2")

						fmt.Println(err)
					}
				}
			} else {
				var _, err = s.ChannelMessageSend(channelId, "```ansi\n"+message[(i*2000):])
				if err != nil {
					fmt.Println("Check3")

					fmt.Println(err)
				}
			}
		}
	} else {
		_, _ = s.ChannelMessageSend(channelId, message)
	}
	//_, _ = s.ChannelMessageSend(channelId, message+"```")

}
