package main

import (
	"fmt"
	"time"
)

const help = `pomodoro - run an interactive pomodoro session in a terminal

    Runs a 25 minute timer designated for working, then a 5 minute timer designated for resting.

    Usage:
        pomodoro

    Mandatory arguments:
        none

    Optional arguments:
        none

    author: (c) 2024 https://gralin.ski
    license: MIT`

// const workPeriodSeconds = 25 * 60
// const restPeriodSeconds = 5 * 60
const workPeriodSeconds = 25
const restPeriodSeconds = 5
const totalPeriodSeconds = workPeriodSeconds + restPeriodSeconds

func print_timer(timer int) {
	if timer < workPeriodSeconds {
		fmt.Println("Work")
		fmt.Println(fmt.Sprintf("[%d/%d]", timer, workPeriodSeconds))
	} else {
		fmt.Println("Rest")
		fmt.Println(fmt.Sprintf("[%d/%d]", timer-workPeriodSeconds, restPeriodSeconds))
	}
	fmt.Println(timer)
}

func print_help() {
	fmt.Println(help)
}

func main() {
	for timer := 0; timer < totalPeriodSeconds; timer++ {
		print_timer(timer)
		time.Sleep(time.Second)
	}
	fmt.Println("Done!")
}
