//go:build e2e
// +build e2e

package tests

import (
	"goland-discord-bot/bot/business/query/builder"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	//Testing type and 1 color request
	expected := "https://api.scryfall.com/cards/search?q=t%3Asquirrel+c%3Ab+"
	input := "type:squirrel, color:b"
	output, err := builder.MtgQueryBuilder(input)
	if expected != output {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
	if err != nil {
		t.Fatalf("%q", err)
	}
}
