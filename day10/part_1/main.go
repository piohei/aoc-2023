package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

type Node struct {
	x, y int
}

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	m := make(map[Node][]Node)
	distMap := make(map[Node]int)

	var startingNode Node
	for y, line := range inputAsLines {
		for x, c := range line {
			cn := Node{x,y}
			// fmt.Printf("x = %v, y = %v, c = %v\n", x, y, string(c))
			switch c {
			case '|':
				m[cn] = []Node{{x, y-1},{x, y+1}}
			case '-':
				m[cn] = []Node{{x-1, y},{x+1, y}}
			case 'L':
				m[cn] = []Node{{x, y-1},{x+1, y}}
			case 'J':
				m[cn] = []Node{{x, y-1},{x-1, y}}
			case '7':
				m[cn] = []Node{{x, y+1},{x-1, y}}
			case 'F':
				m[cn] = []Node{{x, y+1},{x+1, y}}
			case 'S':
				startingNode = cn
			}
			// fmt.Printf("nodes = %v\n", m[cn])
		}
	}

	{
		var nodes []Node
		potentialNodes := []Node{
			{startingNode.x, startingNode.y - 1},
			{startingNode.x, startingNode.y + 1},
			{startingNode.x - 1, startingNode.y},
			{startingNode.x + 1, startingNode.y},
		}
		for _, n := range potentialNodes {
			if val, ok := m[n]; ok {
				fmt.Printf("startingNode = %v, n = %v, val = %v\n", startingNode, n, val)
				if contains(val, startingNode) {
					nodes = append(nodes, n)
				}
			}
		}
		m[startingNode] = nodes
	}

	fmt.Printf("m = %v\n", m)
	fmt.Printf("startingNode = %v\n", startingNode)

	for n, _ := range m {
		distMap[n] = -1
	}
	fmt.Printf("distMap = %v\n", distMap)

	var queue []Node
	distMap[startingNode] = 0
	queue = append(queue, startingNode)
	for true {
		if len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]
		for _, next := range m[current] {
			if distMap[next] < 0 {
				distMap[next] = distMap[current] + 1
				fmt.Printf("current = %v, next = %v, distMap[next] = %v\n", current, next, distMap[next])
				queue = append(queue, next)
			} else {
				fmt.Printf("current = %v, next = %v, distMap[next] = %v [visited]\n", current, next, distMap[next])
			}
		}
	}

	fmt.Printf("distMap = %v\n", distMap)

	res := 0
	for _, v := range distMap {
		if v > res {
			res = v
		}
	}

	fmt.Printf("res = %v\n", res)
}

func contains(list []Node, elem Node) bool {
	for _, n := range list {
		if n.x == elem.x && n.y == elem.y {
			return true
		}
	}
	return false
}