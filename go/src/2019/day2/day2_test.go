package main

import (
	"os"
	"testing"
)

func TestRunProg(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, 3500},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 30},
	}

	for _, test := range tests {
		if output := runProg(test.input); output != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, output)
		}
	}
}

func TestRun(t *testing.T) {
	var tests = []struct {
		input     string
		expected  int
		expected2 int
		expected3 int
	}{
		{os.Getenv("ROOT") + "/2019/day2/input", 4023471, 80, 51},
	}

	for _, test := range tests {
		if o1, o2, o3 := run(test.input); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v %v %v expected, recieved: %v %v %v", test.input, test.expected, test.expected2, test.expected3, o1, o2, o3)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day2/input"
	main()
}
