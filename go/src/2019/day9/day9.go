package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

func getValue(memory []int, param, mode, relativeBase int) int {
	if mode == 0 {
		return memory[param]
	} else if mode == 1 {
		return param
	} else if mode == 2 {
		return memory[relativeBase+param]
	}
	return 0
}

func setValue(memory []int, param, mode, value, relativeBase int) {
	if mode == 2 {
		memory[relativeBase+param] = value
	} else {
		memory[param] = value
	}
}

func runProg(memory []int, input []int, i int) (o, stop int) {
	rb := 0
	if i != 0 {
		input = input[1:]
	}
	for true {
		op := memory[i] % 100
		m1 := (memory[i] / 100) % 10
		m2 := (memory[i] / 1000) % 10
		m3 := (memory[i] / 10000) % 10
		if op == 1 {
			setValue(memory, memory[i+3], m3, getValue(memory, memory[i+1], m1, rb)+getValue(memory, memory[i+2], m2, rb), rb)
			i += 4
		} else if op == 2 {
			setValue(memory, memory[i+3], m3, getValue(memory, memory[i+1], m1, rb)*getValue(memory, memory[i+2], m2, rb), rb)
			i += 4
		} else if op == 3 {
			if len(input) == 0 {
				stop = i
				break
			}
			setValue(memory, memory[i+1], m1, input[0], rb)
			input = input[1:] // dequeue
			i += 2
		} else if op == 4 {
			o = getValue(memory, memory[i+1], m1, rb)
			i += 2
		} else if op == 5 {
			if getValue(memory, memory[i+1], m1, rb) != 0 {
				i = getValue(memory, memory[i+2], m2, rb)
			} else {
				i += 3
			}
		} else if op == 6 {
			if getValue(memory, memory[i+1], m1, rb) == 0 {
				i = getValue(memory, memory[i+2], m2, rb)
			} else {
				i += 3
			}
		} else if op == 7 {
			if getValue(memory, memory[i+1], m1, rb) < getValue(memory, memory[i+2], m2, rb) {
				setValue(memory, memory[i+3], m3, 1, rb)
			} else {
				setValue(memory, memory[i+3], m3, 0, rb)
			}
			i += 4
		} else if op == 8 {
			if getValue(memory, memory[i+1], m1, rb) == getValue(memory, memory[i+2], m2, rb) {
				setValue(memory, memory[i+3], m3, 1, rb)
			} else {
				setValue(memory, memory[i+3], m3, 0, rb)
			}
			i += 4
		} else if op == 9 {
			rb += getValue(memory, memory[i+1], m1, rb)
			i += 2
		} else if op == 99 {
			stop = 0
			break
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

	prog := make([]int, 10000)
	for i, p := range intcode {
		prog[i] = p
	}
	result, _ = runProg(prog, []int{1}, 0)

	prog = make([]int, 10000)
	for i, p := range intcode {
		prog[i] = p
	}
	result2, _ = runProg(prog, []int{2}, 0)

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
