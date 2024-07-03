package main

import (
	"bufio"
	"educationalsp/rpc"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {
}
