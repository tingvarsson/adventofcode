package day11

import (
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
		{"cqjxjnds", "cqjxxyzz"},
		{"cqjxxyzz", "cqkaabcc"},
	}
	for _, test := range tests {
		if output := run(test.input); output != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, output)
		}
	}
}
