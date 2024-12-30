package main

import "fmt"

func basics() {
	type person struct {
		name string
		age  int
		pet  string
	}

	johndoe := person{}
	alice := person{
		"Alice",
		40,
		"cat",
	}
	bob := person{
		age:  46,
		name: "Bob",
	}

	fmt.Printf("`johndoe` (zero-initialized): %v\n", johndoe) // empty string, 0, empty string
	fmt.Printf("`alice` (custom): %v\n", alice)               // "Alice", 40, "cat"
	fmt.Printf("`bob` (custom): %v\n", bob)                   // "Bob", 40, empty string
}

func main() {
	basics()
}
