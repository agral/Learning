package main

import "fmt"

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	// Note: this goroutine will probably not return when main() returns,
	// and as a result this line won't likely be printed.
	go printSomething("Hello, amazing world of Golang's concurrent programming (with goroutines)!")

	printSomething("Hello, nice world of Golang's sequential programming!")
}
