package main

import "fmt"

func main() {

	// Variables

	var card string = "Ace of spades"

	// short assigment type is infered from the initalised value
	card2 := "Nine of Hearts"

	// types can be explicit or inferred if the variable is initalised with a value

	// un-initalised variables
	//var card3 string
	//card3 = "card3"

	// bulk variable declaration
	var (
		card3 string = "three"
		card4 string = "four"
		card5 string = "five"
	)

	fmt.Println(card, card2, card3, card4, card5)
}
