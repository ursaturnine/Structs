package main

import "fmt"

// Define type, person
type person struct {
	firstName string
	lastName  string
}

// Create a value of type, person
func main() {
	Isaac := person{firstName: "Isaac", lastName: "Clarke"}
	fmt.Println(Isaac)
}
