package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	housesCovered := run("../../../day3/example", 1)
	t.Logf("[Example] Houses covered with one santa: %v\n", housesCovered)
	housesCovered = run("../../../day3/example2", 1)
	t.Logf("[Example2] Houses covered with one santa: %v\n", housesCovered)
	housesCovered = run("../../../day3/input", 1)
	t.Logf("[Input] Houses covered with one santa: %v\n", housesCovered)

	housesCovered = run("../../../day3/example", 2)
	t.Logf("[Example] Houses covered with two santas: %v\n", housesCovered)
	housesCovered = run("../../../day3/example2", 2)
	t.Logf("[Example2] Houses covered with two santas: %v\n", housesCovered)
	housesCovered = run("../../../day3/input", 2)
	t.Logf("[Input] Houses covered with two santas: %v\n", housesCovered)
}
