package main

import (
	"fmt"
	"math/rand"
)

func ch4_ex1() []int {
	// Write a for loop that puts 100 random numbers between 0 and 100 into an int slice.
	numbers := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(100+1))
	}
	fmt.Printf("len=%d, capacity=%d, first ten numbers: %v\n", len(numbers), cap(numbers), numbers[:10])
	return numbers
}

func ch4_ex2(n []int) {
	// Loop over the slice you created in exercise 1. For each value in the slice, apply the
	// following rules:
	// a. If the value is divisible by 2, print “Two!”
	// b. If the value is divisible by 3, print “Three!”
	// c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
	// d. Otherwise, print “Never mind”.
	for _, v := range n {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}

func ch4_ex3() {
	//Start a new program. In main, declare an int variable called total. Write a for
	// loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
	// The body of the for loop should be as follows:
	// total := total + i
	// fmt.Println(total)
	// After the for loop, print out the value of total. What is printed out? What is the
	// likely bug in this code?

	// the bug is that total is shadowed. It should be `total = total + i` instead.
	total := 0
	for i := 0; i < 10; i++ {
		total := total + i
		//total = total + i
		fmt.Println(total)
	}
	fmt.Printf("After loop, `total`=%d\n", total)
}

func main() {
	n := ch4_ex1()
	ch4_ex2(n)
	ch4_ex3()
}
