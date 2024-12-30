package main

import "fmt"

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

func main() {
	//try_nil_maps()
	try_map_literals()
}
