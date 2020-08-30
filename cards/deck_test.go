package main

import (
	"os"
	"testing"
)

// This fuction asserts some default properties of the starting deck. 
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Verify the length of the deck. 
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// Verify the first card in the deck. 
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])

	}

	// Verify the last card in the deck. 
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}



}

// This function asserts whether the saveToFile() and newDeckFromFile()
// functions are working properly.
func TestSaveToDeckAndNewDeckFromFile(t *testing.T)  {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}