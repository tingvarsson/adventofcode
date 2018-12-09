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

func determineRootParents(children map[string][]string, parents map[string][]string) (rootParents []string) {
	for parent := range children {
		if _, ok := parents[parent]; !ok {
			rootParents = append(rootParents, parent)
		}
	}
	sort.Strings(rootParents)
	return
}

func determineWorkOrder(inputFilepath string) (order string) {
	lines := readFileToLines(inputFilepath)
	children, parents := parseInstructions(lines)
	available := determineRootParents(children, parents)
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
	return
}

func all(vs []worker, f func(worker) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

type worker struct {
	currInstr int
	doneTime  int
}

func workOrder(inputFilepath string, numWorkers int, workTime int) (time int) {
	lines := readFileToLines(inputFilepath)
	children, parents := parseInstructions(lines)
	available := determineRootParents(children, parents)
	order := ""
	workers := make([]worker, numWorkers)
	for time = 0; ; time++ {
		for i, w := range workers {
			if w.currInstr != 0 && w.doneTime < time {
				order += string(w.currInstr)
				for _, new := range children[string(w.currInstr)] {
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
				workers[i] = worker{0, 0}
			}
		}

		removeBeforeIndex := -1
		for i, a := range available {
			for j, w := range workers {
				if w.currInstr == 0 {
					job := a[0]
					workers[j] = worker{int(job), time + workTime + int(job) - int('A')}
					removeBeforeIndex = i
					break
				}
			}
		}
		available = available[removeBeforeIndex+1:]

		if all(workers, func(w worker) bool {
			return w.currInstr == 0
		}) {
			fmt.Printf("Time: %d\n", time)
			return
		}
	}
}

func main() {
	determineWorkOrder("../input")
	workOrder("../input", 5, 60)
}
