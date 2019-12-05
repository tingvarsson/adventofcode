package day12

import (
	"fmt"
	"regexp"
	"strings"
	"utils"
)

var digitRegexp = regexp.MustCompile("(-?\\d+)")

func sumNumbersInString(input string) (sumNumbers int) {
	match := digitRegexp.FindAllStringSubmatch(input, -1)
	for _, n := range match {
		sumNumbers += utils.Atoi(n[1])
	}
	return
}

func run(filepath string) (sumNumbers int, sumNoRed int) {
	input := utils.ReadFileToString(filepath)
	sumNumbers = sumNumbersInString(input)

	sumNoRed = sumNumbers
	curlyRe := regexp.MustCompile("{[^{}]*\"red\"[^{}}]*}")
	bracketRe := regexp.MustCompile(`\[[^\[\]]*\"red\"[^\[\]]*\]`)
	curlyMatches := curlyRe.FindAllString(input, -1)
	for _, curlyMatch := range curlyMatches {
		fmt.Println(curlyMatch)
		bracketMatches := bracketRe.FindAllString(curlyMatch, -1)
		for _, bracketMatch := range bracketMatches {
			curlyMatch = strings.Replace(curlyMatch, bracketMatch, "", -1)
		}
		fmt.Println(curlyMatch)
		sumNoRed -= sumNumbersInString(curlyMatch)
	}

	return
}
