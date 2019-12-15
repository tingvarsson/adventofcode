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

	p := program.New(intcode)
	p.Run([]int{1})
	o := p.GetOutput()
	result = o[len(o)-1]

	p = program.New(intcode)
	p.Run([]int{5})
	result2 = p.PopOutput()

	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
