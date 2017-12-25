package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"log"
	"strings"
)

type beam struct {
	port1 int
	port2 int
}

type beamStock struct {
	piece beam
	available bool
}

type bridgeAndStrength struct {
	bridge []beam
	strength int
}

var strongestBridge = bridgeAndStrength{}

func main() {
	stockpile := parseInput(readInput())
	//fmt.Println("beam stockpile is:")
	//fmt.Println(stockpile)
	// initialize our bridge with a 0/0 beam, which won't alter its strength, but will make the
	// recursion easier
	startingBeam := beam{port1: 0, port2: 0}
	startingBridge := []beam{startingBeam}
	buildStrongest(stockpile, startingBridge)
	// 1287 is too low
	fmt.Println("Part 1: strongest bridge is", strongestBridge)
	fmt.Println("Part 1: length of strongest bridge is", len(strongestBridge.bridge) - 1)
	fmt.Println("Part 1: strength of strongest bridge is", strongestBridge.strength)
}

// if nextBeam fits the bridge, return it (in the correct direction) and true
// otherwise return an empty beam and false
func matchingBeam(bridge []beam, nextBeam beam) (beam, bool) {
	lastPort := bridge[len(bridge)-1].port2
	if nextBeam.port1 == lastPort {
		return nextBeam, true
	} else if nextBeam.port2 == lastPort {
		return beam{port1: nextBeam.port2, port2: nextBeam.port1}, true
	}
	// doesn't match either way, return empty beam and false
	return beam{}, false
}

func buildStrongest(stockpile []beamStock, bridge []beam) {
	for i, b := range stockpile {
		// if this next beam hasn't been used already
		if b.available {
			// and it fits
			if nextBeam, fits := matchingBeam(bridge, b.piece); fits {
				// extend the bridge with this beam, mark it unavailable in the stockpile, and recursively
				// call this function again
				updatedStockpile := append(stockpile)
				updatedStockpile[i].available = false
				buildStrongest(updatedStockpile, append(bridge, nextBeam))
			}
		}
	}
	if bridgeStrength := strength(bridge); bridgeStrength > strongestBridge.strength {
		strongestBridge = bridgeAndStrength{bridge: bridge, strength: bridgeStrength}
		fmt.Println("Strongest bridge now is:", strongestBridge)
	}
}

func strength(bridge []beam) int {
	answer := 0
	for _, b := range bridge {
		answer += b.port1 + b.port2
	}
	return answer
}

func parseInput(lines []string) []beamStock {
	var b []beamStock
	for _, line := range lines {
		//append each beam to the stockpile as available
		tokens := strings.Split(line, "/")
		b = append(b, beamStock{piece: beam{port1: Atoi(tokens[0]), port2: Atoi(tokens[1])}, available: true})
	}
	return b
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}

func Atoi(s string) int {
	answer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Could not convert string", s, "to integer")
	}
	return answer
}