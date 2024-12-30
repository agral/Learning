package main

import "fmt"

func for_range_on_strings_uses_runes_not_bytes() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		// Iterating with a for-range loop over a string works on runes, not bytes:
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func for_range_works_on_copy() {
	odds := []int{1, 3, 5, 7, 9, 11, 13}
	for _, v := range odds {
		v *= 2
	}
	fmt.Println(odds)
}

func main() {
	for_range_on_strings_uses_runes_not_bytes()
	for_range_works_on_copy()
}
