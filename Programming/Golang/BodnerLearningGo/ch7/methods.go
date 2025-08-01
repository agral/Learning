package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s (%d years old)", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("count: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	fred := Person{"Fred", "Gilliam", 42}
	fmt.Println(fred.String())

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}
