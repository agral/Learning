package main

import "fmt"

func confusingAppend() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	fmt.Printf("`x` is %v, `y (:=x[:2])` is %v\n", x, y)
	fmt.Printf("`cap(x)`=%v, `cap(y)`=%v\n", cap(x), cap(y))
	y = append(y, "z")
	fmt.Printf("After `y = append(y, \"z\")`:\nx=%v, y=%v\n\n", x, y)
}

func confusingSlices() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Printf("x=%v, `y:=x[:2]`=%v, `z:=x[2:]`=%v\n", x, y, z)
	fmt.Printf("`cap(x)`=%d, `cap(y)`=%d, `cap(z)`=%d\n", cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "z")
	fmt.Printf("x=%v, y=%v, z=%v\n", x, y, z)
}

func main() {
	confusingAppend()
	confusingSlices()
}
