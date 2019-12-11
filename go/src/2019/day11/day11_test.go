package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input     string
		expected  int
		expected2 int
	}{
		{os.Getenv("ROOT") + "/2019/day11/input", 3765554916, 76642},
	}

	for _, test := range tests {
		if o1, o2 := run(test.input); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v %v expected, recieved: %v %v", test.input, test.expected, test.expected2, o1, o2)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day11/input"
	main()
}
