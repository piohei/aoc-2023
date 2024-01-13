package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

const (
	debug = false
	steps = 6 * 131 + 65
	// steps = 26501365

	repeats = 20

	MAX_INT = 2147483647
)

// [y][x]
var garden [][]string
var origGarden [][]string
var gardenDistance map[int][][]int

type Pair struct {
	x, y int
}

func toKey(p Pair) int {
	return p.x + 100000 * p.y
}

func main() {
	// fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		var gardenLine []string
		for _, c := range line {
			gardenLine = append(gardenLine, string(c))
		}
		origGarden = append(origGarden, gardenLine)
	}

	lenY := len(origGarden)
	lenX := len(origGarden[0])

	for y := 0; y < repeats * lenY; y++ {
		var gardenLine []string
		for x := 0; x < repeats * lenX; x++ {
			gardenLine = append(gardenLine, origGarden[y % lenY][x % lenX])
		}
		garden = append(garden, gardenLine)
	}

	var startX, startY int

	for y := (repeats - 1) / 2 * lenY; y < (repeats + 1) / 2 * lenY; y++ {
		for x := (repeats - 1) / 2 * lenX; x < (repeats + 1) / 2 * lenX; x++ {
			if garden[y][x] == "S" {
				startX = x
				startY = y
				break
			}
		}
	}

	gardenDistance = make(map[int][][]int)
	calculateGardenDistance(startX, startY)

	// for y := 0; y < len(garden); y++ {
	// 	calculateGardenDistance(0, y)
	// 	calculateGardenDistance(len(garden[0]) - 1, y)
	// }
	// for x := 1; x < len(garden[0]) - 1; x++ {
	// 	calculateGardenDistance(x, 0)
	// 	calculateGardenDistance(x, len(garden) - 1)
	// }

	// if debug {
	// 	for y := 0; y < len(garden); y++ {
	// 		for x := 0; x < len(garden[0]); x++ {
	// 			if gardenDistance[toKey(Pair{startX, startY})][y][x] == -2 {
	// 				fmt.Printf("  #")	
	// 			} else if gardenDistance[toKey(Pair{startX, startY})][y][x] == -1 {
	// 				fmt.Printf("  .")
	// 			} else {
	// 				if y % lenY == lenY/2 && x % lenX == lenX/2 {
	// 					fmt.Printf("%3d", gardenDistance[toKey(Pair{startX, startY})][y][x])
	// 				} else {
	// 					fmt.Printf("  .")
	// 				}
	// 			}
	// 		}
	// 		fmt.Printf("\n")
	// 	}

	// 	fmt.Printf("\n")
	// 	fmt.Printf("\n")
	// }

	if debug {
		fmt.Printf("\n")
		fmt.Printf("\n")

		for y := 0; y < len(garden); y++ {
			if y % lenY == lenY/2 {
				for x := 0; x < len(garden[0]); x++ {
					if x % lenX == lenX/2 {
						fmt.Printf("%5d  ", gardenDistance[toKey(Pair{startX, startY})][y][x])
					}
				}
				fmt.Printf("\n")
				fmt.Printf("\n")
			}
		}

		fmt.Printf("\n")
		fmt.Printf("\n")
	}

	res := int64(0)
	for y := 0; y < len(garden); y++ {
		for x := 0; x < len(garden[0]); x++ {
			if gardenDistance[toKey(Pair{startX, startY})][y][x] == -2 {
				if debug {
					fmt.Printf("#")	
				}
			} else if gardenDistance[toKey(Pair{startX, startY})][y][x] == -1 || gardenDistance[toKey(Pair{startX, startY})][y][x] > steps {
				if debug {
					fmt.Printf(".")
				}
			} else if gardenDistance[toKey(Pair{startX, startY})][y][x] % 2 == steps % 2 {
				if debug {
					fmt.Printf("O")
				}
				res++
			} else {
				if debug {
					fmt.Printf(".")
				}
			}
		}
		if debug {
			fmt.Printf("\n")
		}
	}

	if debug {
		fmt.Printf("\n")
		fmt.Printf("\n")
	}
	fmt.Printf("res = %v\n", res)
}

func calculateGardenDistance(startX, startY int) {
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
			}
		}
		gardenDistance[toKey(Pair{startX, startY})] = append(gardenDistance[toKey(Pair{startX, startY})], gardenDistanceLine)
	}

	gardenDistance[toKey(Pair{startX, startY})][startX][startY] = 0
	var queue []Pair
	queue = append(queue, Pair{startX, startY})
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.x - 1 >= 0 {
			if gardenDistance[toKey(Pair{startX, startY})][p.y][p.x - 1] == -1 {
				gardenDistance[toKey(Pair{startX, startY})][p.y][p.x - 1] = gardenDistance[toKey(Pair{startX, startY})][p.y][p.x] + 1
				queue = append(queue, Pair{p.x - 1, p.y})
			}
		}

		if p.x + 1 < len(gardenDistance[toKey(Pair{startX, startY})][0]) {
			if gardenDistance[toKey(Pair{startX, startY})][p.y][p.x + 1] == -1 {
				gardenDistance[toKey(Pair{startX, startY})][p.y][p.x + 1] = gardenDistance[toKey(Pair{startX, startY})][p.y][p.x] + 1
				queue = append(queue, Pair{p.x + 1, p.y})
			}
		}

		if p.y - 1 >= 0 {
			if gardenDistance[toKey(Pair{startX, startY})][p.y - 1][p.x] == -1 {
				gardenDistance[toKey(Pair{startX, startY})][p.y - 1][p.x] = gardenDistance[toKey(Pair{startX, startY})][p.y][p.x] + 1
				queue = append(queue, Pair{p.x, p.y - 1})
			}
		}

		if p.y + 1 < len(gardenDistance[toKey(Pair{startX, startY})]) {
			if gardenDistance[toKey(Pair{startX, startY})][p.y + 1][p.x] == -1 {
				gardenDistance[toKey(Pair{startX, startY})][p.y + 1][p.x] = gardenDistance[toKey(Pair{startX, startY})][p.y][p.x] + 1
				queue = append(queue, Pair{p.x, p.y + 1})
			}
		}
	}
}
