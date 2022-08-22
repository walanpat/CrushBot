package tests

import (
	"fmt"
	"goland-discord-bot/bot/services"
	"testing"
)

type TestCase struct {
	input    string
	expected string
	actual   bool
}

func TestScryfallQueryErrorResponse(t *testing.T) {
	t.Run("no cards found response (active ping to scryfall)", func(t *testing.T) {
		testCase := TestCase{
			input:    "https://api.scryfall.com/cards/search?q=t%3Avampire+o%3A%27rancor%27+",
			expected: "scryfall returned an error object, either nothing was found or there is a bad input",
		}
		output, err := services.GetQueryService(testCase.input)
		fmt.Println(output)
		if err.Error() != testCase.expected {
			fmt.Println("Output:   ", err)
			fmt.Println("Expected: ", testCase.expected)
		}
		if err == nil {
			t.Fail()
		}
	})
}
