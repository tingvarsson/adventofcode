package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"utils"
)

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
	println(utils.Sum(freqChanges))

	var seenFreqs []int
	freq := 0
	for i := 0; ; i++ {
		freq += freqChanges[i%len(freqChanges)]
		pos := sort.Search(len(seenFreqs),
			func(i int) bool { return seenFreqs[i] >= freq })

		if pos < len(seenFreqs) && seenFreqs[pos] == freq {
			break // done, found an already seen frequency
		} else {
			utils.Insert(&seenFreqs, pos, freq)
		}
	}
	println(freq)
}
