package main

import (
	"utils"
)

func run(filepath string, numSantas int) (housesCovered int) {
	input := utils.ReadFileToString(filepath)
	gridSize := len(input) + 10
	houseGrid := make([][]int, gridSize)
	for i := range houseGrid {
		houseGrid[i] = make([]int, gridSize)
	}

	var santas [][]int
	for i := 0; i < numSantas; i++ {
		santas = append(santas, []int{gridSize / 2, gridSize / 2})
	}

	currentSanta := 0
	houseGrid[santas[currentSanta][1]][santas[currentSanta][0]]++
	for _, r := range input {
		switch r {
		case '^':
			santas[currentSanta][1]--
		case '>':
			santas[currentSanta][0]++
		case 'v':
			santas[currentSanta][1]++
		case '<':
			santas[currentSanta][0]--
		}
		houseGrid[santas[currentSanta][1]][santas[currentSanta][0]]++
		currentSanta = (currentSanta + 1) % numSantas
	}

	for y := range houseGrid {
		for x := range houseGrid[y] {
			if houseGrid[y][x] > 0 {
				housesCovered++
			}
		}
	}
	return housesCovered
}

func main() {

}
