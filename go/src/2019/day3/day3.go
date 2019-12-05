package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"utils"
)

type step struct {
	x int
	y int
	d int
}

var stepRegex = regexp.MustCompile("([A-Z])(\\d+)")

type wire struct {
	x    int
	y    int
	d    int
	path []step
}

func traverse(w wire, direction string, distance int) wire {
	dx := 0
	dy := 0
	if direction == "U" {
		dx = 1
	} else if direction == "R" {
		dy = 1
	} else if direction == "D" {
		dx = -1
	} else if direction == "L" {
		dy = -1
	}

	for i := 0; i < distance; i++ {
		w.x += dx
		w.y += dy
		w.d++
		w.path = append(w.path, step{w.x, w.y, w.d})
	}
	return w
}

func intersection(a, b []step) (c []step) {
	for _, i := range a {
		for _, j := range b {
			if i.x == j.x && i.y == j.y {
				i.d += j.d
				c = append(c, i)
			}
		}
	}
	return
}

func run(filepath string) (closest int, shortest int) {
	input := utils.ReadFileToLines(filepath)
	var wires []wire
	for _, line := range input {
		paths := strings.Split(line, ",")
		w := wire{0, 0, 0, []step{}}
		for _, path := range paths {
			match := stepRegex.FindStringSubmatch(path)
			direction := match[1]
			distance := utils.Atoi(match[2])
			w = traverse(w, direction, distance)
		}
		wires = append(wires, w)
	}

	intersectionPoints := intersection(wires[0].path, wires[1].path)

	for _, p := range intersectionPoints {
		d := utils.Abs(p.x) + utils.Abs(p.y)
		if d < closest || closest == 0 {
			closest = d
		}
		l := p.d
		if l < shortest || shortest == 0 {
			shortest = l
		}
	}

	return
}

func main() {
	closest, shortest := run(os.Args[1])
	fmt.Printf("closest: %v\n shortest: %v\n", closest, shortest)
}
