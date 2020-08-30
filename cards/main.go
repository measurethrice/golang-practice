package main	

// import "fmt"	

// card := "Ace of Spades"	// cannot assign outside of function body
// var deckSize int 		// we can declare a variable, however

func main() {
	// Call newDeck() and create the deck
	// 'cards'.
	cards := newDeck()
	cards.saveToFile("my_cards")


}

