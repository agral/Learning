package main

import "testing"

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name        string
		testNum     int
		expected    bool
		expectedMsg string
	}{
		{"prime", 13, true, "13 is a prime number."},
		{"not prime", 8, false, "8 is not a prime number - divisible by 2."},
	}

	for _, test := range primeTests {
		result, msg := isPrime(test.testNum)
		if test.expected && !result {
			t.Errorf("%s: expected `true`, got `false`", test.name)
		} else if !test.expected && result {
			t.Errorf("%s: expected `false`, got `true`", test.name)
		}
		if test.expectedMsg != msg {
			t.Errorf("%s: expected `%s`, got `%s`", test.name, test.expectedMsg, msg)
		}
	}
}
