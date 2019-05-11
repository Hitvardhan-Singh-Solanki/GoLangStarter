package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStat(link, c)
	}

	for l := range c {
		go func(linkage string) {
			time.Sleep(2 * time.Second)
			checkStat(linkage, c)
		}(l)
	}
}

func checkStat(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- "500 Server error"
		return
	}
	fmt.Println(link, "is live")
	c <- link
}
