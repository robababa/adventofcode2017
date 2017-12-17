package main

import (
	"container/ring"
	"fmt"
)

func main() {
	part1()
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
		//fmt.Println("s.Value is", s.Value)
		r.Link(s)
		r = r.Next()
		//fmt.Println("r.Value is", r.Value)
	}
	fmt.Println(r.Next().Value)
}
