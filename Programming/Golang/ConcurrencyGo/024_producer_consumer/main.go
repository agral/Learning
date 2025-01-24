package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

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

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delaySeconds := rand.Intn(5) + 1
		fmt.Printf("Received order #%d.\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		isSuccess := false

		if rnd < 5 {
			totalPizzasFailed++
		} else {
			totalPizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzaNumber, delaySeconds)
		time.Sleep(time.Duration(delaySeconds) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			isSuccess = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			isSuccess:   isSuccess,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// Keep track of which pizzas are being made:
	i := 0

	// run forever or until a quit notification is received
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza; sent something to the data channel
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
		// decide whether this pizza is good enough.
	}
}

func main() {
	// seed the pseudorandom number generator
	// Note: no need to do this from Golang 1.20 - the default random source is seeded automatically.

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background (as its own goroutine)
	go pizzeria(pizzaJob)

	// create and run a consumer

	// print out the ending message
}
