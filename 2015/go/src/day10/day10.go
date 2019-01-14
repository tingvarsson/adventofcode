package main

import (
	"utils"
)

func run(input string, iterations int) (length int) {
	result := utils.AtoiSlice(input)
	length = len(result)
	for iter := 0; iter < iterations; iter++ {
		var tmp []int
		for i := 0; i < length; {
			number := result[i]
			count := 1
			for ; i+count < length; count++ {
				if result[i+count] != number {
					break
				}
			}
			tmp = append(tmp, count, number)
			i += count
		}
		result = tmp
		length = len(result)
	}
	return
}

func main() {}
