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
	wires = make(map[string]string)
	for _, l := range lines {
		m := re.FindStringSubmatch(l)
		wires[m[2]] = m[1]
	}
	return
}

func getValue(wires wireMap, wire string) uint16 {
	if i, err := strconv.Atoi(wire); err == nil {
		return uint16(i)
	}
	opRegex := regexp.MustCompile("^(.+) ([A-Z]+) (.+)$")
	notRegex := regexp.MustCompile("^NOT (.*)$")

	connection := wires[wire]
	if i, err := strconv.Atoi(connection); err == nil {
		return uint16(i)
	} else if m := opRegex.FindStringSubmatch(connection); len(m) != 0 {
		var value uint16
		switch m[2] {
		case "AND":
			value = getValue(wires, m[1]) & getValue(wires, m[3])
		case "OR":
			value = getValue(wires, m[1]) | getValue(wires, m[3])
		case "LSHIFT":
			value = getValue(wires, m[1]) << getValue(wires, m[3])
		case "RSHIFT":
			value = getValue(wires, m[1]) >> getValue(wires, m[3])
		default:
			panic("unknown OP")
		}
		wires[wire] = strconv.Itoa(int(value))
		return value
	} else if m := notRegex.FindStringSubmatch(connection); len(m) != 0 {
		value := ^getValue(wires, m[1])
		wires[wire] = strconv.Itoa(int(value))
		return value
	} else {
		return getValue(wires, connection)
	}
}

func run(filepath string, lookup string) uint16 {
	wires := parseWires(filepath)
	return getValue(wires, lookup)
}

func main() {}
