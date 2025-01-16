package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilon",
		"eta",
		"theta",
		"iota",
		"kappa",
		"lambda",
	}

	for num, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", num, word))
	}

	time.Sleep(500 * time.Millisecond)
	printSomething("That's the end of main().")
}
