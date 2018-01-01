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

var nextOut = 0

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

// return true if we are sending out the value that we should, false otherwise
func out(arg1 argument) bool {
	if arg1.isValue && arg1.value == nextOut {
		nextOut = 1 - nextOut
		return true
	}
	if !arg1.isValue && registers[arg1.variable] == nextOut {
		nextOut = 1 - nextOut
		return true
	}
	return false
}


func main() {
	parseInput(readInput())
	for i := 0; i < 40000; i++ {
		initializeRegisters()
		nextOut = 0
		registers["a"] = i
		if processInstructions() {
			fmt.Println("success with register a =", i)
			break
		}
	}
	registers["a"] = 0
	processInstructions()
}

func initializeRegisters() {
	registers["a"] = 0
	registers["b"] = 0
	registers["c"] = 0
	registers["d"] = 0
}

func processInstructions() bool {
	instructionNumber := 0
	weAreGood := true
	valuesOut := 0
	goodEnough := 2000
	for weAreGood && instructionNumber < len(instructions) && valuesOut < goodEnough {
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
		case "out": {
			weAreGood = out(thisInstruction.arg1)
			instructionNumber++
			valuesOut++
		}
		}
	}
	if valuesOut == goodEnough {
		return true
	} else {
		return false
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