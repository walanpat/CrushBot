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
	for i := 0; i < len(variablesArr); i++ {
		total := 0
		if len(variablesArr) == 1 {
			// This IF statement only occurs if there is a single "roll" (x^n)d(y^n)+(z^n) occurring
			numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
			diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

			// Basic math, modifier regexp
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
			message += m.Author.Username + " LETS ROLL\n\n"
			message += "\u001B[3m" + strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "\u001B[0m	"

			// Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				if diceToBeRolled == 0 {
					return "```ansi\nDice cannot be 0 sided.```", nil
				}
				rollValueInt := rand.Intn(diceToBeRolled) + 1
				rollValueStr := strconv.Itoa(rollValueInt)
				total += rollValueInt
				if rollValueInt == 1 {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[31m" + rollValueStr + "\u001B[0m]	"
				} else if rollValueInt == diceToBeRolled {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[32m" + rollValueStr + "\u001B[0m]	"
				} else {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[37m" + rollValueStr + "\u001B[0m]	"
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
				message += "\n\n	Modifiers: [\u001B[37m" + strconv.Itoa(arithmeticResult) + "\u001B[0m]\n"
			}
			message += "\n\n	Total Roll: " + "[\u001B[37m" + strconv.Itoa(total) + "\u001B[0m]\n\n"
		} else if i == 0 && len(variablesArr) > 1 {
			// this code block occurs only if (x^n)d(y^n)+(z^n) occurs more than once.
			numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
			diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

			// used to find any modifiers added or subtracted to various dice roll
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)

			message += m.Author.Username + " LETS ROLL\n\n"
			message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

			// Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				rollValueInt := rand.Intn(diceToBeRolled) + 1
				rollValueStr := strconv.Itoa(rollValueInt)
				total += rollValueInt
				sumTotal += rollValueInt
				if rollValueInt == 1 {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[31m" + rollValueStr + "\u001B[0m]	"
				} else if rollValueInt == diceToBeRolled {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[32m" + rollValueStr + "\u001B[0m]	"
				} else {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[37m" + rollValueStr + "\u001B[0m]	"
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
				message += "\n\n	Modifiers: [\u001B[37m" + strconv.Itoa(arithmeticResult) + "\u001B[0m]\n"
				sumTotal += arithmeticResult
			}
			message += "\n	Total Roll: " + "[\u001B[37m" + strconv.Itoa(total) + "\u001B[0m]\n\n"
		} else {
			// This code block is for ALL code after the original (x^n)d(y^n)+(z^n) amount of dice/things rolled.
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])

			if i != 0 {
				initialArray := basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)
				numbOfRoll, _ = strconv.Atoi(initialArray[len(initialArray)-1][0])
			}

			// this has to be initialized weirdly.  So.  Here's what we do
			diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
			// message now tells us what we're rolling
			message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

			// Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				// This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				rollValueInt := rand.Intn(diceToBeRolled) + 1
				rollValueStr := strconv.Itoa(rollValueInt)
				total += rollValueInt
				sumTotal += rollValueInt
				if rollValueInt == 1 {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[31m" + rollValueStr + "\u001B[0m]	"
				} else if rollValueInt == diceToBeRolled {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[32m" + rollValueStr + "\u001B[0m]	"
				} else {
					message += "Roll " + strconv.Itoa(j+1) + ": [\u001B[37m" + rollValueStr + "\u001B[0m]	"
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
				message += "\n\n	Modifiers: [\u001B[37m" + strconv.Itoa(arithmeticResult) + "\u001B[0m]\n"
			}
			message += "\n	Total Roll: " + "[\u001B[37m" + strconv.Itoa(total) + "\u001B[0m]\n\n"
			if i+1 == len(variablesArr) {
				message += "\n Sum Total of All Values: \u001B[37m" + strconv.Itoa(sumTotal) + "\u001B[0m"
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

	//fmt.Printf("\n\nvariablesARr: %v", variablesArr)
	//
	//fmt.Printf("\n\nMod: %v", mod)
	//fmt.Printf("\nDC: %v", dc)

	message := "```ansi\n" // \u001B[0m"

	ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(mod, dc)
	strCritSuccess := strconv.Itoa(ChanceCritSuccess)
	strNormSuccess := strconv.Itoa(ChanceNormalSuccess)
	strNormFail := strconv.Itoa(ChanceNormalFail)
	strCritFail := strconv.Itoa(ChanceCritFail)

	//fmt.Print("\nChance to Crit Succeed: 	 " + strCritSuccess)
	//fmt.Print("\nChance to Succeed:		 " + strNormSuccess)
	//fmt.Print("\nChance to Fail:			" + strNormFail)
	//fmt.Print("\nChance to Crit Fail:   		" + strCritFail)

	message += "\nChance to Crit Succeed: 	" + strCritSuccess + "%\n"
	message += "\nChance to Succeed:		  " + strNormSuccess + "%\n"
	message += "\nChance to Fail:			" + strNormFail + "%\n"
	message += "\nChance to Crit Fail:   	" + strCritFail + "%\n"

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

	var diceVar = float64(20)

	var critSVar = float64(10)
	var critFVar = float64(10)
	var succVar = float64(0)
	var failVar = float64(1)

	if mod == 10 {
		diceVar = 19
		critSVar = 10
		critFVar = critSVar

	} else if mod == 11 {
		diceVar = 19
		critSVar = 8
		critFVar = 9

		succVar = 0
		failVar = 2
	} else if mod >= 12 {
		diceVar = 19
		critSVar = 8
		critFVar = 8

		succVar = 1
		failVar = 2
	}

	ChanceCritSuccess = (diceVar - (dc + critSVar - mod)) * 5
	ChanceNormalSuccess = (diceVar - (dc - mod + succVar)) * 5

	ChanceNormalFail = ((diceVar + failVar) - (dc - mod)) * 5

	ChanceCritFail = ((diceVar + critFVar) - (dc - mod)) * 5
	ChanceCritFail = 100 - ChanceCritFail

	// Norm fail check
	if ChanceNormalFail > 0 {
		ChanceNormalFail = 100 - ChanceNormalFail
	} else {
		ChanceNormalFail = 0
	}

	// Crit fail interacting with normal fail
	if ChanceCritFail > 0 && ChanceNormalFail > 0 {
		ChanceNormalFail -= ChanceCritFail
	}

	// Check to see if regular hit >0
	if ChanceNormalSuccess >= 0 && ChanceCritSuccess <= 0 {
		ChanceCritSuccess = 5
	} else if ChanceNormalSuccess < 0 {
		ChanceCritSuccess = 0
		if ChanceCritFail < 100 {
			ChanceNormalSuccess = 5
		} else {
			ChanceNormalSuccess = 0
		}
	}

	// If chance for failure hard hits 0, needs to rework the formula to just subtract from what is correct.
	if ChanceNormalFail == 0 && ChanceNormalSuccess > 0 && ChanceCritFail > 0 && ChanceCritSuccess <= 0 {
		ChanceNormalFail = 100 - ChanceNormalSuccess - ChanceCritFail
	}

	if ChanceCritFail == 0 && ChanceNormalSuccess > 0 && ChanceNormalFail > 0 && ChanceCritSuccess > 0 {
		ChanceCritFail = 100 - ChanceNormalSuccess - ChanceNormalFail - ChanceCritSuccess
		if ChanceCritFail < 0 {
			ChanceCritFail = 0
		}
	}

	critSuccess = divisibleBy5Rounder(ChanceCritSuccess)
	normalSuccess = divisibleBy5Rounder(ChanceNormalSuccess)
	normalFailure = divisibleBy5Rounder(ChanceNormalFail)
	critFailure = divisibleBy5Rounder(ChanceCritFail)

	return critSuccess, normalSuccess, normalFailure, critFailure
}

//func round(num float64) int {
//	return int(num + math.Copysign(0.5, num))
//}
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
	integer := int(math.Round(n))
	if integer%5 != 0 {
		if (integer+1)%5 == 0 {
			integer++
		} else if (integer-1)%5 == 0 {
			integer--
		}
	}
	return integer
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
