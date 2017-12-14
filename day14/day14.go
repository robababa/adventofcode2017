package main

import (
	"os"
	"fmt"
	"strconv"
	"strings"
)

const ListSize = 256
const Rounds = 64
const GridHeight = 128

var skip int
var startAt int

func main() {
	input := os.Args[1]
	numberOfOnes := 0
	var stringArray []string
	for j := 0; j < GridHeight; j++ {
		skip = 0
		startAt = 0
		hashInput := parseInput(input + "-" + strconv.Itoa(j))
		intArray := buildIntArray()
		for i := 0; i < Rounds; i++ {
			intArray = applyTwists(intArray, hashInput)
		}
		myDenseHash := denseHash(intArray)
		//fmt.Println("Number of ones:", countOnes(myDenseHash))
		numberOfOnes += countOnes(myDenseHash)
		stringArray = append(stringArray, toBinaryString(myDenseHash))
	}
	fmt.Println("Part 1: Number of ones is", numberOfOnes)
	regionGrid := createRegionGrid(stringArray)
	//fmt.Println("Part 2: Grid")
	//fmt.Println(regionGrid)
	consolidateRegions(&regionGrid)
	//fmt.Println(regionGrid)
	fmt.Println("Part 2: Number of regions is", countRegions(&regionGrid))
}

func countRegions(regionGrid *[128][128]int) int {
	m := make(map[int]bool)
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			val := regionGrid[i][j]
			if val != 0 {
				m[val] = true
			}
		}
	}
	return len(m)
}

func consolidateRegions(regionGrid *[128][128]int) {
	updatedGrid := true
	for updatedGrid {
		updatedGrid = false
		for i := 0; i < 128; i++ {
			for j := 0; j < 128; j++ {
				here := regionGrid[i][j]
				if here == 0 {
					continue
				}
				var above, below, left, right int
				if i > 0 {
					above = regionGrid[i-1][j]
				}
				if j > 0 {
					left = regionGrid[i][j-1]
				}
				if i < 127 {
					below = regionGrid[i+1][j]
				}
				if j < 127 {
					right = regionGrid[i][j+1]
				}
				if above != 0 && above < here {
					regionGrid[i][j] = above
					updatedGrid = true
				} else if left != 0 && left < here {
					regionGrid[i][j] = left
					updatedGrid = true
				} else if right != 0 && right < here {
				regionGrid[i][j] = right
				updatedGrid = true
				} else if below != 0 && below < here {
				regionGrid[i][j] = below
				updatedGrid = true
				}
			}
		}
	}
}

func createRegionGrid(stringArray []string) [128][128]int {
	var answer [128][128]int
	regionNumber := 1
	for i, s := range stringArray {
		for j, c := range s {
			if c == '1' {
				answer[i][j] = regionNumber
				regionNumber++
			}
		}
	}
	return answer
}

func toBinaryString(input [16]int) string {
	var answer string
	for _, n := range input {
		s := strconv.FormatInt(int64(n), 2)
		// we need to prefix the string with zeros to make it 8 places long
		answer += strings.Repeat("0", 8 - len(s)) + s
	}
	return answer
}

func countOnes(input [16]int) int {
	answer := 0
	for _, n := range input {
		answer += len(strings.Replace(strconv.FormatInt(int64(n), 2), "0", "", -1))
	}
	return answer
}

// everything below here was copy/pasted from day 10

func denseHash(input [ListSize]int) [16]int {
	var answer [16]int
	for group:=0; group<16; group++ {
		answer[group] = input[group * 16]
		for elementWithinGroup := 1; elementWithinGroup < 16; elementWithinGroup++ {
			answer[group] ^= input[group * 16 + elementWithinGroup]
		}
	}
	return answer
}

func twist(arr [ListSize]int, startAt int, twistSize int) [ListSize]int {
	var answer [ListSize]int
	for i := range answer {
		answer[i] = arr[i]
	}
	//fmt.Println()
	for i := 0; i < twistSize/2; i++ {
		firstElement := (startAt + i) % ListSize
		secondElement := (startAt + (twistSize - 1) - i) % ListSize
		//fmt.Println("swapping element", firstElement, "with value", answer[firstElement],
		//	"with element", secondElement, "with value", answer[secondElement])
		answer[firstElement], answer[secondElement] = answer[secondElement], answer[firstElement]
	}
	//fmt.Println("After twist starting at", startAt, "with size", twistSize, "array is:", answer)
	return answer
}

func applyTwists(arr [ListSize]int, instructions []int) [ListSize]int {
	for _, twistSize := range instructions {
		// do the twist
		arr = twist(arr, startAt, twistSize)
		// determine the next starting point and advance the skip size
		startAt = (startAt + twistSize + skip) % ListSize
		skip++
	}
	return arr
}

func buildIntArray() [ListSize]int {
	var answer [ListSize]int
	for i := range answer {
		answer[i] = i
	}
	return answer
}

func parseInput(s string) []int {
	var answer []int
	for _, b := range s {
		answer = append(answer, int(b))
	}
	// append our final sequence
	answer = append(answer, 17, 31, 73, 47, 23)
	return answer
}
