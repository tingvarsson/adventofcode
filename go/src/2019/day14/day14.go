package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"utils"
)

const cargoCapacity = 1000000000000

type chem struct {
	name string
	n    int
}

type reactionMap map[string][]chem

type wasteMap map[string]int

var reactionRegex = regexp.MustCompile("(.*) => (\\d+) (\\w+)")
var chemicalRegex = regexp.MustCompile("(\\d+) (\\w+)")

func parseReactions(data []string) (reactions reactionMap) {
	reactions = make(reactionMap)
	for _, line := range data {
		rMatch := reactionRegex.FindStringSubmatch(line)
		inData := strings.Split(rMatch[1], ",")
		outCnt := utils.Atoi(rMatch[2])
		out := rMatch[3]
		o := chem{out, outCnt}
		reactions[out] = append(reactions[out], o)
		for _, inLine := range inData {
			cMatch := chemicalRegex.FindStringSubmatch(inLine)
			inCnt := utils.Atoi(cMatch[1])
			in := cMatch[2]
			i := chem{in, inCnt}
			reactions[out] = append(reactions[out], i)
		}
	}
	return
}

func oreForChem(r reactionMap, w wasteMap, c string, n int) (ore int) {
	if c == "ORE" {
		return n
	}

	avail := w[c]
	if avail >= n {
		w[c] = avail - n
		return
	}

	n = n - avail
	o := r[c]
	mult := n / o[0].n
	if n%o[0].n != 0 {
		mult++
	}
	w[c] = mult*o[0].n - n

	for i := 1; i < len(o); i++ {
		ore += oreForChem(r, w, o[i].name, mult*o[i].n)
	}
	return
}

func oreForFuel(reactions reactionMap, nFuel int) (nOre int) {
	waste := make(wasteMap)
	return oreForChem(reactions, waste, "FUEL", nFuel)
}

func run(filepath string) (result, result2 int) {
	data := utils.ReadFileToLines(filepath)
	reactions := parseReactions(data)
	result = oreForFuel(reactions, 1)

	low := cargoCapacity / result
	high := low * 2
	for low != high {
		mid := 1 + (low+high-1)/2
		ore := oreForFuel(reactions, mid)
		if ore > cargoCapacity {
			high = mid - 1
		} else {
			low = mid
		}
	}
	result2 = low
	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
