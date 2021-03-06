package main

import (
	"os"
	"sort"
	"utils"
)

func sumFrequencies(filepath string) int {
	lines := utils.ReadFileToLines(filepath)
	freqChanges := utils.StringsToInts(lines)
	return utils.Sum(freqChanges)
}

func sameFrequency(filepath string) (freq int) {
	lines := utils.ReadFileToLines(filepath)
	freqChanges := utils.StringsToInts(lines)
	var seenFreqs []int
	for i := 0; ; i++ {
		freq += freqChanges[i%len(freqChanges)]
		pos := sort.Search(len(seenFreqs),
			func(i int) bool { return seenFreqs[i] >= freq })

		if pos < len(seenFreqs) && seenFreqs[pos] == freq {
			return // done, found an already seen frequency
		}

		utils.Insert(&seenFreqs, pos, freq)
	}
}

func main() {
	println(sumFrequencies(os.Args[1]))
	println(sameFrequency(os.Args[1]))
}
