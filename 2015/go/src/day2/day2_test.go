package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	wrapping, ribbon := run("../../../day2/input")
	t.Logf("Wrapping: %v Ribbon: %v\n", wrapping, ribbon)
}
