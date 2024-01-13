package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

const (
	debug = true
	steps = 65

	MAX_INT = 2147483647
)

// [y][x]
var garden [][]string
var gardenDistance [][]int

type Pair struct {
	x, y int
}

func main() {
	
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		var gardenLine []string
		for _, c := range line {
			gardenLine = append(gardenLine, string(c))
		}
		garden = append(garden, gardenLine)
	}

	var startX, startY int

	for y := 0; y < len(garden); y++ {
		var gardenDistanceLine []int
		for x := 0; x < len(garden[0]); x++ {
			switch garden[y][x] {
			case "#":
				gardenDistanceLine = append(gardenDistanceLine, -2)
			case ".":
				gardenDistanceLine = append(gardenDistanceLine, -1)
			case "S":
				gardenDistanceLine = append(gardenDistanceLine, -1)
				startX = x
				startY = y
			}
		}
		gardenDistance = append(gardenDistance, gardenDistanceLine)
	}

	gardenDistance[startX][startY] = 0
	var queue []Pair
	queue = append(queue, Pair{startX, startY})
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.x - 1 >= 0 {
			if gardenDistance[p.y][p.x - 1] == -1 {
				gardenDistance[p.y][p.x - 1] = gardenDistance[p.y][p.x] + 1
				if gardenDistance[p.y][p.x - 1] < steps {
					queue = append(queue, Pair{p.x - 1, p.y})
				}
			}
		}

		if p.x + 1 < len(gardenDistance[0]) {
			if gardenDistance[p.y][p.x + 1] == -1 {
				gardenDistance[p.y][p.x + 1] = gardenDistance[p.y][p.x] + 1
				if gardenDistance[p.y][p.x + 1] < steps {
					queue = append(queue, Pair{p.x + 1, p.y})
				}
			}
		}

		if p.y - 1 >= 0 {
			if gardenDistance[p.y - 1][p.x] == -1 {
				gardenDistance[p.y - 1][p.x] = gardenDistance[p.y][p.x] + 1
				if gardenDistance[p.y - 1][p.x] < steps {
					queue = append(queue, Pair{p.x, p.y - 1})
				}
			}
		}

		if p.y + 1 < len(gardenDistance) {
			if gardenDistance[p.y + 1][p.x] == -1 {
				gardenDistance[p.y + 1][p.x] = gardenDistance[p.y][p.x] + 1
				if gardenDistance[p.y + 1][p.x] < steps {
					queue = append(queue, Pair{p.x, p.y + 1})
				}
			}
		}
	}

	// for y := 0; y < len(garden); y++ {
	// 	for x := 0; x < len(garden[0]); x++ {
	// 		if gardenDistance[y][x] == -2 {
	// 			fmt.Printf("#")	
	// 		} else if gardenDistance[y][x] == -1 {
	// 			fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf("%v", gardenDistance[y][x])
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }

	fmt.Printf("\n")
	fmt.Printf("\n")

	res := 0
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[0]); x++ {
			if gardenDistance[y][x] == -2 {
				fmt.Printf("#")	
			} else if gardenDistance[y][x] == -1 || gardenDistance[y][x] > steps {
				fmt.Printf(".")
			} else if gardenDistance[y][x] % 2 == 1 {
				fmt.Printf("O")
				res++
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("res = %v\n", res)
}
