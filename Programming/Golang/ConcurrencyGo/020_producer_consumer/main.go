package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	// calling the program with `-race` switch reports the data race problem.
	wg.Wait()
	fmt.Println(msg)
}
