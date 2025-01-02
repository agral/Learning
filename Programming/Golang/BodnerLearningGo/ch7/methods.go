package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s (%d years old)", p.FirstName, p.LastName, p.Age)
}

func main() {
	fred := Person{"Fred", "Gilliam", 42}
	fmt.Println(fred.String())
}
