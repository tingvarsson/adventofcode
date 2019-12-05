package main

import (
	"fmt"
	"strings"
	"utils"
)

func getValue(memory []int, param int, mode int) int {
	if mode == 0 {
		return memory[param]
	}

	return param
}

func runProg(memory []int, input int) (output int) {
	i := 0
	for true {
		op := memory[i] % 100
		m1 := (memory[i] / 100) % 10
		m2 := (memory[i] / 1000) % 10
		if op == 1 {
			memory[memory[i+3]] = getValue(memory, memory[i+1], m1) + getValue(memory, memory[i+2], m2)
			i += 4
		} else if op == 2 {
			memory[memory[i+3]] = getValue(memory, memory[i+1], m1) * getValue(memory, memory[i+2], m2)
			i += 4
		} else if op == 3 {
			memory[memory[i+1]] = input // assumed position mode
			i += 2
		} else if op == 4 {
			output = getValue(memory, memory[i+1], m1)
			i += 2
		} else if op == 5 {
			if getValue(memory, memory[i+1], m1) != 0 {
				i = getValue(memory, memory[i+2], m2)
			} else {
				i += 3
			}
		} else if op == 6 {
			if getValue(memory, memory[i+1], m1) == 0 {
				i = getValue(memory, memory[i+2], m2)
			} else {
				i += 3
			}
		} else if op == 7 {
			if getValue(memory, memory[i+1], m1) < getValue(memory, memory[i+2], m2) {
				memory[memory[i+3]] = 1
			} else {
				memory[memory[i+3]] = 0
			}
			i += 4
		} else if op == 8 {
			if getValue(memory, memory[i+1], m1) == getValue(memory, memory[i+2], m2) {
				memory[memory[i+3]] = 1
			} else {
				memory[memory[i+3]] = 0
			}
			i += 4
		} else if op == 99 {
			return
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

	memory := append([]int(nil), intcode...)
	result = runProg(memory, 1)

	memory = append([]int(nil), intcode...)
	result2 = runProg(memory, 5)

	return
}

func main() {
	result, result2 := run("../../../day5/input")
	fmt.Printf("result: %v\nresult2: %v\n", result, result2)
}
