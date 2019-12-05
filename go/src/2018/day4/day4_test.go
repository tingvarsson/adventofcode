package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args[1] = os.Getenv("ROOT") + "/2018/day4/sortedinput"
	main()
}
