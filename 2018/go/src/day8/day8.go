package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"utils"
)

type node struct {
	Length   int
	Children []node
	MetaData []int
}

func (n *node) addChild(c node) {
	n.Children = append(n.Children, c)
	n.Length += c.Length
}

func (n *node) addMetadata(m int) {
	n.MetaData = append(n.MetaData, m)
	n.Length++
}

func readFileToIntSlice(filepath string) (data []int) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		data = append(data, utils.Atoi(scanner.Text()))
	}
	return
}

func parseNode(data []int) (n node) {
	n.Length = 2
	for i := 0; i < data[0]; i++ {
		n.addChild(parseNode(data[n.Length:]))
	}
	for i := 0; i < data[1]; i++ {
		n.addMetadata(data[n.Length])
	}
	return
}

func sumMetadata(n node) (sum int) {
	for _, c := range n.Children {
		sum += sumMetadata(c)
	}
	sum += utils.Sum(n.MetaData)
	return
}

func sumMetadata2(n node) (sum int) {
	if len(n.Children) == 0 {
		sum += utils.Sum(n.MetaData)
	} else {
		for _, m := range n.MetaData {
			if m <= len(n.Children) {
				sum += sumMetadata2(n.Children[m-1])
			}
		}
	}
	return
}

func main() {
	data := readFileToIntSlice("../input")
	rootNode := parseNode(data)
	fmt.Printf("%d\n", sumMetadata(rootNode))
	fmt.Printf("%d\n", sumMetadata2(rootNode))
}
