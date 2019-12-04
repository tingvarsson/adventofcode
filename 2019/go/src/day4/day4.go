package main

import (
	"strconv"
)

func ruleCheck(input string) bool {
	adjacentFound := false
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			adjacentFound = true
		} else if input[i] > input[i+1] {
			return false
		}
		if i >= len(input)-1 {
			break
		}
	}
	return adjacentFound
}

func ruleCheck2(input string) bool {
	twoAdjacentFound := false
	for i := 0; i < len(input)-1; i++ {
		if !twoAdjacentFound && input[i] == input[i+1] {
			twoAdjacentFound = true
			if i > 0 && input[i-1] == input[i] {
				twoAdjacentFound = false
			}
			if i < len(input)-2 && input[i+1] == input[i+2] {
				twoAdjacentFound = false
			}
		} else if input[i] > input[i+1] {
			return false
		}
		if i >= len(input)-1 {
			break
		}
	}
	return twoAdjacentFound
}

func run(start, stop int) (found, found2 int) {
	for i := start; i <= stop; i++ {
		password := strconv.Itoa(i)
		if ruleCheck(password) {
			found++
		}
		if ruleCheck2(password) {
			found2++
		}
	}
	return
}
