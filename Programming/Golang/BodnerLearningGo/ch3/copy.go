package main

import "fmt"

func try_copy() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num_copied := copy(y, x) // copy(destination, source)
	fmt.Printf("after copy, y=%v. Copied %d items.\n", y, num_copied)

	just_two := make([]int, 2)
	num_copied = copy(just_two, x) // copy(destination, source)
	fmt.Printf("after copy, just_two=%v. Copied %d items.\n", just_two, num_copied)

	from_the_middle := make([]int, 2)
	// it's OK to not assign the number of copied elements to a variable if it's not intended to be used.
	copy(from_the_middle, x[1:])
	fmt.Printf("after copy, from_the_middle=%v.\n", from_the_middle)

	// copy() can be used for copying between two slices that cover overlapping sections of a parent slice:
	num_copied = copy(x[:3], x[1:])
	fmt.Printf("after copy, x=%v. Copied %d items.\n", x, num_copied)
}

func try_panicking_slice_to_array_conversion() {
	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice) // array created by using type conversion.

	// Note: using [...] is disallowed (compile time error) with conversions. Have to specify the actual size.
	//xArray := [...]int(xSlice)

	// The size of array can be smaller than the size of slice, but can not be bigger.
	// This is not detected at compile time, and results in a runtime panic.
	okArray := [3]int(xSlice)
	//panicArray := [5]int(xSlice)

	fmt.Println(xSlice, xArray, okArray)
}

func main() {
	try_copy()
	try_panicking_slice_to_array_conversion()
}
