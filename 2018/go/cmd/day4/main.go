package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("../sortedinput")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	const guardPattern = "#(\\d+)"
	var guardReg = regexp.MustCompile(guardPattern)
	const asleepPattern = ":(\\d+)] falls asleep"
	var asleepReg = regexp.MustCompile(asleepPattern)
	const awakePattern = ":(\\d+)] wakes up"
	var awakeReg = regexp.MustCompile(awakePattern)

	var guards = make(map[int][60]int)
	var guardID = 0
	var asleepMinute = 0
	var awakeMinute = 0
	var sleepRecord [60]int
	for _, line := range lines {
		guardMatch := guardReg.FindStringSubmatch(line)
		if guardMatch != nil {
			n, err := strconv.Atoi(guardMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			guardID = n
			if record, ok := guards[guardID]; ok {
				sleepRecord = record
			} else {
				sleepRecord = [60]int{}
			}
		}
		asleepMatch := asleepReg.FindStringSubmatch(line)
		if asleepMatch != nil {
			n, err := strconv.Atoi(asleepMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			asleepMinute = n
		}
		awakeMatch := awakeReg.FindStringSubmatch(line)
		if awakeMatch != nil {
			n, err := strconv.Atoi(awakeMatch[1])
			if err != nil {
				log.Fatal(err)
			}
			awakeMinute = n
			for i := asleepMinute; i < awakeMinute; i++ {
				sleepRecord[i]++
			}
			guards[guardID] = sleepRecord
		}
	}

	var scenarioOneGuard = 0
	var scenarioOneMinute = 0
	var scenarioTwoGuard = 0
	var scenarioTwoMinute = 0
	var scenarioTwoMaxMinute = 0
	var maxSleep = 0
	for guardID, sleepRecord := range guards {
		var sleepSum = 0
		var mostActiveMinute = 0
		var maxMinute = 0
		for idx, i := range sleepRecord {
			sleepSum += i
			if i > maxMinute {
				mostActiveMinute = idx
				maxMinute = i
			}
		}
		if sleepSum > maxSleep {
			scenarioOneGuard = guardID
			scenarioOneMinute = mostActiveMinute
			maxSleep = sleepSum
		}
		if maxMinute > scenarioTwoMaxMinute {
			scenarioTwoGuard = guardID
			scenarioTwoMinute = mostActiveMinute
			scenarioTwoMaxMinute = maxMinute
		}
	}
	fmt.Println(scenarioOneGuard * scenarioOneMinute)
	fmt.Println(scenarioTwoGuard * scenarioTwoMinute)
}
