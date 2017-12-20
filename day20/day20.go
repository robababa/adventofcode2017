package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"math"
)

const NoSolutionValue = -1

type Coordinate struct {
	x,y,z int
}

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

var latestCollisionRound = 0

func main() {
	fmt.Println("Starting day 20 part 2...")
	input := readInput()
	particles := parseInput(input)
	fmt.Println(particles[0])
	loadCollisions(particles)
	removeParticles(particles)
	fmt.Println("The number of particles remaining is", countRemainingParticles(particles))
}

func countRemainingParticles(particles []Particle) int {
	answer := 0
	for _, p := range particles {
		if p.zappedRound == 0 {
			answer++
		}
	}
	return answer
}

func removeParticles(particles []Particle) {
	// loop through the collisions each round
	for i := 1; i <= latestCollisionRound; i++ {
		for _, collision := range collisions[i] {
			if (particles[collision.p1num].zappedRound == 0 || particles[collision.p1num].zappedRound == i) &&
				(particles[collision.p2num].zappedRound == 0 || particles[collision.p2num].zappedRound == i) {
				particles[collision.p1num].zappedRound = i
				particles[collision.p2num].zappedRound = i
			}
		}
	}
}

// returns the positive integer solutions to the quadratic equation
// ax^2 + bx + c = 0
// if there are two solutions, this function returns both of them
// if there is only one positive integer solution, it returns that solution and NoSolutionValue
// if there are no positive integer solutions, the function returns NoSolutionValue, NoSolutionValue
func quadraticPositiveIntegerSolutions(a int, b int, c int) (int, int) {
	discriminant := b * b - 4 * a * c
	// imaginary solutions don't work here
	if discriminant < 0 {
		return NoSolutionValue, NoSolutionValue
	}

	// the discriminant is nonnegative.  However, if its root is not an integer, we have no solutions
	discriminantRoot := int(math.Sqrt(float64(discriminant)))
	if discriminantRoot * discriminantRoot != discriminant {
		return NoSolutionValue, NoSolutionValue
	}

	couldBe := [2]int{((-1 * b) + discriminantRoot) / (2 * a), ((-1 * b) - discriminantRoot) / (2 * a)}
	
	// if the possible solutions don't actually work, set them to NoSolutionValue
	if a * couldBe[0] * couldBe[0] + b * couldBe[0] + c != 0 {
		couldBe[0] = NoSolutionValue
	}

	if a * couldBe[1] * couldBe[1] + b * couldBe[1] + c != 0 {
		couldBe[1] = NoSolutionValue
	}

	// if the solutions are valid and the same, set the second one to NoSolutionValue and return them
	if couldBe[0] >= 0 && couldBe[1] == couldBe[0] {
		couldBe[1] = NoSolutionValue
		return couldBe[0], couldBe[1]
	}

	// if we have two solutions, and they are in descending order, then reorder them
	if couldBe[0] >= 0 && couldBe[1] >= 0 && couldBe[0] > couldBe[1] {
		couldBe[0], couldBe[1] = couldBe[1], couldBe[0]
	}
	
	// now just return what we have
	return couldBe[0], couldBe[1]
}

// a particle is at the origin at time t when its x, y and z positions are all zero
func whenAtOrigin(pcl Particle) (int, int) {
	// the position along an axis is a*t*(t-1) + v*t + p = a*t^2 + (v - a) * t + p
	xRoot1, xRoot2 := quadraticPositiveIntegerSolutions(pcl.a.x, pcl.v.x - pcl.a.x, pcl.p.x)
	yRoot1, yRoot2 := quadraticPositiveIntegerSolutions(pcl.a.y, pcl.v.y - pcl.a.y, pcl.p.y)
	zRoot1, zRoot2 := quadraticPositiveIntegerSolutions(pcl.a.z, pcl.v.z - pcl.a.z, pcl.p.z)

	answer1, answer2 := NoSolutionValue, NoSolutionValue
	if xRoot1 > 0 && (xRoot1 == yRoot1 || xRoot1 == yRoot2) && (xRoot1 == zRoot1 || xRoot1 == zRoot2) {
		answer1 = xRoot1
	}
	if xRoot2 > 0 && (xRoot2 == yRoot1 || xRoot2 == yRoot2) && (xRoot2 == zRoot1 || xRoot2 == zRoot2) {
		answer2 = xRoot2
	}
	return answer1, answer2
}

func whenCollide(p1, p2 Particle) (int, int) {
	difference := Particle{
		p: Coordinate{x: p1.p.x - p2.p.x, y: p1.p.y - p2.p.y, z: p1.p.z - p2.p.z},
		v: Coordinate{x: p1.v.x - p2.v.x, y: p1.v.y - p2.v.y, z: p1.v.z - p2.v.z},
		a: Coordinate{x: p1.a.x - p2.a.x, y: p1.a.y - p2.a.y, z: p1.a.z - p2.a.z},
	}
	return whenAtOrigin(difference)
}

func firstCollide(p1, p2 Particle) int {
	answer1, answer2 := whenCollide(p1, p2)
	if answer1 == NoSolutionValue && answer2 == NoSolutionValue {
		return NoSolutionValue
	} else if answer2 == NoSolutionValue {
		return answer1
	} else if answer1 == NoSolutionValue {
		return answer2
	} else if answer1 < answer2 {
		// return the least of the two answers
		return answer1
	} else {
		return answer2
	}
}

func loadCollisions(particles []Particle) {
	for i, p := range particles {
		for j, p2 := range particles[(i+1):] {
			collisionRound := firstCollide(p, p2)
			if collisionRound != NoSolutionValue {
				collisions[collisionRound] = append(collisions[collisionRound], Collision{i, j})
				if collisionRound > latestCollisionRound {
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