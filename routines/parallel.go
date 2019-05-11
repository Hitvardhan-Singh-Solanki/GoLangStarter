package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
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
	c <- "OK"
}
