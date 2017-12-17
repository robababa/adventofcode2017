package main

import (
	"container/ring"
	"fmt"
)

func main() {
	part1()
	part2()
}

func part1() {
	r := ring.New(1)
	r.Value = 0
	iterations := 2017
	steps := 386
	for val := 1; val <= iterations; val++ {
		r = r.Move(steps)
		s := ring.New(1)
		s.Value = val
		r.Link(s)
		r = r.Next()
	}
	fmt.Println("Part 1: Next value is:", r.Next().Value)
}

func part2() {
	r := ring.New(1)
	r.Value = 0
	iterations := 50000000
	steps := 386
	for val := 1; val <= iterations; val++ {
		r = r.Move(steps)
		s := ring.New(1)
		s.Value = val
		r.Link(s)
		r = r.Next()
	}
	for {
		if r.Value == 0 {
			fmt.Println("Part 2: Value after 0 is:", r.Next().Value)
			break
		} else {
			r = r.Next()
		}
	}
}
