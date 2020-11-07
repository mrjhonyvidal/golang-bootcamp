package main

// Main will be used to create and manipulate a deck of cards
func main() {
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
