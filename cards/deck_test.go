package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	if len(deck) != 48 {
		t.Errorf("Expected deck length of 20, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", deck[len(deck)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// TODO
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 48 {
		t.Errorf("Expected 48 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
