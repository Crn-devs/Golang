package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Error: deck length must be 52 got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got: %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs but got: %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {

	// remove any test file if it exists
	os.Remove("_deckTesting")
	// try to create a deck
	deck := newDeck()
	// save the deck to a file
	deck.saveDeckToFile("_deckTesting")
	// read deck from test file
	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != len(deck) {
		t.Errorf("expected a deck of length: %v but got length:%v", len(deck), len(loadedDeck))
	}
	// remove the file
	os.Remove("_deckTesting")
}
