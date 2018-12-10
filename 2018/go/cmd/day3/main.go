package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"utils"
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

	var claims []map[string]int
	pattern := "#(?P<id>\\d+) @ (?P<x>\\d+),(?P<y>\\d+): (?P<xsize>\\d+)x(?P<ysize>\\d+)"
	re := regexp.MustCompile(pattern)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		var claim = make(map[string]int)
		for i, name := range re.SubexpNames() {
			if i > 0 && i <= len(match) {
				claim[name] = utils.Atoi(match[i])
			}
		}
		claims = append(claims, claim)
	}

	const inches = 1000
	fabric := [inches][inches]uint{}

	for _, claim := range claims {
		for x := claim["x"]; x < claim["x"]+claim["xsize"]; x++ {
			for y := claim["y"]; y < claim["y"]+claim["ysize"]; y++ {
				fabric[x][y]++
			}
		}
	}

	var sumOfMultiple int
	for _, x := range fabric {
		for _, y := range x {
			if y >= 2 {
				sumOfMultiple++
			}
		}
	}
	fmt.Printf("Number of square inches within multiple claims: %d\n", sumOfMultiple)

	for _, claim := range claims {
		var found = true
		for x := claim["x"]; x < claim["x"]+claim["xsize"]; x++ {
			for y := claim["y"]; y < claim["y"]+claim["ysize"]; y++ {
				if fabric[x][y] >= 2 {
					found = false
				}
			}
		}
		if found {
			fmt.Printf("Claim that doesn't overlap: %v\n", claim)
		}
	}
}
