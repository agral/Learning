package main

const NumberOfPizzas = 10

var totalPizzasMade, totalPizzasFailed, total int

type PizzaOrder struct {
	pizzaNumber int
	message     string
	isSuccess   bool
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error // this channel holds channel of errors
}

func main() {
	// seed the pseudorandom number generator

	// print out a message

	// create a producer

	// run the producer in the background (as its own goroutine)

	// create and run a consumer

	// print out the ending message
}
