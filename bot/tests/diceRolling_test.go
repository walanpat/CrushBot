package tests

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"goland-discord-bot/bot/business/dicerolling"
	"strconv"
	"testing"
)

//Testing random number generation
func TestFiveEStatCreation(t *testing.T) {
	t.Run("Returns sorted", func(t *testing.T) {
		testCase := TestCase{
			input:    "",
			expected: "",
		}
		output, err := dicerolling.FiveEStats()
		fmt.Println(output)
		if output != testCase.expected {
			fmt.Println("Output:   ", err)
			fmt.Println("Expected: ", testCase.expected)
		}
		if err != nil {
			t.Fail()
		}
	})
}

//Basic Dice Rolling Test
func TestDiceRolling(t *testing.T) {
	t.Run("xdy any values (no modifiers) ", func(t *testing.T) {
		//Here's the inputs.
		timesRolledInt := 2
		amountOfSidesInt := 100

		timesRolledString := strconv.Itoa(timesRolledInt)
		amountOfSidesString := strconv.Itoa(amountOfSidesInt)

		testCase := TestCase{
			input:    "!roll " + timesRolledString + "d" + amountOfSidesString,
			expected: "number between " + timesRolledString + " and " + amountOfSidesString,
		}
		//Setup
		testAuthor := discordgo.User{Username: "CrushTestUserName"}
		testMessage := discordgo.Message{Content: testCase.input, Author: &testAuthor}
		testMessageCreate := discordgo.MessageCreate{Message: &testMessage}
		//Function Call
		output, err := dicerolling.DiceRollGeneric(&testMessageCreate)

		if output[46:46+len(timesRolledString)] != timesRolledString {
			t.Fail()
		} else {
			fmt.Println("Times Rolled is Correct")
		}
		if output[48:48+len(amountOfSidesString)] != amountOfSidesString {
			t.Fail()
		} else {
			fmt.Println("Amount of Sides is Correct")
		}

		//Check the value generated
		for i := 1; i <= len(amountOfSidesString); i++ {
			if output[70+i-1:70+i] != "\u001B" {
				value, err := strconv.Atoi(output[70 : 70+i])
				if err != nil {
					t.Fail()
				}
				if amountOfSidesInt < value {
					t.Fail()
				}
				if value < 1 {
					t.Fail()
				}
			}
		}

		if err != nil {
			t.Fail()
		}
	})
}

func TestDiceRollingWithMods(t *testing.T) {
	t.Run("xdy any values (no modifiers) ", func(t *testing.T) {
		//Here's the inputs.
		timesRolledInt := 1
		amountOfSidesInt := 10
		modifierArray := []int{1, 2, 3, 4}

		timesRolledString := strconv.Itoa(timesRolledInt)
		amountOfSidesString := strconv.Itoa(amountOfSidesInt)
		modifierString := ""

		totalModInt := 0
		for i := 0; i < len(modifierArray); i++ {
			totalModInt += modifierArray[i]
			if 0 <= modifierArray[i] {
				modifierString += "+" + strconv.Itoa(modifierArray[i])
			} else {
				modifierString += strconv.Itoa(modifierArray[i])
			}
		}

		testCase := TestCase{
			input:    "!roll " + timesRolledString + "d" + amountOfSidesString + modifierString + "",
			expected: "number between " + timesRolledString + " and " + amountOfSidesString + "plus modifier of " + strconv.Itoa(totalModInt),
		}
		fmt.Println(testCase.input)
		//Setup
		testAuthor := discordgo.User{Username: "CrushTestUserName"}
		testMessage := discordgo.Message{Content: testCase.input, Author: &testAuthor}
		testMessageCreate := discordgo.MessageCreate{Message: &testMessage}
		//Function Call
		output, err := dicerolling.DiceRollGeneric(&testMessageCreate)

		if output[46:46+len(timesRolledString)] != timesRolledString {
			t.Fail()
		}
		if output[48:48+len(amountOfSidesString)] != amountOfSidesString {
			t.Fail()
		}

		totalModOutput, err := strconv.Atoi(output[96:98])
		if err != nil {
			t.Fail()
		}
		if totalModInt != totalModOutput {
			t.Fail()
		}

		//Check the value generated
		for i := 1; i <= len(amountOfSidesString); i++ {
			fmt.Println(output[125:127])
			//Total Output starts at [125:127~]
			if output[125+i-1:125+i] != "\u001B" {
				value, err := strconv.Atoi(output[125+i-1 : 125+i])
				if err != nil {
					t.Fail()
				}
				if amountOfSidesInt+totalModInt < value {
					t.Fail()
				}
			}
		}

		if err != nil {
			t.Fail()
		}
	})
}
