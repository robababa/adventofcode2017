package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"encoding/csv"
)

const Letters = "abcdefghijklmnop"
const LetterCount = len(Letters)
const Dances = 1000000000

var firstPosition = 0
var letterMap = make(map[string]int)
var positionMap = make(map[int]string)

func init() {
	for i, c := range Letters {
		letterMap[string(c)] = i
		positionMap[i] = string(c)
	}
}

func main() {
	printResults()
	instructions := readInput()
	part1(instructions)
	fmt.Println("Part 1 answer:")
	printResults()

	cyclesAt := 0
	for i := 2;; i++ {
		part1(instructions)
		if danceResults() == Letters {
			fmt.Println("It's a cycle at iteration", i)
			cyclesAt = i
			break
		}
	}
	for j := 1; j <= Dances % cyclesAt; j++ {
		part1(instructions)
	}
	fmt.Println("Part 2 answer:")
	printResults()
}

func part1(instructions []string) {
	for _, instruction := range instructions {
		operation := instruction[0]
		operands := instruction[1:]
		switch operation {
		case 's': {
			spinSize, _ := strconv.Atoi(operands)
			spin(spinSize)
		}
		case 'x': {
			positions := strings.Split(operands, "/")
			position1, _ := strconv.Atoi(positions[0])
			position2, _ := strconv.Atoi(positions[1])
			exchange(position1, position2)

		}
		case 'p': {
			partners := strings.Split(operands, "/")
			partner(partners[0], partners[1])
		}
		}
	}
}

func danceResults() string {
	var answer string
	for i := 0; i < LetterCount; i++ {
		answer += positionMap[(i + firstPosition) % LetterCount]
	}
	return answer
}

func printResults() {
	fmt.Println(danceResults())
}

func spin(spinSize int) {
	firstPosition = (firstPosition + LetterCount - spinSize) % LetterCount
}

func exchange(position1 int, position2 int) {
	actual1 := (position1 + firstPosition) % LetterCount
	actual2 := (position2 + firstPosition) % LetterCount
	positionMap[actual1], positionMap[actual2] = positionMap[actual2], positionMap[actual1]
	letterMap[positionMap[actual1]] = actual1
	letterMap[positionMap[actual2]] = actual2
}

func partner(letter1 string, letter2 string) {
	letterMap[letter1], letterMap[letter2] = letterMap[letter2], letterMap[letter1]
	positionMap[letterMap[letter1]] = letter1
	positionMap[letterMap[letter2]] = letter2
}

func readInput() []string {
	reader := csv.NewReader(os.Stdin)
	answer, _ := reader.Read()
	return answer
}