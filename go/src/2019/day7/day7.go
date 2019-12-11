package main

import (
	"fmt"
	"os"
	"program"
	"strings"
	"utils"
)

func runAmplifiers(intcode, input []int) (o int) {
	for _, i := range input {
		p := program.New(intcode)
		p.Input = []int{i}
		p.Run([]int{o})
		o = p.PopOutput()
	}
	return
}

func runAmplifiersLoop(intcode, input []int) (o int) {
	var programs []program.Program
	for _, i := range input {
		p := program.New(intcode)
		p.Input = []int{i}
		programs = append(programs, p)
	}
	for true {
		for i := range programs {
			if programs[i].Halted {
				return
			}
			programs[i].Run([]int{o})
			o = programs[i].PopOutput()
		}
	}
	return
}

func run(filepath string) (result, result2 int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	perm := utils.Permutations([]int{0, 1, 2, 3, 4})

	for _, p := range perm {
		if res := runAmplifiers(intcode, p); res > result {
			result = res
		}
	}

	perm = utils.Permutations([]int{5, 6, 7, 8, 9})
	for _, p := range perm {
		if res := runAmplifiersLoop(intcode, p); res > result2 {
			result2 = res
		}
	}

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
