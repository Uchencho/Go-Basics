package main

func main() {
	// cards := newDeck()

	// hand, remaingCards := deal(cards, 5)

	// hand.print()
	// remaingCards.print()

	// Type conversion

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_car")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
