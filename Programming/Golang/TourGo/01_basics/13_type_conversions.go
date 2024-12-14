package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Note: assignments between instances of different types require an explicit conversion.
	// So e.g. the following with the conversion removed won't compile:
	// var wrong float64 = math.Sqrt(x*x + y*y)
	// This is unlike C++ with its standard (implicit) conversions and promotions.
}
