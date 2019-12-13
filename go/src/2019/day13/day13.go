package main

import (
	"fmt"
	"os"
	"program"
	"strings"
	"utils"
)

type object int

const (
	empty object = iota
	wall
	block
	paddle
	ball
)

func run(filepath string) (result, result2 int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	p := program.New(intcode)
	p.RunWoArgs()
	output := p.GetOutput()
	
	for i := 0; i < len(output); i+=3 {
		if object(output[i+2]) == block {
			result++
		}
	}


	intcode[0] = 2
	p = program.New(intcode)
	joystick := 0
	paddleX := 0
	ballX := 0
	for true {
		if paddleX > ballX {
			joystick = -1
		} else if paddleX < ballX {
			joystick = 1
		} else {
			joystick = 0
		}

		p.Run([]int{joystick})

		output = p.GetOutput()
		for i := 0; i < len(output); i+=3 {
			if output[i] == -1 && output[i+1] == 0 {
				result2 = output[i+2]
			} else if object(output[i+2]) == paddle {
				paddleX = output[i]
			} else if object(output[i+2]) == ball {
				ballX = output[i]
			}
		}
		if p.Halted {
			return
		}
	}

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
