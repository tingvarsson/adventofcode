package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input      string
		iterations int
		expected   int
	}{
		{"1", 5, 6},
		{"1113222113", 40, 252594},
		{"1113222113", 50, 3579328},
	}
	for _, test := range tests {
		if output := run(test.input, test.iterations); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
