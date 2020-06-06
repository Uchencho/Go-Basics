package main

import "fmt"

type bot interface {
	// This allows two types access the same func
	// As long as both types have a function called
	// getGreeting which returns a type string
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	// Create an english and a spanish type bot
	eb := englishBot{}
	sb := spanishBot{}

	// Call the printGreeting function on the bot interface
	// Passing in the structs, englishBot and spanishBot
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	// printGreeting function that takes in an interface type
	// called bot
	// uses that to call the getGreeting function as a receiver

	fmt.Println(b.getGreeting())
}

// func printGreeting2(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string {
	// A receiver function that implements a unique logic
	// and returns a string
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	// A receiver function that implements a unique logic
	// and returns a string
	return "Hola"
}
