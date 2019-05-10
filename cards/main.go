package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.printEl()
	// cards.saveToFile("cards.log")
}
