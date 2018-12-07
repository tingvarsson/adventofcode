package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func readFileToLines(filepath string) []string {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(array []int) (index int, value int) {
	index = 0
	value = 0
	for i, v := range array {
		if v > value {
			index = i
			value = v
		}
	}
	return
}

func main() {
	lines := readFileToLines("../input")

	var xCoords []int
	var yCoords []int
	var coordArea []int

	const coordPattern = "(\\d+), (\\d+)"
	coordReg := regexp.MustCompile(coordPattern)
	for _, line := range lines {
		coordMatch := coordReg.FindStringSubmatch(line)
		if coordMatch != nil {
			if n, err := strconv.Atoi(coordMatch[1]); err != nil {
				log.Fatal(err)
			} else {
				xCoords = append(xCoords, n)
			}
			if n, err := strconv.Atoi(coordMatch[2]); err != nil {
				log.Fatal(err)
			} else {
				yCoords = append(yCoords, n)
			}
			coordArea = append(coordArea, 0)
		}
	}

	_, xMax := max(xCoords)
	_, yMax := max(yCoords)

	closeToAllArea := 0
	for x := 0; x < xMax+1; x++ {
		for y := 0; y < yMax+1; y++ {
			var closestCoord int
			closestCoordDistance := math.MaxInt64
			coordDistanceSum := 0
			for i := 0; i < len(coordArea); i++ {
				distance := abs(x-xCoords[i]) + abs(y-yCoords[i])
				if distance < closestCoordDistance {
					closestCoord = i
					closestCoordDistance = distance
				} else if distance == closestCoordDistance {
					closestCoord = -1 // No ONE is closest
				}
				coordDistanceSum += distance
			}

			if coordDistanceSum < 10000 {
				closeToAllArea++
			}

			if closestCoord == -1 || coordArea[closestCoord] == -1 {
				continue
			}

			if x == 0 || x == xMax || y == 0 || y == yMax {
				coordArea[closestCoord] = -1
			} else {
				coordArea[closestCoord]++
			}
		}
	}
	maxIndex, maxArea := max(coordArea)
	fmt.Printf("%d %d %d %d\n", maxIndex, xCoords[maxIndex], yCoords[maxIndex], maxArea)
	fmt.Printf("%d\n", closeToAllArea)
}
