package main

import (
	"fmt"
	"slices"
)

func main() {
	// Array of three default-initialized elements:
	var z [3]int
	fmt.Printf("Array `var z [3]int`: %v\n", z)

	// Array of three specified values:
	var y = [3]int{10, 20, 30}
	fmt.Printf("Array with initial values provided `var y [3]int{10, 20, 30}`: %v\n", y)

	// Sparse array (an array where most elements are set to zero) are defined
	// by specifying the indices of nonzero elements:
	var x = [12]int{1, 5: 4, 6, 10: 100, 15 /*,13*/} // Note that it is a compile error to go out of bounds, nice.
	fmt.Printf("Sparse array: %v\n", x)

	// '...' can be used in place of length. The length of the array is then inferred from the arguments list.
	var w = [...]int{10, 20, 30}
	fmt.Printf("Array with initial values provided `var w [...]int{10, 20, 30}`: %v\n", w)
	// It works for sparse arrays, too!
	// but note that the extra value (13) will now not be catched at compile time...
	var v = [...]int{1, 5: 4, 6, 10: 100, 15, 13}
	fmt.Printf("Sparse array, length inferred: %v\n", v)

	fmt.Println("Arrays are considered equal if they are the same length and contain equal values")
	fmt.Printf("w == y: %v\n", w == y)

	// Multi-dimensional arrays are actually arrays of one-dimensional arrays. Like in C (sort of), unlike in Julia.
	var u [2][5]int
	fmt.Printf("`var u [2][5]int` has length %d, and each element has length %d\n", len(u), len(u[0]))

	// An out of bounds read using a constant value is detected at compile time:
	//_ = z[5]
	// Negative indices are disallowed in a similar way (unlike in Python):
	//_ = z[-1]
	// An out of bounds read using an index from a variable is not detected at compile time.
	// The following lines would result in a runtime panic:
	//var idx int = 5
	//_ = z[idx]

	// The size of the array is part of the array type. So e.g. [3]int is different type than [4]int.
	// This is why slices are much more common.

	// Using [...] makes ar array; using [] makes a slice.

	var a = []int{10, 20, 30}
	fmt.Printf("\nSlice `var a []int{10, 20, 30}`: %v\n", a)
	var b = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Printf("Sparse slice `var b = []int{1, 5: 4, 6, 10: 100, 15}`: %v\n", b)

	// Slices can be declared without a literal (it would be analogous to a declaration -
	// as opposed to a definition - from C language). Such slice is zero-initialized to nil.
	var c []int
	fmt.Printf("Zero-initialized slice: `var c []int`: %v\n", c)
	fmt.Printf("`c == nil`: %v\n", c == nil)

	// Unlike arrays, slices can not be directly compared to each other.
	// Slices can only be directly compared to nil. So the following line is illegal Go code:
	//fmt.Printf("`a == b`: %v\n", a == b)

	// Since Go 1.21 (released in 2024-08; almost 1.5 years ago), there's slices.Equal function
	// to compare comparable slices; and slices.EqualFunc can be used to compare almost anything.
	fmt.Printf("`slices.Equal(a, b)`: %v\n", slices.Equal(a, b))
	fmt.Printf("`slices.Equal(a, a)`: %v\n", slices.Equal(a, a))
	// this won't compile; int and string are not comparable:
	//var d = []string{"Mary", "had", "a", "little", "lamb"}
	//fmt.Printf("`slices.Equal(a, d)`: %v\n", slices.Equal(a, d))
	// But such comparison can be done with slices.EqualFunc.

	// The built-in append() function is used to grow slices:
	a = append(a, 40)
	fmt.Printf("\nafter `a=append(a, 40)`, slice a is now: %v\n", a)
	// append() can add many values at once:
	a = append(a, 50, 60, 70)
	fmt.Printf("after `a=append(a, 50, 60, 70)`, slice a is now: %v\n", a)
	// One slice can be appended onto the other by using the ... operator
	// to expand the source slice into individual values.
	d := []int{80, 90}
	a = append(a, d...)
	fmt.Printf("after `a=append(a, d...)`, slice a is now: %v\n", a)
}
