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

type coord struct {
	X int
	Y int
}

func find(cs []coord, a coord) bool {
	for _, c := range cs {
		if c == a {
			return true
		}
	}
	return false
}

func run(filepath string, i int) (result int) {
	input := utils.ReadFileToString(filepath)
	inputData := strings.Split(input, ",")
	var intcode []int
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}

	var visited []coord
	colors := make(map[coord]int)
	current := coord{0, 0}
	colors[current] = i
	dir := up
	p := program.New(intcode)
	for true {
		p.Output = []int{}
		p.Run([]int{colors[current]})
		if p.Halted {
			break
		}
		colors[current] = p.Output[0]
		if !find(visited, current) {
			visited = append(visited, current)
		}
		dir = turn(dir, p.Output[1])
		x := current.X
		y := current.Y
		if dir == up {
			y--
		} else if dir == right {
			x++
		} else if dir == down {
			y++
		} else if dir == left {
			x--
		}
		current = coord{x, y}
	}

	result = len(visited)

	if i == 1 {
		hull := make([][]int, 6)
		for i := range hull {
			hull[i] = make([]int, 50)
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
