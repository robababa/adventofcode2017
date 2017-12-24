package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"log"
	"strings"
)

type beam struct {
	port1 int
	port2 int
}

func main() {
	beams := parseInput(readInput())
	fmt.Println("beams are:")
	fmt.Println(beams)
	strongestBridge := buildStrongest(beams)
	fmt.Println("Part 1: strongest bridge is", strongestBridge)
	fmt.Println("Part 1: strength is", strength(strongestBridge))
}

func buildStrongest(beams []beam) []beam {
	return []beam{}
}

func strength(bridgeBeams []beam) int {
	answer := 0
	for _, b := range bridgeBeams {
		answer += b.port1 + b.port2
	}
	return answer
}

func parseInput(lines []string) []beam {
	var b []beam
	for _, line := range lines {
		tokens := strings.Split(line, "/")
		b = append(b, beam{port1: Atoi(tokens[0]), port2: Atoi(tokens[1])})
	}
	return b
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}

func Atoi(s string) int {
	answer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Could not convert string", s, "to integer")
	}
	return answer
}