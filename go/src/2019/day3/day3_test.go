package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		input string
		e1    int
		e2    int
	}{
		{os.Getenv("ROOT") + "/2019/day3/example", 6, 30},
		{os.Getenv("ROOT") + "/2019/day3/example2", 159, 610},
		{os.Getenv("ROOT") + "/2019/day3/example3", 135, 410},
		{os.Getenv("ROOT") + "/2019/day3/input", 3229, 32132},
	}

	for _, test := range tests {
		if o1, o2 := run(test.input); o1 != test.e1 || o2 != test.e2 {
			t.Errorf("Test Failed: %v input, %v %v expected, recieved: %v %v", test.input, test.e1, test.e2, o1, o2)
		}
	}
}

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2019/day3/input"
	main()
}
