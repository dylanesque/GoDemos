package main

// Required to format strings, one of the standard library packages
// Strings must be in double quotes

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of deck

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"ace", "two", "three"}

	// The underscore takes the place of an index that we don't need to use in the loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// Receiver functions take in a type, and sets up methods to be performed on them
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// The first argument to this method is an example of type coersion in Go
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
