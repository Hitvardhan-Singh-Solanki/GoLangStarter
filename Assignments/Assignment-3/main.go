package main

import (
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	check(err)
	io.Copy(os.Stdout, f)
}
