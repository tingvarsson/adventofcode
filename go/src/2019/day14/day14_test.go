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
		{os.Getenv("ROOT") + "/2019/day14/example", 31, 34482758620},
		{os.Getenv("ROOT") + "/2019/day14/example2", 165, 6323777403},
		{os.Getenv("ROOT") + "/2019/day14/example3", 13312, 82892753},
		{os.Getenv("ROOT") + "/2019/day14/example4", 180697, 5586022},
		{os.Getenv("ROOT") + "/2019/day14/example5", 2210736, 460664},
		{os.Getenv("ROOT") + "/2019/day14/input", 399063, 4215654},
	}

	for _, test := range tests {
		if o1, o2 := run(test.input); o1 != test.expected || o2 != test.expected2 {
			t.Errorf("FAILED: %v input, %v %v expected, recieved: %v %v", test.input, test.expected, test.expected2, o1, o2)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day14/input"
	main()
}
