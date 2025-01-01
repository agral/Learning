package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
	fmt.Printf("in failedUpdate(), assigned g=%v\n", g)
}

func goodUpdate(g *int) {
	*g = 10
	fmt.Printf("in goodUpdate(), assigned g=%v\n", *g)
}

func main() {
	var f *int
	// f is initially nil.
	fmt.Printf("before failedUpdate(), f=%v\n", f)

	// calling this results in a COPY of the value of f (nil) into the parameter g;
	// this does not result in any meaningful update to f.
	// At no point in time does f change from `nil` to anything else.
	failedUpdate(f)
	fmt.Printf("after failedUpdate(), f=%v\n", f)

	x := 1
	failedUpdate(&x)
	fmt.Printf("after different failedUpdate(), x=%v\n", x)

	goodUpdate(&x)
	fmt.Printf("after goodUpdate(), x=%v\n", x)
}
