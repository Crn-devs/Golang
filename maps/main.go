package main

import (
	"fmt"
)

func main() {

	// map literal syntax
	// 1 name := map[Type1]Type2 {type1: type2,}
	colors := map[string]string{
		"red":  "#FF0000",
		"blue": "#5555",
	}
	// 2 var name = map[string]string // Declares a map but leaves it uninitialized (i.e., it’s nil).
	var colors2 map[string]string
	// 3 colors := make(map[string]string)
	colors2 = make(map[string]string)

	// manipulating maps

	// adding values name[key] = value
	colors2["White"] = "#fff"
	colors2["Black"] = "#000"

	// deleting values
	delete(colors2, "White")

	// iterating over maps

	fmt.Println(colors)
	fmt.Println(colors2)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hex for color", color, "is", hex)
	}
}
