package main

import "fmt"

// Type - contactInfo
// email - The persons email address
// zipCode - The persons postal address
type contactInfo struct {
	email   string
	zipCode int
}

// Custom Type - person
// Fields:
// firstName
// lastName
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Creation Of Structs

	// Option #1 - Creation by field order
	p1 := person{
		"John",
		"Smith",
		contactInfo{
			"john@test.com",
			55555,
		},
	}

	// Option #2 - Define the fields during creation
	p2 := person{
		firstName: "Bob",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "bob@test.com",
			zipCode: 55555,
		},
	}

	// Option #3 - Instantiate before assigning properties
	var p3 person

	p3.firstName = "Alex"
	p3.lastName = "Anderson"
	p3.contactInfo.email = "alex@test.com"
	p3.contactInfo.zipCode = 55555

	p1.printDetails()
	p2.printDetails()
	p3.printDetails()

	p1.updateName("Sally")
	p1.printDetails()
}

func (p person) printDetails() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}