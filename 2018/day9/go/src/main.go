package main

import (
	"fmt"
)

type node struct {
	Value int
	Next  *node
	Prev  *node
}

func max(array []int) (value int) {
	for _, v := range array {
		if v > value {
			value = v
		}
	}
	return
}

func play(players int, lastMarble int) int {
	score := make([]int, players)
	currentPlayer := 0
	currentMarble := &node{Value: 0}
	currentMarble.Next = currentMarble
	currentMarble.Prev = currentMarble
	for marble := 1; marble < lastMarble; marble++ {
		if marble%23 != 0 {
			newMarble := &node{Value: marble}
			currentMarble = currentMarble.Next
			newMarble.Prev = currentMarble
			newMarble.Next = currentMarble.Next
			currentMarble.Next.Prev = newMarble
			currentMarble.Next = newMarble
			currentMarble = newMarble
		} else {
			score[currentPlayer] += marble
			currentMarble = currentMarble.Prev.Prev.Prev.Prev.Prev.Prev.Prev
			score[currentPlayer] += currentMarble.Value
			currentMarble = currentMarble.Next
			currentMarble.Prev.Prev.Next = currentMarble
			currentMarble.Prev = currentMarble.Prev.Prev
		}
		currentPlayer = (currentPlayer + 1) % totalPlayers
	}
	return max(score)
}

const totalPlayers = 418
const lastMarble = 70769
const lastMarble2 = lastMarble * 100

func main() {
	fmt.Printf("High score #1: %d\n", play(totalPlayers, lastMarble))
	fmt.Printf("High score #2: %d\n", play(totalPlayers, lastMarble2))
}
