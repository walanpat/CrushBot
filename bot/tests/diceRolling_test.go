package tests

import (
	"fmt"
	"goland-discord-bot/bot/business/dicerolling"
	"testing"
)

//Number generation test
func FiveECharacterCreation(t *testing.T) {
	t.Run("Returns sorted", func(t *testing.T) {
		testCase := TestCase{
			input:    "",
			expected: "eh",
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
