package main

import (
	"fmt"
	"os"
	"utils"
)

func createImage(pixels []int) (image [][][]int) {
	width := 25
	height := 6
	i := 0
	for i < len(pixels) {
		layer := make([][]int, height)
		for y := 0; y < height; y++ {
			layer[y] = make([]int, width)
			for x := 0; x < width; x++ {
				layer[y][x] = pixels[i]
				i++
			}
		}
		image = append(image, layer)
	}
	return
}

func countDigits(layer [][]int) (dict map[int]int) {
	dict = make(map[int]int)
	for _, row := range layer {
		for _, n := range row {
			dict[n] = dict[n] + 1
		}
	}
	return
}

func calcChecksum(image [][][]int) int {
	var fewest0 map[int]int
	for _, layer := range image {
		count := countDigits(layer)
		if fewest0 == nil || count[0] < fewest0[0] {
			fewest0 = count
		}
	}
	return fewest0[1] * fewest0[2]
}

func renderImage(image [][][]int) {
	width := len(image[0][0])
	height := len(image[0])
	output := make([][]int, height)
	for y := range output {
		output[y] = make([]int, width)
		for _, layer := range image {
			for x := range output[y] {
				if output[y][x] == 0 && layer[y][x] != 2 {
					output[y][x] = layer[y][x] + 1
				}
			}
		}
	}

	for row := range output {
		fmt.Println(output[row])
	}
}

func run(filepath string) (result, result2 int) {
	data := utils.ReadFileToString(filepath)
	pixels := utils.AtoiSlice(data)
	image := createImage(pixels)

	result = calcChecksum(image)
	renderImage(image)
	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
