package builder

import (
	"regexp"
	"strings"
)

var TypeRe = regexp.MustCompile(`type:([a-zA-Z ]+)?`)
var ColorRe = regexp.MustCompile(`color:([rgbuw -]+)?`)
var CmcRe = regexp.MustCompile(`cmc:([>=<\d ]+)?`)
var PowerRe = regexp.MustCompile(`power:([>=<\d ]+)?`)
var ToughnessRe = regexp.MustCompile(`toughness:([>=<\d ]+)?`)
var TextRe = regexp.MustCompile(`text:([a-zA-Z' ]+)?`)
var RarityRe = regexp.MustCompile(`rarity:([mruc ]+)?`)
var ArtRe = regexp.MustCompile(`art:([a-zA-Z ]+)?`)
var FunctionRe = regexp.MustCompile(`function:([a-zA-Z ]+)?`)
var IsRe = regexp.MustCompile(`is:([a-zA-Z ]+)?`)

var QueryURL = "https://api.scryfall.com/cards/search?q="

type UrlBuilderObject struct {
	isValue        string
	functionValue  string
	artValue       string
	rarityValue    string
	textValue      string
	toughnessValue string
	powerValue     string
	colorValue     string
	cmcValue       string
	typeValue      string
	finalValue     string
}

func MtgQueryBuilder(query string) (string, error) {

	isArr := IsRe.FindStringSubmatch(query)
	functionArr := FunctionRe.FindStringSubmatch(query)
	artArr := ArtRe.FindStringSubmatch(query)
	rarityArr := RarityRe.FindStringSubmatch(query)
	textArr := TextRe.FindStringSubmatch(query)
	toughnessArr := ToughnessRe.FindStringSubmatch(query)
	powerArr := PowerRe.FindStringSubmatch(query)
	colorArr := ColorRe.FindStringSubmatch(query)
	cmcArr := CmcRe.FindStringSubmatch(query)
	typeArr := TypeRe.FindStringSubmatch(query)

	//If nothing found
	if len(typeArr) == 0 &&
		len(functionArr) == 0 &&
		len(isArr) == 0 &&
		len(artArr) == 0 &&
		len(rarityArr) == 0 &&
		len(textArr) == 0 &&
		len(toughnessArr) == 0 &&
		len(powerArr) == 0 &&
		len(colorArr) == 0 &&
		len(cmcArr) == 0 {
		return "", nil
	}
	QueryObject := UrlBuilderObject{
		isValue:        "",
		functionValue:  "",
		artValue:       "",
		rarityValue:    "",
		textValue:      "",
		toughnessValue: "",
		powerValue:     "",
		colorValue:     "",
		cmcValue:       "",
		typeValue:      "",
		finalValue:     QueryURL,
	}
	if len(typeArr) > 0 {
		QueryObject.typeValue += "t%3A" + typeArr[0][5:len(typeArr[0])]
		QueryObject.typeValue = strings.ReplaceAll(QueryObject.typeValue, " ", "+t%3A")
		QueryObject.finalValue += QueryObject.typeValue + "+"
	}
	if len(colorArr) > 0 {
		QueryObject.colorValue += "c%3A" + colorArr[0][6:len(colorArr[0])]
		QueryObject.colorValue = strings.ReplaceAll(QueryObject.colorValue, " -", "+-c%3A")
		QueryObject.colorValue = strings.ReplaceAll(QueryObject.colorValue, " ", "+c%3A")
		QueryObject.colorValue = strings.ReplaceAll(QueryObject.colorValue, "+c%3A+-", "")
		QueryObject.finalValue += QueryObject.colorValue
	}
	if len(functionArr) > 0 {
		QueryObject.functionValue += "function%3A" + functionArr[0][9:len(functionArr[0])]
		QueryObject.functionValue = strings.ReplaceAll(QueryObject.functionValue, " ", "+function%3A")
		QueryObject.finalValue += QueryObject.functionValue + "+"
	}
	if len(isArr) > 0 {
		if isArr[0][3] == ' ' {
			QueryObject.isValue += "is%3A" + isArr[0][4:len(isArr[0])] + "%27"

		} else {
			QueryObject.isValue += "is%3A" + isArr[0][3:len(isArr[0])] + "%27"
		}
		QueryObject.isValue = strings.ReplaceAll(QueryObject.isValue, " ", "+")
		QueryObject.finalValue += QueryObject.isValue
	}
	if len(textArr) > 0 {
		if textArr[0][5] == ' ' {
			QueryObject.textValue += "o%3A%27" + textArr[0][6:len(textArr[0])] + "%27"

		} else {
			QueryObject.textValue += "o%3A%27" + textArr[0][5:len(textArr[0])] + "%27"
		}
		QueryObject.textValue = strings.ReplaceAll(QueryObject.textValue, " ", "+")
		QueryObject.finalValue += QueryObject.textValue
	}
	if len(cmcArr) > 0 {
		//You need different checks for
		//>
		//>=
		//<
		//<=
		//just the number (=)
		//Inequalities (2<cost<5)
		if cmcArr[0][4] == ' ' {
			QueryObject.cmcValue += "is%3A" + cmcArr[0][5:len(cmcArr[0])] + "%27"

		} else {
			QueryObject.cmcValue += "is%3A" + cmcArr[0][4:len(cmcArr[0])] + "%27"
		}
		QueryObject.cmcValue = strings.ReplaceAll(QueryObject.cmcValue, " ", "+")
		QueryObject.finalValue += QueryObject.cmcValue
	}
	if len(toughnessArr) > 0 {

	}
	if len(powerArr) > 0 {

	}
	if len(rarityArr) > 0 {
		if rarityArr[0][7] == ' ' {
			QueryObject.rarityValue += "r%3A" + rarityArr[0][8:len(rarityArr[0])] + "%27"
		} else {
			QueryObject.rarityValue += "r%3A" + rarityArr[0][7:len(rarityArr[0])] + "%27"
		}
		QueryObject.rarityValue = strings.ReplaceAll(QueryObject.rarityValue, " ", "+")
		QueryObject.finalValue += QueryObject.rarityValue
	}
	if len(artArr) > 0 {
		if artArr[0][4] == ' ' {
			QueryObject.artValue += "art%3A" + artArr[0][5:len(artArr[0])] + "%27"
		} else {
			QueryObject.artValue += "art%3A" + artArr[0][4:len(artArr[0])] + "%27"
		}
		QueryObject.artValue = strings.ReplaceAll(QueryObject.artValue, " ", "+")
		QueryObject.finalValue += QueryObject.artValue
	}

	return QueryObject.finalValue, nil

}
