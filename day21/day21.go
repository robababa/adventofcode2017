package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// an enhancement is a mapping like this
// .#/#. => ##./###/...
// the key is the string on the left, and the value is the string on the right
var enhancements = make(map[string]string)

// the initial value is
// .#.
// ..#
// ###
var InitialGrid = []string{".#.", "..#", "###"}

func main() {
	parseInput(readInput())
	//fmt.Println(enhancements)
	fmt.Println(enhance(InitialGrid))
}

func rotateClockwise(grid []string) []string {
	var answer []string
	// "prime" our array of strings with empty strings
	for range grid {
		answer = append(answer, "")
	}
	// now append to each of the empty strings
	// example:  in the grid
	// 1 2
	// 3 4
	// we visit 3 and 1 to fill in the top row,
	// then visit 4 and 2 to fill in the bottom row
	for i := len(grid) -1; i >= 0; i-- {
		for j := 0; j <= len(grid) -1; j++ {
			answer[j] += string(grid[i][j])
		}
	}
	return answer
}

func flipTopAndBottom(grid []string) []string {
	var answer []string
	for _, s := range grid {
		answer = append([]string{s}, answer...)
	}
	return answer
}

func flipLeftAndRight(grid []string) []string {
	var answer []string
	for _, s := range grid {
		var thisLine string
		for i := len(s) - 1; i >= 0; i-- {
			thisLine += string(s[i])
		}
		answer = append(answer, thisLine)
	}
	return answer
}

func enhance(grid []string) []string {
	var answer []string
	enhancementKey := grid[0] + "/" + grid[1]
	enhancementValue := enhancements[enhancementKey]
	for _, str := range strings.Split(enhancementValue, "/") {
		answer = append(answer, str)
	}
	return answer
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
