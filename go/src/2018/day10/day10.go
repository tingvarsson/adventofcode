package main

import (
	"fmt"
	"os"
	"regexp"
	"utils"
)

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func parseInput(filepath string) (points []point) {
	lines := utils.ReadFileToLines(filepath)
	pattern := "position=< *(-?\\d+), *(-?\\d+)> velocity=< *(-?\\d+), *(-?\\d)>"
	regex := regexp.MustCompile(pattern)
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		p := point{utils.Atoi(match[1]),
			utils.Atoi(match[2]),
			utils.Atoi(match[3]),
			utils.Atoi(match[4])}
		points = append(points, p)
	}
	return
}

func minCoords(points []point) (minX int, minY int) {
	minX = points[0].x
	minY = points[0].y
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
	}
	return
}

func maxCoords(points []point) (maxX int, maxY int) {
	maxX = points[0].x
	maxY = points[0].y
	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return
}

func normalizePoints(points []point) {
	minX, minY := minCoords(points)
	for i := range points {
		points[i].x -= minX
		points[i].y -= minY
	}
}

func printSky(points []point) {
	maxX, maxY := maxCoords(points)
	maxX++
	maxY++
	sky := make([][]string, maxY)
	for y := 0; y < maxY; y++ {
		sky[y] = make([]string, maxX)
		for x := 0; x < maxX; x++ {
			sky[y][x] = "."
		}
	}
	for _, p := range points {
		sky[p.y][p.x] = "#"
	}
	for y := 0; y < maxY; y++ {
		output := ""
		for x := 0; x < maxX; x++ {
			output += sky[y][x]
		}
		println(output)
	}
}

func minimizeBoundingBox(points []point) (time int) {
	normalizePoints(points)
	maxX, maxY := maxCoords(points)
	minBoundingBox := maxX * maxY
	for ; ; time++ {
		for i := range points {
			points[i].x += points[i].vx
			points[i].y += points[i].vy
		}
		normalizePoints(points)
		maxX, maxY := maxCoords(points)
		newBoundingBox := maxX * maxY
		if newBoundingBox <= minBoundingBox {
			minBoundingBox = newBoundingBox
		} else { // go back to previous and return result
			for i := range points {
				points[i].x -= points[i].vx
				points[i].y -= points[i].vy
			}
			normalizePoints(points)
			printSky(points)
			return
		}
	}
}

func main() {
	points := parseInput(os.Args[1])
	fmt.Printf("time: %d", minimizeBoundingBox(points))
}
