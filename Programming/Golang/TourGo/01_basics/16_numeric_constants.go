package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int           { return 10*x + 1 }
func needFloat(x float64) float64 { return 0.1 * x }

func main() {
	// Note that Big can not be converted to int - it would be a compile-time error
	// (Golang catches that this constant overflows an int)
	//fmt.Println(needInt(Big))

	// But Small, derived directly from Big, can be used to instantiate an int:
	fmt.Println(needInt(Small))

	fmt.Println(needFloat(Small))

	// Note that Big can easily be used as a floating point value:
	fmt.Println(needFloat(Big))
}
