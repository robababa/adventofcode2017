package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

const UpDown = '|'
//const LeftRight = '-'
const Corner = '+'
const Empty = ' '
var letters string

type spot struct {
	direction string
	row int
	column int
}

func main() {
	grid := readInput()
	//fmt.Println("grid:", grid)
	//fmt.Println("grid[0][5]:", string(grid[0][5]))
	steps := 0
	//fmt.Println("Starting column is", findStartingColumn(grid[0]))
	previousSpot := spot{direction: "DOWN", row: -2, column: findStartingColumn(grid[0])}
	currentSpot := spot{direction: "DOWN", row: -1, column: findStartingColumn(grid[0])}
	for currentSpot != previousSpot {
		//fmt.Println("moving...")
		newSpot := move(grid, currentSpot)
		previousSpot = currentSpot
		currentSpot = newSpot
		steps += 1
		//fmt.Println("currentSpot is", currentSpot)
	}
	fmt.Println("Part 1: The letters are", letters)
	// subtract one from steps because the last step kept us in the same place
	fmt.Println("Part 2: Number of steps is", steps - 1)
}

func turn(grid []string, current spot) (spot) {
	var returnValue = current
	switch returnValue.direction {
	case "UP", "DOWN": {
		if returnValue.column > 0 && grid[returnValue.row][returnValue.column - 1] != Empty {
			returnValue.direction = "LEFT"
		} else {
			returnValue.direction = "RIGHT"
		}
	}
	case "LEFT", "RIGHT": {
		if returnValue.row > 0 && grid[returnValue.row - 1][returnValue.column] != Empty {
			returnValue.direction = "UP"
		} else {
			returnValue.direction = "DOWN"
		}
	}
	}
	return returnValue
}

// given the current spot, move to the next spot.
// if we can't move at all, return the same spot as before
func move(grid []string, current spot) (spot) {
	var returnValue spot
	// first, move in the current direction
	switch current.direction {
	case "DOWN": {
			returnValue = spot{direction: "DOWN", row: current.row + 1, column: current.column}
		}
	case "UP": {
		returnValue = spot{direction: "UP", row: current.row - 1, column: current.column}
	}
	case "LEFT": {
		returnValue = spot{direction: "LEFT", row: current.row, column: current.column - 1}
	}
	case "RIGHT": {
		returnValue = spot{direction: "RIGHT", row: current.row, column: current.column + 1}
	}
	}
	//fmt.Println("checking for move off grid...")
	// if we moved off the grid, that's a no-no, so return the original spot
	if returnValue.row < 0 || returnValue.column < 0 {
		return current
	}
	newGridChar := grid[returnValue.row][returnValue.column]
	// if the grid is empty in the new spot, we're done, so return the old spot
	//fmt.Println("checking for empty...")
	if newGridChar == Empty {
		return current
	}
	// if we hit a corner, we need to change direction
	//fmt.Println("checking for corner...")
	if newGridChar == Corner {
		returnValue = turn(grid, returnValue)
	}
	//fmt.Println("checking for letter...")
	// finally, if we landed on a letter, add it to our letters string
	if newGridChar >= 'A' && newGridChar <= 'Z' {
		letters += string(newGridChar)
	}
	//fmt.Println("returnValue is", returnValue)
	return returnValue
}

func findStartingColumn(topLine string) int {
	return strings.Index(topLine, string(UpDown))
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}