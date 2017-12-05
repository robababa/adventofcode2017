package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	textInput := readInput()
	//fmt.Println("Text input is", textInput)
	intInput1 := convertSlice(textInput)
	//fmt.Println("Integer input is", intInput)
	jumps1 := jumpAround1(intInput1)
	fmt.Println("Part 1 total number of jumps:", jumps1)
	intInput2 := convertSlice(textInput)
	//fmt.Println("Integer input is", intInput)
	jumps2 := jumpAround2(intInput2)
	fmt.Println("Part 2 total number of jumps:", jumps2)
}

func printOffsets(intInput []int, curPosition int) {
	fmt.Print("Current state of jump instructions:")
	for idx, num := range intInput {
		if curPosition == idx {
			fmt.Print(" " + "(" + strconv.Itoa(num) + ")")
		} else {
			fmt.Print(" " + strconv.Itoa(num))
		}
	}
	fmt.Println("")
}

func jumpAround1(intInput []int) int {
	position := 0
	jumps := 0
	for position < len(intInput) {
		//printOffsets(intInput, position)
		position, intInput[position] = position+intInput[position], intInput[position]+1
		jumps++
	}
	return jumps
}

func jumpAround2(intInput []int) int {
	position := 0
	jumps := 0
	for position < len(intInput) {
		//printOffsets(intInput, position)
		prevPosition := position
		position = position + intInput[position]
		if intInput[prevPosition] < 3 {
			intInput[prevPosition]++
		} else {
			intInput[prevPosition]--
		}
		jumps++
	}
	return jumps
}

func convertSlice(textInput []string) []int {
	var answer []int
	for _, str := range textInput {
		num, _ := strconv.Atoi(str)
		answer = append(answer, num)
	}
	return answer
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
