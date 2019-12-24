package main

import (
	"fmt"
	"os"
	"program"
	"strings"
	"utils"
)

func parseIntcode(file string) (intcode []int) {
	input := utils.ReadFileToString(file)
	inputData := strings.Split(input, ",")
	for _, code := range inputData {
		intcode = append(intcode, utils.Atoi(code))
	}
	return
}

func run(filepath string) (result, result2 int) {
	intcode := parseIntcode(filepath)

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			p := program.New(intcode)
			p.Run([]int{x, y})
			result += p.PopOutput()
		}
	}

	prevX := 0
	for y := 100; true; y++ {
		for x := prevX; true; x++ {
			p := program.New(intcode)
			p.Run([]int{x, y})
			found := p.PopOutput()
			if found == 0 {
				continue
			}
			prevX = x
			p = program.New(intcode)
			p.Run([]int{x+99, y-99})
			foundOpposite := p.PopOutput()
			if foundOpposite == 1 {
				result2 = x * 10000 + y-99
				return
			}
			break
		}
	}
	return
}

func main() {
	fmt.Println(run(os.Args[1]))
}
