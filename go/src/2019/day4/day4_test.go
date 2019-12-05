package main

import (
	"testing"
)

func TestRuleCheck(t *testing.T) {
	var tests = []struct {
		input string
		e1    bool
	}{
		{"111111", true},
		{"223450", false},
		{"123789", false},
		{"111123", true},
		{"135679", false},
		{"122345", true},
	}

	for _, test := range tests {
		if o1 := ruleCheck(test.input); o1 != test.e1 {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.e1, o1)
		}
	}
}

func TestRuleCheck2(t *testing.T) {
	var tests = []struct {
		input string
		e1    bool
	}{
		{"112233", true},
		{"123444", false},
		{"111122", true},
	}

	for _, test := range tests {
		if o1 := ruleCheck2(test.input); o1 != test.e1 {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.e1, o1)
		}
	}
}

func TestRun(t *testing.T) {
	var tests = []struct {
		i1 int
		i2 int
		e1 int
		e2 int
	}{
		{278384, 824795, 921, 603},
	}

	for _, test := range tests {
		if o1, o2 := run(test.i1, test.i2); o1 != test.e1 || o2 != test.e2 {
			t.Errorf("Test Failed: %v %v input, %v %v expected, recieved: %v %v", test.i1, test.i2, test.e1, test.e2, o1, o2)
		}
	}
}

func TestMain(t *testing.T) {
	main()
}
