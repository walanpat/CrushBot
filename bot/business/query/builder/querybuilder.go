package builder

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	TypeRe      = regexp.MustCompile(`type:([a-zA-Z |]+)?`)
	ColorRe     = regexp.MustCompile(`color:([rgbuwc -]+)+(([|])*([rgbuwc -]*)*)*`)
	CommanderRe = regexp.MustCompile(`commander:([rgbuwc])*`)
	CmcRe       = regexp.MustCompile(`cmc:(\d?=?[><]?=?\d?)?m?(=?[><]?=?\d?)?`)
	PowerRe     = regexp.MustCompile(`power:(\d?=?[><]?=?\d?)?p?(=?[><]?=?\d?)?`)
	ToughnessRe = regexp.MustCompile(`toughness:(\d?=?[><]?=?\d?)?t?(=?[><]?=?\d?)?`)
	TextRe      = regexp.MustCompile(`text:([-\da-zA-Z' ]+)?([|]([-\da-zA-Z' ]+)?)*`)
	RarityRe    = regexp.MustCompile(`rarity:(([mruc ]+)?(([|])*([mruc ]+)?)*)*`)
	ArtRe       = regexp.MustCompile(`art:([a-zA-Z ]+)?`)
	FunctionRe  = regexp.MustCompile(`function:([a-zA-Z ]+)?`)
	IsRe        = regexp.MustCompile(`is:([a-zA-Z ]+)?`)
	loyaltyRe   = regexp.MustCompile(`loyalty:(\d?=?[><]?=?\d?)?l?(=?[><]?=?\d?)?`)
)

var QueryURL = "https://api.scryfall.com/cards/search?q="

type URLBuilderObject struct {
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
	// Start with REGEX
	if len(query) < 5 {
		err := errors.New("input is less than 7 characters in length")
		return "Not a long enough query", err
	}

	commaCheck := 0
	isArr := IsRe.FindStringSubmatch(query)
	if len(isArr) > 0 {
		commaCheck++
	}
	functionArr := FunctionRe.FindStringSubmatch(query)
	if len(functionArr) > 0 {
		commaCheck++
	}
	artArr := ArtRe.FindStringSubmatch(query)
	if len(artArr) > 0 {
		commaCheck++
	}
	rarityArr := RarityRe.FindStringSubmatch(query)
	if len(rarityArr) > 0 {
		commaCheck++
	}
	textArr := TextRe.FindStringSubmatch(query)
	if len(textArr) > 0 {
		commaCheck++
	}
	toughnessArr := ToughnessRe.FindStringSubmatch(query)
	if len(toughnessArr) > 0 {
		commaCheck++
	}
	powerArr := PowerRe.FindStringSubmatch(query)
	if len(powerArr) > 0 {
		commaCheck++
	}
	colorArr := ColorRe.FindStringSubmatch(query)
	if len(colorArr) > 0 {
		commaCheck++
	}
	commanderArr := CommanderRe.FindStringSubmatch(query)
	if len(commanderArr) > 0 {
		commaCheck++
	}
	cmcArr := CmcRe.FindStringSubmatch(query)
	if len(cmcArr) > 0 {
		commaCheck++
	}
	typeArr := TypeRe.FindStringSubmatch(query)
	if len(typeArr) > 0 {
		commaCheck++
	}
	loyaltyArr := loyaltyRe.FindStringSubmatch(query)
	if len(loyaltyArr) > 0 {
		commaCheck++
	}
	if commaCheck >= 2 && !strings.Contains(query, ",") {
		err := errors.New("error: more than 2 search modifiers entered, but no comma detected")
		return "", err
	}
	// If nothing found
	if len(typeArr) == 0 &&
		len(functionArr) == 0 &&
		len(isArr) == 0 &&
		len(artArr) == 0 &&
		len(rarityArr) == 0 &&
		len(textArr) == 0 &&
		len(toughnessArr) == 0 &&
		len(powerArr) == 0 &&
		len(colorArr) == 0 &&
		len(commanderArr) == 0 &&
		len(cmcArr) == 0 &&
		len(loyaltyArr) == 0 {
		return "", nil
	}
	// Initialize response object
	QueryObject := URLBuilderObject{
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
		QueryObject.typeValue += "t%3A%27" + strings.TrimSpace(typeArr[0][5:len(typeArr[0])])
		fmt.Println(QueryObject.typeValue)
		QueryObject.typeValue = orFormatting(QueryObject.typeValue)
		QueryObject.finalValue += QueryObject.typeValue + "%27+"
	}
	if len(commanderArr) > 0 {
		innerColorRe := regexp.MustCompile(`([-wubrgc]*)`)
		innerColorArr := innerColorRe.FindAllStringSubmatch(strings.TrimSpace(commanderArr[0][8:len(commanderArr[0])]), -1)
		QueryObject.finalValue += "commander%3A" + innerColorArr[1][0] + "+"
	} else if len(colorArr) > 0 {
		innerColorRe := regexp.MustCompile(`([-wubrgc]*)+([|])*([-wubrgc]*)*`)
		innerColorArr := innerColorRe.FindAllStringSubmatch(strings.TrimSpace(colorArr[0][6:len(colorArr[0])]), -1)

		// Handles individual color
		if len(innerColorArr) == 1 {
			QueryObject.finalValue += "c%3D" + innerColorArr[0][0] + "+"
		}

		// Handles 2 color inputs
		if len(innerColorArr) == 2 {
			if !strings.Contains(innerColorArr[0][0], "-") || !strings.Contains(innerColorArr[0][0], innerColorArr[0][1]) {
				QueryObject.finalValue += "c%3D" + orFormatting(innerColorArr[0][0]) + orFormatting(innerColorArr[1][1])
			}
		}

		// Handles 3 or more
		if len(innerColorArr) >= 3 {
			QueryObject.finalValue += "%28"
			for i, value := range innerColorArr {
				if value[0] == "|" {
					QueryObject.finalValue += "or+"
				} else if i != len(innerColorArr)-1 {
					QueryObject.finalValue += "c%3D" + value[0] + "+"
				} else {
					QueryObject.finalValue += "c%3D" + value[0] + "%29+"
				}
			}
		}
	}
	if len(functionArr) > 0 {
		QueryObject.functionValue += "function%3A" + strings.TrimSpace(functionArr[0][9:len(functionArr[0])])
		QueryObject.functionValue = strings.ReplaceAll(QueryObject.functionValue, " ", "+")
		// QueryObject.functionValue = strings.ReplaceAll(QueryObject.functionValue, " ", "+function%3A")
		QueryObject.finalValue += QueryObject.functionValue + "+"
	}
	if len(isArr) > 0 {
		QueryObject.isValue += "is%3A" + strings.TrimSpace(isArr[0][3:len(isArr[0])])
		QueryObject.isValue = strings.ReplaceAll(QueryObject.isValue, " ", "+")
		QueryObject.finalValue += QueryObject.isValue + "+"
	}
	if len(textArr) > 0 {
		QueryObject.textValue += "o%3A%27" + strings.TrimSpace(textArr[0][5:len(textArr[0])]+"%27")
		QueryObject.textValue = orFormatting(QueryObject.textValue)
		QueryObject.textValue += "+"
		QueryObject.finalValue += QueryObject.textValue
	}
	if len(cmcArr) > 0 {
		QueryObject.cmcValue = InequalityReader(cmcArr, "cmc")
		QueryObject.finalValue += QueryObject.cmcValue
	}
	if len(toughnessArr) > 0 {
		fmt.Println(toughnessArr)
		QueryObject.toughnessValue = InequalityReader(toughnessArr, "tou")
		QueryObject.finalValue += QueryObject.toughnessValue
	}
	if len(loyaltyArr) > 0 {
		QueryObject.loyaltyValue = InequalityReader(loyaltyArr, "loy")
		QueryObject.finalValue += QueryObject.loyaltyValue
	}
	if len(powerArr) > 0 {
		QueryObject.powerValue = InequalityReader(powerArr, "pow")
		QueryObject.finalValue += QueryObject.powerValue
	}
	if len(rarityArr) > 0 {
		if len(rarityArr[1]) > 1 {
			rarityArr[1] = strings.TrimSpace(rarityArr[1])
			fmt.Println(rarityArr[1])
			// fmt.Println(rarityArr)

			for i := 0; i <= len(rarityArr[1])-1; i++ {
				if i == 0 {
					QueryObject.rarityValue += "%28"
				}
				if i == len(rarityArr[1])-1 {
					fmt.Println(rarityArr)
					if rarityArr[1][i] == 'c' && !strings.Contains(QueryObject.rarityValue, "r%3Acommon") {
						QueryObject.rarityValue += "r%3Acommon"
					}
					if rarityArr[1][i] == 'u' && !strings.Contains(QueryObject.rarityValue, "uncommon") {
						QueryObject.rarityValue += "r%3Auncommon"
					}
					if rarityArr[1][i] == 'r' && rarityArr[1][i-1] != 'o' && !strings.Contains(QueryObject.rarityValue, "rare") {
						QueryObject.rarityValue += "r%3Arare"
					}
					if rarityArr[1][i] == 'm' && !strings.Contains(QueryObject.rarityValue, "mythic") {
						QueryObject.rarityValue += "r%3Amythic"
					}
				} else {
					if rarityArr[1][i] == 'c' {
						QueryObject.rarityValue += "r%3Acommon+OR+"
					}
					if rarityArr[1][i] == 'u' {
						QueryObject.rarityValue += "r%3Auncommon+OR+"
					}
					if rarityArr[1][i] == 'r' {
						QueryObject.rarityValue += "r%3Arare+OR+"
					}
					if rarityArr[1][i] == 'm' {
						QueryObject.rarityValue += "r%3Amythic+OR+"
					}
				}
				if i == len(rarityArr[1])-1 {
					QueryObject.rarityValue += "%29"
				}
			}
		} else {
			if strings.Contains(rarityArr[1], "c") {
				QueryObject.rarityValue += "r%3Acommon"
			}
			if strings.Contains(rarityArr[1], "u") {
				QueryObject.rarityValue += "r%3Auncommon"
			}
			if strings.Contains(rarityArr[1], "r") {
				QueryObject.rarityValue += "r%3Arare"
			}
			if strings.Contains(rarityArr[1], "m") {
				QueryObject.rarityValue += "r%3Amythic"
			}
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

func InequalityReader(array []string, typeOfInequality string) string {
	inequalityRe := regexp.MustCompile(`(\d{0,2}[><]?=?\d{0,2})?[mtpl]?(\d{0,2}[><]?=?\d{0,2})?`)
	slicingString := ""
	if typeOfInequality == "pow" {
		slicingString = array[0][6:len(array[0])]
	}
	if typeOfInequality == "tou" {
		slicingString = array[0][10:len(array[0])]
	}
	if typeOfInequality == "cmc" {
		slicingString = array[0][4:len(array[0])]
	}
	if typeOfInequality == "loy" {
		slicingString = array[0][8:len(array[0])]
	}

	inequalityArr := inequalityRe.FindStringSubmatch(slicingString)
	slugQuery := typeOfInequality
	finalQuery := ""
	// fmt.Println(inequalityArr[0])
	// fmt.Println(inequalityArr)
	// First check what operators are in the query/ if it's a one-sided inequality
	// Second, act upon what operators are in the query
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
		// Left inequality side number value
		leftSideNumberValue := digitRe.FindStringSubmatch(inequalityArr[1])[0]
		rightSideNumberValue := digitRe.FindStringSubmatch(inequalityArr[2])[0]
		// fmt.Println(leftSideNumberValue)
		// fmt.Println(rightSideNumberValue)
		// fmt.Println(inequalityArr)
		if strings.Contains(inequalityArr[1], "=") {
			if strings.Contains(inequalityArr[1], ">") {
				finalQuery += slugQuery + "%3C%3D" + leftSideNumberValue + "+"
			} else if strings.Contains(inequalityArr[1], "<") {
				finalQuery += slugQuery + "%3E%3D" + leftSideNumberValue + "+"
			}
		} else
		// if it's JUST >
		if strings.Contains(inequalityArr[1], ">") {
			finalQuery += slugQuery + "%3C" + leftSideNumberValue + "+"
		} else // if it's JUST <
		if strings.Contains(inequalityArr[1], "<") {
			finalQuery += slugQuery + "%3E" + leftSideNumberValue + "+"
		}
		// RIGHT HAND SIDE NUMBER VALUES
		if strings.Contains(inequalityArr[2], "=") {
			// if contains >=
			if strings.Contains(inequalityArr[2], ">") {
				finalQuery += slugQuery + "%3E%3D" + rightSideNumberValue
			} else // if contains <=
			if strings.Contains(inequalityArr[2], "<") {
				finalQuery += slugQuery + "%3C%3D" + rightSideNumberValue
			} else { // if it's just =
				finalQuery += slugQuery + "%3D" + rightSideNumberValue
			}
		} else if strings.Contains(inequalityArr[2], ">") {
			// if it's JUST >
			finalQuery += slugQuery + "%3E" + rightSideNumberValue
		} else if strings.Contains(inequalityArr[2], "<") {
			// if it's JUST <
			finalQuery += slugQuery + "%3C" + rightSideNumberValue
		}
	}
	finalQuery += "+"
	return finalQuery
}

// notes:
// %27 is the end of a ' used for creating complex statements

func orFormatting(str string) string {
	uriSlug := "%27+OR+"

	str = strings.ReplaceAll(str, " | ", uriSlug)
	str = strings.ReplaceAll(str, "| ", uriSlug)
	str = strings.ReplaceAll(str, " |", uriSlug)
	str = strings.ReplaceAll(str, "|", uriSlug)
	str = strings.ReplaceAll(str, " ", "+")
	return str
}
