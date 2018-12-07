package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func insert(s *[]int, i int, v int) {
	*s = append(*s, 0)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = v
}

func main() {
	file, err := os.Open("../input")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	var freqChanges []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		freqChanges = append(freqChanges, i)
	}
	println(sum(freqChanges))

	var seenFreqs []int
	freq := 0
	for i := 0; ; i++ {
		freq += freqChanges[i%len(freqChanges)]
		pos := sort.Search(len(seenFreqs),
			func(i int) bool { return seenFreqs[i] >= freq })

		if pos < len(seenFreqs) && seenFreqs[pos] == freq {
			break // done, found an already seen frequency
		} else {
			insert(&seenFreqs, pos, freq)
		}
	}
	println(freq)
}
