package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
	"sync"
)
/*
snd X snd X sends the value of X to the other program. These values wait in a queue until that program is ready to receive them. Each program has its own message queue, so a program can never receive a message it sent.
set X Y sets register X to the value of Y.
add X Y increases register X by the value of Y.
mul X Y sets register X to the result of multiplying the value contained in register X by the value of Y.
mod X Y sets register X to the remainder of dividing the value contained in register X by the value of Y (that is, it sets X to the result of X modulo Y).
rcv X receives the next value and stores it in register X. If no values are in the queue, the program waits for a value to be sent to it. Programs do not continue to the next instruction until they have received a value. Values are received in the order they are sent.
jgz X Y jumps with an offset of the value of Y, but only if the value of X is greater than zero. (An offset of 2 skips the next instruction, an offset of -1 jumps to the previous instruction, and so on.)
*/

type instruction struct {
	name string
	register string
	argIsOtherRegister bool
	argInt int
	argOtherRegister string
}

var chan0to1 = make(chan int, 100000)
var chan1to0 = make(chan int, 100000)
var sends [2]int

func snd(programNum int, registers map[string]int, inst instruction) {
	if programNum == 0 {
		//fmt.Println("P0 send", registers[inst.register], "len(channel) was", len(chan0to1))
		chan0to1 <- registers[inst.register]
	} else {
		//fmt.Println("P1 send", registers[inst.register], "len(channel) was", len(chan1to0))
		chan1to0 <- registers[inst.register]
	}
}

func set(registers map[string]int, inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] = registers[inst.argOtherRegister]
	} else {
		registers[inst.register] = inst.argInt
	}
	//fmt.Println(registers)
}

func add(registers map[string]int, inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] += registers[inst.argOtherRegister]
	} else {
		registers[inst.register] += inst.argInt
	}
}

func mul(registers map[string]int, inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] *= registers[inst.argOtherRegister]
	} else {
		registers[inst.register] *= inst.argInt
	}
}

func mod(registers map[string]int, inst instruction) {
	if inst.argIsOtherRegister {
		registers[inst.register] %= registers[inst.argOtherRegister]
	} else {
		registers[inst.register] %= inst.argInt
	}
}

func rcv(programNum int, registers map[string]int, inst instruction) {
	if programNum == 0 {
		registers[inst.register] = <-chan1to0
		//fmt.Println("P0 received", registers[inst.register], "new length is", len(chan1to0))
	} else {
		registers[inst.register] = <-chan0to1
		//fmt.Println("P1 received", registers[inst.register], "new length is", len(chan0to1))
	}
	//fmt.Println("Channel lengths:", len(chan0to1), len(chan1to0))
}

func jgz(registers map[string]int, inst instruction) int {
	xValue, err := strconv.Atoi(inst.register)
	if err != nil {
		xValue = registers[inst.register]
	}
	// don't jump if the x value is less than or equal to zero,
	// just go to the next instruction, i.e "jump" forward 1, as usual
	if xValue <= 0 {
		return 1
	}
	if inst.argIsOtherRegister {
		return registers[inst.argOtherRegister]
	} else {
		return inst.argInt
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	instructions := parseInput(readInput())
	//fmt.Println(instructions)
	go func() {
		defer wg.Done()
		processInstructions(0, instructions)
	}()
	go func() {
		defer wg.Done()
		processInstructions(1, instructions)
	}()
	wg.Wait()
}

func processSingleInstruction(programNum int, registers map[string]int, inst instruction) int {
	if inst.name == "jgz" {
		return jgz(registers, inst)
	}
	switch inst.name {
	case "snd": {
		snd(programNum, registers, inst)
		sends[programNum]++
	}
	case "set": {
		set(registers, inst)
	}
	case "add": {
		add(registers, inst)
	}
	case "mul": {
		mul(registers, inst)
	}
	case "mod": {
		mod(registers, inst)
	}
	case "rcv": {
		fmt.Println("SENDS:", sends)
		rcv(programNum, registers, inst)
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

func processInstructions(programNum int, instructions []instruction) {
	var registers = make(map[string]int)
	registers["p"] = programNum
	nextInstruction := 0
	//fmt.Println("Program", programNum, "nextInstruction is:", nextInstruction)
	for nextInstruction < len(instructions) {
		fmt.Println(programNum, ":", printInstruction(instructions[nextInstruction]))
		jump := processSingleInstruction(programNum, registers, instructions[nextInstruction])
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