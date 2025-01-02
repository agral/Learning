package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(20)) // should print 30

	funcAddTo := myAdder.AddTo
	fmt.Println(funcAddTo(25)) // should print 35
}
