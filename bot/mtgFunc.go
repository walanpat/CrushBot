package bot

import (
	"encoding/json"
	"fmt"
	"github.com/MagicTheGathering/mtg-sdk-go"
	"io/ioutil"
	"net/http"
)

func getCard(cardName string) string {
	cardGet, err := http.Get("https://api.magicthegathering.io/v1/cards?name=avacyn")
	if err != nil {
		return "Error has occurred"
	}
	body, err := ioutil.ReadAll(cardGet.Body)
	if err != nil {
		return "Error has occurred, Api gave back strange answer"
	}
	//need it to STOP after rulings array..... figure out the type conversion?
	//responseCard := string(body)
	var result mtg.Card
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(result)
	//cardList, _ := mtg.GetTypes()
	returnMessage := ""
	//for _, element := range cardList {
	//	returnMessage += element + " "
	//}
	return returnMessage
}
