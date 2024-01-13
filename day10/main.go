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

	var plane [][]string
	var origPlane [][]string

	maxX := len(inputAsLines[0])
	maxY := len(inputAsLines)

	for j := 0; j < maxX; j++ {
		var line []string
		var origLine []string
		for i := 0; i < maxY; i++ {
			line = append(line, ".")
			origLine = append(origLine, ".")
		}
		plane = append(plane, line)
		origPlane = append(origPlane, origLine)
	}

	var startingNode Node
	for y, line := range inputAsLines {
		for x, c := range line {
			plane[x][y] = string(c)
			origPlane[x][y] = string(c)
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

	nextHop := make(map[Node]Node)

	var queue []Node
	distMap[startingNode] = 0
	distMap[m[startingNode][0]] = 1
	nextHop[startingNode] = m[startingNode][0]
	queue = append(queue, m[startingNode][0])
	for {
		if len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]
		for _, next := range m[current] {
			if distMap[next] < 0 {
				distMap[next] = distMap[current] + 1
				nextHop[current] = next
				fmt.Printf("current = %v, next = %v, distMap[next] = %v\n", current, next, distMap[next])
				queue = append(queue, next)
			} else {
				fmt.Printf("current = %v, next = %v, distMap[next] = %v [visited]\n", current, next, distMap[next])
			}
		}
	}

	fmt.Printf("distMap = %v\n", distMap)

	fmt.Printf("nextHop = %v\n", nextHop)

	res := 0
	for k, v := range distMap {
		// fmt.Printf("k.x = %v, k.y = %v\n", k.x, k.y)
		if v >= 0 {
			plane[k.x][k.y] = origPlane[k.x][k.y]
		} else {
			plane[k.x][k.y] = "."
		}
		if v > res {
			res = v
		}
	}

	fmt.Printf("res = %v\n", res)

	for i := 0; i < 80; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Printf("%v ", plane[x][y])
		}
		fmt.Printf("\n")
	}

	for y := 0; y < maxY; y++ {
		if plane[0][y] == "." {
			plane[0][y] = "O"
			queue = append(queue, Node{0, y})
		}

		if plane[maxX - 1][y] == "." {
			plane[maxX - 1][y] = "O"
			queue = append(queue, Node{maxX - 1, y})
		}
	}

	for x := 0; x < maxX; x++ {
		if plane[x][0] == "." {
			plane[x][0] = "O"
			queue = append(queue, Node{x, 0})
		}

		if plane[x][maxY - 1] == "." {
			plane[x][maxY - 1] = "O"
			queue = append(queue, Node{x, maxY - 1})
		}
	}

	for {
		if len(queue) == 0 {
			break
		}
		current := queue[0]
		queue = queue[1:]

		x := current.x
		y := current.y

		// gora/dol
		if x > 0 && plane[x-1][y] == "." {
			plane[x-1][y] = "O"
			queue = append(queue, Node{x - 1, y})
		}
		if x + 1 < len(plane) && plane[x+1][y] == "." {
			plane[x+1][y] = "O"
			queue = append(queue, Node{x + 1, y})
		}
		if y > 0 && plane[x][y-1] == "." {
			plane[x][y-1] = "O"
			queue = append(queue, Node{x, y - 1})
		}
		if y + 1 < len(plane[0]) && plane[x][y+1] == "." {
			plane[x][y+1] = "O"
			queue = append(queue, Node{x, y + 1})
		}


		// skosy
		if x > 0 && y > 0 && plane[x-1][y-1] == "." {
			plane[x-1][y-1] = "O"
			queue = append(queue, Node{x - 1, y - 1})
		}
		if x + 1 < len(plane) && y > 0 && plane[x+1][y-1] == "." {
			plane[x+1][y-1] = "O"
			queue = append(queue, Node{x + 1, y - 1})
		}
		if x > 0 && y + 1 < len(plane[0]) && plane[x-1][y+1] == "." {
			plane[x-1][y+1] = "O"
			queue = append(queue, Node{x - 1, y + 1})
		}
		if x + 1 < len(plane) && y + 1 < len(plane[0]) && plane[x+1][y+1] == "." {
			plane[x+1][y+1] = "O"
			queue = append(queue, Node{x + 1, y + 1})
		}
	}


	{
		currentNode := startingNode
		for {
			nextNode, ok := nextHop[currentNode]
			if !ok {
				break
			}
			x := currentNode.x
			y := currentNode.y
			dx := nextNode.x - currentNode.x
			dy := nextNode.y - currentNode.y

			if dx == 0 {
				if dy < 0 {
					if x>0 && y>0 && plane[x-1][y-1] == "." {
						plane[x-1][y-1] = "A"
					}
					if x>0 && plane[x-1][y] == "." {
						plane[x-1][y] = "A"
					}
					if x+1<maxX && y>0 && plane[x+1][y-1] == "." {
						plane[x+1][y-1] = "B"
					}
					if x+1<maxX && plane[x+1][y] == "." {
						plane[x+1][y] = "B"
					}
				} else {
					if x+1<maxX && y+1<maxY && plane[x+1][y+1] == "." {
						plane[x+1][y+1] = "A"
					}
					if x+1<maxX && plane[x+1][y] == "." {
						plane[x+1][y] = "A"
					}
					if x>0 && y+1<maxY && plane[x-1][y+1] == "." {
						plane[x-1][y+1] = "B"
					}
					if x>0 && plane[x-1][y] == "." {
						plane[x-1][y] = "B"
					}
				}
			} else {
				if dx < 0 {
					if x>0 && y+1<maxY && plane[x-1][y+1] == "." {
						plane[x-1][y+1] = "A"
					}
					if y+1<maxY && plane[x][y+1] == "." {
						plane[x][y+1] = "A"
					}
					if x>0 && y>0 && plane[x-1][y-1] == "." {
						plane[x-1][y-1] = "B"
					}
					if y>0 && plane[x][y-1] == "." {
						plane[x][y-1] = "B"
					}
				} else {
					if x+1<maxX && y>0 && plane[x+1][y-1] == "." {
						plane[x+1][y-1] = "A"
					}
					if y>0 && plane[x][y-1] == "." {
						plane[x][y-1] = "A"
					}
					if x+1<maxX && y+1<maxY && plane[x+1][y+1] == "." {
						plane[x+1][y+1] = "B"
					}
					if y+1<maxY && plane[x][y+1] == "." {
						plane[x][y+1] = "B"
					}
				}
			}

			currentNode = nextNode
		}
	}

	for {
		found := false
		for y := 0; y < maxY; y++ {
			for x := 0; x < maxX; x++ {
				if plane[x][y] == "." {
					found = true
					// gora/dol
					if x > 0 && plane[x-1][y] == "A" {
						plane[x][y] = "A"
					}
					if x + 1 < len(plane) && plane[x+1][y] == "A" {
						plane[x][y] = "A"
					}
					if y > 0 && plane[x][y-1] == "A" {
						plane[x][y] = "A"
					}
					if y + 1 < len(plane[0]) && plane[x][y+1] == "A" {
						plane[x][y] = "A"
					}


					// skosy
					if x > 0 && y > 0 && plane[x-1][y-1] == "A" {
						plane[x][y] = "A"
					}
					if x + 1 < len(plane) && y > 0 && plane[x+1][y-1] == "A" {
						plane[x][y] = "A"
					}
					if x > 0 && y + 1 < len(plane[0]) && plane[x-1][y+1] == "A" {
						plane[x][y] = "A"
					}
					if x + 1 < len(plane) && y + 1 < len(plane[0]) && plane[x+1][y+1] == "A" {
						plane[x][y] = "A"
					}

					// gora/dol
					if x > 0 && plane[x-1][y] == "B" {
						plane[x][y] = "B"
					}
					if x + 1 < len(plane) && plane[x+1][y] == "B" {
						plane[x][y] = "B"
					}
					if y > 0 && plane[x][y-1] == "B" {
						plane[x][y] = "B"
					}
					if y + 1 < len(plane[0]) && plane[x][y+1] == "B" {
						plane[x][y] = "B"
					}


					// skosy
					if x > 0 && y > 0 && plane[x-1][y-1] == "B" {
						plane[x][y] = "B"
					}
					if x + 1 < len(plane) && y > 0 && plane[x+1][y-1] == "B" {
						plane[x][y] = "B"
					}
					if x > 0 && y + 1 < len(plane[0]) && plane[x-1][y+1] == "B" {
						plane[x][y] = "B"
					}
					if x + 1 < len(plane) && y + 1 < len(plane[0]) && plane[x+1][y+1] == "B" {
						plane[x][y] = "B"
					}

					if plane[x][y] == "." {
						found = false
					}
				}
			}
		}
		if !found {
			break
		}
	}





	for i := 0; i < 80; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Printf("%v ", origPlane[x][y])
		}
		fmt.Printf("\n")
	}

	for i := 0; i < 80; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Printf("%v ", plane[x][y])
		}
		fmt.Printf("\n")
	}

	res = 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if plane[x][y] == "." {
				res = res + 1
			}
		}
	}

	fmt.Printf("res = %v\n", res)

	resA := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if plane[x][y] == "A" {
				resA = resA + 1
			}
		}
	}

	fmt.Printf("resA = %v\n", resA)

	resB := 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if plane[x][y] == "B" {
				resB = resB + 1
			}
		}
	}

	fmt.Printf("resB = %v\n", resB)
}

func contains(list []Node, elem Node) bool {
	for _, n := range list {
		if n.x == elem.x && n.y == elem.y {
			return true
		}
	}
	return false
}

func hasNeighbour(plane [][]string, x, y int, character string) bool {
	if x > 0 && plane[x-1][y] == character {
		return true
	}
	if x + 1 < len(plane) && plane[x+1][y] == character {
		return true
	}
	if y > 0 && plane[x][y-1] == character {
		return true
	}
	if y + 1 < len(plane[0]) && plane[x][y+1] == character {
		return true
	}
	return false
}