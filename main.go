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
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// print instance of person
func (p person) print() {
	fmt.Println("%+v", p)
}
