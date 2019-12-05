package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {
	input := utils.ReadFileToString("day1/input")
	fmt.Printf("Floor number: %d\n", strings.Count(input, "(")-strings.Count(input, ")"))

	floorCount := 0
	for i, r := range input {
		if r == '(' {
			floorCount++
		} else if r == ')' {
			floorCount--
		}
		if floorCount == -1 {
			fmt.Printf("Reached basement at position: %d\n", i+1)
			break
		}
	}
}
