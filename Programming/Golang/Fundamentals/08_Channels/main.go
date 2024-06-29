package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://amazon.com",
		"https://facebook.com",
		"https://golang.org",
		"https://google.com",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	fmt.Print(time.Now().Format("<15:04:05> "))
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up.")
	c <- link
}
