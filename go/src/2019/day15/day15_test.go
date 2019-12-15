package main

import (
	"os"
	"testing"
	"utils"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input    string
		expected utils.Dim2
		expected2 int
		expected3 int
	}{
		{os.Getenv("ROOT") + "/2019/day15/input", utils.Dim2{12,14}, 318, 390},
	}

	for _, test := range tests {
		if o1, o2, o3 := run(test.input); o1 != test.expected || o2 != test.expected2 || o3 != test.expected3 {
			t.Errorf("Test Failed: %v input, %v %v %v expected, recieved: %v %v %v", test.input, test.expected, test.expected2, test.expected3, o1, o2, o3)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day15/input"
	main()
}
