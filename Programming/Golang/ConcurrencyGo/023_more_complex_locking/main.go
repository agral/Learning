package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// holds a bank balance
	var bankBalance int
	var mtxBalance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks (1 year) and print out how much is made.
	// Keep a running total.
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				mtxBalance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				mtxBalance.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00\n", bankBalance)
}
