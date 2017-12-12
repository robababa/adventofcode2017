package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

const InputSize = 2000

type vertex struct {
	updated bool
	nextTo []int
	lowestReachable int
}

var vertices [InputSize]vertex

func main() {
	plainInput := readInput()
	//fmt.Println(plainInput)
	loadInput(plainInput)
	//fmt.Println(vertices)
	mergeAdjacent()
	//fmt.Println(vertices)
	inGroupZero := 0
	var groups = make(map[int]bool)
	for _, v := range vertices {
		if v.lowestReachable == 0 {
			inGroupZero++
		}
		groups[v.lowestReachable] = true
	}
	fmt.Println("Part 1: Number in group zero is", inGroupZero)
	fmt.Println("Part 2: Number of groups is", len(groups))
}

func updateVertex(v vertex) (bool, int) {
	updated := false
	lowest := v.lowestReachable
	for _, n := range v.nextTo {
		nLowest := vertices[n].lowestReachable
		if nLowest < lowest {
			updated = true
			lowest = nLowest
		}
	}
	return updated, lowest
}

func mergeAdjacent() {
	updatedSome := true
	for updatedSome {
		updatedSome = false
		for i, v := range vertices {
			updatedThis, newLowest := updateVertex(v)
			vertices[i].updated = updatedThis
			if updatedThis {
				updatedSome = true
				vertices[i].lowestReachable = newLowest
			}
		}
		//fmt.Println(vertices)
	}
}

func createVertex(line string) (int, vertex) {
	vertexNumber := -1
	v := vertex{updated: true}
	tokens := strings.Split(strings.Replace(line, ",", "", -1), " ")
	vertexNumber, _ = strconv.Atoi(tokens[0])
	v.lowestReachable = vertexNumber
	for _, n := range tokens[2:] {
		neighbor, _ := strconv.Atoi(n)
		v.nextTo = append(v.nextTo, neighbor)
		if neighbor < v.lowestReachable {
			v.lowestReachable = neighbor
		}
	}
	return vertexNumber, v
}

func loadInput(input []string) {
	for _, line := range input {
		vertexNumber, v := createVertex(line)
		vertices[vertexNumber] = v
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
