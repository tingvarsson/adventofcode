package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input        string
		boostEnabled bool
		expected     int
	}{
		{"../../../day6/example", false, 1000000 - 1000 - 4},
		{"../../../day6/input", false, 377891},
		{"../../../day6/input", true, 14110788},
	}
	for _, test := range tests {
		if output := run(test.input, test.boostEnabled); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
