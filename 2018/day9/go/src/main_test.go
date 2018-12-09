package main

import (
	"testing"
)

func TestPlay(t *testing.T) {
	var tests = []struct {
		numPlayers int
		lastMarble int
		expected   int
	}{
		{9, 25, 32},
		{418, 70769, 402398},
		{418, 7076900, 3426843186},
	}
	for _, test := range tests {
		if output := play(test.numPlayers, test.lastMarble); output != test.expected {
			t.Error("Test Failed: {} numPlayers, {} lastMarble, {} expected, recieved: {}", test.numPlayers, test.lastMarble, test.expected, output)
		}
	}
}
