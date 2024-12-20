package main

import "fmt"

// creating an interface
type bot interface {
	getGreeting() string // function signiture that must be satisfied
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	englishBot, spanishBot := englishBot{}, spanishBot{}

	printGreeting(englishBot)
	printGreeting(spanishBot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi there." // satisfied the interfaces function signatures, becomes a "bot" interface
}

func (spanishBot) getGreeting() string {
	return "Hola." // satisfied the interfaces function signatures, becomes a "bot" interface
}
