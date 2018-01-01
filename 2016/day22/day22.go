package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

type node struct {
	size int
	used int
	hasTargetData bool
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
	fmt.Println(readInput())
}

func transformLine(s string) string {
	// each line is of the form
	// /dev/grid/node-x1-y10    90T   71T    19T   78%
	// we want to
	// (1) get rid of the /dev/grid/node- prefix
	// (2) remove the x, y and Ts
	// (3) replace the - with a single space
	// (4) consolidate multiple spaces with a single space
	almost := strings.Replace(
		strings.Replace(
			strings.Replace(
				strings.Replace(
					strings.Replace(s, "/dev/grid/node-", "", 1),
					"x", "", 1),
				"y", "", 1),
			"T", "", -1),
		"-", " ", 1)
	return strings.Join(strings.Fields(almost), " ")
}

//func parseInput(lines []string) {
//	for _, line := range lines[2:] {
//		tokens := strings.Split(line, " ")
//	}
//}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}
