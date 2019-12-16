package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{os.Getenv("ROOT") + "/2019/day16/example2", "24176176"},
		{os.Getenv("ROOT") + "/2019/day16/example3", "73745418"},
		{os.Getenv("ROOT") + "/2019/day16/example4", "52432133"},
		{os.Getenv("ROOT") + "/2019/day16/input", "40580215"},
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
		expected string
	}{
		{os.Getenv("ROOT") + "/2019/day16/example5", "84462026"},
		{os.Getenv("ROOT") + "/2019/day16/example6", "78725270"},
		{os.Getenv("ROOT") + "/2019/day16/example7", "53553731"},
		{os.Getenv("ROOT") + "/2019/day16/input", "22621597"},
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
