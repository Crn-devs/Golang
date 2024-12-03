package main

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Error: deck length must be 52 got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got: %v", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected Ace of Spades but got: %v", d[len(d)-1])
	}
}