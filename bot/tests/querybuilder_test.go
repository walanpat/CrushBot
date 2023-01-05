package tests

import (
	"fmt"
	"goland-discord-bot/bot/business/query/builder"
	"testing"
)

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

//Testing color inputs
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
			input:    "type:squirrel, color:b or g",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+%28c%3Db+or+c%3Dg%29+",
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
			input:    "type:vampire, color:ub or b or u",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+%28c%3Dub+or+c%3Db+or+c%3Du%29+",
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
	t.Run("1 sided inequality toughness", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, toughness:<4, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+tou%3C4+",
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
	t.Run("2 sided inequality toughness", func(t *testing.T) {
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

	t.Run("1 sided inequality cmc", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, cmc:<=4, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+cmc%3C=4+",
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
	t.Run("1 sided inequality cmc <=10", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, cmc:<=10, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+cmc%3C=10+",
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
	t.Run("2 sided inequality cmc", func(t *testing.T) {
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
	t.Run("1 sided inequality power", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, power:<=2, cmc:<=4, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+cmc%3C=4+pow%3C=2+",
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
	t.Run("2 sided inequality power", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:vampire, power:2<=p<=3, cmc:<=4, color:bu",
			expected: "https://api.scryfall.com/cards/search?q=t%3Avampire+c%3Dbu+cmc%3C=4+pow%3E%3D2+pow%3C%3D3+",
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

//testing type inputs
func TestTypesQueryBuilder(t *testing.T) {
	t.Run("testing for  goblin sorcery", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:goblin sorcery",
			expected: "https://api.scryfall.com/cards/search?q=t%3Agoblin+t%3Asorcery+",
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
	t.Run("testing for fae instant", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:fae instant",
			expected: "https://api.scryfall.com/cards/search?q=t%3Afae+t%3Ainstant+",
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
	t.Run("testing Orc Pirate Wizard", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:orc pirate wizard",
			expected: "https://api.scryfall.com/cards/search?q=t%3Aorc+t%3Apirate+t%3Awizard+",
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
	t.Run("testing for Zombie Minotaur Warrior", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:minotaur zombie warrior",
			expected: "https://api.scryfall.com/cards/search?q=t%3Aminotaur+t%3Azombie+t%3Awarrior+",
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
	t.Run("testing for egg or ooze mutant", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:egg or ooze mutant",
			expected: "https://api.scryfall.com/cards/search?q=t%3Aegg+OR+t%3Aooze+t%3Amutant+",
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
	t.Run("testing for egg artifact or ooze mutant or Phelddagrif", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:egg artifact or ooze mutant or Phelddagrif ",
			expected: "https://api.scryfall.com/cards/search?q=t%3Aegg+t%3Aartifact+OR+t%3Aooze+t%3Amutant+OR+t%3APhelddagrif+",
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

//testing is inputs
func TestIsQueryBuilder(t *testing.T) {
	t.Run("is etb ", func(t *testing.T) {
		testCase := TestCase{
			input:    "is:etb",
			expected: "https://api.scryfall.com/cards/search?q=is%3Aetb+",
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

//testing functional tag inputs
func TestFunctionQueryBuilder(t *testing.T) {
	t.Run("function removal ", func(t *testing.T) {
		testCase := TestCase{
			input:    "function:removal",
			expected: "https://api.scryfall.com/cards/search?q=function%3Aremoval+",
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
	t.Run("function flyer hate ", func(t *testing.T) {
		testCase := TestCase{
			input:    "function:flyer hate",
			expected: "https://api.scryfall.com/cards/search?q=function%3Aflyer+function%3Ahate+",
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

//testing art tag inputs
func TestArtQueryBuilder(t *testing.T) {
	t.Run("", func(t *testing.T) {
		testCase := TestCase{
			input:    "art:squirrel, type:squirrel",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+art%3Asquirrel+",
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

//testing rarity inputs
func TestRarityQueryBuilder(t *testing.T) {
	t.Run("Common Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "art:squirrel, type:squirrel, rarity:c",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+r%3Acommon+art%3Asquirrel+",
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
	t.Run("unCommon Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "art:squirrel, type:squirrel, rarity:u",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+r%3Auncommon+art%3Asquirrel+",
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
	t.Run("rare Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "art:squirrel, type:squirrel, rarity:r",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+r%3Arare+art%3Asquirrel+",
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
	t.Run("Mythic test", func(t *testing.T) {
		testCase := TestCase{
			input:    "art:squirrel, type:squirrel, rarity:m",
			expected: "https://api.scryfall.com/cards/search?q=t%3Asquirrel+r%3Amythic+art%3Asquirrel+",
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
	t.Run("Simplified uncommon or common Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "rarity:uc",
			expected: "https://api.scryfall.com/cards/search?q=%28r%3Auncommon+OR+r%3Acommon%29+",
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

	t.Run("Simplified uncommon or common or rare Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "rarity:ucr",
			expected: "https://api.scryfall.com/cards/search?q=%28r%3Auncommon+OR+r%3Acommon+OR+r%3Arare%29+",
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
	t.Run("Simplified uncommon or common or rare or mythic Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "rarity:ucrm",
			expected: "https://api.scryfall.com/cards/search?q=%28r%3Auncommon+OR+r%3Acommon+OR+r%3Arare+OR+r%3Amythic%29+",
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
	t.Run("Uncommon or Common Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "rarity:u or c",
			expected: "https://api.scryfall.com/cards/search?q=%28r%3Auncommon+OR+r%3Acommon%29+",
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
	t.Run("Uncommon or Common or Mythic Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "rarity:u or c or m",
			expected: "https://api.scryfall.com/cards/search?q=%28r%3Auncommon+OR+r%3Acommon+OR+r%3Amythic%29+",
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
	t.Run("Uncommon or Common or Mythic or Rare Test", func(t *testing.T) {
		testCase := TestCase{
			input:    "rarity:u or c or m or r",
			expected: "https://api.scryfall.com/cards/search?q=%28r%3Auncommon+OR+r%3Acommon+OR+r%3Amythic+OR+r%3Arare%29+",
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

//testing text search input
func TestTextQueryBuilder(t *testing.T) {
	t.Run("test text rancor", func(t *testing.T) {
		testCase := TestCase{
			input:    "text:rancor",
			expected: "https://api.scryfall.com/cards/search?q=o%3A%27rancor%27+",
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

//testing query builder return statement
func TestQueryBuilder(t *testing.T) {
	t.Run("Type, 3 color, and text query", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:insect, color:g or b or gb, text:landfall",
			expected: "https://api.scryfall.com/cards/search?q=t%3Ainsect+%28c%3Dg+or+c%3Db+or+c%3Dgb%29+o%3A%27landfall%27+",
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

//Testing for edge case inputs/error handling
func TestBadInputHandling(t *testing.T) {
	t.Run("no comma but multiple input types handling", func(t *testing.T) {
		testCase := TestCase{
			input:    "type:goblin sorcery instant is:etb function:removal",
			expected: "error: more than 2 search modifiers entered, but no comma detected",
		}
		output, err := builder.MtgQueryBuilder(testCase.input)
		if err.Error() != testCase.expected {
			fmt.Println("Output:   ", err.Error())
			fmt.Println("Expected: ", testCase.expected)
			t.Fail()
		}
		if output != "" {
			fmt.Println("a response that was not an error was given")
			t.Fail()
		}

	})
}

//personal

//func TestPersonalQueryBuilder(t *testing.T) {
//	t.Run("TestPersonalQueryBuilder", func(t *testing.T) {
//		testCase := TestCase{
//			input:    "text:roll a , color:r or g or rg",
//			expected: "https://api.scryfall.com/cards/search?q=",
//		}
//		output, err := builder.MtgQueryBuilder(testCase.input)
//		if testCase.expected != output {
//			fmt.Println("Output:   ", output)
//			fmt.Println("Expected: ", testCase.expected)
//			t.Fail()
//		}
//		if err != nil {
//			t.Fail()
//		}
//	})
//}
