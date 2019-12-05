package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"../../../day8/example", 12},
		{"../../../day8/input", 1333},
	}
	for _, test := range tests {
		if output := run(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestRun2(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{"../../../day8/example", 19},
		{"../../../day8/input", 2046},
	}
	for _, test := range tests {
		if output := run2(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
