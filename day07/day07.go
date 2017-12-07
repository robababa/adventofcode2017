package main

import (
	"bufio"
	"fmt"
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

func main() {
	nodes := make(map[string]*NodeProperties)
	inputStrings := readInput()
	//fmt.Println(inputStrings)
	loadNodes(inputStrings, nodes)
	//fmt.Println(nodes)
	rootNode := findRootNode(nodes)
	fmt.Println("Part 1: root node is", rootNode)
}

func findRootNode(nodes map[string]*NodeProperties) string {
	for k, v := range nodes {
		if !v.isChild {
			return k
		}
	}
	return ""
}

func loadNodes(inputStrings []string, nodes map[string]*NodeProperties) {
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
