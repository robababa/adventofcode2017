package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:	day03 <target>")
		os.Exit(0)
	}
	inputNum, _ := strconv.Atoi(os.Args[1])
	fmt.Println("The input is", inputNum)
	myPath := beginPath()
	foundLargerValue := false
	for pointNum := 2; pointNum <= inputNum; pointNum++ {
		myPath.move()
		if !foundLargerValue && myPath.visitedPoints[myPath.currentPoint] > inputNum {
			foundLargerValue = true
			fmt.Println("Part 2: First value larger than input is at point", myPath.currentPoint, "with value",
				myPath.visitedPoints[myPath.currentPoint])
		}
	}
	fmt.Println("Part 1: Final point is", myPath.currentPoint, ", whose Manhattan distance is",
		int(math.Abs(float64(myPath.currentPoint.x)))+int(math.Abs(float64(myPath.currentPoint.y))))
}

type Point struct {
	x int
	y int
}

func (p Point) left() Point {
	return Point{x: p.x - 1, y: p.y}
}

func (p Point) right() Point {
	return Point{x: p.x + 1, y: p.y}
}

func (p Point) up() Point {
	return Point{x: p.x, y: p.y + 1}
}

func (p Point) down() Point {
	return Point{x: p.x, y: p.y - 1}
}

type Path struct {
	currentPoint  Point
	direction     string
	visitedPoints map[Point]int // the int value is the points value, i.e. the sum of it and its visited neighbors
}

func (p *Path) addPointToVisited() {
	// if some of the neighboring points haven't been visited, their map values will be
	// zero, which is fine
	currentPointValue := p.visitedPoints[p.currentPoint.left()] +
		p.visitedPoints[p.currentPoint.right()] +
		p.visitedPoints[p.currentPoint.up()] +
		p.visitedPoints[p.currentPoint.down()] +
		p.visitedPoints[p.currentPoint.left().up()] +
		p.visitedPoints[p.currentPoint.left().down()] +
		p.visitedPoints[p.currentPoint.right().up()] +
		p.visitedPoints[p.currentPoint.right().down()]
	p.visitedPoints[p.currentPoint] = currentPointValue
}

func (p *Path) updateDirection() {
	switch {
	case p.direction == "Up" && p.visitedPoints[p.currentPoint.left()] == 0:
		p.direction = "Left"
	case p.direction == "Down" && p.visitedPoints[p.currentPoint.right()] == 0:
		p.direction = "Right"
	case p.direction == "Left" && p.visitedPoints[p.currentPoint.down()] == 0:
		p.direction = "Down"
	case p.direction == "Right" && p.visitedPoints[p.currentPoint.up()] == 0:
		p.direction = "Up"
	}
}

func (p *Path) move() {
	// change our direction heading if we need to as we spiral out from the origin
	p.updateDirection()
	switch {
	case p.direction == "Up":
		p.currentPoint = p.currentPoint.up()
	case p.direction == "Down":
		p.currentPoint = p.currentPoint.down()
	case p.direction == "Left":
		p.currentPoint = p.currentPoint.left()
	case p.direction == "Right":
		p.currentPoint = p.currentPoint.right()
	}
	p.addPointToVisited()
}

func beginPath() Path {
	answer := Path{
		currentPoint:  Point{x: 0, y: 0},
		direction:     "Down",
		visitedPoints: make(map[Point]int),
	}
	answer.visitedPoints[Point{x: 0, y: 0}] = 1
	return answer
}
