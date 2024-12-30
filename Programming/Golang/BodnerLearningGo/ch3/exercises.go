package main

import "fmt"

func ch3_ex1() {
	// Write a program that defines a variable named greetings of type slice of
	// strings with the following values: "Hello", "Hola", "नमस्कार", "こんにちは",
	// and "Привіт". Create a subslice containing the first two values; a second subslice
	// with the second, third, and fourth values; and a third subslice with the fourth and
	// fifth values. Print out all four slices.
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	gr01 := greetings[:2]
	gr13 := greetings[1:4]
	gr34 := greetings[3:]
	fmt.Printf("`greetings`=%v\n", greetings)
	fmt.Printf("`gr01`=%v\n", gr01)
	fmt.Printf("`gr13`=%v\n", gr13)
	fmt.Printf("`gr34`=%v\n", gr34)
}

func ch3_ex2() {
	// Write a program that defines a string variable called message with the value
	// "Hi 👩 and 👨"
	// and prints the fourth rune in it as a character, not a number.
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Printf("%v\n", message)
	fmt.Printf("%v (rune value: %v)\n", string(runes[3]), runes[3])
	fmt.Printf("%v (rune value: %v)\n", string(runes[9]), runes[9])
}

func ch3_ex3() {
	// Write a program that defines a struct called Employee with three fields:
	// firstName, lastName, and id. The first two fields are of type string, and the
	// last field (id) is of type int. Create three instances of this struct using whatever
	// values you’d like. Initialize the first one using the struct literal style without
	// names, the second using the struct literal style with names, and the third with a
	// var declaration. Use dot notation to populate the fields in the third struct. Print
	// out all three structs.
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	alice := Employee{"Alice", "Johnson", 4321} // struct literal style without names
	bob := Employee{                            // struct literal style with names (also in custom order)
		id:        1234,
		lastName:  "Carter",
		firstName: "Bob",
	}
	var charlie = Employee{"Charles", "Smith", 1337} // var declaration
	fmt.Println(alice)
	fmt.Println(bob)
	fmt.Println(charlie)
}

func main() {
	ch3_ex1()
	ch3_ex2()
	ch3_ex3()
}
