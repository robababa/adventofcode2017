package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
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

func noop(grid[] string) []string {
	return append(grid)
}

// rotate the grid clockwise
func rotate(grid []string, rotations int) []string {
	if rotations < 0 {
		log.Fatal("Called rotate() with negative rotation count")
	}
	if rotations == 0 {
		// like a no-op, just return the original grid
		return noop(grid)
	}
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
	return rotate(answer, rotations - 1)
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

func gridToKey(grid []string) string {
	var answer string
	for _, s := range grid {
		answer += s + "/"
	}
	return answer[:len(answer) - 1]
}

func findKey(grid []string) string {
	var answer string
	// for each way to flip, including not flipping at all
	for _, flipper := range []func([]string) []string{noop, flipTopAndBottom, flipLeftAndRight} {
		// and each way to rotate, including not rotating at all
		for rotations := 0; rotations < 4; rotations++ {
			// see if the resulting grid is a key in our enhancement rules mapping
			//fmt.Println("Looking for key")
			answer = gridToKey(rotate(flipper(grid), rotations))
			// and if it is, return key
			if enhancements[answer] != "" {
				//fmt.Println("Found the key! Its value is", answer)
				return answer
			}
		}
	}
	return answer
}

func enhance(grid []string) []string {
	var answer []string
	key := findKey(grid)
	fmt.Println("Enhancement key is", key)
	val := enhancements[key]
	fmt.Println("Enhancement value is", val)
	for _, str := range strings.Split(val, "/") {
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
