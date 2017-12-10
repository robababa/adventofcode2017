package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

const LIST_SIZE = 256

func main() {
	instructions := parseInput(readInput())
	//fmt.Println("Instructions are:", instructions)
	arr := buildIntArray()
	//fmt.Println("Before twists, array is:", arr)
	arr = applyTwists(arr, instructions)
	fmt.Println("Part 1: Product of first two elements is:", arr[0] * arr[1])
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
	skip := 0
	startAt := 0
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

func parseInput(s string) []int {
	var answer []int
	for _, s := range strings.Split(s, ",") {
		i, _ := strconv.Atoi(s)
		answer = append(answer, i)
	}
	return answer
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	answerWithNewLine, _ := reader.ReadString('\n')
	return answerWithNewLine[0 : len(answerWithNewLine)-1]
}
