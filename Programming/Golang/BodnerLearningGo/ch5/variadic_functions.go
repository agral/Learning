package main

import "fmt"

// This is a function with variadic parameters.
func addTo(base int, values ...int) []int {
	ans := make([]int, 0, len(values))
	for _, value := range values {
		ans = append(ans, base+value)
	}
	return ans
}

func main() {
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 3))
	fmt.Println(addTo(2, 0, 2, 4, 6, 8))
	v := []int{1, 2, 3}
	fmt.Println(addTo(3, v...))                 // variadic part of parameters can also be supplied
	fmt.Println(addTo(3, []int{1, 2, 3, 4}...)) // as a slice with `...` at the end.
}
