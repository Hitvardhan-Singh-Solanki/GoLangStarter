package main

import (
	"fmt"
	"math"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		height: 2.0,
		base:   3.0,
	}
	s := square{
		sideLength: 4.0,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func printArea(s shape) {
	area := s.getArea()
	st := fmt.Sprintf("%f", area)
	fmt.Println(st)
}
