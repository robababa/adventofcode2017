package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

var m = make(map[string]int)

func main() {
	parseInput(readInput())
	printResults()
	cancelOpposites("n", "s")
	cancelOpposites("ne", "sw")
	cancelOpposites("nw", "se")
	printResults()
}

func cancelOpposites(direction1 string, direction2 string) {
	if m[direction1] >= m[direction2] {
		m[direction1] -= m[direction2]
		m[direction2] = 0
	} else {
		cancelOpposites(direction2, direction1)
	}
}

func printResults() {
	fmt.Println(m)
}

func parseInput(str string) {
	for _, c := range strings.Split(str, ",") {
		m[c]++
	}
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return answer[0:len(answer)-1]
}
