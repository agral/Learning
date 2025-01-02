package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func main() {
	// Write a program that builds a []Person with 10,000,000 entries (they could all
	// be the same names and ages). See how long it takes to run. Change the value
	// of GOGC and see how that affects the time it takes for the program to complete.
	// Set the environment variable GODEBUG=gctrace=1 to see when garbage collections
	// happen and see how changing GOGC changes the number of garbage collections.
	// What happens if you create the slice with a capacity of 10,000,000?

	//var db []Person
	db := make([]Person, 0, 10_000_000)

	// Without preallocated space for 10M Person instances
	// (line 21 uncommented, line 22 commented out):
	// $ time ./exercise_06_03
	// real    0m1.600s
	// user    0m2.443s
	// sys     0m0.577s

	// With preallocated space (line 21 commented out, line 22 uncommented):
	// $ time ./exercise_06_03
	// real    0m0.307s
	// user    0m0.229s
	// sys     0m0.210s
	// That's about 5 times faster in real clock time. Nice!

	for i := 0; i < 10_000_000; i++ {
		db = append(db, MakePerson("Alice", "Johnson", 32))
	}
}
