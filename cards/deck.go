package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck := slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, num := range cardValues {
			cards = append(cards, num+" of "+suit)
		}
	}

	return cards

}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",\n")
}

func (d deck) printEl() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
