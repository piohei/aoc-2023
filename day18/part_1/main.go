package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

const (
	debug = true

	MAX_INT = 2147483647
)

// An Item is something we manage in a priority queue.
type Edge struct {
	fromX, fromY int
	toX, toY int
	color string
}

func (e *Edge) String() string {
	return fmt.Sprintf("(%v,%v) -> (%v,%v) [%v]", e.fromX, e.fromY, e.toX, e.toY, e.color)
}

type Ground [][]string

func NewGround(maxX, maxY int) Ground {
	res := make([][]string, maxX + 1)
	for x := 0; x <= maxX; x++ {
		res[x] = make([]string, maxY + 1)
		for y := 0; y <= maxY; y++ {
			res[x][y] = "."
		}
	}
	return res
}

func (g Ground) String() string {
	res := ""
	for y := 0; y < len(g[0]); y++ {
		for x := 0; x < len(g); x++ {
			res += g[x][y]
		}
		res += "\n"
	}
	return res
}

var edges []*Edge

var ground Ground
var groundY Ground
var groundX Ground


func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	cX, cY := 0, 0
	for _, line := range inputAsLines {
		s := strings.Split(line, " ")

		direction := s[0]
		length, _ := strconv.Atoi(s[1])
		color := strings.ReplaceAll(strings.ReplaceAll(s[2], "(", ""), ")", "")

		var nX, nY int
		switch direction {
		case "U":
			nX, nY = cX, cY - length
		case "D":
			nX, nY = cX, cY + length
		case "L":
			nX, nY = cX - length, cY
		case "R":
			nX, nY = cX + length, cY
		}

		edges = append(edges, &Edge{fromX: cX, fromY: cY, toX: nX, toY: nY, color: color})
		cX, cY = nX, nY
	}

	fmt.Printf("edges = %v\n", edges)

	minX, minY := MAX_INT, MAX_INT
	for _, e := range edges {
		if e.fromX < minX {
			minX = e.fromX
		}
		if e.fromY < minY {
			minY = e.fromY
		}
	}

	fmt.Printf("minX = %v, minY = %v\n", minX, minY)

	maxX, maxY := 0, 0
	for _, e := range edges {
		if e.fromX > maxX {
			maxX = e.fromX
		}
		if e.fromY > maxY {
			maxY = e.fromY
		}
	}

	fmt.Printf("maxX = %v, maxY = %v\n", maxX, maxY)

	moveVectorX, moveVectorY := -1 * minX, -1 * minY

	for _, e := range edges {
		e.fromX += moveVectorX
		e.fromY += moveVectorY
		e.toX += moveVectorX
		e.toY += moveVectorY
	}

	minX += moveVectorX
	minY += moveVectorY

	maxX += moveVectorX
	maxY += moveVectorY


	fmt.Printf("edges = %v\n", edges)
	fmt.Printf("minX = %v, minY = %v\n", minX, minY)
	fmt.Printf("maxX = %v, maxY = %v\n", maxX, maxY)

	ground = NewGround(maxX, maxY)

	fmt.Printf("ground:\n%v\n", ground)

	for _, e := range edges {
		if e.fromX == e.toX {
			if e.fromY < e.toY {
				for y := e.fromY; y <= e.toY; y++ {
					ground[e.fromX][y] = "#"
				}
			} else {
				for y := e.toY; y <= e.fromY; y++ {
					ground[e.fromX][y] = "#"
				}
			}
		} else {
			if e.fromX < e.toX {
				for x := e.fromX; x <= e.toX; x++ {
					ground[x][e.fromY] = "#"
				}
			} else {
				for x := e.toX; x <= e.fromX; x++ {
					ground[x][e.fromY] = "#"
				}
			}
		}
	}

	fmt.Printf("ground:\n%v\n", ground)

	var edgyEdge *Edge
	for _, e := range edges {
		if e.fromY == 0 && e.toY == 0 {
			edgyEdge = e
		}
	}

	fmt.Printf("edgyEdge = %v\n", edgyEdge)

	fillGround(edgyEdge)

	fmt.Printf("ground:\n%v\n", ground)

	res := 0
	for x := 0; x < len(ground); x++ {
		for y := 0; y < len(ground[0]); y++ {
			if ground[x][y] == "#" {
				res++
			}
		}
	}
	fmt.Printf("res = %v\n", res)
}

func fillGround(edgyEdge *Edge) {
	var cX, cY int
	if edgyEdge.fromX < edgyEdge.toX {
		cX = edgyEdge.fromX + 1
		cY = 1
	} else {
		cX = edgyEdge.toX + 1
		cY = 1
	}

	var listX []int
	var listY []int

	listX = append(listX, cX)
	listY = append(listY, cY)

	for len(listX) > 0 {
		cX, cY = listX[0], listY[0]
		listX, listY = listX[1:], listY[1:]

		if ground[cX][cY] == "." {
			ground[cX][cY] = "#"

			if cX - 1 > 0 {
				listX, listY = append(listX, cX - 1), append(listY, cY)
			}
			if cX + 1 < len(ground) {
				listX, listY = append(listX, cX + 1), append(listY, cY)
			}
			if cY - 1 > 0 {
				listX, listY = append(listX, cX), append(listY, cY - 1)
			}
			if cY + 1 < len(ground[0]) {
				listX, listY = append(listX, cX), append(listY, cY + 1)
			}
		}
	}
}