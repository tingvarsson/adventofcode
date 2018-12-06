package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func reactPoly(input string) string {
	output := ""
	for _, newChar := range input {
		lastIdx := len(output) - 1
		if output != "" &&
			rune(output[lastIdx]) != newChar &&
			strings.EqualFold(string(output[lastIdx]), string(newChar)) {
			output = output[:lastIdx]
		} else {
			output += string(newChar)
		}
	}
	return output
}

func main() {
	file, err := os.Open("../input")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	scenarioOnePoly := reactPoly(lines[0])
	println(len(scenarioOnePoly))

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	reduceCharLenghts := make(map[rune]int)
	for _, l := range alphabet {
		newInput := strings.Replace(lines[0], string(l), "", -1)
		newInput = strings.Replace(newInput, strings.ToUpper(string(l)), "", -1)
		tempPoly := reactPoly(newInput)
		reduceCharLenghts[l] = len(tempPoly)
	}

	var minChar rune
	minCharLength := 10000
	for char, length := range reduceCharLenghts {
		if length < minCharLength {
			minChar = char
			minCharLength = length
		}
	}
	println(string(minChar))
	println(minCharLength)
}
