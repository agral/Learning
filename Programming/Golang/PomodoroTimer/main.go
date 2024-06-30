package main

import (
	"fmt"
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

func print_help() {
	fmt.Println(help)
}

func main() {
	print_help()
}
