package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func UpdateSlice(slice []string, change string) {
	slice[len(slice)-1] = change
	fmt.Printf("  Before returning from UpdateSlice, `slice`=%v\n", slice)
}

func GrowSlice(slice []string, change string) {
	slice = append(slice, change)
	fmt.Printf("  Before returning from GrowSlice, `slice`=%v\n", slice)
}

func ch6_ex1() {
	// Create a struct named Person with three fields: FirstName and LastName of
	// type string and Age of type int. Write a function called MakePerson that
	// takes in firstName, lastName, and age and returns a Person. Write a second
	// function MakePersonPointer that takes in firstName, lastName, and age and
	// returns a *Person. Call both from main. Compile your program with go build
	// -gcflags="-m". This both compiles your code and prints out which values
	// escape to the heap. Are you surprised about what escapes?
	alice := MakePerson("Alice", "Smith", 32)
	bob := MakePersonPointer("Bob", "Johnson", 64)
	fmt.Println(alice, bob)
	// The results:
	// ./exercises.go:11:6: can inline MakePerson
	// ./exercises.go:15:6: can inline MakePersonPointer
	// ./exercises.go:32:6: can inline ch6_ex2
	// ./exercises.go:35:6: can inline ch6_ex3
	// ./exercises.go:38:6: can inline main
	// ./exercises.go:27:21: inlining call to MakePerson
	// ./exercises.go:28:26: inlining call to MakePersonPointer
	// ./exercises.go:29:13: inlining call to fmt.Println
	// ./exercises.go:11:17: leaking param: firstName to result ~r0 level=0
	// ./exercises.go:11:35: leaking param: lastName to result ~r0 level=0
	// ./exercises.go:15:24: leaking param: firstName
	// ./exercises.go:15:42: leaking param: lastName
	// ./exercises.go:16:9: &Person{...} escapes to heap
	// ./exercises.go:28:26: &Person{...} escapes to heap
	// ./exercises.go:29:13: ... argument does not escape
	// ./exercises.go:29:14: alice escapes to heap

	// Note that the last line indicates that arguments to fmt.Println escape to the heap,
	// same as pointers. So these are either passed by pointer internally, or maybe (more likely)
	// as interface{}.
}

func ch6_ex2() {
	// Write two functions. The UpdateSlice function takes in a []string and a
	// string. It sets the last position in the passed-in slice to the passed-in string. At
	// the end of UpdateSlice, print the slice after making the change. The GrowSlice
	// function also takes in a []string and a string. It appends the string onto the
	// slice. At the end of GrowSlice, print the slice after making the change. Call these
	// functions from main. Print out the slice before each function is called and after
	// each function is called. Do you understand why some changes are visible in main
	// and why some changes are not?
	sl1 := []string{"Alice"}
	fmt.Printf("Before calling UpdateSlice(), `sl1`=%v\n", sl1)
	UpdateSlice(sl1, "Bob")
	fmt.Printf("After calling UpdateSlice(), `sl1`=%v\n", sl1)

	sl2 := []string{"Charlie"}
	fmt.Printf("Before calling GrowSlice(), `sl2`=%v\n", sl2)
	GrowSlice(sl2, "Derek")
	fmt.Printf("After calling GrowSlice(), `sl2`=%v\n", sl2)
}

func ch6_ex3() {
	// Write a program that builds a []Person with 10,000,000 entries (they could all
	// be the same names and ages). See how long it takes to run. Change the value
	// of GOGC and see how that affects the time it takes for the program to complete.
	// Set the environment variable GODEBUG=gctrace=1 to see when garbage collections
	// happen and see how changing GOGC changes the number of garbage collections.
	// What happens if you create the slice with a capacity of 10,000,000?
	db := make([]Person
}

func main() {
	ch6_ex1()
	ch6_ex2()
	ch6_ex3()
}
