package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ch5_ex1 is already done in calculator.go

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: could not open the file.\n%s\n", err)
		return 0, err
	}
	defer f.Close()
	buffer := make([]byte, 4096)
	length := 0
	for {
		count, err := f.Read(buffer)
		length += count
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Input error.\n%s\n", err)
				return 0, err
			}
			break
		}
	}
	return length, nil
}

func prefixer(prefix string) func(string) string {
	var f = func(input string) string {
		return fmt.Sprintf("%s %s", prefix, input)
	}
	return f
}

func ch5_ex2() {
	// Write a function called fileLen that has an input parameter of type string and
	// returns an int and an error. The function takes in a filename and returns the
	// number of bytes in the file. If there is an error reading the file, return the error.
	// Use defer to make sure the file is closed properly.
}

func ch5_ex3() {
	// Write a function called prefixer that has an input parameter of type string
	// and returns a function that has an input parameter of type string and returns a
	// string. The returned function should prefix its input with the string passed into
	// prefixer. Use the following main function to test prefixer: (...)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a filename for fileLen (exercise #2)")
	}
	l, err := fileLen(os.Args[1])
	if err != nil {
		fmt.Println("There were errors.")
	}
	fmt.Printf("Length of %s is %d bytes.\n", os.Args[1], l)

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
