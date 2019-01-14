package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// Abs returns the absolut value of x
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Atoi returns the integer value of string s
func Atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// AtoiSlice returns a slice of integer values given as string s
func AtoiSlice(s string) (ints []int) {
	for _, r := range s {
		ints = append(ints, Atoi(string(r)))
	}
	return
}

// Insert inserts v at index i into []int s
func Insert(s *[]int, i int, v int) {
	*s = append(*s, 0)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = v
}

// Max returns the maximum value, and its index, in []int s
func Max(s []int) (index int, value int) {
	for i, v := range s {
		if v > value {
			index = i
			value = v
		}
	}
	return
}

// MaxIndex returns the index of the maximum value in []int s
func MaxIndex(array []int) (index int) {
	value := 0
	for i, v := range array {
		if v > value {
			index = i
		}
	}
	return
}

// MaxValue returns the maximum value in []int s
func MaxValue(array []int) (value int) {
	for _, v := range array {
		if v > value {
			value = v
		}
	}
	return
}

// ReadFileToLines reads filepath and returns the content as a []string
func ReadFileToLines(filepath string) (lines []string) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return
}

// ReadFileToString reads filepath and returns the content as a string
func ReadFileToString(filepath string) string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

// StringsToInts takes a []string s and converts it to a []int i with Atoi()
func StringsToInts(s []string) (i []int) {
	for _, s := range s {
		i = append(i, Atoi(s))
	}
	return
}

// Sum returns the sum of all values in []int s
func Sum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}
