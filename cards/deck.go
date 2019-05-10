package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))

}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) printEl() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {

	s := time.Now().UnixNano()

	source := rand.NewSource(s)

	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
