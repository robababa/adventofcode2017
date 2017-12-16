package main

import "fmt"

const Letters = "abcde"
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
	print()
	spin(1)
	print()
	exchange(3, 4)
	print()
	swap("e", "b")
	print()
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

func swap(letter1 string, letter2 string) {
	letterMap[letter1], letterMap[letter2] = letterMap[letter2], letterMap[letter1]
	positionMap[letterMap[letter1]] = letter1
	positionMap[letterMap[letter2]] = letter2
}