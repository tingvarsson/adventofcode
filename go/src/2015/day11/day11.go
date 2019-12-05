package day11

import (
	"strings"
)

func increment(in string) (out string) {
	tmp := []rune(in)
	// quick fix any non allowed chars while incrementing
	if i := strings.IndexAny(in, "iol"); i != -1 {
		tmp[i] = tmp[i] + 1
		for i++; i < len(tmp); i++ {
			tmp[i] = 'a'
		}
		return string(tmp)
	}

	for i := len(tmp) - 1; i > 0; i-- {
		if tmp[i] == 'z' {
			tmp[i] = 'a'
		} else {
			tmp[i] = tmp[i] + 1
			break
		}
	}
	return string(tmp)
}

func validate(pass string) bool {
	if strings.ContainsAny(pass, "iol") {
		return false
	}

	includesSeq := false
	tmp := []rune(pass)
	for i := 0; i < len(tmp)-2; i++ {
		if tmp[i]+1 == tmp[i+1] && tmp[i]+2 == tmp[i+2] {
			includesSeq = true
			break
		}
	}
	if !includesSeq {
		return false
	}

	numPair := 0
	foundPair := ' '
	for i := 0; i < len(tmp)-1; i++ {
		if tmp[i] == tmp[i+1] && tmp[i] != foundPair {
			numPair++
			if numPair == 2 {
				return true
			}
			foundPair = tmp[i]
			i++
		}
	}
	return false
}

func run(pass string) (newPass string) {
	newPass = increment(pass)
	for !validate(newPass) {
		newPass = increment(newPass)
	}
	return
}
