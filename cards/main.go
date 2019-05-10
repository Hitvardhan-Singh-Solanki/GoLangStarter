package main

func main() {
	cards := newDeck()
	hand, remainingDeck := cards.deal(5)
	hand.printEl()
	remainingDeck.printEl()
}
