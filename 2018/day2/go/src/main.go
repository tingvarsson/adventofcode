package main

import (
	"bufio"
	"log"
	"os"
)

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

	sum2 := 0
	sum3 := 0
	for _, line := range lines {
		m := make(map[rune]int)
		for _, c := range line {
			m[c]++
		}

		found2 := false
		found3 := false
		for _, count := range m {
			if !found2 && count == 2 {
				found2 = true
				sum2++
			} else if !found3 && count == 3 {
				found3 = true
				sum3++
			}
		}
	}
	println("checksum:")
	println(sum2 * sum3)
	for i, line := range lines {
		for _, secondline := range lines[i:] {
			out := ""
			for k, c := range line {
				if c == rune(secondline[k]) {
					out += string(c)
				}
			}
			if len(line)-1 == len(out) {
				println("same line:")
				println(out)
			}
		}
	}
}
