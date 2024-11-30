package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// custom types
// custom types can be created using type keyword
// the name of the type and what data it will hold
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spade", "King", "Queen", "Hearts"}
	cardValues :=
		[]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" "+"of"+" "+cardSuit)
		}
	}
	return cards
}

// reciever functions
// these are like methods on classes
// they allow for the function to be called on any type that is of deck
// this can be seen in the 7th line in main.go
// d is similar to using 'this' it references the object that the function is being called on
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// below is an example of returning multiple values from a function

func deal(d deck, handSize int) (hand deck, remainingCards deck) {
	return d[:handSize], d[handSize:]
}

// reciever func toString
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// saving to HD using os.writeFile
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// convert the byte slice back to a string
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
