package main

import "fmt"

// custom types
// custom types can be created using type keyword
// the name of the type and what data it will hold
type deck []string

// reciever functions
// these are like methods on classes
// they allow for the function to be called on any type of deck
// this can be seen in the 7th line in main.go
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
