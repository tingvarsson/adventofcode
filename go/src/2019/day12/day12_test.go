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
		expected2 int
	}{
		{os.Getenv("ROOT") + "/2019/day12/example", 10, 179, 2772},
		{os.Getenv("ROOT") + "/2019/day12/example2", 100, 1940, 4686774924},
		//{os.Getenv("ROOT") + "/2019/day12/input", 1000, 14606, 0},
	}

	for _, test := range tests {
		if o1, o2 := run(test.input, test.input2); o1 != test.expected {
			t.Errorf("FAILED: %v input, %v %v expected, recieved: %v %v", test.input, test.expected, test.expected2, o1, o2)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day11/input"
	main()
}
