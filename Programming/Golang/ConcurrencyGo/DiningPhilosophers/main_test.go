package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second
	for i := 0; i < 10; i++ {
		finishedText = []string{}
		dine()
		if len(finishedText) != 5 {
			t.Errorf("Incorrect length of finishedText slice. Want 5, got %d", len(finishedText))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range theTests {
		finishedText = []string{}
		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay
		dine()
		if len(finishedText) != 5 {
			t.Errorf("%s: Incorrect length of finishedText slice. Want 5, got %d", e.name, len(finishedText))
		}
	}
}
