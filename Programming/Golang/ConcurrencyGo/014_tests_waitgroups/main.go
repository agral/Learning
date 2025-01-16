package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
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

	wg.Add(len(words))

	for num, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", num, word), &wg)
	}

	wg.Wait()

	fmt.Println("That's the end of main().")
}
