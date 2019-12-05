package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		numZeros int
		expected int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
		{"iwrupvqb", 5, 346386},
		{"iwrupvqb", 6, 9958218},
	}
	for _, test := range tests {
		if output := run(test.input, test.numZeros); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
