package responses

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

type PriceObj struct {
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
	Prices          PriceObj        `json:"prices"`
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
