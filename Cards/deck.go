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

// Create a new type 'inheriting' from string, called deck
type deck []string

func newDeck() deck {
	// Creates a new deck of cards

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
	// Receiver type function
	// It can be called as a method, on a deck
	// i.e deck.print()

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// Function that takes in a handsize and a deck
	// and returns a two decks split at the handsize

	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Converts deck type that was split on a comma to
	// A full string, like a full sentence
	// Receiver function, like a method
	// This allows dot notation (deck.toString())
	// The string package import above has a join method

	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// Receiver function that saves to file
	// Savetofile saves a deck file on the hard drive
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
	// Shuffle the positions of a deck
	// Generates a resource so that each shuffle is different
	// Like NOT setting a random seed in python

	// Create a new resource based on the current time &
	// generate an integer with UnicNano
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// Loop through the index of each deck
		// Generate an integer between o and the length of deck
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] // Swap postions
	}
}
