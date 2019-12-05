package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"utils"
)

func parsePotsAndPatterns(filepath string) (pots string, rules []string) {
	lines := utils.ReadFileToLines(filepath)
	initialStatePattern := "initial state: (.*)"
	initialStateRegexp := regexp.MustCompile(initialStatePattern)
	rulePattern := "(.*) => #"
	ruleRegexp := regexp.MustCompile(rulePattern)
	for _, line := range lines {
		initialStateMatch := initialStateRegexp.FindStringSubmatch(line)
		if initialStateMatch != nil {
			pots = initialStateMatch[1]
		}
		ruleMatch := ruleRegexp.FindStringSubmatch(line)
		if ruleMatch != nil {
			rules = append(rules, ruleMatch[1])
		}
	}
	return
}

func applyRules(input string, rules []string) string {
	output := strings.Repeat(".", len(input))
	for _, r := range rules {
		ruleRegex := regexp.MustCompile(regexp.QuoteMeta(r))
		i := 0
		for ruleMatch := ruleRegex.FindStringIndex(input); ruleMatch != nil; ruleMatch = ruleRegex.FindStringIndex(input[i:]) {
			i += ruleMatch[0] + 1
			output = output[:i+1] + "#" + output[i+1+1:]
		}
	}
	return output
}

func sumPots(pots string, offset int) (sum int) {
	for i, p := range pots {
		if string(p) == "#" {
			sum += i - offset
		}
	}
	return
}

func main() {
	pots, rules := parsePotsAndPatterns(os.Args[1])
	offset := 10
	pots = strings.Repeat(".", offset) + pots + strings.Repeat(".", 150)
	for g := 0; ; g++ {
		tempPots := applyRules(pots, rules)

		if g == 19 {
			fmt.Printf("sum #20: %v\n", sumPots(tempPots, offset-g))
		}

		if pots == tempPots[1:]+"." {
			break // Pattern is now re-occuring but with an offset
		}

		pots = tempPots
		offset++
	}
	fmt.Printf("sum #50000000000: %v\n", sumPots(pots, offset-5000000000))
}
