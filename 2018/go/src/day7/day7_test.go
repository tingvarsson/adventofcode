package main

import (
	"testing"
	"utils"
)

func TestDetermineOrder(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"../../../day7/example", utils.ReadFileToLines("../../../day7/example_answer")[0]},
		{"../../../day7/input", utils.ReadFileToLines("../../../day7/input_answer")[0]},
		{"../../../day7/input2", utils.ReadFileToLines("../../../day7/input2_answer")[0]},
	}
	for _, test := range tests {
		if output := determineWorkOrder(test.input); output != test.expected {
			t.Error("Test Failed: {} input, {} expected, recieved: {}", test.input, test.expected, output)
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
		{"../../../day7/example", 2, 0, 15},
		{"../../../day7/input", 5, 60, 1017},
		{"../../../day7/input2", 5, 60, 914},
	}
	for _, test := range tests {
		if output := workOrder(test.input, test.workers, test.workTime); output != test.expected {
			t.Error("Test Failed: {} input, {} workers, {} workTime, {} expected, recieved: {}", test.input, test.workers, test.workTime, test.expected, output)
		}
	}
}
