package main

import "fmt"

// creating structs type person
type person struct {
	firstname string
	lastname  string
}

func main() {
	fmt.Println("Hello World")

	// ordered declaration
	// each value is assigned in order to the structs values
	// can cause issues if the structs fields are changed around
	alex := person{"alex", "anderson"}

	// Explicit initalisation / keyed initialization.
	// each value is assigned by its name
	// this approach allows un-ordered and ordered declaration
	// alexis := person{firstname: "alexis", lastname: "anderson"}
	alexis := person{lastname: "anderson", firstname: "alexis"}

	// Zero value initalisation
	// charlie is initalised with its types zero value for each field
	// firstname: empty string etc
	var charlie person
	charlie.firstname = "charlie"

	fmt.Printf("%+v", charlie)
	fmt.Printf("%v %v %v %v", alex.firstname, alex.lastname, alexis.firstname, alexis.lastname)

}
