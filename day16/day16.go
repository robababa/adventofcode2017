package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const Letters = "abcdefghijklmnop"
const LetterCount = len(Letters)

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
	//print()
	instructions := readInput()
	part1(instructions)
	fmt.Println("Part 1 answer:")
	print()
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

func print() {
	for i := 0; i < LetterCount; i++ {
		fmt.Print(positionMap[(i + firstPosition) % LetterCount])
	}
	fmt.Println()
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
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
	answer = append(answer, scanner.Text())
	}
	return answer
}