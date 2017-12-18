package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)
/*
snd X plays a sound with a frequency equal to the value of X.
set X Y sets register X to the value of Y.
add X Y increases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
rcv X recovers the frequency of the last sound played, but only when the value of X is not zero. (If it is zero, the command does nothing.)
jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
*/

type instruction struct {
	name string
	register string
	argIsOtherRegister bool
	argInt int
	argOtherRegister string
}

var lastSound = 0
var recoveredSound = 0
var haveRecoveredASound = false

var registers = make(map[string]int)

func snd(inst instruction) {
	//fmt.Println("Ring!", registers[inst.register])
	lastSound = registers[inst.register]
}

func set(inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] = registers[inst.argOtherRegister]
	} else {
		registers[inst.register] = inst.argInt
	}
	//fmt.Println(registers)
}

func add(inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] += registers[inst.argOtherRegister]
	} else {
		registers[inst.register] += inst.argInt
	}
}

func mul(inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] *= registers[inst.argOtherRegister]
	} else {
		registers[inst.register] *= inst.argInt
	}
}

func mod(inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] %= registers[inst.argOtherRegister]
	} else {
		registers[inst.register] %= inst.argInt
	}
}

func rcv(inst instruction) {
	if registers[inst.register] > 0 {
		recoveredSound = lastSound
		if !haveRecoveredASound {
			haveRecoveredASound = true
			fmt.Println("Part 1: first recovered sound:", recoveredSound)
		}
	}
}

func jgz(inst instruction) int {
	// don't jump if the register value is less than or equal to zero,
	// just go to the next instruction, i.e "jump" forward 1, as usual
	if registers[inst.register] <= 0 {
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
}

func processSingleInstruction(inst instruction) int {
	if inst.name == "jgz" {
		return jgz(inst)
	}
	switch inst.name {
	case "snd": {
		snd(inst)
	}
	case "set": {
		set(inst)
	}
	case "add": {
		add(inst)
	}
	case "mul": {
		mul(inst)
	}
	case "mod": {
		mod(inst)
	}
	case "rcv": {
		rcv(inst)
	}
	}
	return 1
}

func processInstructions(instructions []instruction) {
	nextInstruction := 0
	for nextInstruction < len(instructions) {
		jump := processSingleInstruction(instructions[nextInstruction])
		//fmt.Println(registers)
		nextInstruction += jump
		//fmt.Println("nextInstruction is:", nextInstruction)
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