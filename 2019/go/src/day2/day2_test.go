package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"../../../day2/example", 3500},
		{"../../../day2/example2", 30},
		{"../../../day2/input", 4023471},
	}

	for _, test := range tests {
		if o1, o2, o3 := run(test.input); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v %v %v", test.input, test.expected, o1, o2, o3)
		}
	}
}
