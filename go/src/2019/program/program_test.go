package program

import (
	"testing"
)

func TestRunProg(t *testing.T) {
	var tests = []struct {
		prog     []int
		input    int
		expected int
	}{
		{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 8, 1},
		{[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, 7, 0},
		{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 7, 1},
		{[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, 9, 0},
		{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 8, 1},
		{[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, 5, 0},
		{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 4, 1},
		{[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, 8, 0},
		{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 7, 999},
		{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 8, 1000},
		{[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, 9, 1001},
		{[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, 0, 99},
		{[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, 0, 1219070632396864},
		{[]int{104, 1125899906842624, 99}, 0, 1125899906842624},
	}

	for _, test := range tests {
		p := New(test.prog)
		if _ = p.Run([]int{test.input}); p.Output[len(p.Output)-1] != test.expected {
			t.Errorf("Test Failed: %v input, %v expected, recieved: %v", test.input, test.expected, p.Output[len(p.Output)-1])
		}
	}
}
