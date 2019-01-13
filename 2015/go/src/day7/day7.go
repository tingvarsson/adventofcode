package main

import (
	"regexp"
	"strconv"
	"utils"
)

type wireMap map[string]string

func parseWires(filepath string) (wires wireMap) {
	re := regexp.MustCompile("^(.*) -> (.*)$")
	lines := utils.ReadFileToLines(filepath)
	wires = make(wireMap)
	for _, l := range lines {
		m := re.FindStringSubmatch(l)
		wires[m[2]] = m[1]
	}
	return
}

var opRegex = regexp.MustCompile("^(\\w*) ?(\\b[A-Z]+\\b) (\\w+)$")

func getValue(wires wireMap, wire string) uint16 {
	if i, err := strconv.Atoi(wire); err == nil {
		// endpoint reached and an actual value is found
		return uint16(i)
	}

	// Either an expression (AND, OR, L/RSHIFT, NOT) or a reference
	var value uint16
	expr := wires[wire]
	if m := opRegex.FindStringSubmatch(expr); len(m) != 0 {
		switch m[2] {
		case "AND":
			value = getValue(wires, m[1]) & getValue(wires, m[3])
		case "OR":
			value = getValue(wires, m[1]) | getValue(wires, m[3])
		case "LSHIFT":
			value = getValue(wires, m[1]) << getValue(wires, m[3])
		case "RSHIFT":
			value = getValue(wires, m[1]) >> getValue(wires, m[3])
		case "NOT":
			value = ^getValue(wires, m[3])
		default:
			panic("unknown OP")
		}
	} else {
		value = getValue(wires, expr)
	}
	// Write back the found wire value to speed up future lookups
	wires[wire] = strconv.Itoa(int(value))
	return value
}

func run(filepath string, lookup string) uint16 {
	wires := parseWires(filepath)
	return getValue(wires, lookup)
}

func main() {}
