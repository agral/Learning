package main

import "fmt"

// Return values can be named, as below:
func split(sum int) (x, y int) {
	// with named return values, a return statement without arguments
	// returns all the named return values.
	// This is called a "naked return".
	// This harms readability but can be encountered in some codebases.
	// This is encouraged for short functions by the golang authors.
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
