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
	input := utils.AtoiSlice(data)

	for n := 0; n < 100; n++ {
		phase(input)
	}

	return utils.IntsToString(input[:8])
}

func phase2(input []int) {
	sum := 0
	for i := len(input) - 1; i >= 0; i-- {
		sum += input[i]
		input[i] = utils.Abs(sum) % 10
	}
	return
}

func run2(filepath string) string {
	data := utils.ReadFileToString(filepath)
	input := utils.AtoiSlice(data)

	offset := 0
	for n := 0; n < 7; n++ {
		offset += input[6-n] * utils.Pow(10, n)
	}

	normOffset := offset % (len(input))
	longInput := make([]int, len(input)*10000-offset)
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
