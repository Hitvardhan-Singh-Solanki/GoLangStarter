package main

import "fmt"

type colorsMap map[string]string

func (m colorsMap) printMap() {
	for color, hex := range m {
		fmt.Println("key: " + color + "\nhex code: " + hex)
	}
}

func main() {
	colors := colorsMap{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
		"black": "#000000",
	}

	colors.printMap()

	fmt.Println("*********************")

	delete(colors, "red")

	colors.printMap()
}
