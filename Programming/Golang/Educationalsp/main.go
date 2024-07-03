package main

import (
	"bufio"
	"educationalsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/tmp/educationalsp.txt")
	logger.Println("Educationalsp is starting")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("Bad log file given")
	}
	return log.New(logfile, "[educationalsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
