package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"utils"
)

func main() {
	lines := utils.ReadFileToLines("day6/input")

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

	xMax := utils.MaxValue(xCoords)
	yMax := utils.MaxValue(yCoords)

	closeToAllArea := 0
	for x := 0; x < xMax+1; x++ {
		for y := 0; y < yMax+1; y++ {
			var closestCoord int
			closestCoordDistance := math.MaxInt64
			coordDistanceSum := 0
			for i := 0; i < len(coordArea); i++ {
				distance := utils.Abs(x-xCoords[i]) + utils.Abs(y-yCoords[i])
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
	maxIndex, maxArea := utils.Max(coordArea)
	fmt.Printf("%d %d %d %d\n", maxIndex, xCoords[maxIndex], yCoords[maxIndex], maxArea)
	fmt.Printf("%d\n", closeToAllArea)
}
