package services

import (
	"encoding/json"
	"fmt"
	response "goland-discord-bot/bot/services/responses"
	"io"
	"io/ioutil"
	"net/http"
)

type MTGService interface {
	GetSetsService(s string) (response.SetListResponse, error)
	GetCardRulingService(s string) (response.RulingResponse, error)
	GetQueryService(s string) (response.QueryResponse, error)
	GetCardService(s string) (response.CardResponse, error)
}

func GetSetsService(s string) (response.SetListResponse, error) {
	res, err := http.Get(s)
	if err != nil {
		return response.SetListResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response.SetListResponse{}, err
	}
	var data response.SetListResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
		return response.SetListResponse{}, err
	}
	if data.Object == "error" {
		return response.SetListResponse{}, fmt.Errorf("get request resulted in returned error")
	}
	return data, nil
}

func GetCardRulingService(s string) (response.RulingResponse, error) {
	res, err := http.Get(s)
	if err != nil {
		return response.RulingResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response.RulingResponse{}, err
	}
	var data response.RulingResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	return data, nil
}

func GetCardService(s string) (response.CardResponse, error) {

	res, err := http.Get("https://api.scryfall.com/cards/named?fuzzy=" + s)
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

func GetCardImageService(s string) (response io.ReadCloser, err error) {
	res, err := http.Get(s)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func GetQueryService(s string) (response.QueryResponse, error) {
	res, err := http.Get(s)
	if err != nil {
		return response.QueryResponse{}, fmt.Errorf("get query service error")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response.QueryResponse{}, fmt.Errorf("error reading response body")
	}
	var data response.QueryResponse
	if err := json.Unmarshal(body, &data); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println(err)
	}
	if data.Object == "error" {
		return response.QueryResponse{}, fmt.Errorf("scryfall returned an error object")
		//_, _ = s.ChannelMessageSend(channelId, "```ansi\n ```")
	}
	return data, nil
}
