package main

func main() {
	cards := newDeckFromFile("cards.log")
	cards.printEl()
	// cards.saveToFile("cards.log")
}
