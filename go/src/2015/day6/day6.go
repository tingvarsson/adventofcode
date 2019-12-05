package main

import (
	"regexp"
	"strings"
	"utils"
)

const gridSize = 1000

type grid [gridSize][gridSize]int
type mode int

const (
	off    mode = 0
	on     mode = 1
	toggle mode = 2
)

func parseMode(s string) (m mode) {
	if strings.Contains(s, "turn on") {
		m = on
	} else if strings.Contains(s, "toggle") {
		m = toggle
	} else if strings.Contains(s, "turn off") {
		m = off
	}
	return
}

var coordRegex = regexp.MustCompile("(\\d+),(\\d+) .* (\\d+),(\\d+)")

func parseCoords(s string) (int, int, int, int) {
	m := coordRegex.FindStringSubmatch(s)
	coords := utils.StringsToInts(m[1:])
	return coords[0], coords[1], coords[2], coords[3]
}

func applyMode(prev int, m mode, brightnessControl bool) (new int) {
	if brightnessControl {
		switch m {
		case toggle:
			new = prev + 2
		case on:
			new = prev + 1
		case off:
			new = prev - 1
			if new < 0 {
				new = 0
			}
		}
	} else {
		switch m {
		case toggle:
			new = prev ^ 1
		case on, off:
			new = int(m)
		}
	}
	return
}

func run(filepath string, brightnessControl bool) (lit int) {
	var g grid

	lines := utils.ReadFileToLines(filepath)
	for _, l := range lines {
		selectedMode := parseMode(l)
		x1, y1, x2, y2 := parseCoords(l)
		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				g[y][x] = applyMode(g[y][x], selectedMode, brightnessControl)
			}
		}
	}
	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			lit += g[y][x]
		}
	}
	return
}

func main() {

}
