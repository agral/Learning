package main

import "fmt"

// the `var` statement declares a list of variables.
// the type is last, same as with function argument lists.
// These are package level variables:
var c, python, java bool

func main() {
	// vars can also be declared at function level.
	var i int
	fmt.Println(i, c, python, java)
}
