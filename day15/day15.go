package main

import "fmt"

const StartA = 277 // 65
const StartB = 349 // 8921

const FactorA = 16807
const FactorB = 48271
const Modulus = 2147483647
const Part1Rounds = 40000000
const Part2Rounds = 5000000
const CompareModulus = 65536

func main() {
	part1()
	part2()
}

func part2() {
	numA := StartA
	numB := StartB
	similarCount := 0
	for generatorRound := 1; generatorRound <= Part2Rounds; generatorRound++ {
		numA = generateNextA(numA, true)
		numB = generateNextB(numB, true)
		//fmt.Println("A is", numA, "and B is", numB)
		if (numA - numB) %CompareModulus == 0 {
			//fmt.Println("A and B are similar at the end")
			similarCount++
		}
	}
	fmt.Println("Part 1: similar random number count is", similarCount)
}


func part1() {
	numA := StartA
	numB := StartB
	similarCount := 0
	for generatorRound := 1; generatorRound <= Part1Rounds; generatorRound++ {
		numA = generateNextA(numA, false)
		numB = generateNextB(numB, false )
		//fmt.Println("A is", numA, "and B is", numB)
		if (numA - numB) %CompareModulus == 0 {
			//fmt.Println("A and B are similar at the end")
			similarCount++
		}
	}
	fmt.Println("Part 2: similar random number count is", similarCount)
}

func generateNextA(currentA int, beChoosy bool) int {
	if !beChoosy {
		return currentA * FactorA % Modulus
	}
	answer := currentA * FactorA % Modulus
	for {
		if answer % 4 == 0 {
			return answer
		} else {
			answer = answer * FactorA % Modulus
		}
	}
}

func generateNextB(currentB int, beChoosy bool) int {
	if !beChoosy {
		return currentB * FactorB % Modulus
	}
	answer := currentB * FactorB % Modulus
	for {
		if answer % 8 == 0 {
			return answer
		} else {
			answer = answer * FactorB % Modulus
		}
	}
}