package main

import "fmt"

// embed struct inside another -contact in person

type contactInfo struct {
	email   string
	zipCode int
}

// Define type, person
type person struct {
	firstName string
	lastName  string
	contactInfo
}

// Create a value of type, person
func main() {
	Isaac := person{
		firstName: "Isaac",
		lastName:  "Clarke",
		contactInfo: contactInfo{
			email:   "Isaacatsomethingdotcom",
			zipCode: 98112,
		},
	}

	// "&" gets the memory address
	// IsaacPointer points to the address
	// in memory of the variable, Isaac
	// | 8934 | {firstName: "Isaac"}
	IsaacPointer := &Isaac
	IsaacPointer.updateName("Bunny")
	Isaac.print()
}

// update name
// this code won't update person, Isaac, bc of pass
// by reference
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// update name
// "*" gets the value in the memory address
// *pointerToPerson unpacks the value stored at the address the pointer points at
// *person - a star before a type; this is a description of a type; * is not an operator in this context
// *person says pointerToPerson is a variable containing a value stored in memory
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// print instance of person
func (p person) print() {
	fmt.Println("%+v", p)
}
