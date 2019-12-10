package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input     string
		expected  coord
		expected2 int
		expected3 coord
	}{
		{os.Getenv("ROOT") + "/2019/day10/example", coord{3, 4}, 8, coord{0, 0}},
		{os.Getenv("ROOT") + "/2019/day10/example2", coord{5, 8}, 33, coord{0, 0}},
		{os.Getenv("ROOT") + "/2019/day10/example3", coord{1, 2}, 35, coord{0, 0}},
		{os.Getenv("ROOT") + "/2019/day10/example4", coord{6, 3}, 41, coord{0, 0}},
		{os.Getenv("ROOT") + "/2019/day10/example5", coord{11, 13}, 210, coord{8, 2}},
		{os.Getenv("ROOT") + "/2019/day10/input", coord{20, 18}, 280, coord{7, 6}},
	}

	for _, test := range tests {
		if o1, o2, o3 := run(test.input); o1 != test.expected || o2 != test.expected2 || o3 != test.expected3 {
			t.Errorf("Test Failed: %v input, %v %v %v expected, recieved: %v %v %v", test.input, test.expected, test.expected2, test.expected3, o1, o2, o3)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day10/input"
	main()
}
