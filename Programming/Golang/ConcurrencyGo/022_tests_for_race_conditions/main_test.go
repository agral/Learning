package main

import "testing"

// Running this with `go test .` results with a pass.
// However, there's a race condition.
// Runnning theis test with `go test -race .` highlights the problem.
func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"
	wg.Add(2)
	go updateMessage("zzz")
	go updateMessage("Goodbye, world!")
	wg.Wait()

	if msg != "Goodbye, world!" && msg != "zzz" {
		t.Errorf("Incorrect value in msg; got: %s", msg)
	}
}
