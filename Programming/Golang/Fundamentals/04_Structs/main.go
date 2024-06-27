package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (ptrPerson *person) updateName(newFirstName string) {
	(*ptrPerson).firstName = newFirstName
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@party.com",
			zipCode: 94000,
		},
	}
	jim.updateName("jimmy")
	jim.print()
}
