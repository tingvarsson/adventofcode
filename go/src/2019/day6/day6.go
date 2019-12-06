package main

import (
	"fmt"
	"os"
	"strings"
	"utils"
)

type node struct {
	Name     string
	Distance int
	Children []node
}

func parseMap(m map[string][]string, name string, dist int) (n node) {
	n = node{name, dist, []node{}}
	for _, child := range m[name] {
		n.Children = append(n.Children, parseMap(m, child, dist+1))
	}
	return
}

func sumDistance(n node) (sum int) {
	for _, c := range n.Children {
		sum += sumDistance(c)
	}
	sum += n.Distance
	return
}

func travelDistance(r node, a, b string) (d int) {
	if r.Name == a || r.Name == b {
		return r.Distance - 1
	}
	if len(r.Children) == 0 {
		return 0
	}
	for _, c := range r.Children {
		tmp := travelDistance(c, a, b)
		if d == 0 && tmp != 0 {
			d = tmp
		} else if d != 0 && tmp != 0 {
			return d + tmp - (2 * r.Distance)
		}
	}
	return
}

func run(filepath string) (orbitCount, orbitTransfers int) {
	input := utils.ReadFileToLines(filepath)
	m := make(map[string][]string)
	for _, line := range input {
		objs := strings.Split(line, ")")
		m[objs[0]] = append(m[objs[0]], objs[1])
	}

	root := parseMap(m, "COM", 0)
	orbitCount = sumDistance(root)
	orbitTransfers = travelDistance(root, "YOU", "SAN")
	return
}

func main() {
	r1, r2 := run(os.Args[1])
	fmt.Println(r1)
	fmt.Println(r2)
}
