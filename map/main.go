package main

import "fmt"

func main() {

	// Creating a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Declaring a map without values
	// var colors map[string]string

	// colors := make(map[string]string)

	// fmt.Println(colors)
	printMap(colors)

	// Deleting a key value pair in map
	// delete(colors, "white")

}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("Hex code for", key, "is", value)
	}
}
