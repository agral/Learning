package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	// when iterating over map's key-value pairs with for-range,
	// the iteration order varies - keys are randomly accessed.
	// So we're likely to see differently ordered outputs in each loop:
	for i := 0; i < 3; i++ {
		fmt.Printf("Loop #%d\n", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	// (this is a security feature, against Hash DoS and other issues)
}