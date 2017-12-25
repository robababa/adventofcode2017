package main

import "fmt"

//Begin in state A.
//Perform a diagnostic checksum after 12386363 steps.

const Iterations = 12386363
const A = 'A'
const B = 'B'
const C = 'C'
const D = 'D'
const E = 'E'
const F = 'F'

var state = A
var tape = []int{0}
var position = 0

func computeChecksum() int {
	answer := 0
	for i := 0; i < len(tape); i++ {
		if tape[i] == 1 {answer++}
	}
	return answer
}

func currentState() rune {return state}

func currentValue() int {return tape[position]}

func moveLeft() {
	if position > 0 {
		position--
		return
	}
	// position is at zero, so we have to insert a slot on the tape at the beginning, then move to it
	tape = append([]int{0}, tape...)
	// "moving" is a no-op, because our position was zero, and still is, just now at the new position 0
}

func moveRight() {
	if position < len(tape) - 1 {
		position++
		return
	}
	// position is at the end of the tape, so we have to extend it and move there
	tape = append(tape, 0)
	position++
}

func writeValue(val int) {tape[position] = val}

func setState(newState rune) {state = newState}

func main() {
	for i := 0; i < Iterations; i++ {
		switch currentState() {
		case A:
			if currentValue() == 0 {
				writeValue(1)
				moveRight()
				setState(B)
			} else {
				writeValue(0)
				moveLeft()
				setState(E)
			}
		case B:
			if currentValue() == 0 {
				writeValue(1)
				moveLeft()
				setState(C)
			} else {
				writeValue(0)
				moveRight()
				setState(A)
			}
		case C:
			if currentValue() == 0 {
				writeValue(1)
				moveLeft()
				setState(D)
			} else {
				writeValue(0)
				moveRight()
				setState(C)
			}
		case D:
			if currentValue() == 0 {
				writeValue(1)
				moveLeft()
				setState(E)
			} else {
				writeValue(0)
				moveLeft()
				setState(F)
			}
		case E:
			if currentValue() == 0 {
				writeValue(1)
				moveLeft()
				setState(A)
			} else {
				writeValue(1)
				moveLeft()
				setState(C)
			}
		case F:
			if currentValue() == 0 {
				writeValue(1)
				moveLeft()
				setState(E)
			} else {
				writeValue(1)
				moveRight()
				setState(A)
			}
		}
	}
	fmt.Println("Part 1: the checksum is", computeChecksum())
}