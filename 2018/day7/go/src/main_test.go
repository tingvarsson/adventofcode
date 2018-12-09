package main

import (
	"testing"
)

func TestDetermineOrder(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"../../example", "CABDFE"},
		{"../../input", "BGJCNLQUYIFMOEZTADKSPVXRHW"},
		{"../../input2", "CQSWKZFJONPBEUMXADLYIGVRHT"},
	}
	for _, test := range tests {
		if output := determineWorkOrder(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

func TestWorkOrder(t *testing.T) {
	var tests = []struct {
		input    string
		workers  int
		workTime int
		expected int
	}{
		{"../../example", 2, 0, 15},
		{"../../input", 5, 60, 1017},
		{"../../input2", 5, 60, 914},
	}
	for _, test := range tests {
		if output := workOrder(test.input, test.workers, test.workTime); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}
