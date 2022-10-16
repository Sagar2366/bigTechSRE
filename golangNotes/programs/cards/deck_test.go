package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected length of 20 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs but got %v", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeckFromFile := newDeckFromFile("_deckTesting")

	if len(loadedDeckFromFile) != 20 {
		t.Errorf("Expected 20 cards in deck, but got %v", len(loadedDeckFromFile))
	}

	os.Remove("_deckTesting")
}
