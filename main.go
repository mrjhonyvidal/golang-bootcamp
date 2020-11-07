package main

// Main will be used to create and manipulate a deck of cards
func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	// append() creates a new slice and returns a new one
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
