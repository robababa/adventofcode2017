package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)
/*
set X Y sets register X to the value of Y.
sub X Y decreases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
jnz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
*/

type instruction struct {
	name string
	register string
	argIsOtherRegister bool
	argInt int
	argOtherRegister string
}

var mulCalls = 0

func set(registers map[string]int, inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] = registers[inst.argOtherRegister]
	} else {
		registers[inst.register] = inst.argInt
	}
	//fmt.Println(registers)
}

func sub(registers map[string]int, inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] -= registers[inst.argOtherRegister]
	} else {
		registers[inst.register] -= inst.argInt
	}
}

func mul(registers map[string]int, inst instruction) {
	mulCalls++
	if inst.argIsOtherRegister {
		registers[inst.register] *= registers[inst.argOtherRegister]
	} else {
		registers[inst.register] *= inst.argInt
	}
}

func jnz(registers map[string]int, inst instruction) int {
	xValue, err := strconv.Atoi(inst.register)
	if err != nil {
		xValue = registers[inst.register]
	}
	// don't jump if the x value is zero,
	// just go to the next instruction, i.e "jump" forward 1, as usual
	if xValue == 0 {
		return 1
	}
	if inst.argIsOtherRegister {
		return registers[inst.argOtherRegister]
	} else {
		return inst.argInt
	}
}

func main() {
	instructions := parseInput(readInput())
	//fmt.Println(instructions)
	processInstructions(instructions)
	fmt.Println("Part 1: mul calls is", mulCalls)
}

func processSingleInstruction(registers map[string]int, inst instruction) int {
	if inst.name == "jnz" {
		return jnz(registers, inst)
	}
	switch inst.name {
	case "set": {
		set(registers, inst)
	}
	case "sub": {
		sub(registers, inst)
	}
	case "mul": {
		mul(registers, inst)
	}
	}
	return 1
}

func printInstruction(inst instruction) string {
	answer := inst.name + " " + inst.register
	if inst.name == "snd" || inst.name == "rcv" {
		return answer
	}
	var lastArg string
	if inst.argIsOtherRegister {
		lastArg = inst.argOtherRegister
	} else {
		lastArg = strconv.Itoa(inst.argInt)
	}
	return answer + " " + lastArg
}

func processInstructions(instructions []instruction) {
	var registers = make(map[string]int)
	nextInstruction := 0
	//fmt.Println("nextInstruction is:", nextInstruction)
	for nextInstruction < len(instructions) {
		fmt.Println(printInstruction(instructions[nextInstruction]))
		jump := processSingleInstruction(registers, instructions[nextInstruction])
		//fmt.Println(registers)
		nextInstruction += jump
		//fmt.Println("Program", programNum, "nextInstruction is:", nextInstruction)
	}
}

func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func parseInput(lines []string) []instruction {
	var instructions []instruction
	for _, s := range lines {
		tokens := strings.Split(s, " ")
		argIsOtherRegister := false
		argInt := 0
		argOtherRegister := ""
		if len(tokens) > 2 {
			if isInteger(tokens[2]) {
				argInt, _ = strconv.Atoi(tokens[2])
			} else {
				argIsOtherRegister = true
				argOtherRegister = tokens[2]
			}
		}
		myInstruction := instruction{
			name: tokens[0],
			register: tokens[1],
			argIsOtherRegister: argIsOtherRegister,
			argInt: argInt,
			argOtherRegister: argOtherRegister,
		}
		instructions = append(instructions, myInstruction)
	}
	return instructions
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}