package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("The loop has completed normally")
done:
	fmt.Println("This gets printed no matter how the loop was left")
	fmt.Println(a)
}
