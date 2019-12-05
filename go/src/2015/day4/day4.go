package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func test(secretKey string, n int, numZeros int) bool {
	hasher := md5.New()
	hasher.Write([]byte(secretKey + strconv.Itoa(n)))
	result := hex.EncodeToString(hasher.Sum(nil))
	if result[:numZeros] == strings.Repeat("0", numZeros) {
		return true
	}
	return false
}

func run(secretKey string, numZeros int) (n int) {
	n = 1
	for !test(secretKey, n, numZeros) {
		n++
	}
	return n
}

func main() {
}
