package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Algorithm:
Read lines of input into an array of strings until we read a . character on a line by itself
Convert the array of strings into an two-dimensional array of integers
Read each row of integers to find the smallest and largest
Add up these differences for the answer
*/

func main() {
	stringInput := readInput()
	//fmt.Println("stringInput is", stringInput)
	//fmt.Println("len(stringInput) is", len(stringInput))
	spreadsheet := convertInput(stringInput)
	fmt.Println("spreadsheet is", spreadsheet)
	fmt.Println("row-based checksum is", computeRowBasedChecksum(spreadsheet))
}

func computeRowBasedChecksum(spreadsheet [][]int) int {
	answer := 0
	for _, row := range spreadsheet {
		answer += computeRowChecksum(row)
	}
	return answer
}

func computeRowChecksum(row []int) int {
	starting := true
	var minValue int
	var maxValue int
	for _, num := range row {
		if starting {
			minValue, maxValue = num, num
			starting = false
		} else if num < minValue {
			minValue = num
		} else if num > maxValue {
			maxValue = num
		}
	}
	return maxValue - minValue
}

func convertInput(inputSlice []string) [][]int {
	var answer [][]int
	for _, el := range inputSlice {
		stringNums := strings.Split(el, "\t")
		var intLine []int
		for _, el := range stringNums {
			intNum, _ := strconv.Atoi(el)
			intLine = append(intLine, intNum)
		}
		answer = append(answer, intLine)
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
