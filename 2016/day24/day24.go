package main

import (
	"bufio"
	"os"
	"strconv"
	"log"
	"fmt"
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
const PhonyDistance = -1

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

var outpostsToAdd = make(map[outpost]bool)

// the shortest one-way trips between two digits
var shortestOneWays = make(map[digitPair]int)

var allDigits = make(map[int]bool)

var permutations [][]int

func main() {
	lines := readInput()
	buildGrid(lines)
	round := 1
	for !doneWithPaths() {
		extendPaths(round)
		round++
	}
	//fmt.Println("Shortest one-way distances:", shortestOneWays)
	//fmt.Println(allDigitsExceptZero())
	parts1and2(1)
	parts1and2(2)
}

func parts1and2(part int) {
	var bestOrder []int
	bestOrderDistance := -1
	//fmt.Println(digitPermutations(allDigitsExceptZero()))
	for _, permutation := range digitPermutations(allDigitsExceptZero()) {
		firstDigit := 0
		distance := 0
		for _, digit := range permutation {
			if firstDigit < digit {
				distance+= shortestOneWays[digitPair{lower: firstDigit, higher: digit}]
			} else {
				distance+= shortestOneWays[digitPair{lower: digit, higher: firstDigit}]
			}
			firstDigit = digit
		}
		// in part 2, add the trip back to 0
		if part == 2 {
			distance += shortestOneWays[digitPair{lower: 0, higher: permutation[len(permutation) - 1]}]
		}
		if bestOrderDistance == -1 || distance < bestOrderDistance {
			bestOrder = append([]int{0}, permutation...)
			bestOrderDistance = distance
		}
	}
	fmt.Println("Part", part, ": Best ordering is", bestOrder, "with distance", bestOrderDistance)
}

func digitPermutations(digits []int) [][]int {
	//fmt.Println("digitPermutations() called with", digits)
	if len(digits) == 1 {
		return [][]int{{digits[0]}}
	}

	// an inner function that we use to build up the permutations
	var combine func(n int, partials [][]int) [][]int
	combine = func(n int, partials [][]int) [][]int {
		var combineAnswer [][]int
		for _, partial := range partials {
			combineAnswer = append(combineAnswer, append([]int{n}, partial...))
		}
		return combineAnswer
	}

	var answer [][]int
	for i, d := range digits {
		//fmt.Println("digitPermutations() range index and digit", i, "and", d, "from digits", digits)
		answer = append(answer, combine(d, digitPermutations(append(append([]int{}, digits[:i]...), digits[i+1:]...)))...)
		//fmt.Println("digitPermutations() answer is", answer)
	}
	return answer
}

func allDigitsExceptZero() []int {
	var answer []int
	for key := range allDigits {
		if key != 0 {
			answer = append(answer, key)
		}
	}
	return answer
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
		// adjust for the phony distance value if necessary
		if pathLengthFromKey == PhonyDistance {pathLengthFromKey = 0}
		// and make the link
		//fmt.Println("Linking", lowestDigit, "to", highestDigit, "at coordinate", spot,
		//	"on digit", digit, "'s turn in round", round, "with distance", round + pathLengthFromKey)
		shortestOneWays[digitPair{lower: lowestDigit, higher: highestDigit}] = round + pathLengthFromKey
	}
}

func tryNewOutpost(round int, possibility outpost) {
	//fmt.Println("tryNewOutpost() round", round, "digit", possibility.digit, "at", possibility.spot)
	possibleSpot := possibility.spot
	// if this part of the grid is not passable, bail
	if !grid[possibleSpot].passable {
		return
	}
	// if the digit has already visited this part of the grid, bail
	if grid[possibleSpot].bestPaths[possibility.digit] != 0 {
		return
	}
	//fmt.Println("...creating outpost")
	// so this is a good new outpost.  Add it! (to the lists of outposts to add)
	outpostsToAdd[possibility] = true
	// and mark the distance for this digit
	grid[possibleSpot].bestPaths[possibility.digit] = round
	// make new connections to other digits if we can
	markNewOneWays(round, possibility.digit, possibleSpot)
}

func removeOldOutposts(deleteList map[outpost]bool) {
	for key := range deleteList {
		//fmt.Println("removeOutposts() deleting outpost", key)
		delete(outposts, key)
	}
}

func addNewOutposts() {
	for key := range outpostsToAdd {
		//fmt.Println("addNewOutposts() adding outpost", key)
		outposts[key] = true
	}
	for key := range outpostsToAdd {
		delete(outpostsToAdd, key)
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
	removeOldOutposts(outpostsToRemove)
	// and add the new ones
	addNewOutposts()
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
			grid[coord].bestPaths[digit] = PhonyDistance
			//fmt.Println("buildRow() creating outpost", outpost{digit: digit, spot: coord})
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