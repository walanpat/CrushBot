package dicerolling

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func FiveEStats() (string, error) {
	message := "```ansi\n                             Stats:\n"
	statHolder := [5]int{}
	messageHolder := [6]string{}
	totalHolder := [6]int{}
	extraMessage := ""

	for i := 0; i <= 5; i++ {
		extraMessage = ""
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		for j := 0; j < 4; j++ {
			cachedCardTimer := time.NewTimer(1 * time.Millisecond)
			<-cachedCardTimer.C
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
		//message += messageHolder[i]

	}
	sort.Ints(totalHolder[:])

	for i := 5; i >= 0; i-- {
		for j := 0; j < 6; j++ {
			x, err := strconv.Atoi(messageHolder[j][len(messageHolder[j])-3 : len(messageHolder[j])-1])
			if err != nil {
				fmt.Errorf("error detected on str conversion: %e", err)
			}
			if x == 0 {
				x, err = strconv.Atoi(messageHolder[j][len(messageHolder[j])-2 : len(messageHolder[j])-1])
			} else if x == 0 {
				continue
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
	//Initializing our "base" query expression
	re := regexp.MustCompile(`([+\-]?\d+)*d(\d+)([+\-]?\d*[^\dd][^d]+)*`)
	variablesArr := re.FindAllStringSubmatch(m.Content, -1)
	message := "```ansi\n \u001B[0m"
	sumTotal := 0
	for i := 0; i < len(variablesArr); i++ {
		total := 0
		if len(variablesArr) == 1 {
			//This IF statement only occurs if there is a single "roll" (x^n)d(y^n)+(z^n) occurring
			numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
			diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

			//Basic math, modifier regexp
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
			message += m.Author.Username + " LETS ROLL\n\n"
			message += "\u001B[3m" + strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "\u001B[0m	"

			//Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
				if diceToBeRolled == 0 {
					return "Dice cannot be 0 sided.", nil
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
			//Adds any and all modifiers
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
			//this code block occurs only if (x^n)d(y^n)+(z^n) occurs more than once.
			numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])
			diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])

			//used to find any modifiers added or subtracted to various dice roll
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)

			message += m.Author.Username + " LETS ROLL\n\n"
			message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

			//Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
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
			//Adds any and all modifiers, makes sure we don't add the next dice roll
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
			//This code block is for ALL code after the original (x^n)d(y^n)+(z^n) amount of dice/things rolled.
			basicMathRE := regexp.MustCompile(`([+\-]?\d*)`)
			numbOfRoll, _ := strconv.Atoi(variablesArr[i][1])

			if i != 0 {
				initialArray := basicMathRE.FindAllStringSubmatch(variablesArr[i-1][3], -1)
				numbOfRoll, _ = strconv.Atoi(initialArray[len(initialArray)-1][0])
			}

			//this has to be initialized weirdly.  So.  Here's what we do
			diceToBeRolled, _ := strconv.Atoi(variablesArr[i][2])
			arithmetic := basicMathRE.FindAllStringSubmatch(variablesArr[i][3], -1)
			//message now tells us what we're rolling
			message += strconv.Itoa(numbOfRoll) + "d" + strconv.Itoa(diceToBeRolled) + "	"

			//Actual for loop for the "dice rolls"
			for j := 0; j < numbOfRoll; j++ {
				//This has to be diceToBeRolled +1 because rand.intn uses [0,n) noninclusive n.
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

			//If, we are NOT at the last modifier/read value
			if i+1 != len(variablesArr) {
				for j := 0; j < len(arithmetic)-1; j++ {
					x, _ := strconv.Atoi(arithmetic[j][0])
					arithmeticResult += x
				}
			} else {
				//If it is the last value, then add it if there's any modifiers/arithmetic values.
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
