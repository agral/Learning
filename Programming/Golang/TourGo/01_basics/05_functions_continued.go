package main

import "fmt"

// it used to be (in 04_functions.go):
// func add(x int, y int) int {
// when two or more consecutive params share a type, they can be joined:
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
