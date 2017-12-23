package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"github.com/robababa/quadratic"
)

type Coordinate struct {
	x,y,z int
}

const NotYet = -1

type Particle struct {
	// p=position, v=velocity, a=acceleration
	p, v, a Coordinate
	// which round this particle collided with other particle(s)
	zappedRound int
}

type Collision struct {
	p1num int
	p2num int
}

var collisions = make(map[int][]Collision)

var latestCollisionRound = NotYet

func main() {
	fmt.Println("Starting day 20 part 2...")
	input := readInput()
	particles := parseInput(input)
	//fmt.Println(particles[0])
	loadCollisions(particles)
	removeParticles(particles)
	fmt.Println("The number of particles remaining is", countRemainingParticles(particles))
	//fmt.Println("Particles are:")
	//fmt.Println(particles)
	//fmt.Println("Collisions are:")
	//fmt.Println(collisions)
}

func countRemainingParticles(particles []Particle) int {
	answer := 0
	for _, p := range particles {
		if p.zappedRound == NotYet {
			answer++
		}
	}
	return answer
}

func removeParticles(particles []Particle) {
	// loop through the collisions each round
	for i := 0; i <= latestCollisionRound; i++ {
		fmt.Println("Removing particles in round", i)
		for _, collision := range collisions[i] {
			// if both particles either have NEVER been zapped, or are zapped THIS round, then they can
			// zap each other
			if (particles[collision.p1num].zappedRound == NotYet || particles[collision.p1num].zappedRound == i) &&
				(particles[collision.p2num].zappedRound == NotYet || particles[collision.p2num].zappedRound == i) {
					fmt.Println("Zapping particles", collision.p1num, "and", collision.p2num)
				particles[collision.p1num].zappedRound = i
				particles[collision.p2num].zappedRound = i
			}
		}
	}
}

func netParticle(p1, p2 Particle) Particle {
	return Particle{
		p: Coordinate{x: p1.p.x - p2.p.x, y: p1.p.y - p2.p.y, z: p1.p.z - p2.p.z},
		v: Coordinate{x: p1.v.x - p2.v.x, y: p1.v.y - p2.v.y, z: p1.v.z - p2.v.z},
		a: Coordinate{x: p1.a.x - p2.a.x, y: p1.a.y - p2.a.y, z: p1.a.z - p2.a.z},
		zappedRound: NotYet,
		}
}

func positionFormula(pt Particle, dimension string) (int, int, int) {
	// the position of the particle is given by the function
	// (a)(t)(t+1)/2 + (v)(t) + p
	// we can multiply this whole expression by 2, and the roots will be the same,
	// and all we care about is the roots.  When we do that, we get
	// (a)(t)(t + 1) + (2)(v)(t) + (2)(p)
	// or
	// a*t*t + [2v + a]*t + 2p
	switch dimension {
	case "x": {return pt.a.x, 2 * pt.v.x + pt.a.x, 2 * pt.p.x}
	case "y": {return pt.a.y, 2 * pt.v.y + pt.a.y, 2 * pt.p.y}
	case "z": {return pt.a.z, 2 * pt.v.z + pt.a.z, 2 * pt.p.z}
	// stupid default, but I need to return something
	default: {return pt.a.z, 2 * pt.v.z + pt.a.z, 2 * pt.p.z}
	}
}

// loop over all of the pairs of particles to determine their earliest collision, if one even exists
func loadCollisions(particles []Particle) {
	for i, p1 := range particles {
		for j, p2 := range particles[(i+1):] {
			//fmt.Println("Comparing particles", p1, "and", p2)
			// d = distance, a function of time
			d := netParticle(p1, p2)
			collisionSolution := quadratic.CombineSolutions(
				quadratic.QuadraticPositiveIntegerSolutions(positionFormula(d, "x")),
				quadratic.QuadraticPositiveIntegerSolutions(positionFormula(d, "y")),
				quadratic.QuadraticPositiveIntegerSolutions(positionFormula(d, "z")),
			)
			if collisionSolution.Never() {continue}
			// this special case should never happen, because the input data has no particles at the same
			// starting point
			if collisionSolution.Always() {
				fmt.Println("Round ZERO collision for", p1, p2)
				// j = 0 at particle i +  1
				collisions[0] = append(collisions[0], Collision{p1num: i, p2num: i + 1 + j})
				if latestCollisionRound == NotYet {
					latestCollisionRound = 0
				}
			}
			// sometimes, so we have a single-valued solution
			if collisionSolution.Sometimes() {
				collisionRound := collisionSolution.Values()[0]
				// j = 0 at particle i +  1
				fmt.Println("Round", collisionRound, "collision for particles", i, "and", i + 1 + j)
				collisions[collisionRound] = append(collisions[collisionRound], Collision{p1num: i, p2num: i + 1 + j})
				if latestCollisionRound < collisionRound {
					latestCollisionRound = collisionRound
				}
			}
		}
	}
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
	var pt Particle
	// s is of the form p=<1199,-2918,1457>, v=<-13,115,-8>, a=<-7,8,-10>
	// the ", " is the separator between p, v and a
	for i, coordinates := range strings.Split(s, ", ") {
		// strip off the p=< and > (or v=< and >, or a=< and >)
		strippedCoordinates := coordinates[3:len(coordinates) - 1]
		switch i {
		case 0: pt.p = parseCoordinates(strippedCoordinates)
		case 1: pt.v = parseCoordinates(strippedCoordinates)
		case 2: pt.a = parseCoordinates(strippedCoordinates)
		}	
	}
	pt.zappedRound = NotYet
	return pt
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

// generic helper function that BrianB inspired me to use
func Atoi(s string) int {
	answer, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Could not convert string", s, "to integer")
	}
	return answer
}