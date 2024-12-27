package main

import (
	"fmt"
	"math"
)

func ex01() {
	// Write a program that declares an integer variable called i with the value 20.
	// Assign i to a floating-point variable named f. Print out i and f.

	// This can not be done; it is not legal to use an integer variable to form a float.
	// It would have been possible if i was an untyped constant; but not with a concrete integer variable.
	//var i int = 20
	//var f float64 = i;
}

func ex02() {
	// Write a program that declares a constant called value that can be assigned to
	// both an integer and a floating-point variable. Assign it to an integer called i and a
	// floating-point variable called f. Print out i and f.
	const meters_in_kilometer = 1000
	var i int = 2 * meters_in_kilometer
	var f float64 = 3.141592 * meters_in_kilometer
	fmt.Printf("Two kilometers is equal to %d meters.\n", i)
	fmt.Printf("Pi kilometers is roughly %.2f meters.\n", f)
}

func ex03() {
	// Write a program with three variables, one named b of type byte, one named
	// smallI of type int32, and one named bigI of type uint64. Assign each variable
	// the maximum legal value for its type; then add 1 to each variable. Print out their
	// values.
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64
	fmt.Println("Before adding one:")
	fmt.Printf("byte: %d\nint32: %d\nuint64: %d\n", b, smallI, bigI)

	b += 1
	smallI += 1
	bigI += 1
	fmt.Println("\nAfter adding one:")
	fmt.Printf("byte: %d\nint32: %d\nuint64: %d\n", b, smallI, bigI)
}

func main() {
	ex01()
	ex02()
	ex03()
}
