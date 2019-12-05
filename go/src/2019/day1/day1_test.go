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
		{os.Getenv("ROOT") + "/2019/day1/example", 33583, 50346},
		{os.Getenv("ROOT") + "/2019/day1/input", 3339288, 5006064},
	}

	for _, test := range tests {
		if output, output2 := run(test.input); output != test.expected || output2 != test.expected2 {
			t.Errorf("Test Failed: %v input, %v, %v expected, recieved: %v, %v", test.input, test.expected, test.expected2, output, output2)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day1/input"
	main()
}
