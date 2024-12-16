package main

import "fmt"

func main() {
	int_val := 42
	fmt.Printf("int_val is of type %T\n", int_val)

	double_val := 42.0
	fmt.Printf("double_val is of type %T\n", double_val)

	complex_val := 42.0 + 42i
	fmt.Printf("complex_val is of type %T\n", complex_val)

	bool_val := false
	fmt.Printf("bool_val is of type %T\n", bool_val)

	byte_val := byte(42)
	fmt.Printf("byte_val is of type %T\n", byte_val)
	fmt.Println("  (note: `byte` is just an alias for `uint8`)")

	string_val := "forty-two"
	fmt.Printf("string_val is of type %T\n", string_val)

	rune_val := rune(42)
	fmt.Printf("rune_val is of type %T\n", rune_val)
	fmt.Println("  (note: `rune` is just an alias for `int32`)")
}
