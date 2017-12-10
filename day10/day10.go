package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
	"encoding/hex"
)

const LIST_SIZE = 256
const PART2_ROUNDS = 64

var skip = 0
var startAt = 0

func main() {
	input := readInput()
	part1instructions := part1ParseInput(input)
	//fmt.Println("Part 1 instructions are:", part1instructions)
	arrPart1 := buildIntArray()
	arrPart1 = applyTwists(arrPart1, part1instructions)
	fmt.Println("Part 1: Product of first two elements is:", arrPart1[0] * arrPart1[1])
	// reset skip and startAt for part 2
	skip = 0
	startAt = 0
	part2instructions := part2ParseInput(input)
	//fmt.Println("Part 2 instructions are:", part2instructions)
	arrPart2 := buildIntArray()
	for i := 0; i < PART2_ROUNDS; i++ {
		arrPart2 = applyTwists(arrPart2, part2instructions)
	}
	//fmt.Println("Part 2: sparse hash is:", arrPart2)
	part2DenseHash := denseHash(arrPart2)
	//fmt.Println("Part 2: dense hash is:", part2DenseHash)
	fmt.Println("Part 2: Knot Hash is:", denseHashToHexString(part2DenseHash))
}

func denseHashToHexString(input [16]int) string {
	var bytes []byte
	for i := 0; i<16; i++ {
		bytes = append(bytes, byte(input[i]))
	}
	return hex.EncodeToString(bytes)
}

func denseHash(input [LIST_SIZE]int) [16]int {
	var answer [16]int
	for group:=0; group<16; group++ {
		answer[group] = input[group * 16]
		for elementWithinGroup := 1; elementWithinGroup < 16; elementWithinGroup++ {
			answer[group] ^= input[group * 16 + elementWithinGroup]
		}
	}
	return answer
}

func twist(arr [LIST_SIZE]int, startAt int, twistSize int) [LIST_SIZE]int {
	var answer [LIST_SIZE]int
	for i := range answer {
		answer[i] = arr[i]
	}
	//fmt.Println()
	for i := 0; i < twistSize/2; i++ {
		firstElement := (startAt + i) % LIST_SIZE
		secondElement := (startAt + (twistSize - 1) - i) % LIST_SIZE
		//fmt.Println("swapping element", firstElement, "with value", answer[firstElement],
		//	"with element", secondElement, "with value", answer[secondElement])
		answer[firstElement], answer[secondElement] = answer[secondElement], answer[firstElement]
	}
	//fmt.Println("After twist starting at", startAt, "with size", twistSize, "array is:", answer)
	return answer
}

func applyTwists(arr [LIST_SIZE]int, instructions []int) [LIST_SIZE]int {
	for _, twistSize := range instructions {
		// do the twist
		arr = twist(arr, startAt, twistSize)
		// determine the next starting point and advance the skip size
		startAt = (startAt + twistSize + skip) % LIST_SIZE
		skip++
	}
	return arr
}

func buildIntArray() [LIST_SIZE]int {
	var answer [LIST_SIZE]int
	for i := range answer {
		answer[i] = i
	}
	return answer
}

func part1ParseInput(s string) []int {
	var answer []int
	for _, s := range strings.Split(s, ",") {
		i, _ := strconv.Atoi(s)
		answer = append(answer, i)
	}
	return answer
}

func part2ParseInput(s string) []int {
	var answer []int
	for _, b := range s {
		answer = append(answer, int(b))
	}
	// append our final sequence
	answer = append(answer, 17, 31, 73, 47, 23)
	return answer
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	answerWithNewLine, _ := reader.ReadString('\n')
	return answerWithNewLine[0 : len(answerWithNewLine)-1]
}
