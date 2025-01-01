package main

import "fmt"

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Printf("first: %d\n", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Printf("second: %d\n", val)
	}(a)
	a = 30
	fmt.Printf("exiting: %d\n", a)
	return a
}

func main() {
	deferExample()
}
