package main

import "fmt"

func main() {
	fmt.Printf("Slice:\t\t\tlen:\t\tcap:\n")
	var x []int
	fmt.Printf("%v\t\t\t%d\t\t%d\n", x, len(x), cap(x))
	x = append(x, 10)
	fmt.Printf("%v\t\t\t%d\t\t%d\n", x, len(x), cap(x))
	x = append(x, 20)
	fmt.Printf("%v\t\t\t%d\t\t%d\n", x, len(x), cap(x))
	x = append(x, 30)
	fmt.Printf("%v\t\t%d\t\t%d\n", x, len(x), cap(x))
	x = append(x, 40)
	fmt.Printf("%v\t\t%d\t\t%d\n", x, len(x), cap(x))
	x = append(x, 50)
	fmt.Printf("%v\t%d\t\t%d\n", x, len(x), cap(x))
}
