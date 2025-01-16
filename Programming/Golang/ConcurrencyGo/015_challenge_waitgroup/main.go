package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	// Challenge: modify this code so that the calls to updateMessage()
	// are executed as goroutines. Implement waitgroups so that the program
	// runs properly and predictably, i.e. the output is always the same
	// between runs (no race conditions).
	// Then write a test for all three functions.
	var wg sync.WaitGroup
	msg = "Hello, world!"
	wg.Add(1)
	go updateMessage("Hello, universe!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!", &wg)
	wg.Wait()
	printMessage()
}
