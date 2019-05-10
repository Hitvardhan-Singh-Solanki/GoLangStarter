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

func (p *person) updateName(newFName string) {
	p.firstName = newFName
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
	p.updateName("Jane")
	fmt.Println("********************")
	p.getDetails()
}
