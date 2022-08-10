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
var loyaltyRe = regexp.MustCompile(`loyalty:([>=<\d ]+)?`)

var QueryURL = "https://api.scryfall.com/cards/search?q="

type UrlBuilderObject struct {
	isValue        string
	functionValue  string
	artValue       string
	rarityValue    string
	textValue      string
	toughnessValue string
	powerValue     string
	loyaltyValue   string
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
	loyaltyArr := loyaltyRe.FindStringSubmatch(query)

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
		len(cmcArr) == 0 &&
		len(loyaltyArr) == 0 {
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
		loyaltyValue:   "",
		typeValue:      "",
		finalValue:     QueryURL,
	}
	if len(typeArr) > 0 {
		QueryObject.typeValue += "t%3A" + strings.TrimSpace(typeArr[0][5:len(typeArr[0])])
		QueryObject.typeValue = strings.ReplaceAll(QueryObject.typeValue, " ", "+t%3A")
		QueryObject.finalValue += QueryObject.typeValue + "+"
	}
	if len(colorArr) > 0 {
		//Think
		//we need OR
		//AND (exclusive)
		//NOT
		//and can just be bguwr formatting
		// or bg b g

		//Analyze if there is space.  Then, if there's multiple things grouped together,
		//put in the AND symbol (at most color<%3DWU for WHITE and BLUE)
		//More regex lmao

		innerColorRe := regexp.MustCompile(`([-wubrg]*)*`)

		//Alan, you have to work with this.
		//This is the trick here.  Handle it within this innercolorArr, and you can do it.  just stick to what you got and handle it from there.
		innerColorArr := innerColorRe.FindAllStringSubmatch(strings.TrimSpace(colorArr[0][6:len(colorArr[0])]), -1)
		if len(innerColorArr) == 1 {
			QueryObject.finalValue += "c%3D" + innerColorArr[0][0] + "+"
		}
		if len(innerColorArr) == 2 {
			if !strings.Contains(innerColorArr[0][0], "-") || !strings.Contains(innerColorArr[0][0], innerColorArr[0][1]) {
				QueryObject.finalValue += "c%3C%3D" + innerColorArr[0][0] + innerColorArr[1][0] + "+"

			}
		}
		//checking to see if we have a OR statement (BU B U) stating I want
		//Blue and Black, or Black, or Blue cards

		lastIndex := len(innerColorArr) - 1
		if len(innerColorArr) >= 3 {
			for i, value := range innerColorArr {
				if value[0] != innerColorArr[0][0] {
					if strings.Contains(innerColorArr[0][0], value[0]) {
						if i == lastIndex {
							QueryObject.finalValue += "c%3C%3D" + innerColorArr[0][0] + "+"
						}
					} else {
						break
					}
				}
			}
		}
	}
	if len(functionArr) > 0 {
		QueryObject.functionValue += "function%3A" + strings.TrimSpace(functionArr[0][9:len(functionArr[0])])
		QueryObject.functionValue = strings.ReplaceAll(QueryObject.functionValue, " ", "+function%3A")
		QueryObject.finalValue += QueryObject.functionValue + "+"
	}
	if len(isArr) > 0 {

		QueryObject.isValue += "is%3A" + strings.TrimSpace(isArr[0][3:len(isArr[0])])
		QueryObject.isValue = strings.ReplaceAll(QueryObject.isValue, " ", "+")
		QueryObject.finalValue += QueryObject.isValue + "+"
	}
	if len(textArr) > 0 {

		QueryObject.textValue += "o%3A%27" + strings.TrimSpace(textArr[0][5:len(textArr[0])]+"%27")

		QueryObject.textValue = strings.ReplaceAll(QueryObject.textValue, " ", "+")
		QueryObject.textValue += "+"
		QueryObject.finalValue += QueryObject.textValue
	}
	if len(cmcArr) > 0 {
		QueryObject.cmcValue = InequalityReader(cmcArr, "cmc")
		QueryObject.finalValue += QueryObject.cmcValue
	}
	if len(toughnessArr) > 0 {
		QueryObject.toughnessValue = InequalityReader(toughnessArr, "tou")
		QueryObject.finalValue += QueryObject.toughnessValue + "+"
	}
	if len(loyaltyArr) > 0 {
		QueryObject.loyaltyValue = InequalityReader(loyaltyArr, "loy")
		QueryObject.finalValue += QueryObject.loyaltyValue + "+"
	}
	if len(powerArr) > 0 {
		QueryObject.powerValue = InequalityReader(powerArr, "pow")
		QueryObject.finalValue += QueryObject.powerValue + "+"
	}
	if len(rarityArr) > 0 {
		if strings.Contains(rarityArr[1], "c") {
			QueryObject.rarityValue += "r%3Acommon+"
		}
		if strings.Contains(rarityArr[1], "u") {
			QueryObject.rarityValue += "r%3Auncommon+"
		}
		if strings.Contains(rarityArr[1], "r") {
			QueryObject.rarityValue += "r%3Arare+"
		}
		if strings.Contains(rarityArr[1], "m") {
			QueryObject.rarityValue += "r%3Amythic+"
		}
		QueryObject.finalValue += QueryObject.rarityValue + "+"
	}
	if len(artArr) > 0 {
		QueryObject.artValue += "art%3A" + strings.TrimSpace(artArr[0][4:len(artArr[0])])
		QueryObject.artValue = strings.ReplaceAll(QueryObject.artValue, " ", "+")
		QueryObject.finalValue += QueryObject.artValue + "+"
	}
	fmt.Println("https://scryfall.com/search?q=" + QueryObject.finalValue[40:] + "\n")

	return QueryObject.finalValue, nil

}

//Rework Inequality Reader.
//things are backwards,
//implementation does not work.
//You're on the right track, but you need to fix it.
//regex needs to be fixed,
//the number, if it's 2 digits throws an error
func InequalityReader(array []string, typeOfInequality string) string {
	inequalityRe := regexp.MustCompile(`(\d?[><]?=?\d?)?m?(\d?[><]?=?\d?)?`)
	slicingString := ""
	//power = 4:
	if typeOfInequality == "pow" {
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
	if typeOfInequality == "loy" {
		slicingString = array[0][8:len(array[0])]
	}

	inequalityArr := inequalityRe.FindStringSubmatch(slicingString)

	slugQuery := typeOfInequality
	finalQuery := ""

	//First check what operators are in the query/ if it's a one-sided inequality
	//Second, act upon what operators are in the query
	if inequalityArr[0] == inequalityArr[1] {
		if strings.Contains(inequalityArr[0], "=") {
			if strings.Contains(inequalityArr[0], ">") {
				finalQuery = slugQuery + "%3E%3D" + inequalityArr[0][2:]
			} else if strings.Contains(inequalityArr[0], "<") {
				finalQuery = slugQuery + "%3C%3D" + inequalityArr[0][2:]
			} else {
				finalQuery = slugQuery + "%3D" + inequalityArr[0][1:]
			}
		}
		if strings.Contains(inequalityArr[0], ">") {
			finalQuery = slugQuery + "%3E" + inequalityArr[0][1:]
		} else if strings.Contains(inequalityArr[0], "<") {
			finalQuery = slugQuery + "%3C" + inequalityArr[0][1:]
		} else {
			finalQuery = slugQuery + "%3D" + inequalityArr[0]
		}
	} else {
		//This is for checking if there are 2 digits or 1
		//(on the left hand side of our inequality)
		digitRe := regexp.MustCompile(`(\d)+`)
		//Left inequality side number value
		leftSideNumberValue := digitRe.FindStringSubmatch(inequalityArr[1])[0]
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
	fmt.Println(finalQuery)
	return finalQuery
}
