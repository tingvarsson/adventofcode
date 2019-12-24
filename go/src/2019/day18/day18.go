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

func searchPath(graph map[rune]map[rune]int, availableKeys string) (shortest int) {
	start := node{'@', ""}
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
		if len(n.keys) == len(availableKeys) {
			if shortest == 0 || distance[n] < shortest {
				shortest = distance[n]
			}
			continue
		}
		for neighbor, dist := range graph[n.obj] {
			if unicode.IsUpper(neighbor) {
				neededKey := unicode.ToLower(neighbor)
				if strings.ContainsRune(availableKeys, neededKey) && !strings.ContainsRune(n.keys, neededKey) {
					continue
				}
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
return
}

func run(filepath string) int {
	data := utils.ReadFileToLines(filepath)
	area, objs := parseArea(data)
	neighbors := make(map[rune]map[rune]int)
	var keys string
	for pos, obj := range objs {
		neighbors[obj] = searchNeighbors(area, pos)
		if unicode.IsLower(obj) {
			keys += string(obj)
		}
	}
	return searchPath(neighbors, keys)
}

type quadRange struct {
	minX, minY, maxX, maxY int
}

func parseQuads(data []string) (quads [][][]rune, quadObjs []map[utils.Dim2]rune) {
	area, objs := parseArea(data)
	halfX := len(area[0])/2
	fullX := len(area[0])
	halfY := len(area)/2
	fullY := len(area)
	ranges := []quadRange{{0, 0, halfX+1, halfY+1},{0, halfY, halfX+1, fullY},{halfX, 0, fullX, halfY+1},{halfX, halfY, fullX, fullY}}
	for _, r := range ranges {
		quad := make([][]rune, halfY+1)
		for y := r.minY; y < r.maxY; y++ {
			quad[y-r.minY] = make([]rune, halfX+1)
			for x := r.minX; x < r.maxX; x++ {
				quad[y-r.minY][x-r.minX] = area[y][x]
			}
		}
		quads = append(quads, quad)
		qObjs := make(map[utils.Dim2]rune) 
		for p, o := range objs {
			if p.X > r.minX && p.X < r.maxX && p.Y > r.minY && p.Y < r.maxY {
				p.X -= r.minX
				p.Y -= r.minY
				qObjs[p] = o
			}
		}
		quadObjs = append(quadObjs, qObjs)
	}
	return
}

func run2(filepath string) (result int) {
	data := utils.ReadFileToLines(filepath)
	quads, quadObjs := parseQuads(data)
	for i := range quads {
		neighbors := make(map[rune]map[rune]int)
		var keys string
		for pos, obj := range quadObjs[i] {
			neighbors[obj] = searchNeighbors(quads[i], pos)
			if unicode.IsLower(obj) {
				keys += string(obj)
			}
		}
		result += searchPath(neighbors, keys)
	}
	return
}

func main() {
	res := run(os.Args[1])
	fmt.Println(res)
}
