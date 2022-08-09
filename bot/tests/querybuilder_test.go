//go:build e2e
// +build e2e

package tests

import (
	"fmt"
	"goland-discord-bot/bot/business/query/builder"
	"testing"
)

type TestCase struct {
	input    string
	expected string
	actual   bool
}

func TestQueryBuilder(t *testing.T) {
	//Testing type and 1 color request
	t.Run("Query Builder: type and basic color test", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:squirrel, color:b",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+c%3Ab+",
		}
		output, err := builder.MtgQueryBuilder(testCase.input)
		if testCase.expected != output {
			fmt.Println("Output:   ", output)
			fmt.Println("Expected: ", testCase.expected)
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}
	})
	t.Run("test for multicolor query (ub || b || u)", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, color:ub b u",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3C%3Dub+",
		}
		output, err := builder.MtgQueryBuilder(testCase.input)
		if testCase.expected != output {
			fmt.Println("Output:   ", output)
			fmt.Println("Expected: ", testCase.expected)
			t.Fail()
		}
		if err != nil {
			t.Fail()
		}
	})

}
