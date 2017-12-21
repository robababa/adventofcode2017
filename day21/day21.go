package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// an enhancement is a mapping like this
// .#/#. => ##./###/...
var enhancements = make(map[string]string)

func main() {
	parseInput(readInput())
	fmt.Println(enhancements)
}

func parseInput(lines []string) {
	for _, line := range lines {
		tokens := strings.Split(line, " => ")
		enhancements[tokens[0]] = tokens[1]
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
