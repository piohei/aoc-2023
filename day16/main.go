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

	north = 0
	east = 1
	south = 2
	west = 3
)

type Point struct {
	character string
	beamDirection []bool
}

type Move struct {
	x, y int
	direction int
}

var origContraption [][]*Point
var contraption [][]*Point

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for y, line := range inputAsLines {
		if origContraption == nil {
			origContraption = make([][]*Point, len(line))
		}

		for x, char := range strings.Split(line, "") {
			if origContraption[x] == nil {
				origContraption[x] = make([]*Point, len(inputAsLines))
			}
			origContraption[x][y] = &Point{
				character: char,
				beamDirection: []bool{false, false, false, false},
			}
		}
	}

	maxSum := 0
	x := 0
	y := 0
	for x = 0; x < len(origContraption); x++ {
		y = 0

		copyContraption()
		moveBeam(x, y, south)

		sum := 0
		for y := range contraption[0] {
			for x := range contraption {
				p := contraption[x][y]
				if p.beamDirection[north] || p.beamDirection[east] || p.beamDirection[south] || p.beamDirection[west] {
					sum++
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	for x = 0; x < len(origContraption); x++ {
		y = len(origContraption[0]) - 1

		copyContraption()
		moveBeam(x, y, north)

		sum := 0
		for y := range contraption[0] {
			for x := range contraption {
				p := contraption[x][y]
				if p.beamDirection[north] || p.beamDirection[east] || p.beamDirection[south] || p.beamDirection[west] {
					sum++
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	for y = 0; y < len(origContraption[0]); y++ {
		x = 0
		
		copyContraption()
		moveBeam(x, y, east)

		sum := 0
		for y := range contraption[0] {
			for x := range contraption {
				p := contraption[x][y]
				if p.beamDirection[north] || p.beamDirection[east] || p.beamDirection[south] || p.beamDirection[west] {
					sum++
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	for y = 0; y < len(origContraption[0]); y++ {
		x = len(origContraption) - 1

		copyContraption()
		moveBeam(x, y, west)

		sum := 0
		for y := range contraption[0] {
			for x := range contraption {
				p := contraption[x][y]
				if p.beamDirection[north] || p.beamDirection[east] || p.beamDirection[south] || p.beamDirection[west] {
					sum++
				}
			}
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	fmt.Printf("maxSum = %v\n", maxSum)
}

func copyContraption() {
	contraption = nil
	for _, contrLine := range origContraption {
		var line []*Point
		for _, c := range contrLine {
			line = append(line, &Point{character: c.character, beamDirection: []bool{false, false, false, false}})
		}
		contraption = append(contraption, line)
	}
}

func moveBeam(startX, startY, direction int) {
	var queue []Move
	queue = append(queue, Move{x: startX, y: startY, direction: direction})

	for len(queue) > 0 {
		move := queue[0]
		queue = queue[1:]

		if debug {
			fmt.Printf("move: %v\n", move)
		}

		c := contraption[move.x][move.y]

		var moves []Move
		switch c.character {
		case ".":
			switch move.direction {
			case north:
				moves = append(moves, moveNorth(move.x, move.y))
			case east:
				moves = append(moves, moveEast(move.x, move.y))
			case south:
				moves = append(moves, moveSouth(move.x, move.y))
			case west:
				moves = append(moves, moveWest(move.x, move.y))
			}
		case "\\":
			switch move.direction {
			case north:
				moves = append(moves, moveWest(move.x, move.y))
			case east:
				moves = append(moves, moveSouth(move.x, move.y))
			case south:
				moves = append(moves, moveEast(move.x, move.y))
			case west:
				moves = append(moves, moveNorth(move.x, move.y))
			}
		case "/":
			switch move.direction {
			case north:
				moves = append(moves, moveEast(move.x, move.y))
			case east:
				moves = append(moves, moveNorth(move.x, move.y))
			case south:
				moves = append(moves, moveWest(move.x, move.y))
			case west:
				moves = append(moves, moveSouth(move.x, move.y))
			}
		case "|":
			switch move.direction {
			case north:
				moves = append(moves, moveNorth(move.x, move.y))
			case east:
				moves = append(moves, moveNorth(move.x, move.y))
				moves = append(moves, moveSouth(move.x, move.y))
			case south:
				moves = append(moves, moveSouth(move.x, move.y))
			case west:
				moves = append(moves, moveNorth(move.x, move.y))
				moves = append(moves, moveSouth(move.x, move.y))
			}
		case "-":
			switch move.direction {
			case north:
				moves = append(moves, moveEast(move.x, move.y))
				moves = append(moves, moveWest(move.x, move.y))
			case east:
				moves = append(moves, moveEast(move.x, move.y))
			case south:
				moves = append(moves, moveEast(move.x, move.y))
				moves = append(moves, moveWest(move.x, move.y))
			case west:
				moves = append(moves, moveWest(move.x, move.y))
			}
		}

		for _, m := range moves {

			if !c.beamDirection[m.direction] {
				if debug {
					fmt.Printf("next move: %v\n", m)
				}
				c.beamDirection[m.direction] = true

				if m.x < 0 || m.x >= len(contraption) || m.y < 0 || m.y >= len(contraption[0]) {
					if debug {
						fmt.Printf("skipping move: %v, maxX=%v, maxY=%v, %v %v %v %v\n", m, len(contraption), len(contraption[0]), m.x < 0, m.x >= len(contraption), m.y < 0, m.y >= len(contraption[0]))
					}
				} else {
					queue = append(queue, m)	
				}
			}
		}

	}
}

func moveNorth(x, y int) Move {
	return Move{x: x, y: y - 1, direction: north}
}
func moveEast(x, y int) Move {
	return Move{x: x + 1, y: y, direction: east}
}
func moveSouth(x, y int) Move {
	return Move{x: x, y: y + 1, direction: south}
}
func moveWest(x, y int) Move {
	return Move{x: x - 1, y: y, direction: west}
}