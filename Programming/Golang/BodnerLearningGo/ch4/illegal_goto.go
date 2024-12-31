package main

import "fmt"

func main() {
	a := 10
	goto skip // illegal, goto can not be used to skip over variable declarations.
	b := 20
skip:
	c := 30
	fmt.Println(a, b, c)
	if c > a {
		goto inner // illegal, goto can not be used to jump into inner blocks.
	}
	if a < b {
	inner:
		fmt.Println("a is less than b")
	}
}
