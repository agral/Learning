package main

import (
	"fmt"
)

const help = `pomodoro - run an interactive pomodoro session in a terminal

	options: TBD

	author: (c) 2024 https://gralin.ski
	license: MIT`

func print_help() {
	fmt.Println(help)
}

func main() {
	print_help()
}
