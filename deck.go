package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// Generating a [Seed] or [Source] on every new random generate round
	source := rand.NewSource(time.Now().UnixNano()) //int64
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// interchange values
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
