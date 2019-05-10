package main

func main() {
	cards := deck{newCardVal(), newCardVal()}
	cards = append(cards, "Six of Spades")
	cards.printEl()
}

/**
Need to tell the return value
much like solidity
*/
func newCardVal() string {
	return "Five of diamonds"
}
