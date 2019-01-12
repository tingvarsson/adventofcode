package main

import (
	"strings"
	"utils"
)

func run(filepath string) (sumNonVisibleChar int) {
	lines := utils.ReadFileToLines(filepath)

	for _, l := range lines {
		sumNonVisibleChar += strings.Count(l, "\\\\")
		l = strings.Replace(l, "\\\\", "", -1)
		sumNonVisibleChar += strings.Count(l, "\\\"")
		l = strings.Replace(l, "\\\"", "", -1)
		sumNonVisibleChar += strings.Count(l, "\"")
		l = strings.Replace(l, "\"", "", -1)
		sumNonVisibleChar += strings.Count(l, "\\x") * 3
	}
	return sumNonVisibleChar
}

func run2(filepath string) (sumEncodeChar int) {
	lines := utils.ReadFileToLines(filepath)

	for _, l := range lines {
		sumEncodeChar += strings.Count(l, "\"")
		sumEncodeChar += strings.Count(l, "\\")
		sumEncodeChar += 2
	}
	return sumEncodeChar
}

func main() {}
