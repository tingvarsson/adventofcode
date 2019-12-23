package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
	"utils"
)

func parseArea(data []string) (area [][]rune, objs map[utils.Dim2]rune) {
	objs = make(map[utils.Dim2]rune)
	area = make([][]rune, len(data))
	for y, line := range data {
		area[y] = make([]rune, len(line))
		for x, r := range line {
			area[y][x] = r
			if unicode.IsLetter(r) || r == '@' {
				objs[utils.Dim2{X: x, Y: y}] = r
			}
		}
	}
	return
}

func searchNeighbors(area [][]rune, start utils.Dim2) (distance map[rune]int) {
	queue := []utils.Dim2{start}
	discovered := make(map[utils.Dim2]int)
	distance = make(map[rune]int)
	discovered[start] = 0
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		obj := area[pos.Y][pos.X]
		if pos != start && (unicode.IsLetter(obj) || obj == '@') {
			distance[obj] = discovered[pos]
			continue
		}
		for _, d := range utils.AllDirections() {
			newPos := utils.Move(d, pos)
			if _, ok := discovered[newPos]; ok {
				continue
			}
			if area[newPos.Y][newPos.X] == '#' {
				continue
			}

			discovered[newPos] = discovered[pos] + 1
			queue = append(queue, newPos)
		}
	}
	return
}

type node struct {
	obj  rune
	keys string
}

func searchPath(graph map[rune]map[rune]int, start node, endCount int) (shortest int) {
	distance := make(map[node]int)
	distance[start] = 0
	queue := []node{start}
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			if len(queue[i].keys) != len(queue[j].keys) {
				return len(queue[i].keys) > len(queue[j].keys)
			}
			return distance[queue[i]] < distance[queue[j]]
		})
		n := queue[0]
		queue = queue[1:]
		if len(n.keys) == endCount {
			continue
		}
		for neighbor, dist := range graph[n.obj] {
			if unicode.IsUpper(neighbor) && !strings.ContainsRune(n.keys, unicode.ToLower(neighbor)) {
				continue
			}

			keys := n.keys
			if unicode.IsLower(neighbor) && !strings.ContainsRune(n.keys, neighbor) {
				keys = utils.SortStringByCharacter(keys)
				keys += string(neighbor)
			}
			newNode := node{neighbor, keys}
			if _, ok := distance[newNode]; ok {
				if distance[n]+dist < distance[newNode] {
					distance[newNode] = distance[n] + dist
					queue = append(queue, newNode)
				}
			} else {
				distance[newNode] = distance[n] + dist
				queue = append(queue, newNode)
			}
		}
	}
	shortest = 1000000
	for n, d := range distance {
		if len(n.keys) == endCount && d < shortest {
			shortest = d
		}
	}
	return
}

func run(filepath string) int {
	data := utils.ReadFileToLines(filepath)
	area, objs := parseArea(data)
	neighbors := make(map[rune]map[rune]int)
	numKeys := 0
	for pos, obj := range objs {
		neighbors[obj] = searchNeighbors(area, pos)
		if unicode.IsLower(obj) {
			numKeys++
		}
	}

	start := node{'@', ""}
	return searchPath(neighbors, start, numKeys)
}

func searchPath2(graph map[rune]map[rune]int, start node, endCount int) (shortest int) {
	distance := make(map[node]int)
	distance[start] = 0
	queue := []node{start}
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			if len(queue[i].keys) != len(queue[j].keys) {
				return len(queue[i].keys) > len(queue[j].keys)
			}
			return distance[queue[i]] < distance[queue[j]]
		})
		n := queue[0]
		queue = queue[1:]
		if len(n.keys) == endCount {
			continue
		}
		for neighbor, dist := range graph[n.obj] {
			if unicode.IsUpper(neighbor) && !strings.ContainsRune(n.keys, unicode.ToLower(neighbor)) {
				continue
			}

			keys := n.keys
			if unicode.IsLower(neighbor) && !strings.ContainsRune(n.keys, neighbor) {
				keys = utils.SortStringByCharacter(keys)
				keys += string(neighbor)
			}
			newNode := node{neighbor, keys}
			if _, ok := distance[newNode]; ok {
				if distance[n]+dist < distance[newNode] {
					distance[newNode] = distance[n] + dist
					queue = append(queue, newNode)
				}
			} else {
				distance[newNode] = distance[n] + dist
				queue = append(queue, newNode)
			}
		}
	}
	shortest = 1000000
	for n, d := range distance {
		if len(n.keys) == endCount && d < shortest {
			shortest = d
		}
	}
	return
}

func run2(filepath string) int {
	data := utils.ReadFileToLines(filepath)
	area, objs := parseArea(data)
	neighbors := make(map[rune]map[rune]int)
	numRobots := 0
	numKeys := 0
	for pos, obj := range objs {
		if obj == '@' {
			numRobots++
			obj = rune(numRobots + int('0'))
			objs[pos] = obj
			area[pos.Y][pos.X] = obj
		}
		neighbors[obj] = searchNeighbors(area, pos)
		if unicode.IsLower(obj) {
			numKeys++
		}
	}

	start := node{'@', ""}
	return searchPath2(neighbors, start, numKeys)
}

func main() {
	res := run(os.Args[1])
	fmt.Println(res)
}
