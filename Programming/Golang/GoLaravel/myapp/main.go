package main

import (
	"fmt"

	"github.com/agral/celeritas"
)

func main() {
	result := celeritas.AddTwoInts(1, 2)
	result = celeritas.MulTwoInts(4, 5)
	fmt.Println(result)
}
