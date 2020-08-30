package main	

import "fmt"	

// card := "Ace of Spades"	// cannot assign outside of function body
// var deckSize int 		// we can declare a variable, however

func main() {
	cards := []string{newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
