package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
	"math"
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
	//fmt.Println(enhanceSubGrid(InitialGrid))
	currentGrid := InitialGrid
	//fmt.Println("Before enhancement, grid starts as")
	//fmt.Println(currentGrid)
	for round := 1; round <= 5; round++ {
		currentGrid = enhanceEntireGrid(currentGrid)
		//fmt.Println("After round", round, "grid is")
		//fmt.Println(currentGrid)
	}
	fmt.Println("Final grid is")
	fmt.Println(currentGrid)
	// 134 is too low
	fmt.Println("Number of hashes is", countHashes(currentGrid))
}

func countHashes(lines []string) int {
	answer := 0
	for _, s := range lines {
		answer += len(strings.Replace(s, ".", "", -1))
	}
	return answer
}

func noop(grid []string) []string {
	return append(grid)
}

// rotate the grid clockwise
func rotate(grid []string, rotations int) []string {
	//fmt.Println("rotate arguments are rotations =", rotations, "and grid =", grid)
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

func enhanceSubGrid(grid []string) []string {
	//fmt.Println("enhanceSubGrid argument is", grid)
	var answer []string
	key := findKey(grid)
	//fmt.Println("Enhancement key is", key)
	val := enhancements[key]
	//fmt.Println("Enhancement value is", val)
	for _, str := range strings.Split(val, "/") {
		answer = append(answer, str)
	}
	return answer
}

func combineGrids(grids [][]string) []string {
	//fmt.Println("combineGrids() input is", grids)
	// if there is only one grid, return it
	if len(grids) == 1 {
		return grids[0]
	}

	var answer []string
	subGridDimension := len(grids[0])
	gridDimension := int(math.Sqrt(float64(len(grids) * subGridDimension * subGridDimension)))
	subGridsAcross := gridDimension / subGridDimension
	subGridsDown := subGridsAcross
	// 0,0 + 1,0
	// 0,1 + 1,1
	// 0,2 + 1,2
	// 2,0 + 3,0
	// 2,1 + 3,1
	// 2,2 + 3,2
	for subGridGroup := 0; subGridGroup < subGridsDown; subGridGroup++ {
		subGridStart := subGridGroup * subGridsAcross
		for j := 0; j < subGridDimension; j++ {
			var nextRow string
			for i := 0; i < subGridsAcross; i++ {
				nextRow += grids[subGridStart+i][j]
			}
			answer = append(answer, nextRow)
		}
	}
	return answer
}

func enhanceAllSubGrids(grids [][]string) [][]string {
	var answer [][]string
	//fmt.Println("enhanceAllSubGrids argument is", grids)
	for _, g := range grids {
		answer = append(answer, enhanceSubGrid(g))
	}
	//fmt.Println("in enhanceAllSubGrids, answer is", answer)
	return answer
}

// example:
// 1234
// 5678
// ABCD
// EFGH
// should return
// [[12 56] [34 78] [AB EF] [CD GH]]
func divideGrid(grid []string) [][]string {
	//fmt.Println()
	//fmt.Println("divideGrid input is", grid)
	var subGrids [][]string
	gridLength := len(grid)
	// nothing to do if grid size is 2 or 3, just return the original grid inside an array
	if gridLength == 2 || gridLength == 3 {
		return append(subGrids, grid)
	}
	subGridSize := 2
	if gridLength % 3 == 0 {
		//fmt.Println("Setting subGridSize to 3")
		subGridSize = 3
	}

	gridDimension := len(grid[0])
	subGridCount := gridDimension * gridDimension / subGridSize	/ subGridSize

	for subGridNum := 0; subGridNum < subGridCount; subGridNum++ {
		rowStart :=  subGridSize * (subGridNum * subGridSize / gridDimension)
		columnStart := (subGridNum * subGridSize) % gridDimension
		var newSubGrid []string
		for r := 0; r < subGridSize; r++ {
			newSubGrid = append(newSubGrid, grid[rowStart + r][columnStart:columnStart+subGridSize])
		}
		subGrids = append(subGrids, newSubGrid)
	}
	return subGrids
}

func enhanceEntireGrid(grid []string) []string {
	// enhance each subgrid individually, and then combine them for the answer
	return combineGrids(enhanceAllSubGrids(divideGrid(grid)))
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
