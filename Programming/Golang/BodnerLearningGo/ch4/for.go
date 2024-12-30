package main

import "fmt"

func completeForStatement() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	// Note: must use `:=`, var is not legal!
	// for var i = 0; i < 4; i++ {
	// 	fmt.Println(i)
	// }
}

func conditionOnlyForStatement() {
	// It is allowed to leave out either initialization and/or the increment.
	// If both are omitted, the semicolons are to be omitted too:
	// This looks very similar to a while loop.
	i := 10
	for i < 14 {
		fmt.Println(i)
		i++
	}
}

func infiniteForStatement() {
	counter := 0 // not really necessary, but I want this loop to end after four iterations too.
	for {
		fmt.Println(20 + counter)
		counter++
		if counter >= 4 {
			break
		}
	}
}

func forRangeStatement() {
	evenValues := []int{2, 4, 6, 8, 10, 12}
	for i, val := range evenValues {
		fmt.Println(i, val)
	}
}

func main() {
	completeForStatement()
	conditionOnlyForStatement()
	infiniteForStatement()
	forRangeStatement()
}
