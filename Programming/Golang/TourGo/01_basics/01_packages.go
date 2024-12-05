package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intn(num) returns a random int in range [0 .. n-1]; or [0, n)
	fmt.Println("My favorite number is", rand.Intn(10))
}
