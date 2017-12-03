package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type gridPosition struct {
	East      int
	North     int
	Direction byte // N,E,S,W
}

var pos gridPosition = gridPosition{East: 0, North: 0, Direction: 'N'}

func (pos *gridPosition) turnLeft() {
	switch pos.Direction {
	case 'N':
		pos.Direction = 'W'
	case 'W':
		pos.Direction = 'S'
	case 'S':
		pos.Direction = 'E'
	case 'E':
		pos.Direction = 'N'
	}
}

func (pos *gridPosition) moveForward(steps int) {
	switch pos.Direction {
	case 'N':
		pos.North += steps
	case 'W':
		pos.East -= steps
	case 'S':
		pos.North -= steps
	case 'E':
		pos.East += steps
	}
}

func (pos *gridPosition) applyInstructions(instructions []string) {
	for _, instruction := range instructions {
		if instruction[0] == 'L' {
			pos.turnLeft()
		} else if instruction[0] == 'R' {
			// three lefts is a right
			pos.turnLeft()
			pos.turnLeft()
			pos.turnLeft()
		} else {
			fmt.Println("Bad instruction, first character is", instructions[0])
			log.Panic("Cannot apply move instruction:", instruction)
		}
		steps, _ := strconv.Atoi(instruction[1:])
		pos.moveForward(steps)
		fmt.Println("After instruction", instruction, ", new position is", pos.North, "North and", pos.East, "East")
	}
}

func main() {
	input := readInput()
	fmt.Println("input is", input)
	instructions := parseInput(input)
	fmt.Println("instructions are", instructions)
	//fmt.Println("number of instructions", len(instructions))
	(&pos).applyInstructions(instructions)
	fmt.Println("Final position is", pos.North, "blocks North and", pos.East, "blocks East")
	fmt.Println("Final distance from origin is", int(math.Abs(float64(pos.North))+math.Abs(float64(pos.East))))
}

func parseInput(input string) []string {
	return strings.Split(strings.Replace(input, " ", "", -1), ",")
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	answerWithNewLine, _ := reader.ReadString('\n')
	return answerWithNewLine[0 : len(answerWithNewLine)-1]
}
