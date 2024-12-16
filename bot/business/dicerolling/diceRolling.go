package dicerolling

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var re = regexp.MustCompile(`([+-]*\d+)*,(\d\d)*`)

func FiveEStats() (string, error) {
	message := "```ansi\n                             Stats:\n"
	statHolder := [5]int{}
	messageHolder := [6]string{}
	totalHolder := [6]int{}
	extraMessage := ""

	for i := 0; i <= 5; i++ {
		extraMessage = ""
		for j := 0; j < 4; j++ {
			cachedCardTimer := time.NewTimer(1 * time.Millisecond)
			<-cachedCardTimer.C
			s1 := rand.NewSource(time.Now().UnixNano())
			r1 := rand.New(s1)
			var bytes int
			for z := 0; z < 10; z++ {
				bytes = r1.Intn(6) + 1
			}
			statHolder[j] = bytes
			extraMessage += strconv.Itoa(statHolder[j])
		}

		statHolder[4] = 7
		sort.Ints(statHolder[:])
		statHolder[4] = 0

		messageHolder[i] += "Rolls: "

		for j := 0; j < 4; j++ {
			if statHolder[j] == 6 {
				messageHolder[i] += "[\u001B[32m" + strconv.Itoa(statHolder[j]) + "\u001B[0m] "
			} else if statHolder[j] == 1 {
				messageHolder[i] += "[\u001B[31m" + strconv.Itoa(statHolder[j]) + "\u001B[0m] "
			} else {
				messageHolder[i] += "[" + strconv.Itoa(statHolder[j]) + "] "
			}
			if j != 0 {
				statHolder[4] += statHolder[j]
			}
			if j == 3 {
				messageHolder[i] += "    Total: " + strconv.Itoa(statHolder[4]) + "\n"
				totalHolder[i] += statHolder[4]
			}
		}
		statHolder[4] = 0
	}
	sort.Ints(totalHolder[:])

	for i := 5; i >= 0; i-- {
		for j := 0; j < 6; j++ {
			x, err := strconv.Atoi(strings.TrimSpace(messageHolder[j][len(messageHolder[j])-3 : len(messageHolder[j])-1]))
			if err != nil {
				oof := fmt.Errorf("error detected on str conversion: %e", err)
				fmt.Println(oof)
				return oof.Error(), oof
			}
			if x == 0 {
				x, err = strconv.Atoi(messageHolder[j][len(messageHolder[j])-2 : len(messageHolder[j])-1])
			}
			if totalHolder[i] == x {
				message += messageHolder[j]
				messageHolder[j] = "0000"
			}
		}
	}

	message += "```"
	return message, nil
}

func DiceRollGeneric(m *discordgo.MessageCreate) (string, error) {
	// Initializing our "base" query expression
	re := regexp.MustCompile(`([+\-]?\d+)*d(\d+)([+\-]?\d*[^\dd][^d]+)*`)
	variablesArr := re.FindAllStringSubmatch(m.Content, -1)
	message := "```ansi\n \u001B[0m"
	sumTotal := 0
	singleDiceType := len(variablesArr) == 1
	multipleDiceType := len(variablesArr) > 1

	for i := 0; i < len(variablesArr); i++ {
		total := 0
		numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
		diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
		if diceToBeRolled == 0 {
			return "```ansi\nDice cannot be 0 sided.```", nil
		}

		message += fmt.Sprintf("%s LETS ROLL\n\n", m.Author.Username)
		message += fmt.Sprintf("\u001B[3m%dd%d\u001B[0m\t\n", numbOfRoll, diceToBeRolled)
		// This IF statement only occurs if there is a single "roll" (x^n)d(y^n)+(z^n) occurring
		if singleDiceType {
			// Basic math, modifier regexp
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)

			// Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				rollValueInt := rand.Intn(diceToBeRolled) + 1
				total += rollValueInt
				if rollValueInt == 1 {
					message += fmt.Sprintf("\tRoll %d: [\u001B[31m%d\u001B[0m]", j+1, rollValueInt)
				} else if rollValueInt == diceToBeRolled {
					message += fmt.Sprintf("\tRoll %d: [\u001B[32m%d\u001B[0m]", j+1, rollValueInt)
				} else {
					message += fmt.Sprintf("\tRoll %d: [\u001B[37m%d\u001B[0m]", j+1, rollValueInt)
				}
				if ((j + 1) % 3) == 0 {
					message += "\n"
				}
			}
			arithmeticResult := 0
			// Adds any and all modifiers
			for j := 0; j < len(arithmetic); j++ {
				x, _ := strconv.Atoi(arithmetic[j][0])
				arithmeticResult += x
			}
			if arithmeticResult != 0 {
				total += arithmeticResult
				message += fmt.Sprintf("\n\n\tModifiers: [\u001B[37m%d\u001B[0m]\n", arithmeticResult)
			}
			message += fmt.Sprintf("\n\n	Total Roll: [\u001B[37m%d\u001B[0m]\n\n", total)
		} else
		// this code block occurs only if (x^n)d(y^n)+(z^n) occurs more than once.
		if i == 0 && multipleDiceType {
			// used to find any modifiers added or subtracted to various dice roll
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)

			// Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				rollValueInt := rand.Intn(diceToBeRolled) + 1
				total += rollValueInt
				sumTotal += rollValueInt
				if rollValueInt == 1 {
					message += fmt.Sprintf("\tRoll %d: [\u001B[31m%d\u001B[0m]", j+1, rollValueInt)
				} else if rollValueInt == diceToBeRolled {
					message += fmt.Sprintf("\tRoll %d: [\u001B[32m%d\u001B[0m]", j+1, rollValueInt)
				} else {
					message += fmt.Sprintf("\tRoll %d: [\u001B[37m%d\u001B[0m]", j+1, rollValueInt)
				}
			}
			arithmeticResult := 0
			// Adds any and all modifiers, makes sure we don't add the next dice roll
			for j := 0; j < len(arithmetic)-1; j++ {
				x, _ := strconv.Atoi(arithmetic[j][0])
				arithmeticResult += x
			}
			if arithmeticResult != 0 {
				total += arithmeticResult
				message += fmt.Sprintf("\n\n\tModifiers: [\u001B[37m%d\u001B[0m]\n", arithmeticResult)
				sumTotal += arithmeticResult
			}
			message += fmt.Sprintf("\n\n	Total Roll: [\u001B[37m%d\u001B[0m]\n\n", total)
		} else {
			// This code block is for ALL code after the original (x^n)d(y^n)+(z^n) amount of dice/things rolled.
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)

			if i != 0 {
				initialArray := basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)
				numbOfRoll, _ = strconv.Atoi(initialArray[len(initialArray)-1][0])
			}

			// this has to be initialized weirdly.  So.  Here's what we do
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
			// message now tells us what we're rolling
			message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

			// Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				rollValueInt := rand.Intn(diceToBeRolled) + 1
				total += rollValueInt
				sumTotal += rollValueInt
				if rollValueInt == 1 {
					message += fmt.Sprintf("\tRoll %d: [\u001B[31m%d\u001B[0m]", j+1, rollValueInt)
				} else if rollValueInt == diceToBeRolled {
					message += fmt.Sprintf("\tRoll %d: [\u001B[32m%d\u001B[0m]", j+1, rollValueInt)
				} else {
					message += fmt.Sprintf("\tRoll %d: [\u001B[37m%d\u001B[0m]", j+1, rollValueInt)
				}
			}
			arithmeticResult := 0

			// If, we are NOT at the last modifier/read value
			if i+1 != len(variablesArr) {
				for j := 0; j < len(arithmetic)-1; j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
			} else {
				// If it is the last value, then add it if there's any modifiers/arithmetic values.
				for j := 0; j < len(arithmetic); j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
			}
			if arithmeticResult != 0 {
				total += arithmeticResult
				sumTotal += arithmeticResult
				message += fmt.Sprintf("\n\n\tModifiers: [\u001B[37m%d\u001B[0m]\n", arithmeticResult)
			}
			message += fmt.Sprintf("\n\tTotal Roll: [\u001B[37m%d\u001B[0m]\n\n", total)
			if i+1 == len(variablesArr) {
				message += fmt.Sprintf("\n Sum Total of All Values: \u001B[37m%d\u001B[0m", sumTotal)
			}
		}
	}
	message += "\n```"
	return message, nil
}

func SaveProbabilityCalculator(m *discordgo.MessageCreate) (string, error) {
	variablesArr := re.FindAllStringSubmatch(m.Content, -1)
	mod, _ := strconv.ParseFloat(variablesArr[0][1], 32)
	dc, _ := strconv.ParseFloat(variablesArr[0][2], 32)

	message := "```ansi\n"

	ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(mod, dc)

	message += "\nChance to Crit Succeed: 	" + strconv.Itoa(ChanceCritSuccess) + "%\n"
	message += "\nChance to Succeed:		  " + strconv.Itoa(ChanceNormalSuccess) + "%\n"
	message += "\nChance to Fail:			" + strconv.Itoa(ChanceNormalFail) + "%\n"
	message += "\nChance to Crit Fail:   	" + strconv.Itoa(ChanceCritFail) + "%\n"

	message += "```"
	return message, nil
}

func saveProbabilityCalculator(mod float64, dc float64) (critSuccess int, normalSuccess int, normalFailure int, critFailure int) {
	//Chance for any success = (((21 - (dc - mod)) / 20) * 100)
	//Chance for any failure = 100 - ((21-(dc-mod))/20)*100

	var ChanceCritSuccess float64
	var ChanceNormalSuccess float64
	var ChanceNormalFail float64
	var ChanceCritFail float64

	// Success Checks
	ChanceCritSuccess = round64((11 - dc + mod) * 5)
	if 0 < dc-mod && dc-mod <= 10 {
		ChanceNormalSuccess = round64((21-dc+mod)*5) - ChanceCritSuccess
	} else if dc-mod <= 0 {
		if ChanceCritSuccess < 100 {
			ChanceNormalSuccess = 100 - ChanceCritSuccess
		} else {
			ChanceNormalSuccess = 5
			ChanceCritSuccess = 95
		}

	} else if 20 < dc-mod && dc-mod <= 29 {
		ChanceNormalSuccess = 5

	} else {
		ChanceNormalSuccess = round64((21 - dc + mod) * 5)
		if ChanceCritSuccess <= 0 && ChanceNormalSuccess > 0 {
			ChanceCritSuccess = 5
			ChanceNormalSuccess -= 5
		}
	}

	// Fail Checks
	ChanceCritFail = round64(100 - ((30 - (dc - mod)) * 5))
	if 0 < dc-mod && dc-mod <= 10 {
		ChanceNormalFail = 100 - ChanceNormalSuccess - ChanceCritSuccess
		ChanceCritFail = 0
	} else if dc-mod <= 0 {

	} else if 20 < dc-mod && dc-mod <= 30 {
		ChanceCritFail -= 5
		ChanceNormalFail = 100 - ChanceNormalSuccess - ChanceCritFail
	} else if dc-mod >= 30 {
		ChanceNormalSuccess = 0
		ChanceCritSuccess = 0
		ChanceNormalFail = 5
		ChanceCritFail = 95

	} else {
		ChanceNormalFail = 100 - ChanceNormalSuccess - ChanceCritSuccess - ChanceCritFail
		if ChanceCritFail <= 0 && ChanceNormalFail > 0 {
			ChanceCritFail = 5
			ChanceNormalFail -= 5
		}
	}

	critSuccess = divisibleBy5Rounder(ChanceCritSuccess)
	normalSuccess = divisibleBy5Rounder(ChanceNormalSuccess)
	normalFailure = divisibleBy5Rounder(ChanceNormalFail)
	critFailure = divisibleBy5Rounder(ChanceCritFail)

	return critSuccess, normalSuccess, normalFailure, critFailure
}

func round64(n float64) float64 {
	if n <= 0 {
		return 0
	}
	n = math.Round(n/5) * 5
	if n > 100 {
		n = 100
	}
	return n
}

//
//func toFixed(num float64, precision int) float64 {
//	output := math.Pow(10, float64(precision))
//	return float64(round(num*output)) / output
//}

//func rounder(n float64) float64 {
//	if n < 0.5 && n > -0.5 {
//		n = math.Ceil(n)
//	} else {
//		math.Round(n)
//	}
//	return n
//}

func divisibleBy5Rounder(n float64) int {
	if n <= 0 {
		return 0
	}
	return int(math.Round(n/5) * 5)
}

// DiceRollBasic returns arr of dice rolled in int, then string, then returns total in int and string
func DiceRollBasic(sidedDice int, timesRolled int) ([]int, []string, int, string) {
	var IntArr []int
	var StrArr []string
	var sum int
	for j := 0; j < timesRolled; j++ {
		// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
		if sidedDice == 0 {
			return nil, nil, 0, ""
		}
		x := rand.Intn(sidedDice) + 1
		IntArr = append(IntArr, x)
		StrArr = append(StrArr, strconv.Itoa(x))
		sum += x
	}

	return nil, nil, 0, ""
}

func randomNumberGenerator(maxNumberPossible int) int {
	cachedCardTimer := time.NewTimer(1 * time.Millisecond)
	<-cachedCardTimer.C
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rollValue := r1.Intn(maxNumberPossible)
	return rollValue
}
