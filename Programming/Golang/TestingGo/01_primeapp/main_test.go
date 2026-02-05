package main

import "testing"

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result == true {
		t.Errorf("With %d as test parameter, got `true`, but expected `false`", 0)
	}
	if msg != "0 is not prime by definition." {
		t.Error("Wrong message returned:", msg)
	}

	result, msg = isPrime(13)
	if result == false {
		t.Errorf("With %d as test parameter, got `false`, but expected `true`", 13)
	}
	if msg != "13 is a prime number." {
		t.Error("Wrong message returned:", msg)
	}

}
