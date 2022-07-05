package bot

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"strconv"
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

var RulingUri string

func getCard(cardName string, channelId string, s *discordgo.Session) (string, string) {
	res, err := http.Get("https://api.scryfall.com/cards/named?fuzzy=" + cardName)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush tried. API said no  :(")
		return "Error", "Error"
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Card database said no to that :(")
		return "Error", "Error"
	}

	var data CardResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelId, data.Details)
		return "Card Error", "Card Error"
	}

	res, err = http.Get(data.ImageUris.Png)
	if err != nil {
		_, err = s.ChannelMessageSend(channelId, "Crush can't GET that card image :(")
		fmt.Println(err)
		return "Error", "Error"
	}
	if res.StatusCode == 200 {
		_, err = s.ChannelFileSend(channelId, data.Name+".png", res.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(data.RulingsUri)
	if len(data.RulingsUri) > 1 {
		RulingUri = data.RulingsUri
	}
	fmt.Println(data.SetUri)
	return "h", "ho"

}

func getRuling(channelId string, s *discordgo.Session) {
	fmt.Println(RulingUri)
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
	if data.Object == "error" {
		_, err = s.ChannelMessageSend(channelId, data.Details)
	}
	for i := 0; i < len(data.Data); i++ {
		_, err = s.ChannelMessageSend(channelId, "```ansi\n"+strconv.Itoa(i+1)+". "+data.Data[i].Comment+"\n```")

	}
}

//func getCard(cardName string, channelId string, s *discordgo.Session) (string, string) {
//	//type Card struct {
//	//	// The card name. For split, double-faced and flip cards, just the name of one side of the card. Basically each ‘sub-card’ has its own record.
//	//	Name string `json:"name"`
//	//	// Only used for split, flip and dual cards. Will contain all the names on this card, front or back.
//	//	Names []string `json:"names"`
//	//	// The mana cost of this card. Consists of one or more mana symbols. (use cmc and colors to query)
//	//	ManaCost string `json:"manaCost"`
//	//	// Converted mana cost. Always a number.
//	//	CMC float64 `json:"cmc"`
//	//	// The card colors. Usually this is derived from the casting cost, but some cards are special (like the back of dual sided cards and Ghostfire).
//	//	Colors []string `json:"colors"`
//	//	// The card colors by color code. [“Red”, “Blue”] becomes [“R”, “U”]
//	//	ColorIdentity []string `json:"colorIdentity"`
//	//	// The card type. This is the type you would see on the card if printed today. Note: The dash is a UTF8 'long dash’ as per the MTG rules
//	//	Type string `json:"type"`
//	//	// The types of the card. These appear to the left of the dash in a card type. Example values: Instant, Sorcery, Artifact, Creature, Enchantment, Land, Planeswalker
//	//	Types []string `json:"types"`
//	//	// The supertypes of the card. These appear to the far left of the card type. Example values: Basic, Legendary, Snow, World, Ongoing
//	//	Supertypes []string `json:"supertypes"`
//	//	// The subtypes of the card. These appear to the right of the dash in a card type. Usually each word is its own subtype. Example values: Trap, Arcane, Equipment, Aura, Human, Rat, Squirrel, etc.
//	//	Subtypes []string `json:"subtypes"`
//	//	// The rarity of the card. Examples: Common, Uncommon, Rare, Mythic Rare, Special, Basic Land
//	//	Rarity string `json:"rarity"`
//	//	// The set the card belongs to (set code).
//	//	Set mtg.SetCode `json:"set"`
//	//	// The set the card belongs to.
//	//	SetName string `json:"setName"`
//	//	// The oracle text of the card. May contain mana symbols and other symbols.
//	//	Text string `json:"text"`
//	//	// The flavor text of the card.
//	//	Flavor string `json:"flavor"`
//	//	// The artist of the card. This may not match what is on the card as MTGJSON corrects many card misprints.
//	//	Artist string `json:"artist"`
//	//	// The card number. This is printed at the bottom-center of the card in small text. This is a string, not an integer, because some cards have letters in their numbers.
//	//	Number string `json:"number"`
//	//	// The power of the card. This is only present for creatures. This is a string, not an integer, because some cards have powers like: “1+*”
//	//	Power string `json:"power"`
//	//	// The toughness of the card. This is only present for creatures. This is a string, not an integer, because some cards have toughness like: “1+*”
//	//	Toughness string `json:"toughness"`
//	//	// The loyalty of the card. This is only present for planeswalkers.
//	//	Loyalty string `json:"loyalty"`
//	//	// The card layout. Possible values: normal, split, flip, double-faced, token, plane, scheme, phenomenon, leveler, vanguard
//	//	Layout string `json:"layout"`
//	//	// The multiverseid of the card on Wizard’s Gatherer web page. Cards from sets that do not exist on Gatherer will NOT have a multiverseid. Sets not on Gatherer are: ATH, ITP, DKM, RQS, DPA and all sets with a 4 letter code that starts with a lowercase 'p’.
//	//	MultiverseId string `json:"multiverseid"`
//	//	// If a card has alternate art (for example, 4 different Forests, or the 2 Brothers Yamazaki) then each other variation’s multiverseid will be listed here, NOT including the current card’s multiverseid.
//	//	Variations []string `json:"variations"`
//	//	// The image url for a card. Only exists if the card has a multiverse id.
//	//	ImageUrl string `json:"imageUrl"`
//	//	// The watermark on the card. Note: Split cards don’t currently have this field set, despite having a watermark on each side of the split card.
//	//	Watermark string `json:"watermark"`
//	//	// If the border for this specific card is DIFFERENT than the border specified in the top level set JSON, then it will be specified here. (Example: Unglued has silver borders, except for the lands which are black bordered)
//	//	Border string `json:"border"`
//	//	// If this card was a timeshifted card in the set.
//	//	Timeshifted bool `json:"timeshifted"`
//	//	// Maximum hand size modifier. Only exists for Vanguard cards.
//	//	Hand int `json:"hand"`
//	//	// Starting life total modifier. Only exists for Vanguard cards.
//	//	Life int `json:"life"`
//	//	// Set to true if this card is reserved by Wizards Official Reprint Policy
//	//	Reserved bool `json:"reserved"`
//	//	// The date this card was released. This is only set for promo cards. The date may not be accurate to an exact day and month, thus only a partial date may be set (YYYY-MM-DD or YYYY-MM or YYYY). Some promo cards do not have a known release date.
//	//	ReleaseDate mtg.Date `json:"releaseDate"`
//	//	// Set to true if this card was only released as part of a core box set. These are technically part of the core sets and are tournament legal despite not being available in boosters.
//	//	Starter bool `json:"starter"`
//	//	// The rulings for the card.
//	//	Rulings []*mtg.Ruling `json:"rulings"`
//	//	// Foreign language names for the card, if this card in this set was printed in another language. An array of objects, each object having 'language’, 'name’ and 'multiverseid’ keys. Not available for all sets.
//	//	ForeignNames []mtg.ForeignCardName `json:"foreignNames"`
//	//	// The sets that this card was printed in, expressed as an array of set codes.
//	//	Printings []mtg.SetCode `json:"printings"`
//	//	// The original text on the card at the time it was printed. This field is not available for promo cards.
//	//	OriginalText string `json:"originalText"`
//	//	// The original type on the card at the time it was printed. This field is not available for promo cards.
//	//	OriginalType string `json:"originalType"`
//	//	// A unique id for this card. It is made up by doing an SHA1 hash of setCode + cardName + cardImageName
//	//	Id mtg.CardId `json:"id"`
//	//	// For promo cards, this is where this card was originally obtained. For box sets that are theme decks, this is which theme deck the card is from.
//	//	Source string `json:"source"`
//	//	// Which formats this card is legal, restricted or banned in. An array of objects, each object having 'format’ and 'legality’.
//	//	Legalities []mtg.Legality `json:"legalities"`
//	//}
//	type cardResponse struct {
//		Card  *Card   `json:"card"`
//		Cards []*Card `json:"cards"`
//	}
//	res, err := http.Get("https://api.scryfall.com/cards/named?fuzzy=aust+com")
//	if err != nil {
//		_, err = s.ChannelMessageSend(channelId, "Crush tried. API said no  :(")
//		return "Error", "Error"
//	}
//	fmt.Println(res.Body)
//
//	defer res.Body.Close()
//	if err != nil {
//		_, err = s.ChannelMessageSend(channelId, "Card database said no to that :(")
//		return "Error", "Error"
//	}
//
//	decoder := json.NewDecoder(res.Body)
//	var data Card
//	err = decoder.Decode(&data)
//	if err != nil {
//		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
//	}
//	fmt.Print(data)
//	//if len(data.Cards) == 0 {
//	//	_, err = s.ChannelMessageSend(channelId, "Crush can't find card :(")
//	//	return "Error", "Error"
//	//}
//	//res, err = http.Get(data.Cards[0].ImageUrl)
//	//if err != nil {
//	//	if len(data.Cards) > 1 {
//	//		res, err = http.Get(data.Cards[1].ImageUrl)
//	//		if err != nil {
//	//			_, err = s.ChannelMessageSend(channelId, "Crush can't GET that card image :(")
//	//			fmt.Println(err)
//	//			return "Error", "Error"
//	//		}
//	//	} else {
//	//		_, err = s.ChannelMessageSend(channelId, "Crush can't GET that card image :(")
//	//		return "Error", "Error"
//	//
//	//	}
//	//
//	//}
//	//if res.StatusCode == 200 {
//	//	_, err = s.ChannelFileSend(channelId, data.Cards[0].Name+".png", res.Body)
//	//	if err != nil {
//	//		fmt.Println(err)
//	//	}
//	//}
//	//var rulings = "```ansi\n"
//	//rulings += data.Cards[0].Name + "\n"
//	//for i := 0; i < len(data.Cards[0].Rulings); i++ {
//	//	rulings += "\n" + "[" + strconv.Itoa(i+1) + "] " + data.Cards[0].Rulings[i].Text + "\n"
//	//}
//	//rulings += "\n```"
//	//setList := ""
//	//for _, v := range data.Cards[0].Printings {
//	//	for _, i := range setCodes {
//	//		if string(v) == i.Code {
//	//			setList += "\n  " + i.Name
//	//		}
//	//	}
//	//}
//	//return rulings, setList
//
//	return "h", "ho"
//}
