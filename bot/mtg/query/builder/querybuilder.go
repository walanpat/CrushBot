package builder

import (
	"fmt"
	"regexp"
	"strings"
)

var TypeRe = regexp.MustCompile(`type:([a-zA-Z ]+)?`)
var ColorRe = regexp.MustCompile(`color:([rgbuw -]+)?`)
var CmcRe = regexp.MustCompile(`cmc:(\d?=?[><]?=?\d?)?m?(=?[><]?=?\d?)?`)
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
	//Start with REGEX
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
	//Initialize response object
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
		QueryObject.finalValue += QueryObject.colorValue + "+"
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
		QueryObject.cmcValue = InequalityReader(cmcArr, "cmc")

		QueryObject.finalValue += QueryObject.cmcValue
	}
	if len(toughnessArr) > 0 {
		QueryObject.toughnessValue = InequalityReader(toughnessArr, "tou")

		QueryObject.finalValue += QueryObject.toughnessValue
	}
	if len(powerArr) > 0 {
		QueryObject.powerValue = InequalityReader(powerArr, "pow")

		QueryObject.finalValue += QueryObject.powerValue
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

func InequalityReader(array []string, typeOfInequality string) string {
	inequalityRe := regexp.MustCompile(`(\d?[><]?=?\d?)?m?(\d?[><]?=?\d?)?`)
	slicingString := ""
	//power = 4:
	if typeOfInequality == "pow" {
		fmt.Println(array[0][6:len(array[0])])
		slicingString = array[0][6:len(array[0])]
	}
	//toughness = 8:
	if typeOfInequality == "tou" {
		slicingString = array[0][10:len(array[0])]
	}
	//cmc 2:
	if typeOfInequality == "cmc" {
		slicingString = array[0][4:len(array[0])]
	}

	inequalityArr := inequalityRe.FindStringSubmatch(slicingString)
	slugQuery := typeOfInequality
	finalQuery := ""

	//First check what operators are in the query/ if it's a one-sided inequality
	//Second, act upon what operators are in the query
	if inequalityArr[0] == inequalityArr[1] {
		if strings.Contains(inequalityArr[0], "=") {
			if strings.Contains(inequalityArr[0], ">") {
				finalQuery = slugQuery + "%3E%3D" + string(inequalityArr[0][2:])
			} else if strings.Contains(inequalityArr[0], "<") {
				finalQuery = slugQuery + "%3C%3D" + string(inequalityArr[0][2:])
			} else {
				finalQuery = slugQuery + "%3D" + inequalityArr[0][1:]
			}
		}
		if strings.Contains(inequalityArr[0], ">") {
			finalQuery = slugQuery + "%3E" + string(inequalityArr[0][1:])
		} else if strings.Contains(inequalityArr[0], "<") {
			finalQuery = slugQuery + "%3C" + string(inequalityArr[0][1:])
		} else {
			finalQuery = slugQuery + "%3D" + string(inequalityArr[0])
		}
	} else {
		//This is for checking if there are 2 digits or 1
		//(on the left hand side of our inequality)
		digitRe := regexp.MustCompile(`(\d)+`)
		//Left inequality side number value
		leftSideNumberValue := digitRe.FindStringSubmatch(inequalityArr[1])[0]
		fmt.Println(inequalityArr[2])
		rightSideNumberValue := digitRe.FindStringSubmatch(inequalityArr[2])[0]

		if strings.Contains(inequalityArr[1], "=") {
			if strings.Contains(inequalityArr[1], ">") {
				finalQuery += slugQuery + "%3C%3D" + leftSideNumberValue + "+"
			} else if strings.Contains(inequalityArr[1], "<") {
				finalQuery += slugQuery + "%3E%3D" + leftSideNumberValue + "+"

			}
		} else
		//if it's JUST >
		if strings.Contains(inequalityArr[1], ">") {
			finalQuery += slugQuery + "%3C" + leftSideNumberValue + "+"

		} else //if it's JUST <
		if strings.Contains(inequalityArr[1], "<") {
			finalQuery += slugQuery + "%3E" + leftSideNumberValue + "+"

		}
		//RIGHT HAND SIDE NUMBER VALUES
		if strings.Contains(inequalityArr[2], "=") {
			// if contains >=
			if strings.Contains(inequalityArr[2], ">") {
				finalQuery += slugQuery + "%3E%3D" + rightSideNumberValue + "+"
			} else // if contains <=
			if strings.Contains(inequalityArr[2], "<") {
				finalQuery += slugQuery + "%3C%3D" + rightSideNumberValue + "+"
			} else { //if it's just =
				finalQuery += slugQuery + "%3D" + rightSideNumberValue + "+"
			}
		} else if strings.Contains(inequalityArr[2], ">") {
			//if it's JUST >
			finalQuery += slugQuery + "%3E" + rightSideNumberValue + "+"
		} else if strings.Contains(inequalityArr[2], "<") {
			//if it's JUST <
			finalQuery += slugQuery + "%3C" + rightSideNumberValue + "+"
		}
	}

	return finalQuery
}
