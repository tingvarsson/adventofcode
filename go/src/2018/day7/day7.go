package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"utils"
)

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
	lines := utils.ReadFileToLines(inputFilepath)
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

type worker struct {
	currInstr int
	doneTime  int
}

func all(workers []worker, f func(worker) bool) bool {
	for _, v := range workers {
		if !f(v) {
			return false
		}
	}
	return true
}

func workOrder(inputFilepath string, numWorkers int, workTime int) (time int) {
	lines := utils.ReadFileToLines(inputFilepath)
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
	determineWorkOrder(os.Args[1])
	workOrder(os.Args[1], 5, 60)
}
