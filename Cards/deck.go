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
// Which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues { // The underscore is to tell go we don't need the index
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// Return two values
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Converts byte to string
	// Receiver function, like a method
	// This allows dot notation (deck.toString())
	// The string package import above has a join method

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// Savetofile saves a desck file on the hard drive
	// The writefile takes 3 arguements, the last of which is permissions
	// 0666 is allowall
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// Creates a deck from a file

	bs, err := ioutil.ReadFile(filename)
	// ReadFile returns a byte slice
	if err != nil {
		// Option #1 - log the error and call newDeck()
		// Option #2 - Log the error and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] // Swap postions
	}
}
