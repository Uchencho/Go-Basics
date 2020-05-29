package main

import "fmt"

// Stephen Grider
// Create a struct within a struct
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Stephen",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94990,
		},
	}

	// jimPointer := &jim             // This returns the mem address of the var jim
	// jimPointer.updateName("Jimmy") // This calls updateName with a pointer add
	jim.updateName("Patrick") // You can skip the
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// * in front of a type is different from star in front of a pointer
	// * in front of a type tells go that the type is a pointer to person (in this case)
	// It could have been pointer to string
	// *type == ponter to the type
	// * operator requests for the value of a pointer at a specific memory add (when used on a var)

	// Necessary to reference pointer in the receiver function becos...
	// ...we are dealing with value type not reference type
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
