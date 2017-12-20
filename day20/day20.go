package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
)

type Coordinate struct {
	x,y,z int
}

type Particle struct {
	position,velocity,acceleration Coordinate
}

func main() {
	fmt.Println("Starting day 20 part 2...")
	input := readInput()
	particles := parseInput(input)
	fmt.Println(particles[0])
}

func Atoi(s string) int {
	answer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Could not convert string", s, "to integer")
	}
	return answer
}

func parseCoordinates(s string) Coordinate {
	var answer Coordinate
	for i, point := range strings.Split(s, ",") {
		switch i {
		case 0: answer.x = Atoi(point)
		case 1: answer.y = Atoi(point)
		case 2: answer.z = Atoi(point)
		}
	}
	return answer
}

func parseParticle(s string) Particle {
	var p Particle
	// s is of the form p=<1199,-2918,1457>, v=<-13,115,-8>, a=<-7,8,-10>
	// the ", " is the separator between position, velocity and acceleration
	for i, coordinates := range strings.Split(s, ", ") {
		// strip off the p=< and > (or v=< and >, or a=< and >)
		strippedCoordinates := coordinates[3:len(coordinates) - 1]
		switch i {
		case 0: p.position = parseCoordinates(strippedCoordinates)
		case 1: p.velocity = parseCoordinates(strippedCoordinates)
		case 2: p.acceleration = parseCoordinates(strippedCoordinates)
		}	
	}
	return p
}

func parseInput(input []string) []Particle {
	var answer []Particle
	for _, s := range input {
		//fmt.Println("Processing line", s)
		answer = append(answer, parseParticle(s))
	}
	return answer
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, scanner.Text())
	}
	return answer
}
