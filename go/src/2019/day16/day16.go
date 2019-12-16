package main

import (
	"fmt"
	"os"
	"utils"
)

func multiplier(index, n int) (m int) {
	n -= index
	n /= (index + 1)
	switch n % 4 {
	case 0:
		m = 1
	case 1:
		m = 0
	case 2:
		m = -1
	case 3:
		m = 0
	}
	return
}

func phase(input []int) {
	for i := range input {
		sum := 0
		for j := i; j < len(input); j++ {
			sum += input[j] * multiplier(i, j)
		}
		input[i] = utils.Abs(sum) % 10
	}
	return
}

func run(filepath string) string {
	data := utils.ReadFileToString(filepath)
	var input []int
	for _, i := range data {
		input = append(input, utils.Atoi(string(i)))
	}

	for n := 0; n < 100; n++ {
		phase(input)
	}

	return utils.IntsToString(input[:8])
}

func phase2(input []int) {
	output := 0
	for i := len(input) - 1; i >= 0; i-- {
		output += input[i]
		input[i] = utils.Abs(output) % 10
	}
	return
}

func run2(filepath string) string {
	data := utils.ReadFileToString(filepath)
	var input []int
	for _, i := range data {
		input = append(input, utils.Atoi(string(i)))
	}

	endIdx := len(input) * 10000
	startIdx := 0
	for n := 0; n < 7; n++ {
		startIdx += input[6-n] * utils.Pow(10, n)
	}

	normOffset := startIdx % (len(input))
	longInput := make([]int, endIdx-startIdx)
	for n := 0; n < len(longInput); n++ {
		longInput[n] = input[(normOffset+n)%len(input)]
	}

	for n := 0; n < 100; n++ {
		phase2(longInput)
	}

	return utils.IntsToString(longInput[:8])
}

func main() {
	res := run(os.Args[1])
	fmt.Println(res)
	res = run2(os.Args[1])
	fmt.Println(res)
}
