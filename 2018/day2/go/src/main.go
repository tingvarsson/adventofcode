package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func hasValue(s map[rune]int, v int) bool {
	for k := range s {
		if s[k] == v {
			return true
		}
	}
	return false
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

	sumOfHasDoublet := 0
	sumOfHasTriplet := 0
	for _, line := range lines {
		runeMap := make(map[rune]int)
		for _, c := range line {
			runeMap[c]++
		}

		if hasValue(runeMap, 2) {
			sumOfHasDoublet++
		}
		if hasValue(runeMap, 3) {
			sumOfHasTriplet++
		}
	}
	fmt.Printf("checksum: %d\n", sumOfHasDoublet*sumOfHasTriplet)

	for i, line := range lines {
		for _, secondline := range lines[i:] {
			out := ""
			for k, c := range line {
				if c == rune(secondline[k]) {
					out += string(c)
				}
			}
			if len(line)-1 == len(out) {
				fmt.Printf("same line except one char: %s", out)
			}
		}
	}
}
