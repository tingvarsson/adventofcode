package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

func getValue(memory []int, param int, mode int) int {
	if mode == 0 {
		return memory[param]
	}

	return param
}

func runProg(memory []int, i1, i2, i int) (o, stop int) {
	o = i2
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
			if i == 0 {
				memory[memory[i+1]] = i1 // assumed position mode
			} else {
				memory[memory[i+1]] = i2 // assumed position mode
			}
			i += 2
		} else if op == 4 {
			o = getValue(memory, memory[i+1], m1)
			i += 2
			stop = i
			break
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
			stop = 0
			break
		}
	}
	return
}

func runAmplifiers(prog, input []int) (o int) {
	for _, i := range input {
		p := append([]int(nil), prog...)
		o, _ = runProg(p, i, o, 0)
	}
	return
}

func runAmplifiersLoop(prog, input []int) (o int) {
	var progs [][]int
	for range input {
		progs = append(progs, append([]int(nil), prog...))
	}
	s := []int{0, 0, 0, 0, 0}
	for true {
		for n, i := range input {
			o, s[n] = runProg(progs[n], i, o, s[n])

			if s[n] == 0 {
				return
			}
		}
	}
	return
}

func permutations(input []int) (output [][]int) {
	var helper func([]int, int)
	helper = func(input []int, n int) {
		if n == 1 {
			tmp := make([]int, len(input))
			copy(tmp, input)
			output = append(output, tmp)
			return
		}
		for i := 0; i < n; i++ {
			helper(input, n-1)
			if n%2 == 1 {
				tmp := input[i]
				input[i] = input[n-1]
				input[n-1] = tmp
			} else {
				tmp := input[0]
				input[0] = input[n-1]
				input[n-1] = tmp
			}
		}
	}
	helper(input, len(input))
	return
}

func run(filepath string) (result, result2 int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	perm := permutations([]int{0, 1, 2, 3, 4})

	prog := append([]int(nil), intcode...)
	for _, p := range perm {
		if res := runAmplifiers(prog, p); res > result {
			result = res
		}
	}

	perm = permutations([]int{5, 6, 7, 8, 9})
	for _, p := range perm {
		if res := runAmplifiersLoop(prog, p); res > result2 {
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
