package main

import (
	"fmt"

	"github.com/agral/celeritas"
)

func main() {
	result := celeritas.TestFunc(1, 2)
	fmt.Println(result)
}
