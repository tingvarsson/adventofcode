package main

import (
	"fmt"
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

func run(filepath string) (pos0 int, noun int, verb int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	memory := append([]int(nil), intcode...)
	memory[1] = 12
	memory[2] = 2
	pos0 = runProg(memory)

	for noun = 0; noun < 100; noun++ {
		for verb = 0; verb < 100; verb++ {
			memory := append([]int(nil), intcode...)
			memory[1] = noun
			memory[2] = verb
			currentPos0 := runProg(memory)
			if currentPos0 == 19690720 {
				return
			}
		}
	}
	return
}

func main() {
	pos0, noun, verb := run("../../../day2/input")
	fmt.Printf("pos0: %v \nnoun: %v \nverb: %v \n", pos0, noun, verb)
}
