package main

import "testing"

func Test_bookworm(t *testing.T) {
	type testCase struct {
		bookwormsFile string
		expected      []Bookworm
		expectedErr   bool
	}

	tests := map[string]testCase{}

	var handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	var oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}

	for i, tc := range testcases {
		actual := bookworm(tc.bookwormsFile)
		if actual != tc.expected {
			t.Errorf("Testcase bookworm#%02d (%v) failed: want %d, got %d",
				i+1, tc.bookwormsFile, tc.expected, actual)
		}
	}
}
