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
	north direction = 1
	south = 2
	west = 3
	east = 4
)

func oppositeDirection(d direction) (o direction) {
	switch d {
	case north:
		o = south
	case south:
		o = north
	case west:
		o = east
	case east:
		o = west
	}
	return
}

func move(d direction, p utils.Dim2) (n utils.Dim2) {
	n = utils.Dim2{p.X, p.Y}
	switch d {
	case north:
		n.X++
	case south:
		n.X--
	case west:
		n.Y--
	case east:
		n.Y++
	}
	return
}

type status int
const (
	wall status = iota
	moved
	oxygen
)

type areaMap map[utils.Dim2]status

func runRobot(a areaMap, p program.Program, pos utils.Dim2, steps int) (g utils.Dim2, s int, ok bool) {
	a[pos] = moved
	for _, d := range []direction{north, south, west, east} {
		newPos := move(d, pos)
		if _, ok := a[newPos]; ok {
			continue
		}

		p.Run([]int{int(d)})
		o := status(p.PopOutput())

		if o == wall {
			a[newPos] = wall
			continue
		} 

		if o == oxygen {
			g, s, ok = newPos, steps+1, true
			// continue to map the whole area even after oxygen system is found
		}

		if gTmp, sTmp, okTmp := runRobot(a, p, newPos, steps+1); okTmp {
			// found the oxygen system along this path
			g, s, ok = gTmp, sTmp, okTmp
		}
		// step back and continue from where we were
		p.Run([]int{int(oppositeDirection(d))})
		_ = p.PopOutput()
	}
	return
}

func oxygenate(a areaMap, start utils.Dim2) (time int) {
	visited := make(map[utils.Dim2]int)
	visited[start] = 0
	queue := []utils.Dim2{start}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		for _, d := range []direction{north, south, west, east} {
			newPos := move(d, pos)
			if _, ok := visited[newPos]; ok {
				continue
			}
			if a[newPos] == wall {
				continue
			}
			visited[newPos] = visited[pos] + 1
			queue = append(queue, newPos)
		}
	}

	for _, t := range visited {
		if t > time {
			time = t
		}
	}
	return
}

func parseIntcode(file string) (intcode []int) {
	input := utils.ReadFileToString(file)
	inputData := strings.Split(input, ",")
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}
	return
}

func run(filepath string) (result utils.Dim2, result2, result3 int) {
	intcode := parseIntcode(filepath)
	area := make(areaMap)
	p := program.New(intcode)
	start := utils.Dim2{0, 0}
	result, result2, _ = runRobot(area, p, start, 0)

	result3 = oxygenate(area, result)

	return
}

func main() {
	r1, r2, r3 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}
