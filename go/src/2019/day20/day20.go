package main

import (
	"fmt"
	"os"
	"sort"
	"unicode"
	"utils"
)

func parseArea(data []string) (area [][]rune, portals map[utils.Dim2]string) {
	portals = make(map[utils.Dim2]string)
	area = make([][]rune, len(data))
	for y, line := range data {
		area[y] = make([]rune, len(line))
		for x, r := range line {
			area[y][x] = r
		}
	}
	for y := range area {
		for x := range area[y] {
			if x < len(area[y])-1 && unicode.IsLetter(area[y][x]) && unicode.IsLetter(area[y][x+1]) {
				portal := string(area[y][x]) + string(area[y][x+1])
				if x > 0 && area[y][x-1] == '.' {
					portals[utils.Dim2{X: x, Y: y}] = portal
					area[y][x+1] = '#'
				} else if area[y][x+2] == '.' {
					portals[utils.Dim2{X: x+1, Y: y}] = portal
					area[y][x] = '#'
				}
			} else if y < len(area)-1 && unicode.IsLetter(area[y][x]) && unicode.IsLetter(area[y+1][x]) {
				portal := string(area[y][x]) + string(area[y+1][x])
				if y > 0 && area[y-1][x] == '.' {
					portals[utils.Dim2{X: x, Y: y}] = portal
					area[y+1][x] = '#'
				} else if area[y+2][x] == '.' {
					portals[utils.Dim2{X: x, Y: y+1}] = portal
					area[y][x] = '#'
				}
			}
		}
	}
	return
}

func searchEdges(area [][]rune, portals map[utils.Dim2]string, start utils.Dim2) (edges map[string]int) {
	queue := []utils.Dim2{start}
	visited := make(map[utils.Dim2]int)
	edges = make(map[string]int)
	visited[start] = 0
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		obj := area[pos.Y][pos.X]
		if pos != start && unicode.IsLetter(obj) {
			edges[portals[pos]] = visited[pos]-1
			continue
		}
		for _, d := range utils.AllDirections() {
			newPos := utils.Move(d, pos)
			if _, ok := visited[newPos]; ok {
				continue
			}
			if area[newPos.Y][newPos.X] == '#' || area[newPos.Y][newPos.X] == ' ' {
				continue
			}

			visited[newPos] = visited[pos] + 1
			queue = append(queue, newPos)
		}
	}
	return
}

func searchPath(graph map[string]map[string]int) (shortest int) {
	source := "AA"
	distance := make(map[string]int)
	distance[source] = 0
	queue := []string{source}
	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			return distance[queue[i]] < distance[queue[j]]
		})
		n := queue[0]
		queue = queue[1:]
		if n == "ZZ" {
			return distance[n]-1
		}
		for edge, dist := range graph[n] {
			if _, ok := distance[edge]; ok {
				if distance[n]+dist < distance[edge] {
					distance[edge] = distance[n] + dist
					queue = append(queue, edge)
				}
			} else {
				distance[edge] = distance[n] + dist
				queue = append(queue, edge)
			}
		}
	}
	return
}

func run(filepath string) int {
	data := utils.ReadFileToLines(filepath)
	area, portals := parseArea(data)
	edges := make(map[string]map[string]int)
	for pos, portal := range portals {
		if edges[portal] == nil {
			edges[portal] = make(map[string]int)
		}
		for k, v := range searchEdges(area, portals, pos) {
			edges[portal][k] = v
		}
	}
	return searchPath(edges)
}

func main() {
	fmt.Println(run(os.Args[1]))
}
