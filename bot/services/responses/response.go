package responses

type ImageURIs struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type LegalitiesObj struct {
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

type RelatedUrisObj struct {
	Gatherer                  string `json:"gatherer"`
	TcgplayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TcgplayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	Edhrec                    string `json:"edhrec"`
}

type PurchaseUrisObj struct {
	Tcgplayer   string `json:"tcgplayer"`
	Cardmarket  string `json:"cardmarket"`
	Cardhoarder string `json:"cardhoarder"`
}

type CardResponse struct {
	Object          string          `json:"object"`
	ID              string          `json:"id"`
	OracleID        string          `json:"oracle_id"`
	MultiverseIDs   []int           `json:"multiverse_ids"`
	TcgplayerID     int             `json:"tcgplayer_id"`
	Name            string          `json:"name"`
	Lang            string          `json:"lang"`
	ReleasedAt      string          `json:"released_at"`
	URI             string          `json:"uri"`
	ScryfallURI     string          `json:"scryfall_uri"`
	Layout          string          `json:"layout"`
	HighresImage    bool            `json:"highres_image"`
	ImageStatus     string          `json:"image_status"`
	ImageURIs       ImageURIs       `json:"image_uris"`
	ManaCost        string          `json:"mana_cost"`
	Cmc             float32         `json:"cmc"`
	TypeLine        string          `json:"type_line"`
	OracleText      string          `json:"oracle_text"`
	Colors          []string        `json:"colors"`
	ColorIdentity   []string        `json:"color_identity"`
	Keywords        []string        `json:"keywords"`
	Legalities      LegalitiesObj   `json:"legalities"`
	Games           []string        `json:"games"`
	Reserved        bool            `json:"reserved"`
	Foil            bool            `json:"foil"`
	Finishes        []string        `json:"finishes"`
	Oversized       bool            `json:"oversized"`
	Promo           bool            `json:"promo"`
	Reprint         bool            `json:"reprint"`
	Variation       bool            `json:"variation"`
	SetID           string          `json:"set_id"`
	Set             string          `json:"set"`
	SetName         string          `json:"set_name"`
	SetType         string          `json:"set_type"`
	SetURI          string          `json:"set_uri"`
	SetSearchURI    string          `json:"set_search_uri"`
	ScryfallSetURI  string          `json:"scryfall_set_uri"`
	CollectorNumber string          `json:"collector_number"`
	Digital         bool            `json:"digital"`
	Rarity          string          `json:"rarity"`
	CardBackID      string          `json:"card_back_id"`
	Artist          string          `json:"artist"`
	ArtistIds       []string        `json:"artist_ids"`
	IllustrationID  string          `json:"illustration_id"`
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
	RelatedURIs     RelatedUrisObj  `json:"related_uris"`
	PurchaseURIs    PurchaseUrisObj `json:"purchase_uris"`
	Details         string          `json:"details"`
	RulingsURI      string          `json:"rulings_uri"`
	PrintsSearchURI string          `json:"prints_search_uri"`
	CardFaces       []CardResponse  `json:"card_faces"`
}

type CardFace struct {
	Object          string    `json:"object"`
	Name            string    `json:"name"`
	URI             string    `json:"uri"`
	ImageURIs       ImageURIs `json:"image_uris"`
	Cmc             float32   `json:"mana_cost"`
	TypeLine        string    `json:"type_line"`
	OracleText      string    `json:"oracle_text"`
	Colors          []string  `json:"colors"`
	Keywords        []string  `json:"keywords"`
	Artist          string    `json:"artist"`
	ArtistIds       []string  `json:"artist_ids"`
	IllustrationID  string    `json:"illustration_id"`
	SecurityStamp   string    `json:"security_stamp"`
	FullArt         bool      `json:"full_art"`
	Textless        bool      `json:"textless"`
	Details         string    `json:"details"`
	RulingsURI      string    `json:"rulings_uri"`
	PrintsSearchURI string    `json:"prints_search_uri"`
	Power           float32   `json:"power"`
	Toughness       float32   `json:"toughness"`
	Watermark       string    `json:"watermark"`
}

type RulingData struct {
	Object        string `json:"object"`
	OracleID      string `json:"oracle_id"`
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
	ID          string         `json:"id"`
	Code        string         `json:"code"`
	MtgoCode    string         `json:"mtg_code"`
	ArenaCode   string         `json:"arena_code"`
	TcgplayerID string         `json:"tcgplayer_id"`
	Name        string         `json:"name"`
	URI         string         `json:"uri"`
	ScryfallURI string         `json:"scryfall_uri"`
	SearchURI   string         `json:"search_uri"`
	ReleasedAt  string         `json:"released_at"`
	SetType     string         `json:"set_type"`
	CardCount   string         `json:"card_count"`
	PrintedSize int            `json:"printed_size"`
	Digital     bool           `json:"digital"`
	NonFoilOnly bool           `json:"nonfoil_only"`
	FoilOnly    bool           `json:"foil_only"`
	IconSvgURI  string         `json:"icon_svg_uri"`
}

type QueryResponse struct {
	Object     string         `json:"object"`
	TotalCards int            `json:"total_cards"`
	HasMore    bool           `json:"has_more"`
	NextPage   string         `json:"next_page"`
	Data       []CardResponse `json:"data"`
	Details    string         `json:"details"`
}
