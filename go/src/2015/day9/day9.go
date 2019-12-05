package main

import (
	"regexp"
	"utils"
)

type distanceMap map[string]int
type pathMap map[string]distanceMap

func addPath(paths pathMap, src string, dst string, dist int) {
	if destinations, ok := paths[src]; ok {
		destinations[dst] = dist
	} else {
		paths[src] = make(distanceMap)
		paths[src][dst] = dist
	}
}

func parsePaths(filepath string) (paths pathMap) {
	paths = make(pathMap)
	lines := utils.ReadFileToLines(filepath)
	for _, l := range lines {
		re := regexp.MustCompile("(\\w+) to (\\w+) = (\\d+)")
		m := re.FindStringSubmatch(l)
		addPath(paths, m[1], m[2], utils.Atoi(m[3]))
		addPath(paths, m[2], m[1], utils.Atoi(m[3]))
	}
	return
}

func permutations(input []string) (output [][]string) {
	var helper func([]string, int)
	helper = func(input []string, n int) {
		if n == 1 {
			tmp := make([]string, len(input))
			copy(tmp, input)
			output = append(output, tmp)
			return
		}
		for i := 0; i < n; i++ {
			helper(input, n-1)
			if n%2 == 1 {
				tmp := input[i]
				input[i] = input[n-1]
				input[n-1] = tmp
			} else {
				tmp := input[0]
				input[0] = input[n-1]
				input[n-1] = tmp
			}
		}
	}
	helper(input, len(input))
	return
}

func run(filepath string) (shortest, longest int) {
	paths := parsePaths(filepath)

	var locations []string
	for src := range paths {
		locations = append(locations, src)
	}

	possibleRoutes := permutations(locations)

	for _, route := range possibleRoutes {
		dist := 0
		for i := 0; i < len(route)-1; i++ {
			dist += paths[route[i]][route[i+1]]
		}
		if shortest == 0 || dist < shortest {
			shortest = dist
		}
		if dist > longest {
			longest = dist
		}
	}
	return
}

func main() {}
