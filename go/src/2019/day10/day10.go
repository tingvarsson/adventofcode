package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"utils"
)

type coord struct {
	X int
	Y int
}

func (c coord) Distance(o coord) int {
	return utils.Abs(utils.GCD(c.X-o.X, c.Y-o.Y))
}

func (c coord) Normalize(o coord) coord {
	g := c.Distance(o)
	return coord{(c.X - o.X) / g, (c.Y - o.Y) / g}
}

func (c coord) Atan2(o coord) float64 {
	return math.Atan2(float64(c.X-o.X), float64(c.Y-o.Y))
}

func testCoord(asteroids []coord, a coord) int {
	var visible []coord

	for _, other := range asteroids {
		if a == other {
			continue
		}
		otherNorm := other.Normalize(a)

		blocked := false
		for _, b := range visible {
			if otherNorm == b.Normalize(a) {
				blocked = true
			}

		}
		if !blocked {
			visible = append(visible, other)
		}

	}
	return len(visible)
}

func vaporize(asteroids []coord, laser coord) (last coord) {
	sort.Slice(asteroids, func(i, j int) bool {
		return asteroids[i].Atan2(laser) > asteroids[j].Atan2(laser)
	})

	var asteroidBins [][]coord
	var bin []coord
	currentBin := asteroids[0]
	for _, a := range asteroids {
		if a == laser {
			continue
		}
		if currentBin.Normalize(laser) != a.Normalize(laser) {
			sort.Slice(bin, func(i, j int) bool {
				return bin[i].Distance(laser) < bin[j].Distance(laser)
			})
			asteroidBins = append(asteroidBins, bin)
			bin = make([]coord, 0)
			currentBin = a
		}
		bin = append(bin, a)
	}

	for i, j := 0, 0; i < 200; j = (j + 1) % len(asteroidBins) {
		if len(asteroidBins[j]) > 0 {
			last = asteroidBins[j][0]
			asteroidBins[j] = asteroidBins[j][1:]
			i++
		}
	}
	return
}

func run(filepath string) (best coord, num int, last coord) {
	data := utils.ReadFileToLines(filepath)
	var asteroids []coord
	for y := range data {
		for x, pos := range data[y] {
			if pos == '#' {
				asteroids = append(asteroids, coord{x, y})
			}
		}
	}

	for _, a := range asteroids {
		if res := testCoord(asteroids, a); res > num {
			best = a
			num = res
		}
	}

	if len(asteroids) < 200 {
		return
	}
	last = vaporize(asteroids, best)

	return
}

func main() {
	r1, r2, r3 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}
