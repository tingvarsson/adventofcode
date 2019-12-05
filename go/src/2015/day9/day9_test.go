package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input     string
		expected  int
		expected2 int
	}{
		{"../../../day9/example", 605, 982},
		{"../../../day9/input", 117, 909},
	}
	for _, test := range tests {
		if output, output2 := run(test.input); output != test.expected || output2 != test.expected2 {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output, test.expected2, output2)
		}
	}
}
