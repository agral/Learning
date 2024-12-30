package main

import "fmt"

func simple_shadowing() {
	x := 10
	if x > 5 {
		fmt.Printf("`x`=%d, which is >5\n", x)
		x := 5
		fmt.Printf("Still inside the if block, `x`=%d\n", x)
	}
	fmt.Printf("Just outside the if block, `x`=%d\n", x)
}

func simple_shadowing_fixed() {
	x := 10
	// A fix for that is simple: use `=` instead of `:=`.
	if x > 5 {
		fmt.Printf("Once again, `x`=%d, which is >5\n", x)
		x = 5
		fmt.Printf("Still inside the if block, `x`=%d\n", x)
	}
	fmt.Printf("Just outside the if block, `x`=%d\n", x)
}

func multiple_assignment_shadowing() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Printf("\nStill inside the if block, `x`=%d, `y`=%d\n", x, y)
	}
	fmt.Printf("Just outside the if block, `x`=%d\n", x)
}

func shadowing_package_names() {
	x := 10
	fmt.Println(x)
	//fmt := "oops!"
	//fmt.Println(fmt)
}

func shadowing_true() {
	fmt.Println(true)
	true := "that's true!"
	fmt.Println(true)
}

func main() {
	simple_shadowing()
	simple_shadowing_fixed()
	multiple_assignment_shadowing()
	// shadowing_package_names()
	shadowing_true()
}
