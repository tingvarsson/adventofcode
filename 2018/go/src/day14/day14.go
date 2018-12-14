package main

import (
	"fmt"
	"reflect"
)

func findLast(s []int, p []int) bool {
	if len(s) < len(p) {
		return false
	}
	if reflect.DeepEqual(s[len(s)-len(p):], p) {
		return true
	}
	return false
}

func main() {
	input := []int{3, 7}
	elves := [2]int{0, 1}
	pattern := []int{7, 6, 5, 0, 7, 1}

	for {
		newValue := input[elves[0]] + input[elves[1]]
		if newValue/10 != 0 {
			input = append(input, newValue/10)
			if findLast(input, pattern) {
				break
			}
		}
		input = append(input, newValue%10)
		if findLast(input, pattern) {
			break
		}
		elves[0] = (elves[0] + 1 + input[elves[0]]) % len(input)
		elves[1] = (elves[1] + 1 + input[elves[1]]) % len(input)
	}
	fmt.Printf("Ten recipes after sequence: %v\n", input[765071:765071+10])
	fmt.Printf("#recipes before sequence: %d", len(input)-len(pattern))
}
