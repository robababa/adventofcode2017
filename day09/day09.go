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
	finalAnswerChan := make(chan int)
	go readInput(cancelChan)
	go applyCancel(cancelChan, garbageChan)
	go removeGarbage(garbageChan, groupCountChan)
	go countGroups(groupCountChan, finalAnswerChan)
	finalCount := <-finalAnswerChan
	fmt.Println("Part 1: Total group count:", finalCount)
}

func countGroups(inputChan <-chan byte, outputChan chan <-int) {
	groupsNested := 0
	groupCount := 0
	for {
		b := <-inputChan
		if b == zeroByte {
			break
		} else if b == startGroup {
			groupsNested++
		} else if b == endGroup {
			groupCount += groupsNested
		} else {
			continue
		}
	}
	outputChan <- groupCount
	fmt.Println("countGroups() finished")
}


func removeGarbage(inputChan <-chan byte, outputChan chan <-byte) {
	insideGarbage := false
	for {
		b := <-inputChan
		if b == zeroByte {
			break
		} else if b == startGarbage {
			insideGarbage = true
			continue
		} else if b == endGarbage {
			insideGarbage = false
			continue
		} else if insideGarbage {
			continue
		} else {
			outputChan <- b
		}
	}
	fmt.Println("removeGarbage() finished")
}

func applyCancel(inputChan <-chan byte, outputChan chan <-byte) {
	doingCancel := false
	for {
		b := <-inputChan
		if b == zeroByte {
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
	fmt.Println("applyCancel() finished")
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
	fmt.Println("readInput() finished")
}