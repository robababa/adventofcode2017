package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type NodeProperties struct {
	weight   int
	isChild  bool
	exists   bool
	children []string
}

var nodes = make(map[string]*NodeProperties)

func main() {
	inputStrings := readInput()
	//fmt.Println(inputStrings)
	loadNodes(inputStrings)
	//fmt.Println(nodes)
	rootNode := findRootNode()
	fmt.Println("Part 1: root node is", rootNode)
	fmt.Println("Part 2: unbalanced node information:")
	weight(rootNode)
}

func findRootNode() string {
	for k, v := range nodes {
		if !v.isChild {
			return k
		}
	}
	return ""
}

func balanced(weights []int) bool {
	firstWeight := weights[0]
	for _, w := range weights {
		if w != firstWeight {
			return false
		}
	}
	return true
}

func weightsOfChildren(baseNode string) []int {
	var answer []int
	for _, child := range nodes[baseNode].children {
		answer = append(answer, nodes[child].weight)
	}
	return answer
}

func weight(baseNode string) int {
	var singleChildrenWeights []int
	var totalChildrenWeight int
	for _, child := range nodes[baseNode].children {
		childWeight := weight(child)
		singleChildrenWeights = append(singleChildrenWeights, childWeight)
		totalChildrenWeight += childWeight
	}
	// if there are children, and they aren't balanced, say so
	if len(singleChildrenWeights) > 0 && !balanced(singleChildrenWeights) {
		fmt.Println("base node:", baseNode)
		fmt.Println("child nodes:", nodes[baseNode].children)
		fmt.Println("child node total weights:", singleChildrenWeights)
		fmt.Println("child node individual weights:", weightsOfChildren(baseNode))
		log.Panic("Halting program!")
	}

	return nodes[baseNode].weight + totalChildrenWeight
}

func loadNodes(inputStrings []string) {
	for _, s := range inputStrings {
		fields := strings.Split(s, " ")
		nodeName := fields[0]
		nodeWeight, _ := strconv.Atoi(fields[1])
		var nodeChildren []string
		if len(fields) > 2 {
			// we have children!
			nodeChildren = fields[2:]
		}
		if nodes[nodeName] != nil {
			// this parent node was visited before as a child  Update its weight and its children
			(nodes[nodeName]).weight = nodeWeight
			(nodes[nodeName]).children = nodeChildren
		} else {
			// we haven't seen this parent node before.  Create it.
			nodes[nodeName] = &NodeProperties{weight: nodeWeight, isChild: false, exists: true, children: nodeChildren}
		}
		// now go through the child nodes, marking existing ones as children, and creating new ones as needed
		for _, child := range nodeChildren {
			if nodes[child] != nil {
				(nodes[child]).isChild = true
			} else {
				nodes[child] = &NodeProperties{weight: 0, isChild: true, exists: true, children: []string{}}
			}
		}
	}
}

func readInput() []string {
	var answer []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		txt = strings.Replace(txt, "-> ", "", -1)
		txt = strings.Replace(txt, ",", "", -1)
		txt = strings.Replace(txt, "(", "", -1)
		txt = strings.Replace(txt, ")", "", -1)
		answer = append(answer, txt)
	}
	return answer
}
