package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
	"strings"
)

var m = make(map[string]int)

var realPart = 0
var imaginaryPart = 0

func main() {
	parseInput(readInput())
	//printParts()
	fmt.Println("Part 1: Distance is", computeDistance())
}

func computeDistance() int {
	absRealPart := int(math.Abs(float64(realPart)))
	absImaginaryPart := int(math.Abs(float64(imaginaryPart)))
	// if the real part is greater, then we can move adjust the imaginary part up and down as needed without making
	// any extra moves
	// else the imaginary part is greater, and we can move along the diagonal "realPart" times, and then move up or down
	// as needed, but each move up or down adjusts our imaginary part by 2, so we divide the difference by 2 to get
	// the number of remaining moves
	if absRealPart > absImaginaryPart {
		return absRealPart
	} else {
		return absRealPart + (absImaginaryPart - absRealPart)/2
	}
}

//func printParts() {
//	fmt.Println("realPart is", realPart, "and imaginaryPart is", imaginaryPart)
//}

func parseInput(str string) {
	greatestDistance := 0
	currentDistance := 0
	for _, c := range strings.Split(str, ",") {
		switch c {
		case "n":
			imaginaryPart += 2
		case "s":
			imaginaryPart -= 2
		case "ne":
			{
				realPart += 1
				imaginaryPart += 1
			}
		case "nw":
			{
				realPart -= 1
				imaginaryPart += 1
			}
		case "sw":
			{
				realPart -= 1
				imaginaryPart -= 1
			}
		case "se":
			{
				realPart += 1
				imaginaryPart -= 1
			}
		}
		currentDistance = computeDistance()
		if currentDistance > greatestDistance {
			greatestDistance = currentDistance
		}
	}
	fmt.Println("Part 2: Greatest distance is", greatestDistance)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	return answer[0:len(answer)-1]
}
