package main

import (
	"fmt"
)

const gridSize = 301 // Adjusted for indexing 1..300

func calculatePowerLevel(x int, y int, serialNumber int) int {
	rackID := x + 10
	initial := rackID * y
	adjusted := (initial + serialNumber) * rackID
	reduced := ((adjusted / 100) % 10) - 5
	return reduced
}

func newGrid(serialNumber int) (grid [][]int) {
	grid = make([][]int, gridSize)
	for y := 1; y < gridSize; y++ {
		grid[y] = make([]int, gridSize)
		for x := 1; x < gridSize; x++ {
			grid[y][x] = calculatePowerLevel(x, y, serialNumber)
		}
	}
	return
}

func newEmptyGrid() (grid [][]int) {
	grid = make([][]int, gridSize)
	for y := range grid {
		grid[y] = make([]int, gridSize)
		for x := range grid[y] {
			grid[y][x] = 0
		}
	}
	return
}

func sumSquare(grid [][]int, sumGrid [][]int, startX int, startY int, size int) (sum int) {
	sum = sumGrid[startY][startX]
	for y := startY; y < startY+size; y++ {
		sum += grid[y][startX+size-1]
	}
	for x := startX; x < startX+size; x++ {
		sum += grid[startY+size-1][x]
	}
	sumGrid[startY][startX] = sum
	return
}

func findLargestSquare(serialNumber int, minSize int, maxSize int) (largest [4]int) {
	grid := newGrid(serialNumber)
	sumGrid := newEmptyGrid()
	for size := 1; size <= maxSize; size++ {
		for y := 1; y < gridSize-(size-1); y++ {
			for x := 1; x < gridSize-(size-1); x++ {
				newSum := sumSquare(grid, sumGrid, x, y, size)
				if size >= minSize && newSum > largest[0] {
					largest = [4]int{newSum, x, y, size}
				}
			}
		}
	}
	return
}

func main() {
	fmt.Printf("%v\n", findLargestSquare(6548, 3, 3))
	fmt.Printf("%v\n", findLargestSquare(6548, 1, 300))
}
