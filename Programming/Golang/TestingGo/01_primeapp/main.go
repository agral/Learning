package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	printIntroduction()
	prompt()

	// create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(doneChan)

	// block until the doneChan gets a value
	<-doneChan

	// close the channel
	close(doneChan)

	// print a goodbye message
	fmt.Println("Goodbye.")
}

func printIntroduction() {
	fmt.Println("Is it Prime?")
	fmt.Println("============")
	fmt.Println("Enter a whole number to check if it's a prime number or not.")
	fmt.Println("Enter `q` to quit.")
}

func prompt() {
	fmt.Print("-> ")
}

func readUserInput(doneChan chan bool) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		res, isDone := checkNumbers(scanner)
		if isDone {
			doneChan <- true
			return
		}

		fmt.Println(res)
		prompt()
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	// read user input:
	scanner.Scan()

	// check if the user wants to quit:
	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	// try to convert the input into an int:
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "Please enter a whole number, or `q` to quit.", false
	}

	_, msg := isPrime(num)
	return msg, false
}

func isPrime(n int) (bool, string) {
	// 0 and 1 are not prime by definition
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime by definition.", n)
	}

	// negative numbers are not prime.
	if n < 0 {
		return false, "Negative numbers are not prime by definition."
	}

	// actually check for primality:
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false, fmt.Sprintf("%d is not a prime number - divisible by %d.", n, i)
		}
	}
	return true, fmt.Sprintf("%d is a prime number.", n)
}
