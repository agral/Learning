package main

import "fmt"

// If an initializer is present, then the type can be omitted.
// This declaration includes the type (int):
var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
