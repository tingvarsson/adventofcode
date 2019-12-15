package main

import (
	"fmt"
	"os"
	"program"
	"strings"
	"utils"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

func move(d direction, p utils.Dim2) (n utils.Dim2) {
	n = utils.Dim2{p.X, p.Y}
	switch d {
	case up:
		n.Y--
	case right:
		n.X++
	case down:
		n.Y++
	case left:
		n.X--
	}
	return
}

func turn(d direction, i int) direction {
	if i == 0 {
		if d == up {
			return left
		}
		return d-1
	} else {
		if d == left {
			return up
		}
		return d+1
	}
}

func parseIntcode(file string) (intcode []int) {
	input := utils.ReadFileToString(file)
	inputData := strings.Split(input, ",")
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}
	return
}

func run(filepath string, i int) (result int) {
	intcode := parseIntcode(filepath)

	var visited []utils.Dim2
	colors := make(map[utils.Dim2]int)
	current := utils.Dim2{0, 0}
	colors[current] = i
	dir := up
	p := program.New(intcode)
	for true {
		p.Run([]int{colors[current]})
		o := p.GetOutput()
		if p.Halted {
			break
		}
		colors[current] = o[0]
		if !utils.FindDim2(visited, current) {
			visited = append(visited, current)
		}
		dir = turn(dir, o[1])
		current = move(dir, current)
	}

	result = len(visited)

	if i == 1 {
		hull := make([][]int, 6)
		for i := range hull {
			hull[i] = make([]int, 43)
		}

		for _, c := range visited {
			hull[c.Y][c.X] = colors[c]
		}

		for row := range hull {
			fmt.Println(hull[row])
		}
	}

	return
}

func main() {
	r1 := run(os.Args[1], 0)
	fmt.Println(r1)
}
