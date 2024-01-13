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

	MAX_INT = 2147483647
)

type Point struct {
	x, y int
	steps int
}

var hikingMap [][]string
var visitedPoints [][]bool

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		var hikingMapLine []string
		var visitedPointsLine []bool
		for _, c := range line {
			hikingMapLine = append(hikingMapLine, string(c))
			visitedPointsLine = append(visitedPointsLine, false)
		}
		hikingMap = append(hikingMap, hikingMapLine)
		visitedPoints = append(visitedPoints, visitedPointsLine)
	}
	hikingMap[0][1] = "S"
	hikingMap[len(hikingMap) - 1][len(hikingMap[0]) - 2] = "E"
	fmt.Printf("hikingMap = %v\n", hikingMap)

	res := findLongestPath(1, 0) - 1
	fmt.Printf("res = %v\n", res)
}

func findLongestPath(x, y int) int {
	fmt.Printf("findLongestPath(%v, %v)\n", x, y)
	if hikingMap[y][x] == "E" {
		return 1
	}

	visitedPoints[y][x] = true

	longestPath := 0
	possibleMoves := getPossibleNextMoves(x, y)
	fmt.Printf("possibleMoves = %v\n", possibleMoves)
	for _, point := range possibleMoves {
		if !visitedPoints[point.y][point.x] {
			res := findLongestPath(point.x, point.y)
			if res > 0 && res + 1 > longestPath {
				longestPath = res + point.steps
			}
		}
	}

	visitedPoints[y][x] = false
	return longestPath
}

func getPossibleNextMoves(x, y int) []Point {
	var res []Point

	// up
	var nx, ny, steps int
	{
		nx = x
		ny = y - 1
		steps = 1
		for {
			if !isWithinMap(nx, ny) {
				break
			}

			if hikingMap[ny][nx] == "." || hikingMap[ny][nx] == "S" || hikingMap[ny][nx] == "E" {
				res = append(res, Point{nx, ny, steps})
				break
			}

			if hikingMap[ny][nx] == "#" {
				break
			}

			switch hikingMap[ny][nx] {
			case "^":
				ny = ny - 1
			case "v":
				ny = ny + 1
			case "<":
				nx = nx - 1
			case ">":
				nx = nx + 1
			}

			steps++
		}
	}
	// down
	{
		nx = x
		ny = y + 1
		steps = 1
		for {
			if !isWithinMap(nx, ny) {
				break
			}

			if hikingMap[ny][nx] == "." || hikingMap[ny][nx] == "S" || hikingMap[ny][nx] == "E" {
				res = append(res, Point{nx, ny, steps})
				break
			}

			if hikingMap[ny][nx] == "#" {
				break
			}

			switch hikingMap[ny][nx] {
			case "^":
				ny = ny - 1
			case "v":
				ny = ny + 1
			case "<":
				nx = nx - 1
			case ">":
				nx = nx + 1
			}

			steps++
		}
	}
	// left
	{
		nx = x - 1
		ny = y
		steps = 1
		for {
			if !isWithinMap(nx, ny) {
				break
			}

			if hikingMap[ny][nx] == "." || hikingMap[ny][nx] == "S" || hikingMap[ny][nx] == "E" {
				res = append(res, Point{nx, ny, steps})
				break
			}

			if hikingMap[ny][nx] == "#" {
				break
			}

			switch hikingMap[ny][nx] {
			case "^":
				ny = ny - 1
			case "v":
				ny = ny + 1
			case "<":
				nx = nx - 1
			case ">":
				nx = nx + 1
			}

			steps++
		}
	}
	// right
	{
		nx = x + 1
		ny = y
		steps = 1
		for {
			if !isWithinMap(nx, ny) {
				break
			}

			if hikingMap[ny][nx] == "." || hikingMap[ny][nx] == "S" || hikingMap[ny][nx] == "E" {
				res = append(res, Point{nx, ny, steps})
				break
			}

			if hikingMap[ny][nx] == "#" {
				break
			}

			switch hikingMap[ny][nx] {
			case "^":
				ny = ny - 1
			case "v":
				ny = ny + 1
			case "<":
				nx = nx - 1
			case ">":
				nx = nx + 1
			}

			steps++
		}
	}

	return res
}

func isWithinMap(x, y int) bool {
	return 0 <= x && x < len(hikingMap[0]) && 0 <= y && y < len(hikingMap)
}