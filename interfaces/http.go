package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	respPointer, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}

	respVal := *respPointer

	lw := logWriter{}

	io.Copy(lw, respVal.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
