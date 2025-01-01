package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(lhs int, rhs int) (int, error) { return lhs + rhs, nil }
func sub(lhs int, rhs int) (int, error) { return lhs - rhs, nil }
func mul(lhs int, rhs int) (int, error) { return lhs * rhs, nil }
func div(lhs int, rhs int) (int, error) {
	if rhs == 0 {
		return 0, errors.New("Division by zero")
	}
	return lhs / rhs, nil
}

var opMap = map[string]func(int, int) (int, error){
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
		{"4", "/", "0"},
		{"0", "/", "0"},
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
		result, err := opFunc(lhs, rhs)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}

func main() {
	tryCalc()
}
