package main

import "fmt"

func main() {

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// loops are created using for and the range keywords
	// when ranging over a slice, array or map 2 values are returned
	// for each iteration First(index) Second(copy of the element at the index)
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	// for _, number := range numbers {
	// 	fmt.Println(number)
	// }
}
