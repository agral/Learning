package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("%v -> %v\n", color, hex)
	}
}
