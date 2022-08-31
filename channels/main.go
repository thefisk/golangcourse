package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Below is a standard serial loop (one at a time)
	// for _, link := range links {
	// 	checkLink(link)
	// }

	// Can be improved by  putting 'go' keyword
	// before the function call to execute it in it's
	// own go routine, perfect for blocking functions.
	// Go routines communicate to spin up when they
	// each get blocked

	// To do this, we also need to make a channel and
	// make sure the receiving function accepts a
	// channel as an argument

	c := make(chan string)

	for _, link := range links {
		go checkLinkOverChannel(link, c)
	}

	// Take output from channel and print it
	// Needs to be looped because main routine
	// Will exit after first response otherwise
	// Below will loop 5 times (len of slice)
	// for i := 0; i < len(links); i++ {
	// blank for loop is infinite
	for l := range c {
		// The loop will wait at each iteration
		// For a response
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkOverChannel(link, c)
		}(l)
	}
}

// Original function
func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

// Function that communicates over a channel
func checkLinkOverChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Send text back through channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}