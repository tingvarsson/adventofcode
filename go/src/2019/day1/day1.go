package main

import (
	"fmt"
	"os"
	"utils"
)

func calculateFuel(mass int) int {
	return int(mass/3) - 2
}
func recursiveFuel(mass int, fuel int) int {
	f := calculateFuel(mass)

	if f <= 0 {
		return fuel
	}

	return recursiveFuel(f, fuel+f)
}

func run(filepath string) (sumFuel int, sumRecursiveFuel int) {
	input := utils.ReadFileToLines(filepath)
	var modules []int
	for _, line := range input {
		modules = append(modules, utils.Atoi(line))
	}

	for _, mass := range modules {
		sumFuel += calculateFuel(mass)
		sumRecursiveFuel += recursiveFuel(mass, 0)
	}
	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
