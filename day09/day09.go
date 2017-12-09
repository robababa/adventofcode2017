package main

import (
	"bufio"
	"os"
	"fmt"
)

var zeroByte byte = '0'
var cancelByte byte = '!'
var startGarbage byte = '<'
var endGarbage byte = '>'
var startGroup byte = '{'
var endGroup byte = '}'

func main() {
	cancelChan := make(chan byte)
	garbageChan := make(chan byte)
	groupCountChan := make(chan byte)
	Part1AnswerChan := make(chan int)
	Part2AnswerChan := make(chan int)
	go readInput(cancelChan)
	go applyCancel(cancelChan, garbageChan)
	go removeGarbage(garbageChan, groupCountChan, Part2AnswerChan)
	go countGroups(groupCountChan, Part1AnswerChan)
	fmt.Println("Part 1: Total group count:", <-Part1AnswerChan)
	fmt.Println("Part 2: Total garbage count:", <-Part2AnswerChan)
}

func countGroups(inputChan <-chan byte, outputChan chan <-int) {
	groupsNested := 0
	groupCount := 0
	for {
		b := <-inputChan
		if b == zeroByte {
			outputChan <- groupCount
			break
		} else if b == startGroup {
			groupsNested++
		} else if b == endGroup {
			groupCount += groupsNested
			groupsNested--
		} else {
			continue
		}
	}
	close(outputChan)
	//fmt.Println("countGroups() finished")
}


func removeGarbage(inputChan <-chan byte, outputChan chan <-byte, part2AnswerChan chan <-int) {
	insideGarbage := false
	garbageCharCount := 0
	for {
		b := <-inputChan
		if b == zeroByte {
			outputChan <- b
			break
		} else if b == startGarbage && !insideGarbage {
			insideGarbage = true
			continue
		} else if b == endGarbage {
			insideGarbage = false
			continue
		} else if insideGarbage {
			garbageCharCount++
			continue
		} else {
			outputChan <- b
		}
	}
	part2AnswerChan <- garbageCharCount
	close(outputChan)
	//fmt.Println("removeGarbage() finished")
}

func applyCancel(inputChan <-chan byte, outputChan chan <-byte) {
	doingCancel := false
	for {
		b := <-inputChan
		if b == zeroByte {
			outputChan <- b
			close(outputChan)
			break
		} else if doingCancel {
			doingCancel = false
			continue
		} else if b == cancelByte {
			doingCancel = true
		} else {
			outputChan <- b
		}
	}
	//fmt.Println("applyCancel() finished")
}

func readInput(nextChan chan <-byte) {
	byteReader := bufio.NewReader(os.Stdin)
	for {
		b, err := byteReader.ReadByte()
		if err != nil {
			nextChan <- zeroByte
			close(nextChan)
			break
		} else {
			nextChan <- b
		}
	}
	//fmt.Println("readInput() finished")
}