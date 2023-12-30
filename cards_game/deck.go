package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filePath string) {
	err := os.WriteFile(filePath, []byte(d.toString()), 0666)
	if err != nil {
		panic(err)
	}
}

func newDeckFromFile(filePath string) deck {
	byteSlice, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// fmt.Printf(string(byteSlice))

	return strings.Split(string(byteSlice), ",")
}

func (d deck) shuffle() {
	rand.NewSource(time.Now().UnixNano())
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
