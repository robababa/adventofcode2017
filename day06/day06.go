package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arrayInput := convertInput(parseInput(readInput()))
	fmt.Println("The input is", arrayInput)
	reallocate(arrayInput)
}

func reallocate(intArray []int) {
	myArray := intArray
	arrangements := make(map[string]int)
	currentArrangement := convertToString(myArray)
	arrangements[currentArrangement] = 1
	iterations := 0
	arrangementFound := false
	for {
		myArray = reallocateCycle(myArray)
		currentArrangement = convertToString(myArray)
		iterations++
		arrangementFound = arrangements[currentArrangement] != 0
		if arrangementFound {
			break
		}
		arrangements[currentArrangement] = iterations
	}
	fmt.Println("Part number of arrangements before repeating:", iterations)
	fmt.Println("Part 2: cycle length of repetition:", iterations-arrangements[currentArrangement])
}

func reallocateCycle(intArray []int) []int {
	answer := intArray
	maxBank := 0
	maxBlocks := 0
	for i, n := range answer {
		if n > maxBlocks {
			maxBlocks = n
			maxBank = i
		}
	}
	answer[maxBank] = 0
	// reallocate the blocks
	for j := 1; maxBlocks > 0; j++ {
		nextIndex := (maxBank + j) % len(answer)
		answer[nextIndex]++
		maxBlocks--
	}
	return answer
}

func convertToString(intArray []int) string {
	answer := ""
	for _, i := range intArray {
		answer += "," + strconv.Itoa(i)
	}
	fmt.Println("Arrangement is now", answer)
	return answer
}

func convertInput(input []string) []int {
	var answer []int
	for _, s := range input {
		num, _ := strconv.Atoi(s)
		answer = append(answer, num)
	}
	return answer
}

func parseInput(input string) []string {
	return strings.Split(input, "\t")
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	answerWithNewLine, _ := reader.ReadString('\n')
	return answerWithNewLine[0 : len(answerWithNewLine)-1]
}
