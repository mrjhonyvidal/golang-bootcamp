package main

// Main will be used to create and manipulate a deck of cards
func main() {
	//cards := newDeck()
	//cards.saveToFile("my_cards")

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
