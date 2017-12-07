package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	node     string
	weight   int
	children []Tree
}

type Forest []Tree

func main() {
	nodes := make(map[string]int)
	rawInput := readInput()
	//fmt.Println(readInput())
	loadNodes(rawInput, nodes)
	//fmt.Println(nodes)
	forest := loadForest(rawInput, nodes)
	//fmt.Println(forest)
}

func loadForest(rawInput []string, nodes map[string]int) Forest {
	var forest Forest
	for _, s := range rawInput {
		fields := strings.Split(s, " ")
		nodeName := fields[0]
		tree := Tree{node: nodeName, weight: nodes[nodeName]}
		// if this node has children, then add them to the tree
		if len(fields) > 3 {
			for _, child := range fields[3:] {
				childName := child[:len(child)-1]
				tree.children = append(tree.children, Tree{node: childName, weight: nodes[childName]})
			}
		}
		forest = append(forest, tree)
	}
	return forest
}

func loadNodes(inputStrings []string, nodes map[string]int) {
	for _, s := range inputStrings {
		fields := strings.Split(s, " ")
		nodeName := fields[0]
		nodeWeight, _ := strconv.Atoi(fields[1][1 : len(fields[1])-1])
		nodes[nodeName] = nodeWeight
	}
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}
