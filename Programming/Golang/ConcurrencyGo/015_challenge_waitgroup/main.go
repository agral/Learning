package main

import "fmt"

var msg string

func updateMessage(s string) {
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
	msg = "Hello, world!"
	updateMessage("Hello, universe!")
	printMessage()

	updateMessage("Hello, cosmos!")
	printMessage()

	updateMessage("Hello, world!")
	printMessage()
}
