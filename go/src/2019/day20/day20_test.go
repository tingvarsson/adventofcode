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
		{os.Getenv("ROOT") + "/2019/day20/example", 23},
		{os.Getenv("ROOT") + "/2019/day20/example2", 58},
		{os.Getenv("ROOT") + "/2019/day20/input", 656},
	}

	for _, test := range tests {
		if o1 := run(test.input); o1 != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, o1)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day20/input"
	main()
}
