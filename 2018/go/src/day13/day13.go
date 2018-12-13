package main

import (
	"fmt"
	"sort"
	"utils"
)

type cart struct {
	x        int
	y        int
	dir      rune
	nextTurn rune
	crashed  bool
}

var cartStartTrack = map[rune]rune{
	'^': '|', '<': '-', 'v': '|', '>': '-',
}
var nextTurn = map[rune]rune{
	'l': 's', 's': 'r', 'r': 'l',
}
var directionOffset = map[rune][]int{
	'^': {0, -1}, '<': {-1, 0}, 'v': {0, 1}, '>': {1, 0},
}
var directionChange = map[rune]map[rune]rune{
	'^': {'/': '>', '|': '^', '\\': '<', '-': 'X'},
	'v': {'/': '<', '|': 'v', '\\': '>', '-': 'X'},
	'<': {'/': 'v', '|': 'X', '\\': '^', '-': '<'},
	'>': {'/': '^', '|': 'X', '\\': 'v', '-': '>'},
}
var intersectionDirection = map[rune]map[rune]rune{
	'^': {'l': '<', 's': '^', 'r': '>'},
	'v': {'l': '>', 's': 'v', 'r': '<'},
	'<': {'l': 'v', 's': '<', 'r': '^'},
	'>': {'l': '^', 's': '>', 'r': 'v'},
}

func parseInputToTracksAndCarts(lines []string) (tracks [][]rune, carts []cart) {
	for y, l := range lines {
		var row []rune
		for x, r := range l {
			if startTrack, ok := cartStartTrack[r]; ok {
				c := cart{x, y, r, 'l', false}
				carts = append(carts, c)
				row = append(row, startTrack)
			} else {
				row = append(row, r)
			}
		}
		tracks = append(tracks, row)
	}
	return
}

func find(carts []cart, x int, y int) int {
	for i, c := range carts {
		if c.x == x && c.y == y {
			return i
		}
	}
	return -1
}

func run(tracks [][]rune, carts []cart) {
	for len(carts) > 1 {
		sort.Slice(carts, func(i, j int) bool {
			if carts[i].y != carts[j].y {
				return carts[i].y < carts[j].y
			}
			return carts[i].x < carts[j].x
		})

		for ci, c := range carts {
			if c.crashed {
				continue
			}
			newX := c.x + directionOffset[c.dir][0]
			newY := c.y + directionOffset[c.dir][1]
			if i := find(carts, newX, newY); i != -1 {
				fmt.Printf("BOOM! @ %d,%d\n", newX, newY)
				carts[i].crashed = true
				carts[ci].crashed = true
			} else { // move the cart
				carts[ci].x = newX
				carts[ci].y = newY
				if t := tracks[newY][newX]; t == '+' {
					carts[ci].dir = intersectionDirection[c.dir][c.nextTurn]
					carts[ci].nextTurn = nextTurn[c.nextTurn]
				} else {
					carts[ci].dir = directionChange[c.dir][t]
				}
			}
		}

		i := 0
		for _, c := range carts {
			if !c.crashed {
				carts[i] = c
				i++
			}
		}
		carts = carts[:i]

		if len(carts) == 1 {
			fmt.Printf("Last cart: %v\n", carts[0])
		}
	}
}

func main() {
	lines := utils.ReadFileToLines("day13/input")
	tracks, carts := parseInputToTracksAndCarts(lines)
	run(tracks, carts)
}
