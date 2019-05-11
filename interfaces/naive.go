// naive approach to write a program
// Will not run has some errors
// implementating the same logic with interfaces
package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreet() string {
	// Some custom logic
	return "Hi there!"
}
func (spanishBot) getGreet() string {
	// Some custom logic
	return "Hola"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreet())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreet())
}
