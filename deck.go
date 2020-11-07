package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heards", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// With a Receiver (d deck) any variable of type "deck" gets access to the methods define on the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Use: hand, remainingCards := deal(cards, 5)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// We'll be saving the cards deck to filesystem and for that we need to have a []byte but first we need a list of strings to the conversion
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save cards deck to filesystem
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
