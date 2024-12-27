package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "Hello"
	fmt.Println(x)
	fmt.Println(y)

	const untyped10 = 10  // type can be omitted resulting in an *untyped constant*.
	var t int = untyped10 // then such untyped constant can be used to form concrete types.
	var ft float64 = untyped10 * 0.3
	print(fmt.Println(t, ft))

	//x = x + 1 // this will not compile.
}
