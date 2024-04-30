package commons

import (
	"log"
	"os"
)

var Help string
var RollDiceInfo string
var ProbInfo string
var QueryScryfallInfo string
var ScryfallExample string
var CardGetExample string

func InitializeResponses() error {

	h, err := os.ReadFile("./bot/commons/Help.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	Help = string(h)

	p, err := os.ReadFile("./bot/commons/ProbInfo.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	ProbInfo = string(p)

	r, err := os.ReadFile("./bot/commons/RollingInfo.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	RollDiceInfo = string(r)

	query, err := os.ReadFile("./bot/commons/ScryfallQueryInfo.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	QueryScryfallInfo = string(query)

	queryExample, err := os.ReadFile("./bot/commons/ScryfallQueryExample.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	ScryfallExample = string(queryExample)

	cardGetExample, err := os.ReadFile("./bot/commons/GetCardInfo.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	CardGetExample = string(cardGetExample)
	return nil
}
