package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	deckLength := 52
	firstCard := "Ace of Spades"
	lastCard := "King of Clubs"

	if len(d) != deckLength {
		t.Errorf("Expected %v cards but %v exists", deckLength, (d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected the first card to be %v\nBut got %v", firstCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected the last card to be %v\nBut got %v", lastCard, d[len(d)-1])
	}
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	deckLength := 52
	test_file := "mydeck.txt"
	os.Remove(test_file)

	d := newDeck()
	d.saveToFile(test_file)

	loadedDeck := newDeckFromFile(test_file)

	if len(loadedDeck) != deckLength {
		t.Errorf("Expected %v cards but %v exists", deckLength, (d))
	}
}
