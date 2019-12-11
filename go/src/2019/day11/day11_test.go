package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input     string
		input2    int
		expected  int
	}{
		{os.Getenv("ROOT") + "/2019/day11/input", 0, 2252},
		{os.Getenv("ROOT") + "/2019/day11/input", 1, 248},
	}

	for _, test := range tests {
		if o1 := run(test.input, test.input2); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, o1)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day11/input"
	main()
}
