package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		numSanta int
		expected int
	}{
		{"../../../day3/example", 1, 4},
		{"../../../day3/example2", 1, 2},
		{"../../../day3/input", 1, 2081},
		{"../../../day3/example", 2, 3},
		{"../../../day3/example2", 2, 11},
		{"../../../day3/input", 2, 2341},
	}
	for _, test := range tests {
		if output := run(test.input, test.numSanta); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
