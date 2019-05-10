package main

import (
	"fmt"
	"strconv"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) getDetails() {
	fmt.Println("Person Name: " + p.firstName + " " + p.lastName)
	fmt.Println("Email: " + p.contactInfo.email)
	fmt.Println("ZipCode: " + strconv.Itoa(p.contactInfo.zip))
}

// passed by reference
func (p *person) updateName(newFName string) {
	// This is pointer to the memory address
	// * will return the value on the memory address
	(*p).firstName = newFName
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Doe",
		contactInfo: contactInfo{
			zip:   12312,
			email: "test@test.com",
		},
	}
	p.getDetails()
	// This will give us the exact memory address of the
	// variable which holds the value
	pPointer := &p
	pPointer.updateName("Jane")
	fmt.Println("++++++++++++++++++++++++++++++++")
	p.getDetails()
}
