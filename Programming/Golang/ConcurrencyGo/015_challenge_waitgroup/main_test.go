package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	msg = "test"
	go updateMessage("updated message", &wg)
	wg.Wait()
	if msg != "updated message" {
		t.Errorf("Expected to see \"updated message\", instead msg=\"%s\"", msg)
	}
}

func Test_printMessage(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "testing"
	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	if !strings.Contains(output, "testing") {
		t.Errorf("Expected to see \"testing\" printed; saw \"%s\" instead", output)
	}

	os.Stdout = stdout
}

func Test_main(t *testing.T) {
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	if !strings.Contains(output, "Hello, universe!\nHello, cosmos!\nHello, world!") {
		t.Errorf("Wrong message printed in main(). Saw:\n%s", output)
	}

	os.Stdout = stdout
}
