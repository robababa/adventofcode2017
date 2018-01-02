package main

import (
	"bufio"
	"os"
	"strconv"
	"log"
	"testing"
)

type coordinate struct {
	x int
	y int
}

type shortestPaths map[int]int

type pathData struct {
	passable bool
	bestPaths shortestPaths
}

var grid = make(map[coordinate]pathData)

const WALL = '#'
const OPEN = '.'

// two different digits.  Store the lower-numbered digit in lower, the higher-numbered one in higher
type digitPair struct {
	lower int
	higher int
}

type outpost struct {
	digit int
	spot coordinate
}

var outposts = make(map[outpost]bool)

// the shortest one-way trips between two digits
var shortestOneWays = make(map[digitPair]int)

var allDigits = make(map[int]bool)

func main() {
	lines := readInput()
	buildGrid(lines)
	round := 0
	for !doneWithPaths() {
		extendPaths(round)
		round++
	}
}

// only return true if all pair-wise shortest paths have been found
func doneWithPaths() bool {
	for key := range allDigits {
		for key2 := range allDigits {
			if key < key2 && shortestOneWays[digitPair{lower: key, higher: key2}] == 0 {
				return false
			}
			if key2 < key && shortestOneWays[digitPair{lower: key2, higher: key}] == 0 {
				return false
			}
		}
	}
	return true
}

func markNewOneWays(round int, digit int, spot coordinate) {
	for key, value := range grid[spot].bestPaths {
		// skip over the digit we are working with
		if key == digit {continue}
		// the key and the digit are different.
		lowestDigit, highestDigit := digit, key
		if lowestDigit > highestDigit {
			lowestDigit, highestDigit = highestDigit, lowestDigit
		}
		// if the connection has already been made, then skip to the next bestPath
		if shortestOneWays[digitPair{lower: lowestDigit, higher: highestDigit}] != 0 {
			continue
		}
		// at this point, we should make the connection!
		pathLengthFromKey := value
		// adjust for the -1 dummy value
		if pathLengthFromKey == -1 {pathLengthFromKey = 0}
		// and make the link
		shortestOneWays[digitPair{lower: lowestDigit, higher: highestDigit}] = round + pathLengthFromKey
	}
}

func tryNewOutpost(round int, possibility outpost) {
	possibleSpot := coordinate{x: possibility.spot.x, y: possibility.spot.y}
	// if this part of the grid is not passable, bail
	if !grid[possibleSpot].passable {
		return
	}
	// if the digit has already visited this part of the grid, bail
	if grid[possibleSpot].bestPaths[possibility.digit] != 0 {
		return
	}
	// so this is a good new outpost.  Add it!
	outposts[possibility] = true
	// and mark the distance for this digit
	grid[possibleSpot].bestPaths[possibility.digit] = round
	// make new connections to other digits if we can
	markNewOneWays(round, possibility.digit, possibleSpot)
}

func removeOutposts(deleteList map[outpost]bool) {
	for key := range deleteList {
		delete(outposts, key)
	}
}

func extendPaths(round int) {
	outpostsToRemove := make(map[outpost]bool)
	for key := range outposts {
		// try to go up, down, left, right to new outposts if possible, then mark this outpost for removal
		tryNewOutpost(round, outpost{digit: key.digit, spot: coordinate{x: key.spot.x - 1, y: key.spot.y}})
		tryNewOutpost(round, outpost{digit: key.digit, spot: coordinate{x: key.spot.x + 1, y: key.spot.y}})
		tryNewOutpost(round, outpost{digit: key.digit, spot: coordinate{x: key.spot.x, y: key.spot.y - 1}})
		tryNewOutpost(round, outpost{digit: key.digit, spot: coordinate{x: key.spot.x, y: key.spot.y + 1}})
		outpostsToRemove[key] = true
	}
	// now remove the old outposts
	removeOutposts(outpostsToRemove)
}

func buildRow(rowNumber int, row string) {
	elements := []byte(row)
	for j, ch := range elements {
		coord := coordinate{x: rowNumber, y: j}
		switch ch {
		case WALL: {
			grid[coord] = pathData{passable: false, bestPaths: make(shortestPaths)}
		}
		case OPEN: {
			grid[coord] = pathData{passable: true, bestPaths: make(shortestPaths)}
		}
		default: {
			// must be a digit.  Record it, enable paths to go through it, mark the distance as zero,
			// and mark it as an outpost
			digit := Atoi(string(ch))
			allDigits[digit] = true
			grid[coord] = pathData{passable: true, bestPaths: make(shortestPaths)}
			grid[coord].bestPaths[digit] = -1
			outposts[outpost{digit: digit, spot: coord}] = true
		}
		}
	}
}

func buildGrid(lines []string) {
	for i, line := range lines {
		buildRow(i, line)
	}
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