package main

import (
	"fmt"
	"os"
	"program"
	"strings"
	"utils"
)

func run(filepath string) (result, result2 int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	prog := make([]int, 10000)
	for i, p := range intcode {
		prog[i] = p
	}
	p := program.New(prog)
	_ = p.Run([]int{1})
	result = p.Output[0]

	prog = make([]int, 10000)
	for i, p := range intcode {
		prog[i] = p
	}

	p = program.New(prog)
	_ = p.Run([]int{2})
	result2 = p.Output[0]

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
