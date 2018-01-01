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
	// guard against invalid instruction caused by toggle
	if arg2.isValue {
		return
	}
	if arg1.isValue {
		registers[arg2.variable] = arg1.value
	} else {
		registers[arg2.variable] = registers[arg1.variable]
	}
}

func inc(arg1 argument) {
	// guard against invalid instruction caused by toggle
	if arg1.isValue {
		return
	}
	registers[arg1.variable]++
}

func dec(arg1 argument) {
	// guard against invalid instruction caused by toggle
	if arg1.isValue {
		return
	}
	registers[arg1.variable]--
}

func jnz(arg1, arg2 argument) int {
	switch {
	case arg1.isValue && arg1.value == 0: {
		return 1
		}
	case !arg1.isValue && registers[arg1.variable] == 0: {
		return 1
		}
	case arg2.isValue: {
		return arg2.value
	}
	default: {
		return registers[arg2.variable]
	}
	}
}

func tgl(instructionNumber int, arg1 argument) {
	toggleTarget := 0
	if arg1.isValue {
		toggleTarget = instructionNumber + arg1.value
	} else {
		toggleTarget = instructionNumber + registers[arg1.variable]
	}
	// if our toggle goes outside the list of instructions, then return
	if toggleTarget < 0 || toggleTarget >= len(instructions) {
		return
	}
	oldInstruction := instructions[toggleTarget].command
	switch oldInstruction {
	case "inc": {instructions[toggleTarget].command = "dec"}
	case "dec": {instructions[toggleTarget].command = "inc"}
	case "tgl": {instructions[toggleTarget].command = "inc"}
	case "jnz": {instructions[toggleTarget].command = "cpy"}
	case "cpy": {instructions[toggleTarget].command = "jnz"}
	}
}

func main() {
	parseInput(readInput())
	//fmt.Println("Instructions are:", instructions)
	registers["a"] = 7
	processInstructions()
	fmt.Println("Part 1: Value of a:", registers["a"])
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
		case "tgl": {
			tgl(instructionNumber, thisInstruction.arg1)
			instructionNumber++
		}
		//fmt.Println("registers are", registers)
		//fmt.Println("instructions are", instructions)
		//fmt.Println("instructionNumber is", instructionNumber)
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