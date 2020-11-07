package main

import "fmt"

// Main will be used to create and manipulate a deck of cards
func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	// append() creates a new slice and returns a new one
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
