package main

import "fmt"

// odd or even challange

// create a slice of ints, iterate through them and return even or odd

func main() {

	oneToTen := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range oneToTen {
		if num%2 == 0 {
			fmt.Printf("element at position %v is even\n", i)
		} else {
			fmt.Printf("element at position %v is odd\n", i)
		}
	}

}
