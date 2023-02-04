package main

// embed struct inside another -contact in person

type contactInfo struct {
	email   string
	zipCode int
}

// Define type, person
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// Create a value of type, person
func main() {
	Isaac := person{
		firstName: "Isaac",
		lastName:  "Party",
		contact: contactInfo{
			email:   "Isaacatsomethingdotcom",
			zipCode: 98112,
		},
	}
}
