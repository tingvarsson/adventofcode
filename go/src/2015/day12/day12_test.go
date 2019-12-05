package day12

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input     string
		expected  int
		expected2 int
	}{
		{"../../../day12/example", 6, 4},
		{"../../../day12/input", 111754, 0},
	}
	for _, test := range tests {
		if output, output2 := run(test.input); output != test.expected || output2 != test.expected2 {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output, test.expected2, output2)
		}
	}
}
