package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("../input")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, i)
	}

	sum := 0
	for _, n := range numbers {
		sum += n
	}
	println(sum)

	var knownsums []int
	sum = 0
	idx := 0
	for {
		sum += numbers[idx%len(numbers)]
		pos := sort.Search(len(knownsums), func(i int) bool { return knownsums[i] >= sum })
		if pos < len(knownsums) && knownsums[pos] == sum {
			break
		} else {
			knownsums = append(knownsums, 0)
			copy(knownsums[pos+1:], knownsums[pos:])
			knownsums[pos] = sum
			idx++
		}
	}
	println(sum)
}
