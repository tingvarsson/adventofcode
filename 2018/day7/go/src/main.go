package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func readFileToLines(filepath string) (lines []string) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

func parseInstructions(instrList []string) (children map[string][]string, parents map[string][]string) {
	children = make(map[string][]string)
	parents = make(map[string][]string)
	p := "Step (.) .* step (.)"
	re := regexp.MustCompile(p)
	for _, instr := range instrList {
		match := re.FindStringSubmatch(instr)
		lhs := match[1]
		rhs := match[2]
		children[lhs] = append(children[lhs], rhs)
		parents[rhs] = append(parents[rhs], lhs)
	}
	return
}

func main() {
	lines := readFileToLines("../input")
	children, parents := parseInstructions(lines)

	var rootParent string
	for parent := range children {
		if _, ok := parents[parent]; !ok {
			rootParent = parent
		}
	}

	available := []string{rootParent}
	order := ""
	for len(available) != 0 {
		order += available[0]
		available = available[1:]
		for _, new := range children[string(order[len(order)-1])] {
			allParentsInOrder := true
			for _, parent := range parents[new] {
				if !strings.Contains(order, parent) {
					allParentsInOrder = false
				}
			}
			if allParentsInOrder {
				available = append(available, new)
			}
		}
		sort.Strings(available)
	}
	fmt.Printf("Theoretical order: %s\n", order)

	available = []string{rootParent}
	order = ""
	var workers [5][2]int
	workOngoing := true
	time := 0
	for time = 0; workOngoing; time++ {
		for i, w := range workers {
			if w[0] != 0 && w[1] < time {
				order += string(w[0])
				for _, new := range children[string(w[0])] {
					allParentsInOrder := true
					for _, parent := range parents[new] {
						if !strings.Contains(order, parent) {
							allParentsInOrder = false
						}
					}
					if allParentsInOrder {
						available = append(available, new)
					}
				}
				sort.Strings(available)
				workers[i] = [2]int{0, 0}
			}
		}

		removeBeforeIndex := -1
		for i, a := range available {
			for j, w := range workers {
				if w[0] == 0 {
					job := a[0]
					workers[j] = [2]int{int(job), time + 60 + int(job) - int('A')}
					removeBeforeIndex = i
					break
				}
			}
		}
		available = available[removeBeforeIndex+1:]

		workOngoing = false
		for _, w := range workers {
			if w[0] != 0 {
				workOngoing = true
				break
			}
		}
	}
	fmt.Printf("Time: %d\n", time-1)
}
