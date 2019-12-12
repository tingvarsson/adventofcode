package main

import (
	"fmt"
	"os"
	"regexp"
	. "utils"
)

type moon struct {
	p Dim3
	v Dim3
}

func (m moon) energy() int {
	pot := Abs(m.p.X) + Abs(m.p.Y) + Abs(m.p.Z)
	kin := Abs(m.v.X) + Abs(m.v.Y) + Abs(m.v.Z)
	return pot * kin
}

func (m moon) gravity(other moon) (g Dim3) {
	g.X = 0
	if m.p.X > other.p.X {
		g.X = -1
	} else if m.p.X < other.p.X {
		g.X = 1
	}
	g.Y = 0
	if m.p.Y > other.p.Y {
		g.Y = -1
	} else if m.p.Y < other.p.Y {
		g.Y = 1
	}
	g.Z = 0
	if m.p.Z > other.p.Z {
		g.Z = -1
	} else if m.p.Z < other.p.Z {
		g.Z = 1
	}
	return
}

func (m *moon) applyGravity(g Dim3) {
	m.v = m.v.Add(g)
}

func (m *moon) applyVelocity() {
	m.p = m.p.Add(m.v)
}

var posRegex = regexp.MustCompile("<x=(-?\\d+), y=(-?\\d+), z=(-?\\d+)>")

func parseMoons(data []string) (moons []moon) {
	for _, line := range data {
		match := posRegex.FindStringSubmatch(line)
		p := Dim3{Atoi(match[1]), Atoi(match[2]), Atoi(match[3])}
		v := Dim3{0, 0, 0}
		moons = append(moons, moon{p, v})
	}
	return
}

func runIteration(moons []moon) {
	for n := range moons {
		g := Dim3{0, 0, 0}
		for _, m2 := range moons {
			if moons[n] == m2 {
				continue
			}
			g = g.Add(moons[n].gravity(m2))
		}
		moons[n].applyGravity(g)
	}

	for n := range moons {
		moons[n].applyVelocity()
	}
}

func run(filepath string, iterations int) (result, result2 int) {
	data := ReadFileToLines(filepath)
	moons := parseMoons(data)

	for i := 0; i < iterations; i++ {
		runIteration(moons)
	}

	for _, m := range moons {
		result += m.energy()
	}

	moons = parseMoons(data)

	var occurs []map[[2*4]int]int
	for i := 0; i < 3; i++ {
		occurs = append(occurs, make(map[[2*4]int]int))
	}
	reoccur := make([][]int, 3)
	for i := 0; true; i++ {
		runIteration(moons)

		var xCoords [2*4]int
		var yCoords [2*4]int
		var zCoords [2*4]int
		for n := range moons {
			xCoords[2*n] = moons[n].p.X
			xCoords[2*n+1] = moons[n].v.X
			yCoords[2*n] = moons[n].p.Y
			yCoords[2*n+1] = moons[n].v.Y
			zCoords[2*n] = moons[n].p.Z
			zCoords[2*n+1] = moons[n].v.Z
		}

		if j, ok := occurs[0][xCoords]; ok {
			if reoccur[0] == nil || reoccur[0][1] != i-j {
				reoccur[0] = []int{j, i - j}

			}
		}
		occurs[0][xCoords] = i

		if j, ok := occurs[1][yCoords]; ok {
			if reoccur[1] == nil || reoccur[1][1] != i-j {
				reoccur[1] = []int{j, i - j}
			}
		}
		occurs[1][yCoords] = i

		if j, ok := occurs[2][zCoords]; ok {
			if reoccur[2] == nil || reoccur[2][1] != i-j {
				reoccur[2] = []int{j, i - j}
			}
		}
		occurs[2][zCoords] = i

		foundAll := true
		for _, r := range reoccur {
			if r == nil {
				foundAll = false
			}
		}
		if foundAll {
			result2 = LCM(reoccur[0][1], reoccur[1][1], reoccur[2][1])
			return
		}
	}

	return
}

func main() {
	r1, r2 := run(os.Args[1], 0)
	fmt.Println(r1)
	fmt.Println(r2)
}
