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

	fmt.Printf("instructions = %v\n", instructions)
	fmt.Printf("nodes = %v\n", nodes)

	steps := 0
	currentNode := nodes["AAA"]
	fmt.Printf("currentNode = %v\n", currentNode)
	done := false
	for !done {
		for _, letter := range instructions {
			steps++
			if letter == 'L' {
				currentNode = nodes[currentNode.left]
			} else {
				currentNode = nodes[currentNode.right]
			}
			fmt.Printf("currentNode = %v\n", currentNode)
			if currentNode.name == "ZZZ" {
				done = true
				break
			}
		}
	}
	fmt.Printf("steps = %v\n", steps)
}