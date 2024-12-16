package main

import "fmt"

const Pi = 3.14 // good enough

func main() {
	const World = "세상" // this program is basically a "안녕, 세상!", but in Go ;-)

	fmt.Println("Hello,", World)
	fmt.Println("Happy", Pi, "Day!")
}
