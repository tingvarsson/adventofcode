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
		{"cqjxxyzz", ""},
	}
	for _, test := range tests {
		if output := run(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
