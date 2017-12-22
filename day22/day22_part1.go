package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	//"time"
	//"os/exec"
)

var grid [][]byte
var infections = 0

type virus struct {
	x int
	y int
	heading byte
}

const Infected = '#'
const Clean = '.'

func turnLeft(v *virus) {
	switch v.heading {
	case 'U': v.heading = 'L'
	case 'L': v.heading = 'D'
	case 'D': v.heading = 'R'
	case 'R': v.heading = 'U'
	}
}

func turnRight(v *virus) {
	// three rights make a left
	turnLeft(v)
	turnLeft(v)
	turnLeft(v)
}

func onInfected(v *virus) bool {
	return grid[v.x][v.y] == Infected
}

func infect(v *virus) {
	grid[v.x][v.y] = Infected
	infections++
}

func clean(v *virus) {
	grid[v.x][v.y] = Clean
}

func expandLeft(v *virus) {
	for i  := range grid {
		grid[i] = append([]byte{Clean}, grid[i]...)
	}
	// when the grid expands left, the virus's column is numbered one higher
	v.y += 1
}

func expandRight() {
	for i  := range grid {
		grid[i] = append(grid[i], Clean)
	}
}

func newCleanRow() []byte {
	return []byte(strings.Repeat(string(Clean), len(grid[0])))
}

func expandUp(v *virus) {
	grid = append([][]byte{newCleanRow()}, grid...)
	// when the grid expands above, the virus's row is numbered one higher
	v.x += 1
}

func expandDown() {
	grid = append(grid, newCleanRow())
}

func move(v *virus) {
	switch {
	case v.y == 0 && v.heading == 'L': expandLeft(v)
	case v.y == len(grid[0])-1 && v.heading == 'R': expandRight()
	case v.x == 0 && v.heading == 'U': expandUp(v)
	case v.x == len(grid)-1 && v.heading == 'D': expandDown()
	}
	switch v.heading {
	case 'L': v.y -= 1
	case 'R': v.y += 1
	case 'U': v.x -= 1
	case 'D': v.x += 1
	}
}

func main() {
	grid = readInput()
	v := &virus{x: len(grid)/2, y: len(grid)/2, heading: 'U'}
	for i := 0; i < 10000; i++ {
		if onInfected(v) {
			turnRight(v)
			clean(v)
			move(v)
		} else {
			turnLeft(v)
			infect(v)
			move(v)
		}
		//exec.Command("clear")
		//printGrid()
		//fmt.Println("Coordinates and heading:", v.x, v.y, string(v.heading))
		//fmt.Println()
		//time.Sleep(1 * time.Millisecond)
	}
	//printGrid()
	fmt.Println("Part 1: infections caused:", infections)
}

func printGrid() {
	for _, bytes := range grid {
		fmt.Println(string(bytes))
	}
}

func readInput() [][]byte {
	var answer [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer = append(answer, []byte(scanner.Text()))
	}
	return answer
}
