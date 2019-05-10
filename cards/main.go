package main

import "fmt"

func main() {
	card := newCardVal()
	fmt.Println(card)
	printState()
}

/**
Need to tell the return value
much like solidity
*/
func newCardVal() string {
	return "Five of diamonds"
}
