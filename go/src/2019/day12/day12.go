package main

import (
	"fmt"
	"os"
	"regexp"
	"utils"
)

type pos struct {
	x, y, z int
}

func addPos(a, b pos) (c pos) {
	c.x = a.x + b.x
	c.y = a.y + b.y
	c.z = a.z + b.z
	return
}

type vel struct {
	x, y, z int
}

func addVel(a, b vel) (c vel) {
	c.x = a.x + b.x
	c.y = a.y + b.y
	c.z = a.z + b.z
	return
}

type moon struct {
	p pos
	v vel
}

func (m moon) energy() int {
	pot := utils.Abs(m.p.x) + utils.Abs(m.p.y) + utils.Abs(m.p.z)
	kin := utils.Abs(m.v.x) + utils.Abs(m.v.y) + utils.Abs(m.v.z)
	return pot * kin
}

func (m moon) gravity(other moon) (g vel) {
	g.x = 0
	if m.p.x > other.p.x {
		g.x = -1
	} else if m.p.x < other.p.x {
		g.x = 1
	}
	g.y = 0
	if m.p.y > other.p.y {
		g.y = -1
	} else if m.p.y < other.p.y {
		g.y = 1
	}
	g.z = 0
	if m.p.z > other.p.z {
		g.z = -1
	} else if m.p.z < other.p.z {
		g.z = 1
	}
	return
}

func (m *moon) applyGravity(g vel) {
	m.v = addVel(m.v, g)
}

func (m *moon) applyVelocity() {
	m.p = addPos(m.p, pos{m.v.x, m.v.y, m.v.z})
}

var posRegex = regexp.MustCompile("<x=(-?\\d+), y=(-?\\d+), z=(-?\\d+)>")

func run(filepath string, iterations int) (result, result2 int) {
	data := utils.ReadFileToLines(filepath)

	var moons []moon
	for _, line := range data {
		match := posRegex.FindStringSubmatch(line)
		p := pos{utils.Atoi(match[1]), utils.Atoi(match[2]), utils.Atoi(match[3])}
		moons = append(moons, moon{p, vel{0, 0, 0}})
	}

	for i := 0; i < iterations; i++ {
		var grav []vel
		for n := range moons {
			g := vel{0, 0, 0}
			for _, m2 := range moons {
				if moons[n] == m2 {
					continue
				}
				g = addVel(g, moons[n].gravity(m2))
			}
			grav = append(grav, g)
		}

		for n := range moons {
			moons[n].applyGravity(grav[n])
			moons[n].applyVelocity()
		}
	}

	for _, m := range moons {
		result += m.energy()
	}

	return
}

func main() {
	r1, r2 := run(os.Args[1], 0)
	fmt.Println(r1)
	fmt.Println(r2)
}
