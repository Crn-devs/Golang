package main

import "fmt"

type contactInfo struct {
	email    string
	postcode string
}

// creating structs type person
type person struct {
	firstname string
	lastname  string
	// embedded struct
	contact contactInfo
}

func main() {
	fmt.Println("Hello World")

	// ordered declaration
	// each value is assigned in order to the structs values
	// can cause issues if the structs fields are changed around
	// alex := person{"alex", "anderson"} // after updating the struct to include new fields this declaration no longer works

	// Explicit initalisation / keyed initialization.
	// each value is assigned by its name
	// this approach allows un-ordered and ordered declaration
	// alexis := person{firstname: "alexis", lastname: "anderson"}
	// alexis := person{lastname: "anderson", firstname: "alexis"}

	// Zero value initalisation
	// charlie is initalised with its types zero value for each field
	// firstname: empty string etc
	var charlie person
	charlie.firstname = "charlie"
	fmt.Printf("%+v", charlie)

	alex := person{
		firstname: "Alex",
		lastname:  "beggins",
		contact: contactInfo{
			email:    "alexbeggins@gmail.com",
			postcode: "SH52 YHG",
		}}
	fmt.Printf("%+v", alex)
	fmt.Println(alex.contact)
}
