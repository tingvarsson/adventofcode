package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{os.Getenv("ROOT") + "/2019/day18/example", 8},
		{os.Getenv("ROOT") + "/2019/day18/example2", 86},
		{os.Getenv("ROOT") + "/2019/day18/example3", 132},
		{os.Getenv("ROOT") + "/2019/day18/example4", 136},
		{os.Getenv("ROOT") + "/2019/day18/example5", 81},
		{os.Getenv("ROOT") + "/2019/day18/input", 4350},
	}

	for _, test := range tests {
		if o1 := run(test.input); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, o1)
		}
	}
}

func TestRun2(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
	}{
		{os.Getenv("ROOT") + "/2019/day18/example6", 8},
		{os.Getenv("ROOT") + "/2019/day18/example7", 24},
		{os.Getenv("ROOT") + "/2019/day18/example8", 32},
		{os.Getenv("ROOT") + "/2019/day18/example9", 72},
		{os.Getenv("ROOT") + "/2019/day18/input2", 0},
	}

	for _, test := range tests {
		if o1 := run2(test.input); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, o1)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day16/input"
	main()
}
