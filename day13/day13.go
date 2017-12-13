package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

var firewall = make(map[int]int)

func main() {
	parseInput(readInput())
	part1Severity, _ := runThroughFirewall(0)
	fmt.Println("Part 1: Severity is", part1Severity)
	fmt.Println("Part 2: Smallest perfect delay is", smallestPerfectDelay())
}

func smallestPerfectDelay() int {
	for i :=0; ; i++ {
		if _, caught := runThroughFirewall(i); !caught {
			return i
		}
	}
}

func runThroughFirewall(delay int) (int, bool) {
	severity := 0
	caught := false
	for k, v := range firewall {
		cycle := 2 * (v - 1)
		if (delay + k) % cycle == 0 {
			caught = true
			severity += k * v
		}
	}
	return severity, caught
}

func parseInput(input []string) {
	for _, line := range input {
		tokens := strings.Split(strings.Replace(line, " ", "", 1), ":")
		depth, _ := strconv.Atoi(tokens[0])
		height, _ := strconv.Atoi(tokens[1])
		firewall[depth] = height
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