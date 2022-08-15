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

func TestColorsQueryBuilder(t *testing.T) {

	//Color testing
	t.Run("test query Type, Color, text contains rancor", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:enchantment, text:rancor, color:g",
			expected: "https://api.scryfall.com/cards/search?q=t%3Aenchantment+c%3Dg+o%3A%27rancor%27+",
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

	t.Run("type and basic color test", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:squirrel, color:b",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+c%3Db+",
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

	//Tests for black and green and black&green
	t.Run("type and 2 basic color test", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:squirrel, color:b g",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+c%3C%3Dbg+",
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

	t.Run("triple color exclusive test", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, color:wbr",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dwbr+",
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

	//Tests for white and blue and white&blue
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

	t.Run("test for multicolor EXCLUSIVE query", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:planeswalker, color:wu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Aplaneswalker+c%3Dwu+",
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

	//t.Run("", func(t *testing.T) {
	//	testCase := TestCase{
	//		input:    "",
	//		expected: "https://api.scryfall.com/cards/search?q=",
	//	}
	//	output, err := builder.MtgQueryBuilder(testCase.input)
	//	if testCase.expected != output {
	//		fmt.Println("Output:   ", output)
	//		fmt.Println("Expected: ", testCase.expected)
	//		t.Fail()
	//	}
	//	if err != nil {
	//		t.Fail()
	//	}
	//})

}

//Tests for Power, Toughness, Loyalty, and CMC
func TestInequalitiesQueryBuilder(t *testing.T) {
	t.Run("1 sided inequality test toughness", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, toughness:<4, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+tou%3C4+",
			//expected: "https://api.scryfall.com/cards/search?q=",
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
	t.Run("2 sided inequality test cmc", func(t *testing.T) {
		testCase := TestCase{
			input:    "cmc:3<=m<6",
			expected: "https://api.scryfall.com/cards/search?q=cmc%3E%3D3+cmc%3C6+",
			//expected: "https://api.scryfall.com/cards/search?q=",
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

	t.Run("test 2 sided query toughness", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, toughness:4<=t<6, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+tou%3E%3D4+tou%3C6+",
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

func TestTypesQueryBuilder(t *testing.T) {
}
func TestIsQueryBuilder(t *testing.T) {
}
func TestFunctionQueryBuilder(t *testing.T) {
}
func TestArtQueryBuilder(t *testing.T) {
}
func TestRarityQueryBuilder(t *testing.T) {
}
func TestTextQueryBuilder(t *testing.T) {
}
func TestQueryBuilder(t *testing.T) {
}
