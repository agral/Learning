package main

import (
	"fmt"
	"math/rand"
)

func basic_if() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Printf("%d is too low!\n", n)
	} else if n > 5 {
		fmt.Printf("%d is too big!\n", n)
	} else {
		fmt.Printf("%d is just right.\n", n)
	}
}

func variable_scoped_to_if_statement() {
	if n := rand.Intn(10); n == 0 {
		fmt.Printf("%d is too low!\n", n)
	} else if n > 5 {
		fmt.Printf("%d is too big!\n", n)
	} else {
		fmt.Printf("%d is just right.\n", n)
	}
	// Note: n is no longer available here:
	// fmt.Printf("`n`=%v\n", n)
}

func main() {
	basic_if()
	variable_scoped_to_if_statement()
}
