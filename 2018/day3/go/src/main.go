package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

	// Parse claims from input to array of maps
	var claims []map[string]int
	for _, line := range lines {
		var pattern = "#(?P<id>\\d+) @ (?P<x>\\d+),(?P<y>\\d+): (?P<xsize>\\d+)x(?P<ysize>\\d+)"
		var re = regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(line)
		var claim = make(map[string]int)
		for i, name := range re.SubexpNames() {
			if i > 0 && i <= len(match) {
				n, err := strconv.Atoi(match[i])
				if err != nil {
					log.Fatal(err)
				}
				claim[name] = n
			}
		}
		claims = append(claims, claim)
	}

	// Create fabric
	const inches = 1000
	fabric := [inches][inches]uint{}

	// Add claims
	for _, claim := range claims {
		for x := claim["x"]; x < claim["x"]+claim["xsize"]; x++ {
			for y := claim["y"]; y < claim["y"]+claim["ysize"]; y++ {
				fabric[x][y]++
			}
		}
	}

	// Sum of multiple claims
	var sumofmultiple int = 0
	for _, x := range fabric {
		for _, y := range x {
			if y >= 2 {
				sumofmultiple++
			}
		}
	}
	fmt.Println(sumofmultiple)

	// test for lonely claims
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
			fmt.Println(claim)
		}
	}
}
