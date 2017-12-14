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
	}
	fmt.Println("Part 1: Number of ones:", numberOfOnes)
}

func countOnes(input [16]int) int {
	answer := 0
	for _, n := range input {
		answer += len(strings.Replace(strconv.FormatInt(int64(n), 2), "0", "", -1))
	}
	return answer
}

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
