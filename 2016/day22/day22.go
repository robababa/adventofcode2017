package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"log"
)

type node struct {
	size int
	used int
	avail int
}

var nodes [][]node

// each line is of the form
// /dev/grid/node-x1-y10    90T   71T    19T   78%
// we want to
// (1) get rid of the /dev/grid/node- prefix
// (2) remove the x, y and Ts
// (3) replace the - with a single space
// (4) consolidate multiple spaces with a single space


func main() {
	lines := readInput()
	parseInput(lines)
	//fmt.Println(nodes)
	fmt.Println("Part 1: viable pair count is", countViablePairs())
}

func countViablePairs() int {
	answer := 0
	for i := range nodes {
		for j, node1 := range nodes[i] {
			for i2 := range nodes {
				for j2, node2 := range nodes[i2] {
					if node1.used != 0 && !(i == i2 && j == j2) && node1.used <= node2.avail {
						answer++
					}
				}
			}
		}
	}
	return answer
}

func transformLine(s string) []string {
	// each line is of the form
	// /dev/grid/node-x1-y10    90T   71T    19T   78%
	// we want to
	// (1) get rid of the /dev/grid/node- prefix
	// (2) remove the x, y and Ts
	// (3) replace the - with a single space
	// and return the fields as an array of strings
	almost := strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(s, "/dev/grid/node-", "", 1),
					"x", "", 1),
				"y", "", 1),
			"T", "", -1),
		"-", " ", 1)
	return strings.Fields(almost)
}

func Atoi(s string) int {
	answer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Could not convert string", s, "into integer")
	}
	return answer
}

func addNode(x, y, size, used int) {
	if x >= len(nodes) {
		nodes = append(nodes, []node{})
	}
	if y >= len(nodes[x]) {
		nodes[x] = append(nodes[x], node{})
	}
	nodes[x][y] = node{size: size, used: used, avail: size - used}
}

func parseInput(lines []string) {
	for _, line := range lines[2:] {
		tokens := transformLine(line)
		addNode(Atoi(tokens[0]), Atoi(tokens[1]), Atoi(tokens[2]), Atoi(tokens[3]))
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
