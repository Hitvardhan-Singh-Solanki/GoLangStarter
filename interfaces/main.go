package main

import "fmt"

type bot interface {
	getGreet() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreet())
}

func (englishBot) getGreet() string {
	// Some custom logic
	return "Hi there!"
}
func (spanishBot) getGreet() string {
	// Some custom logic
	return "Hola"
}
