package main

import "fmt"

func main() {

	card := newCard()

	fmt.Println(card)
}

// Syntax
// defines as a function || define the name || defines return type || function body
func newCard() string {
	return "this is a new card"
}
func newCardDeck() int {
	return 52
}

// functions can be called from other files if they are in the same package
