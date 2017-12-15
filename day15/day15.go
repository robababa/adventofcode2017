package main

import "fmt"

const StartA = 277 // 65
const StartB = 349 // 8921

const FactorA = 16807
const FactorB = 48271
const Modulus = 2147483647
const NumRounds = 40000000
const compareModulus = 65536

func main() {
	numA := StartA
	numB := StartB
	similarCount := 0
	for generatorRound := 1; generatorRound <= NumRounds; generatorRound++ {
		numA = numA * FactorA % Modulus
		numB = numB * FactorB % Modulus
		//fmt.Println("A is", numA, "and B is", numB)
		if (numA - numB) % compareModulus == 0 {
			//fmt.Println("A and B are similar at the end")
			similarCount++
		}
	}
	fmt.Println("Part 1: similar random number count is", similarCount)
}
