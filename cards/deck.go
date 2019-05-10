package main

import "fmt"

// Create a new type of deck := slice of string
type deck []string

func (d deck) printEl() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
