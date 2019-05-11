package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	respPointer, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	respVal := *respPointer

	io.Copy(os.Stdout, respVal.Body)
}
