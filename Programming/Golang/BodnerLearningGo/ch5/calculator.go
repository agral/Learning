package main

import (
	"fmt"
	"strconv"
)

func add(lhs int, rhs int) int { return lhs + rhs }
func sub(lhs int, rhs int) int { return lhs - rhs }
func mul(lhs int, rhs int) int { return lhs * rhs }
func div(lhs int, rhs int) int { return lhs / rhs }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func tryCalc() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Printf("Invalid expression: %v\n", expression)
			continue
		}
		lhs, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Printf("Operator %v not supported.\n", op)
			continue
		}
		rhs, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(lhs, rhs)
		fmt.Println(result)
	}
}

func main() {
	tryCalc()
}
