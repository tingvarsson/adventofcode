package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

func runProg(memory []int) int {
	for i := 0; memory[i] != 99; i += 4 {
		if memory[i] == 1 {
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		} else if memory[i] == 2 {
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		}
	}
	return memory[0]
}

func findInput(intcode []int, output int) (noun, verb int) {
	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			memory := append([]int(nil), intcode...)
			memory[1] = noun
			memory[2] = verb
			if runProg(memory) == output {
				return
			}
		}
	}
	return
}

func run(filepath string) (output int, noun int, verb int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	memory := append([]int(nil), intcode...)
	memory[1] = 12
	memory[2] = 2
	output = runProg(memory)

	noun, verb = findInput(intcode, 19690720)
	return
}

func main() {
	output, noun, verb := run(os.Args[1])
	fmt.Printf("output: %v \nnoun: %v \nverb: %v \n", output, noun, verb)
}
