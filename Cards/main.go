package main

import "fmt"

func main() {

	// array - (fixed length)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// slice - (dynamic array / list)
	var StringNumbers = []string{"1", "2"}
	fmt.Println(StringNumbers)

	// append function
	StringNumbers = append(StringNumbers, "3", "4", "5")

	// numbers = append(numbers, 3, 4, 5) - cannot append to arrays (fixed size)
	fmt.Println(StringNumbers)
	fmt.Println(StringNumbers[0:1]) // gets the elements from posistion 0 - 1 but not including
	fmt.Println(StringNumbers[1:])  // arguments omitted default to either (start) or (length)

}
