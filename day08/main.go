package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

type Node struct {
	name string
	left, right string
}

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	nodes := make(map[string]Node)

	first := true
	instructions := ""
	for _, line := range inputAsLines {
		if first {
			instructions = line
			first = false
			continue
		}

		if line == "" {
			continue
		}

		node := Node{name: line[0:3], left: line[7:10], right: line[12:15]}
		nodes[node.name] = node
	}

	var steps []int64
	for _, node := range nodes {
		if node.name[2] == 'A' {
			steps = append(steps, stepsToEnd(node, instructions, nodes))
		}
	}

	fmt.Printf("steps = %v\n", steps)

	res := nww(steps[0], steps[1])
	for i := 2; i < len(steps); i++ {
		res = nww(res, steps[i])
	}

	fmt.Printf("res = %v\n", res)
}

func stepsToEnd(startingNode Node, instructions string, nodes map[string]Node) int64 {
	steps := int64(0)
	currentNode := startingNode
	done := false
	for !done {
		for _, letter := range instructions {
			steps++
			if letter == 'L' {
				currentNode = nodes[currentNode.left]
			} else {
				currentNode = nodes[currentNode.right]
			}
			// fmt.Printf("currentNode = %v\n", currentNode)
			if currentNode.name[2] == 'Z' {
				done = true
				break
			}
		}
	}
	return steps
}

func nww(a, b int64) int64 {
	return a * b / nwd(a, b)
}

func nwd(a, b int64) int64 {
	for b != 0 {
		c := a % b
		a = b
		b = c
	}
	return a
}