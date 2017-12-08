package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	varToChange   string
	operator      string
	changeAmount  int
	varToCompare  string
	comparison    string
	compareAmount int
}

var registers = make(map[string]int)

func main() {
	lines := readInput()
	//fmt.Println(lines)
	commands := parseCommands(lines)
	//fmt.Println(commands)
	processInstructions(commands)
	//fmt.Println(registers)
	findLargest()
}

func findLargest() {
	var largestName string
	var largestValue int
	for k, v := range registers {
		if largestName == "" {
			largestName = k
			largestValue = v
		} else if v > largestValue {
			largestName = k
			largestValue = v
		}
	}
	fmt.Println("Largest name and value is", largestName, largestValue)
}

func condition(instruction Command) bool {
	v := registers[instruction.varToCompare]
	comparison := instruction.comparison
	amt := instruction.compareAmount
	return (comparison == "!=" && v != amt) ||
		(comparison == "<" && v < amt) ||
		(comparison == "<=" && v <= amt) ||
		(comparison == "==" && v == amt) ||
		(comparison == ">" && v > amt) ||
		(comparison == ">=" && v >= amt)
}

func processInstruction(instruction Command) {
	if !condition(instruction) {
		return
	}
	if instruction.operator == "inc" {
		registers[instruction.varToChange] += instruction.changeAmount
	} else if instruction.operator == "dec" {
		registers[instruction.varToChange] -= instruction.changeAmount

	}
}

func processInstructions(commands []Command) {
	for _, instruction := range commands {
		processInstruction(instruction)
	}
}

func parseCommand(line string) Command {
	//fmt.Println("line is:", line)
	tokens := strings.Split(line, " ")
	changeAmount, _ := strconv.Atoi(tokens[2])
	compareAmount, _ := strconv.Atoi(tokens[6])
	answer := Command{
		varToChange:   tokens[0],
		operator:      tokens[1],
		changeAmount:  changeAmount,
		varToCompare:  tokens[4],
		comparison:    tokens[5],
		compareAmount: compareAmount,
	}
	return answer
}

func parseCommands(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		commands = append(commands, parseCommand(line))
	}
	return commands
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}
