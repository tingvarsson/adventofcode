package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		lookup   string
		expected uint16
	}{
		{"../../../day7/example", "d", 72},
		{"../../../day7/example", "e", 507},
		{"../../../day7/example", "f", 492},
		{"../../../day7/example", "g", 114},
		{"../../../day7/example", "h", 65412},
		{"../../../day7/example", "i", 65079},
		{"../../../day7/example", "x", 123},
		{"../../../day7/example", "y", 456},
		{"../../../day7/input", "a", 16076},
		{"../../../day7/input2", "a", 2797},
	}
	for _, test := range tests {
		if output := run(test.input, test.lookup); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
