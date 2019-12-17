package main

import (
	"fmt"
	"os"
	"program"
	"strconv"
	"strings"
	"utils"
)

func move(d string, p utils.Dim2) (n utils.Dim2) {
	n = utils.Dim2{p.X, p.Y}
	switch d {
	case "^":
		n.Y--
	case "v":
		n.Y++
	case "<":
		n.X--
	case ">":
		n.X++
	}
	return
}

func turnLeft(d string) string {
	switch d {
	case "^":
		return "<"
	case "v":
		return ">"
	case "<":
		return "v"
	case ">":
		return "^"
	}
	return ""
}

func turnRight(d string) string {
	switch d {
	case "^":
		return ">"
	case "v":
		return "<"
	case "<":
		return "^"
	case ">":
		return "v"
	}
	return ""
}

type areaMap map[utils.Dim2]string

func parseIntcode(file string) (intcode []int) {
	input := utils.ReadFileToString(file)
	inputData := strings.Split(input, ",")
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}
	return
}

func generateView(p program.Program) (a areaMap, robot utils.Dim2) {
	a = make(areaMap)
	p.Run([]int{})
	output := p.GetOutput()
	x, y := 0, 0
	for _, o := range output {
		if string(o) == "\n" {
			x = 0
			y++
			continue
		}
		if string(o) == "." {
			x++
			continue
		}

		if string(o) == "^" || string(o) == ">" || string(o) == "v" || string(o) == "<" {
			robot = utils.Dim2{x, y}
		}
		a[utils.Dim2{x, y}] = string(o)
		x++
	}
	return
}

func findIntersections(a areaMap) (intersections []utils.Dim2) {
	for c := range a {
		intersection := true
		for _, d := range []string{"^", ">", "v", "<"} {
			newPos := move(d, c)
			if _, ok := a[newPos]; !ok {
				intersection = false
			}
		}
		if intersection {
			intersections = append(intersections, c)
		}
	}
	return
}

func generatePath(a areaMap, robot utils.Dim2) (path string) {
	for true {
		dir := a[robot]
		
		newPos := move(dir, robot)
		if _, ok := a[newPos]; ok {
			steps := 1
			for true {
				tmpPos := move(dir, newPos)
				if _, ok := a[tmpPos]; ok {
					newPos = tmpPos
					steps++
					continue
				}
				break
			}
			a[robot] = "#"
			robot = newPos
			a[robot] = dir
			path += "," + strconv.Itoa(steps)
			continue
		}

		newPos = move(turnLeft(dir), robot)
		if _, ok := a[newPos]; ok {
			a[robot] = turnLeft(dir)
			path += ",L"
			continue
		}

		newPos = move(turnRight(dir), robot)
		if _, ok := a[newPos]; ok {
			a[robot] = turnRight(dir)
			path += ",R"
			continue
		}
		break
	}
	path = path[1:]
	return
}

func run(filepath string) (result, result2 int) {
	intcode := parseIntcode(filepath)
	p := program.New(intcode)
	area, robot := generateView(p)
	intersections := findIntersections(area)
	for _, i := range intersections {
		result += i.X * i.Y
	}

	path := generatePath(area, robot)
	fmt.Println(path)
	// manually deduced this shit...
	inputString := "A,B,A,B,C,C,B,A,B,C\nL,10,R,10,L,10,L,10\nR,10,R,12,L,12\nR,12,L,12,R,6\nn\n"
	var input []int
	for _, s := range inputString {
		input = append(input, int(s))
	}
	intcode[0] = 2
	p = program.New(intcode)
	p.Run(input)
	o := p.GetOutput()
	result2 = o[len(o)-1]

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1, r2)
}
