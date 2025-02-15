package main

import (
	"fmt"
	"maps"
)

func try_nil_maps() {
	// This is a nil map, with keys of type string and values of type int. It has a length of 0.
	var nilMap map[string]int
	fmt.Printf("%v (len: %d)\n", nilMap, len(nilMap))

	// Attempting to read a value from any key in a nil map is OK, returns a zero value for the map's value type.
	fmt.Printf("`nilMap[\"foo\"]`=%v\n", nilMap["foo"])

	// But attempting to write to a nil map causes a panic.
	nilMap["foo"] = 42
}

func try_map_literals() {
	// This is an empty map, with string keys and int values. It has a length of 0.
	// Both reading and writing to an empty map is legal.
	emptyMap := map[string]int{}
	fmt.Printf("%v (len %d)\n", emptyMap, len(emptyMap))

	fmt.Printf("By default, `emptyMap[\"foo\"]`=%v\n", emptyMap["foo"])
	emptyMap["foo"] = 42
	fmt.Printf("After assigning value of 42, `emptyMap[\"foo\"]`=%v\n", emptyMap["foo"])

	// This is a nonempty map literal, with string keys and slices of string for values.
	animalsMap := map[string][]string{
		"Orcas":    []string{"Zoey", "Waldo", "Xavier"},
		"Zebras":   []string{"Betty", "Charlotte", "Diana"},
		"Giraffes": []string{"Janice", "Kathy", "Lilly"},
	}
	fmt.Printf("`animalsMap`=%v (len: %d)\n", animalsMap, len(animalsMap))

	// make() can be used to create a map with a default size.
	// Useful when it is known how many elements will be there,
	// but don't know the exact key-value pairs yet.
	userData := make(map[int][]string, 10)
	// Maps created with make() still have a length of 0, and can grow past the initially specified size.
	fmt.Printf("`userData`=%v (len: %d)\n", userData, len(userData))

	// The key for a map must be any comparable type. This means that slices or maps can not be used as keys,
	// as these are not comparable.
	//map_as_key := map[map[int]float64]bool{}
	//slice_as_key := map[[]int]bool{}
}

func try_basic_usage() {
	scores := map[string]int{
		"Alice": 1,
		"Bob":   0,
	}
	scores["Bob"]++
	scores["Alice"] = 3
	fmt.Println(scores)

	// built-in delete() function is used to remove keys from a map.
	delete(scores, "Bob")
	fmt.Printf("After deleting \"Bob\", `scores`=%v\n", scores)

	// From version 1.21 there's an Equal function that can be used to compare maps.
	// There's also an option to use custom comparison functions with EqualFunc, similar to slices.
	aliceMap := map[string]int{
		"Alice": 3,
	}
	fmt.Printf("Comparison of scores with aliceMap: %v\n", maps.Equal(scores, aliceMap))
	scores["Eve"] = 4
	fmt.Printf("After adding Eve to scores, comparison with aliceMap: %v\n", maps.Equal(scores, aliceMap))

	// Similar to slices, clear() function can be used to empty a map too.
	clear(scores)
	fmt.Printf("After `clear(scores)`, scores=%v\n", scores)
}

func comma_ok_idiom() {
	m := map[string]int{
		"hello":  1,
		"world":  2,
		"of":     3,
		"Golang": 4,
	}
	// This is the comma ok idiom.
	//  - v is initialized to either a value from the map, or a default zero value.
	//  - ok is set to either true if given key exists in the map, or false otherwise.
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ok = m["OCaml"]
	fmt.Println(v, ok)
}

func map_as_set() {
	// Map can be used as a set in a pinch. This does not require third-party libraries,
	// but on the other hand is somewhat limited with no union, intersection etc. functions.
	vals := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	intSet := map[int]bool{}
	for _, val := range vals {
		intSet[val] = true
	}
	fmt.Printf("`len(vals)`=%d, `len(intSet)`=%d\n", len(vals), len(intSet))
	fmt.Printf("`intSet[5]`=%v, `intSet[8]`=%v\n", intSet[5], intSet[8])
}

func main() {
	//try_nil_maps()
	try_map_literals()
	try_basic_usage()
	comma_ok_idiom()
	map_as_set()
}
