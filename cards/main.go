package main

import (
	"fmt"
)

func main() {
	// cards := newDeck()
	// // cards.print()
	// hand, remaining := deal(cards, 5)
	// hand.print()
	// remaining.print()

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
	fmt.Println("Shuffling cards...")
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }