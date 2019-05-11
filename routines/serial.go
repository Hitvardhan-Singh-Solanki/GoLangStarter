// Serial Approach
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

	for _, link := range links {
		checkStat(link)
	}
}

func checkStat(link string) {
	fmt.Println("Checking stat")
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is live")
}
