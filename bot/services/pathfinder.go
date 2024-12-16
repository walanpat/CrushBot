package services

import (
	"encoding/json"
	"fmt"
	response "goland-discord-bot/bot/services/responses"
	"io/ioutil"
	"net/http"
)

// WIP
func GetMonsterService(s string) (response.CardResponse, error) {
	res, err := http.Get("https://2e.aonprd.com/NPCs.aspx?ID=968&Elite=true&NoRedirect=1" + s)
	fmt.Println("https://2e.aonprd.com/NPCs.aspx?ID=968&Elite=true&NoRedirect=1=" + s)
	if err != nil {
		return response.CardResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response.CardResponse{}, err
	}
	var data response.CardResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	return data, nil
}
