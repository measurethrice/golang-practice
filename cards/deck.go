package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// This function creates and returns a deck 
// of 52 cards from two slices.
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This function takes the original deck 'd',
// then prints each element of the deck.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// This function takes the original deck 'd', 
// splits it by 'handSize', then returns two 
// new decks.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// This helper function takes the 'deck'
// type and converts it into a string. 
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	
}

// This helper function  takes the 'deck' type, 
// converts it to a string, then converts it 
// to a byte-slice. It then write the byte-slice
// to a file. 
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// This helper function opens a file 
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Convert the byte-slice into a string,
	// and split it by the comma delimiter.
	s := strings.Split(string(bs), ",")

	// Convert the split string into the 
	// deck type.
	return deck(s)
}

