package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			zip:   12312,
			email: "test@test.com",
		},
	}
	// var p person
	p.firstName = "Jane"
	p.lastName = "Doe"

	fmt.Println(p.firstName + " " + p.lastName)
	fmt.Println(p)
}
