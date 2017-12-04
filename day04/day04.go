package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)

func main() {
	phrases := readInput()
	fmt.Println("Part 1: number of valid phrases is", countValid(phrases, testPart1))
	fmt.Println("Part 2: number of valid phrases is", countValid(phrases, testPart2))
}

func countValid(phrases []string, validateFunc func(string) bool) int {
	validCount := 0
	for _, phrase := range phrases {
		if validateFunc(phrase) {
			validCount++
		}
	}
	return validCount
}

func testPart1(phrase string) bool {
	words := strings.Split(phrase, " ")
	numWords := len(words)
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}
	distinctWords := len(wordMap)
	return numWords == distinctWords
}

func sortWord(word string) string {
	var myInts []int
	for _, b := range []byte(word) {
		myInts = append(myInts, int(b))
	}
	sort.Ints(myInts)
	var sortedBytes []byte
	for _, i := range myInts {
		sortedBytes = append(sortedBytes, byte(i))
	}
	return string(sortedBytes)
}

func testPart2(phrase string) bool {
	words := strings.Split(phrase, " ")
	numWords := len(words)
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[sortWord(word)] = true
	}
	distinctWords := len(wordMap)
	return numWords == distinctWords
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		answer = append(answer, input)
	}
	return answer
}
