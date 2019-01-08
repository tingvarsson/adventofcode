package main

import (
	"testing"
)

func TestNiceBasic(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}
	for _, test := range tests {
		if output := niceBasic(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestNiceEnhanced(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}
	for _, test := range tests {
		if output := niceEnhanced(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		tester   niceTest
		expected int
	}{
		{"../../../day5/input", niceBasic, 258},
		{"../../../day5/input", niceEnhanced, 53},
	}
	for _, test := range tests {
		if output := run(test.input, test.tester); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
