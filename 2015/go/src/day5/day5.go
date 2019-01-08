package main

import (
	"strings"
	"utils"
)

type niceTest func(string) bool

func niceBasic(s string) bool {
	// It contains at least three vowels (aeiou only)
	vowels := []string{"a", "e", "i", "o", "u"}
	numVowels := 0
	for _, v := range vowels {
		numVowels += strings.Count(s, v)
	}
	if numVowels < 3 {
		return false
	}

	// It does not contain the strings ab, cd, pq, or xy
	illegalStrings := []string{"ab", "cd", "pq", "xy"}
	for _, illegal := range illegalStrings {
		if strings.Contains(s, illegal) {
			return false
		}
	}

	// It contains at least one letter that appears twice in a row
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func niceEnhanced(s string) bool {
	// It contains a pair of any two letters that appears
	// at least twice in the string without overlapping
	foundPair := false
	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Contains(s[i+2:], pair) {
			foundPair = true
			break
		}
	}
	if !foundPair {
		return false
	}

	// It contains at least one letter which repeats
	// with exactly one letter between them
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func run(filepath string, tester niceTest) (numNice int) {
	lines := utils.ReadFileToLines(filepath)
	for _, l := range lines {
		if tester(l) {
			numNice++
		}
	}
	return
}

func main() {
}
