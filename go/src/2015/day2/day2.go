package main

import (
	"fmt"
	"regexp"
	"sort"
	"utils"
)

type box struct {
	dim []int
}

var boxRegex = regexp.MustCompile("(\\d+)x(\\d+)x(\\d+)")

func newBox(s string) box {
	match := boxRegex.FindStringSubmatch(s)
	dim := []int{utils.Atoi(match[1]),
		utils.Atoi(match[2]),
		utils.Atoi(match[3])}
	sort.Ints(dim)
	return box{dim}
}

func (b box) minPerimeter() int {
	return 2*b.dim[0] + 2*b.dim[1]
}

func (b box) minSurface() int {
	return b.dim[0] * b.dim[1]
}

func (b box) totalSurface() int {
	return 2*b.dim[0]*b.dim[1] + 2*b.dim[1]*b.dim[2] + 2*b.dim[2]*b.dim[0]
}

func (b box) volume() int {
	return b.dim[0] * b.dim[1] * b.dim[2]
}

func run(filepath string) (wrappingPaperNeeded int, ribbonNeeded int) {
	input := utils.ReadFileToLines(filepath)
	var boxes []box
	for _, line := range input {
		boxes = append(boxes, newBox(line))
	}

	for _, b := range boxes {
		wrappingPaperNeeded += b.totalSurface() + b.minSurface()
		ribbonNeeded += b.minPerimeter() + b.volume()
	}
	return
}

func main() {
	wrapping, ribbon := run("day2/input")
	fmt.Printf("Wrapping: %v Ribbon: %v\n", wrapping, ribbon)
}
