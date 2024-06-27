package main

import (
	"fmt"
)

func main() {
	for num := range 10 + 1 { // from 0 to 10 inclusively
		if num%2 == 0 {
			fmt.Printf("%d: even\n", num)
		} else {
			fmt.Printf("%d: odd\n", num)
		}
	}
}
