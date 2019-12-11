package main

import (
	"fmt"
	"os"
	"program"
	"strings"
	"utils"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func turn(d Direction, i int) Direction {
	if i == 0 {
		if d == Up {
			return Left
		}
		return d-1
	} else {
		if d == Left {
			return Up
		}
		return d+1
	}
}

type coord struct {
	X int
	Y int
	Color int
}

func run(filepath string) (result, result2 int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	var visited []coord
	current := coord{0, 0, 0}
	var dir Direction = Up
	p := program.New(intcode)
	for p.Run([]int{current.Color}); !p.Halted {
		newColor, turn := p.Output[0], p.Output[1]
		current.Color = newColor
		visited = append(visited, current)
		dir = turn(dir, turn)
		if dir == Up {
			current.Y++
		} else if dir == Right {
			current.X++
		} else if dir == Down {
			current.Y--
		} else if dir == Left {
			current.X--
		}

	}

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
