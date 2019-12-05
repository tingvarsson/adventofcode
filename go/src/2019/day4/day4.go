package main

import (
	"fmt"
	"strconv"
)

func ruleCheck(input string) (ok bool) {
	ok = false
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			ok = true
		} else if input[i] > input[i+1] {
			return false
		}
	}
	return
}

func ruleCheck2(input string) (ok bool) {
	ok = false
	for i := 0; i < len(input)-1; i++ {
		if !ok && input[i] == input[i+1] {
			ok = true
			if i > 0 && input[i-1] == input[i] {
				ok = false
			}
			if i < len(input)-2 && input[i+1] == input[i+2] {
				ok = false
			}
		} else if input[i] > input[i+1] {
			return false
		}
	}
	return
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

func main() {
	r1, r2 := run(278384, 824795)
	fmt.Println(r1)
	fmt.Println(r2)
}
