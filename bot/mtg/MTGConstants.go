package mtg

import "regexp"

var typeRe = regexp.MustCompile(`type:([a-zA-Z ]+)?`)
var colorRe = regexp.MustCompile(`color:([rgbuw -]+)?`)
var cmcRe = regexp.MustCompile(`cmc:([>=<\d ]+)?`)
var powerRe = regexp.MustCompile(`power:([>=<\d ]+)?`)
var toughnessRe = regexp.MustCompile(`toughness:([>=<\d ]+)?`)
var textRe = regexp.MustCompile(`text:([a-zA-Z' ]+)?`)
var rarityRe = regexp.MustCompile(`rarity:([mruc ]+)?`)
var artRe = regexp.MustCompile(`art:([a-zA-Z ]+)?`)
var functionRe = regexp.MustCompile(`function:([a-zA-Z ]+)?`)
var isRe = regexp.MustCompile(`is:([a-zA-Z ]+)?`)

var QueryURL = "https://api.scryfall.com/cards/search?q="
