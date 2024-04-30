package dicerolling

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/bwmarrin/discordgo"
)

type TestCase struct {
	input    string
	expected string
	actual   bool
}

// Testing random number generation
func TestFiveEStatCreation(t *testing.T) {
	t.Run("Returns sorted", func(t *testing.T) {
		testCase := TestCase{
			input:    "",
			expected: "",
		}
		output, err := FiveEStats()
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

// Basic Dice Rolling Test
func TestDiceRolling(t *testing.T) {
	t.Run("xdy any values (no modifiers) ", func(t *testing.T) {
		// Here's the inputs.
		timesRolledInt := 2
		amountOfSidesInt := 100

		timesRolledString := strconv.Itoa(timesRolledInt)
		amountOfSidesString := strconv.Itoa(amountOfSidesInt)

		testCase := TestCase{
			input:    "!roll " + timesRolledString + "d" + amountOfSidesString,
			expected: "number between " + timesRolledString + " and " + amountOfSidesString,
		}
		// Setup
		testAuthor := discordgo.User{Username: "CrushTestUserName"}
		testMessage := discordgo.Message{Content: testCase.input, Author: &testAuthor}
		testMessageCreate := discordgo.MessageCreate{Message: &testMessage}

		// Function Call
		output, err := DiceRollGeneric(&testMessageCreate)

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

		// Check the value generated
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
		// Here's the inputs.
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
		// Setup
		testAuthor := discordgo.User{Username: "CrushTestUserName"}
		testMessage := discordgo.Message{Content: testCase.input, Author: &testAuthor}
		testMessageCreate := discordgo.MessageCreate{Message: &testMessage}
		// Function Call
		output, err := DiceRollGeneric(&testMessageCreate)

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

		// Check the value generated
		for i := 1; i <= len(amountOfSidesString); i++ {
			fmt.Println(output[125:127])
			// Total Output starts at [125:127~]
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

func TestSaveProbabilityCalculator(t *testing.T) {
	type TestCase struct {
		inputMod         float64
		inputDC          float64
		expectedCritSucc int
		expectedSucc     int
		expectedFail     int
		expectedCritFail int
	}
	//Template Test
	//Actual Tests
	t.Run("mod:0, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         0,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     0,
			expectedFail:     45,
			expectedCritFail: 50,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:1, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         1,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     5,
			expectedFail:     45,
			expectedCritFail: 45,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:2, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         2,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     10,
			expectedFail:     45,
			expectedCritFail: 40,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", 100)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:3, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         3,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     15,
			expectedFail:     45,
			expectedCritFail: 35,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:4, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         4,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     20,
			expectedFail:     45,
			expectedCritFail: 30,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:5, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         5,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     25,
			expectedFail:     45,
			expectedCritFail: 25,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:6, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         6,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     30,
			expectedFail:     45,
			expectedCritFail: 20,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:7, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         7,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     35,
			expectedFail:     45,
			expectedCritFail: 15,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:8, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         8,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     40,
			expectedFail:     45,
			expectedCritFail: 10,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:9, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod:         9,
			inputDC:          20,
			expectedCritSucc: 5,
			expectedSucc:     45,
			expectedFail:     45,
			expectedCritFail: 5,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})

	//dc-mod <=10
	t.Run("mod:10, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod: 10,
			inputDC:  20,
			//20 = crit
			//10-19 = succ
			//1-9 = fail
			expectedCritSucc: 5,
			expectedSucc:     50,
			expectedFail:     45,
			expectedCritFail: 0,
		}
		//crit succ on 20
		//norm succ 11-19
		//norm fail 2-9
		//crit fail on 1

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("sum not 100:	%v\n\n", sum)

			t.Fail()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:11, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod: 11,
			inputDC:  20,

			// crit on 19,20 	(10%)
			// succ on 9-18   	(50%)
			//fail on 2-8 		(35%)
			// crit fail on 1 	(5%)
			expectedCritSucc: 10,
			expectedSucc:     50,
			expectedFail:     40,
			expectedCritFail: 0,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:12, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod: 12,
			inputDC:  20,

			// crit on 18,19,20 	(15%) 3/20
			// succ on 8-17   	(50%) 10/20
			//fail on 2-7 		(30%) 6/20
			// crit fail on 1 	(5%)  1/20
			expectedCritSucc: 15,
			expectedSucc:     50,
			expectedFail:     35,
			expectedCritFail: 0,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:13, dc:20", func(t *testing.T) {
		testCase := TestCase{
			inputMod: 13,
			inputDC:  20,

			// crit on 17,18,19,20 	(15%) 4/20
			// succ on 7-16   	(50%) 10/20
			//fail on 2,3,4,5,6 (25%) 5/20
			// crit fail on 1 	(5%)  1/20
			expectedCritSucc: 20,
			expectedSucc:     50,
			expectedFail:     30,
			expectedCritFail: 0,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})

	//dc-mod>=20
	t.Run("mod:0 dc:21", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  21,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     45,
			expectedCritFail: 50,
		}
		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})

	t.Run("mod:0 dc:22", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  22,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     40,
			expectedCritFail: 55,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})

	t.Run("mod:0 dc:23", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  23,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     35,
			expectedCritFail: 60,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:24", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  24,

			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     30,
			expectedCritFail: 65,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:25", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  25,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     25,
			expectedCritFail: 70,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:26", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  26,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     20,
			expectedCritFail: 75,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:27", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  27,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     15,
			expectedCritFail: 80,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:28", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  28,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     10,
			expectedCritFail: 85,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:29", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  29,

			// crit on
			// succ on 20
			//fail on 11-19
			//critfail on 1-10
			expectedCritSucc: 0,
			expectedSucc:     5,
			expectedFail:     5,
			expectedCritFail: 90,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:30", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  30,

			expectedCritSucc: 0,
			expectedSucc:     0,
			expectedFail:     5,
			expectedCritFail: 95,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:0 dc:31", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 0,
			inputDC:  31,

			expectedCritSucc: 0,
			expectedSucc:     0,
			expectedFail:     5,
			expectedCritFail: 95,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})

	//high numbers, within 10 and 20 diff
	t.Run("mod:37 dc:50", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 37,
			inputDC:  50,
			//37-30 = 7
			//50-30 = 20

			expectedCritSucc: 5,
			expectedSucc:     35,
			expectedFail:     45,
			expectedCritFail: 15,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	t.Run("mod:40 dc:50", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 40,
			inputDC:  50,
			//40-30 = 10
			//50-30 = 20
			expectedCritSucc: 5,
			expectedSucc:     50,
			expectedFail:     45,
			expectedCritFail: 0,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})

	t.Run("mod:i dc:20", func(t *testing.T) {

		testCase := TestCase{
			inputMod: 13,
			inputDC:  20,

			// crit on 17,18,19,20 	(15%) 3/20
			// succ on 7-16   	(50%) 10/20
			//fail on 2-6 		(30%) 6/20
			// crit fail on 1 	(5%)  1/20
			expectedCritSucc: 20,
			expectedSucc:     50,
			expectedFail:     30,
			expectedCritFail: 5,
		}

		ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
		sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
		if sum != 100 {
			fmt.Printf("\nsum not 100:		%v\n\n", sum)
			fmt.Printf("ChanceCritSuccess Output:		%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:		%v\n\n", testCase.expectedCritSucc)
			fmt.Printf("ChanceNormalSuccess Output:		%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			fmt.Printf("ChanceNormalFail Output:		%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:		%v\n\n", testCase.expectedFail)
			fmt.Printf("ChanceCritFail Output:			%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:		%v\n\n", testCase.expectedCritFail)
			t.Fail()
			t.FailNow()
		}
		if ChanceCritSuccess != testCase.expectedCritSucc {
			fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
			fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
			t.Fail()
		}
		if ChanceNormalSuccess != testCase.expectedSucc {
			fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
			fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
			t.Fail()
		}
		if ChanceNormalFail != testCase.expectedFail {
			fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
			fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
			t.Fail()
		}
		if ChanceCritFail != testCase.expectedCritFail {
			fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
			fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
			t.Fail()
		}
	})
	//t.Run("mod:, dc:", func(t *testing.T) {
	//	testCase := TestCase{
	//		inputMod:         0,
	//		inputDC:          20,
	//		expectedCritSucc: 0,
	//		expectedSucc:     0,
	//		expectedFail:     0,
	//		expectedCritFail: 0,
	//	}
	//
	//	ChanceCritSuccess, ChanceNormalSuccess, ChanceNormalFail, ChanceCritFail := saveProbabilityCalculator(testCase.inputMod, testCase.inputDC)
	//	sum := ChanceCritSuccess + ChanceNormalSuccess + ChanceNormalFail + ChanceCritFail
	//	if sum != 100 {
	//		fmt.Printf("sum not 100:	%v\n\n", sum)
	//
	//		fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
	//		fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
	//		fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
	//		fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
	//		t.Fail()
	//	}
	//	if ChanceCritSuccess != testCase.expectedCritSucc {
	//		fmt.Printf("ChanceCritSuccess Output:	%v\n", ChanceCritSuccess)
	//		fmt.Printf("ChanceCritSuccess Expected:	%v\n\n", testCase.expectedCritSucc)
	//		t.Fail()
	//	}
	//	if ChanceNormalSuccess != testCase.expectedSucc {
	//		fmt.Printf("ChanceNormalSuccess Output:	%v\n", ChanceNormalSuccess)
	//		fmt.Printf("ChanceNormalSuccess Expected:	%v\n\n", testCase.expectedSucc)
	//		t.Fail()
	//	}
	//	if ChanceNormalFail != testCase.expectedFail {
	//		fmt.Printf("ChanceNormalFail Output:	%v\n", ChanceNormalFail)
	//		fmt.Printf("ChanceNormalFail Expected:	%v\n\n", testCase.expectedFail)
	//		t.Fail()
	//	}
	//	if ChanceCritFail != testCase.expectedCritFail {
	//		fmt.Printf("ChanceCritFail Output:		%v\n", ChanceCritFail)
	//		fmt.Printf("ChanceCritFail Expected:	%v\n\n", testCase.expectedCritFail)
	//		t.Fail()
	//	}
	//})

	//highest possible mod 34 ?
	//highest DC should be 53 (absolute max?)
}
