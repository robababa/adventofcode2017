package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

type argument struct {
	isValue bool
	value int
	variable string
}

type instruction struct {
	command string
	arg1 argument
	arg2 argument
}

var instructions []instruction

// since all registers default to 0, I don't need to initialize the values for a, b, c, d
var registers = make(map [string]int)

func cpy(arg1, arg2 argument) {
	if arg1.isValue {
		registers[arg2.variable] = arg1.value
	} else {
		registers[arg2.variable] = registers[arg1.variable]
	}
}

func inc(arg1 argument) {
	registers[arg1.variable]++
}

func dec(arg1 argument) {
	registers[arg1.variable]--
}

func jnz(arg1, arg2 argument) int {
	if arg1.isValue && arg1.value != 0 {
		return arg2.value
	}
	firstValue := registers[arg1.variable]
	if firstValue != 0 {
		return arg2.value
	}
	return 1
}

//func tgl() {
//
//}

func main() {
	parseInput(readInput())
	fmt.Println("Instructions are:", instructions)
	processInstructions()
	fmt.Println("Part 1: Value in register a:", registers["a"])
	initPart2()
	processInstructions()
	fmt.Println("Part 2: Value in register a:", registers["a"])

}

func initPart2() {
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 1
	registers["d"] = 0
}

func processInstructions() {
	instructionNumber := 0
	for instructionNumber < len(instructions) {
		thisInstruction := instructions[instructionNumber]
		switch thisInstruction.command {
		case "cpy": {
			cpy(thisInstruction.arg1, thisInstruction.arg2)
			instructionNumber++
		}
		case "inc": {
			inc(thisInstruction.arg1)
			instructionNumber++
		}
		case "dec": {
			dec(thisInstruction.arg1)
			instructionNumber++
		}
		case "jnz": {
			instructionNumber = instructionNumber + jnz(thisInstruction.arg1, thisInstruction.arg2)
		}
		}
	}
}

func parseArgument(s string) argument {
	argValue, err := strconv.Atoi(s)
	if err == nil {
		return argument{isValue: true, value: argValue, variable: ""}
	} else {
		return argument{isValue: false, value: 0, variable: s}
	}
}

func parseInstruction(tokens []string) instruction {
	answer := instruction{
		command: tokens[0],
		arg1: parseArgument(tokens[1]),
	}
	if len(tokens) == 3 {
		answer.arg2 = parseArgument(tokens[2])
	}
	return answer
}

func parseInput(lines []string) {
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		instructions = append(instructions, parseInstruction(tokens))
	}
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}