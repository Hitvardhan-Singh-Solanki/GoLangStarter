package main

func main() {
	cards := newDeck()
	cards.saveToFile("cards.log")
}
