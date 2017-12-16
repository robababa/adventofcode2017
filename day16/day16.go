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
const Dances = 1000000000

var firstPosition = 0
var letterMap = make(map[string]int)
var positionMap = make(map[int]string)

// each array slot i is the result of doing 2^i dances
var dances []map[string]string

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
	buildRound0(danceResults())
	fmt.Println(dances)
	buildFutureRounds()
	//buildDances()
}

func buildRound0(danceResults string) {
	//fmt.Println("danceResults is", danceResults)
	dances = append(dances, make(map[string]string))
	for _, b := range Letters {
		c := string(b)
		// after 2^0 = 1 rounds of dance, this character maps to this other character
		//fmt.Println("character", c, "is in danceResults in position", strings.IndexAny(danceResults, c))
		dances[0][c] = string(Letters[strings.IndexAny(danceResults, c)])
	}
}

func buildFutureRounds() {
	d := Dances
	for d > 0 {
		previousRound := len(dances) - 1
		dances = append(dances, make(map[string]string))
		currentRound := previousRound + 1
		for _, b := range Letters {
			c := string(b)
			dances[currentRound][c] = dances[previousRound][dances[previousRound][c]]
		}
		d /= 2
	}
	fmt.Println("After building future rounds, dances are:")
	fmt.Println(dances)
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

func print() {
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
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
	answer = append(answer, scanner.Text())
	}
	return answer
}