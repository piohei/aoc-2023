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
	x, y, length int
}

func (p Point) Key() string {
	return fmt.Sprintf("x%vy%v", p.x, p.y)
}

type Road struct {
	p1, p2 Point
	length int
}

type Edge struct {
	p Point
	w int
}

var hikingMap [][]string
var visitedPoints [][]bool
var crossroads []Point
var roads []Road

var graph map[string][]Edge
var keyToPoint map[string]Point

var dist map[string]int
var prev map[string]string

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
	// if debug {
	// 	fmt.Printf("hikingMap = %v\n", hikingMap)
	// }

	// res := findLongestPath(1, 0) - 1
	// fmt.Printf("res = %v\n", res)

	for y := 0; y < len(hikingMap); y++ {
		for x := 0; x < len(hikingMap[0]); x++ {
			if hikingMap[y][x] != "#" {
				moves := getPossibleNextMoves(x, y)
				if len(moves) != 2 {
					crossroads = append(crossroads, Point{x,y,0})
				}
			}
		}
	}

	fmt.Printf("len=%v crossroads = %v\n", len(crossroads), crossroads)

	graph = make(map[string][]Edge)
	keyToPoint = make(map[string]Point)
	delta := 0
	for _, crossroad := range crossroads {
		x := crossroad.x
		y := crossroad.y
		possibleMoves := getPossibleNextMoves(x, y)

		visitedPoints[y][x] = true
		for _, move := range possibleMoves {
			nx, ny, length := findNextCrossroad(move.x, move.y)
			if length > 0 {
				road := Road{Point{x, y, 0}, Point{nx, ny, 0}, length + 1}
				roads = append(roads, road)
				if road.length > delta {
					delta = road.length
				}
			}
		}
		visitedPoints[y][x] = false
	}

	fmt.Printf("roads = %v\n", roads)

	fmt.Printf("len(roads) = %v\n", len(roads))

	for _, road := range roads {
		graph[road.p1.Key()] = append(graph[road.p1.Key()], Edge{road.p2, road.length})
		keyToPoint[road.p1.Key()] = road.p1
		graph[road.p2.Key()] = append(graph[road.p2.Key()], Edge{road.p1, road.length})
		keyToPoint[road.p2.Key()] = road.p2
	}
	fmt.Printf("graph = %v\n", graph)

	// BellmanFord()
	// test1()

	fmt.Printf("dist = %v\n", dist)
	fmt.Printf("prev = %v\n", prev)

	// endPoint := Point{len(hikingMap[0]) - 2, len(hikingMap) - 1, 0}
	// c := endPoint.Key()
	// steps := 0
	// fmt.Printf("path:\n")
	// fmt.Printf("  %v\n", c)
	// for c != "x1y0" {
	// 	c = prev[c]
	// 	fmt.Printf("  %v\n", c)
	// 	steps++

	// 	if c == "undefined" {
	// 		break
	// 	}
	// }

	// fmt.Printf("steps = %v\n", steps)

	// res := -1 * dist[endPoint.Key()]

	res := findLongestPathWithCrossroads(1, 0) - 1
	fmt.Printf("res = %v\n", res)
}

func test1() {
	dist = make(map[string]int)
	prev = make(map[string]string)

	var q []string
	for k := range keyToPoint {
		dist[k] = -1
		prev[k] = "undefined"
		q = append(q, k)
	}

	dist["x1y0"] = 0
	for len(q) > 0 {
		index := findWithMaxDist(q)
		u := q[index]
		q = remove(q, index)

		for _, e := range graph[u] {
			v := e.p.Key()
			if contains(q, v) {
				alt := dist[u] + e.w
				if alt > dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}
}

func BellmanFord() {
	dist = make(map[string]int)
	prev = make(map[string]string)

	for k := range keyToPoint {
		dist[k] = MAX_INT
		prev[k] = "undefined"
	}

	dist["x1y0"] = 0

	for i := 0; i < len(dist); i++ {
		for u, edges := range graph {
			for _, e := range edges {
				v := e.p.Key()
				if dist[u] + e.w < dist[v] && prev[u] != v {
					dist[v] = dist[u] + e.w
					prev[v] = u
				}
			}
		}
	}

	// for u, edges := range graph {
	// 	for _, e := range edges {
	// 		v := e.p.Key()
	// 	}
	// }
}

func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

func contains(slice []string, v string) bool {
	for _, s := range slice {
		if s == v {
			return true
		}
	}
	return false
}

func findWithMaxDist(q []string) int {
	maxIndex := 0
	maxValue := 0
	for i, v := range q {
		if dist[v] > maxValue {
			maxValue = dist[v]
			maxIndex = i
		}
	}

	return maxIndex
}

func findWithMinDist(q []string) int {
	minIndex := 0
	minValue := MAX_INT
	for i, v := range q {
		if dist[v] < minValue {
			minValue = dist[v]
			minIndex = i
		}
	}

	return minIndex
}


func findNextCrossroad(x, y int) (resx, resy, length int) {
	visitedPoints[y][x] = true

	possibleMoves := getPossibleNextMoves(x, y)
	if len(possibleMoves) != 2 {
		return x, y, 0
	}
	res := 0
	for _, point := range possibleMoves {
		if !visitedPoints[point.y][point.x] {
			x, y, res = findNextCrossroad(point.x, point.y)
			res++
		}
	}

	visitedPoints[y][x] = false
	return x, y, res
}

func findLongestPathWithCrossroads(x, y int) int {
	if debug {
		fmt.Printf("findLongestPathWithCrossroads(%v, %v)\n", x, y)
	}
	if hikingMap[y][x] == "E" {
		return 1
	}

	visitedPoints[y][x] = true

	longestPath := 0
	possibleMoves := getPossibleNextCrossroad(x, y)
	if debug {
		fmt.Printf("getPossibleNextCrossroad(%v, %v) = %v\n", x, y, possibleMoves)
	}
	for _, point := range possibleMoves {
		if !visitedPoints[point.y][point.x] {
			res := findLongestPathWithCrossroads(point.x, point.y)
			if res > 0 && res + point.length > longestPath {
				longestPath = res + point.length
			}
		}
	}

	visitedPoints[y][x] = false
	if debug {
		fmt.Printf("findLongestPathWithCrossroads(%v, %v) -> %v\n", x, y, longestPath)
	}
	return longestPath
}

func findLongestPath(x, y int) int {
	// if debug {
	// 	fmt.Printf("findLongestPath(%v, %v)\n", x, y)
	// }
	if hikingMap[y][x] == "E" {
		return 1
	}

	visitedPoints[y][x] = true

	longestPath := 0
	possibleMoves := getPossibleNextMoves(x, y)
	// if debug {
	// 	fmt.Printf("possibleMoves = %v\n", possibleMoves)
	// }
	for _, point := range possibleMoves {
		if !visitedPoints[point.y][point.x] {
			res := findLongestPath(point.x, point.y)
			if res > 0 && res + 1 > longestPath {
				longestPath = res + 1
			}
		}
	}

	visitedPoints[y][x] = false
	if debug {
		fmt.Printf("findLongestPath(%v, %v) -> %v\n", x, y, longestPath)
	}
	return longestPath
}

func getPossibleNextCrossroad(x, y int) []Point {
	var res []Point
	for _, road := range roads {
		if road.p1.x == x && road.p1.y == y {
			res = append(res, Point{road.p2.x, road.p2.y, road.length})
		}
		if road.p2.x == x && road.p2.y == y {
			res = append(res, Point{road.p1.x, road.p1.y, road.length})
		}
	}
	return res
}

func getPossibleNextMoves(x, y int) []Point {
	var res []Point

	// up
	var nx, ny int
	{
		nx = x
		ny = y - 1
		if isWithinMap(nx, ny) {
			switch hikingMap[ny][nx] {
			case ".", "S", "E", "^", "v", "<", ">":
				res = append(res, Point{nx, ny, 0})
			}
		}
	}
	// down
	{
		nx = x
		ny = y + 1
		if isWithinMap(nx, ny) {
			switch hikingMap[ny][nx] {
			case ".", "S", "E", "^", "v", "<", ">":
				res = append(res, Point{nx, ny, 0})
			}
		}
	}
	// left
	{
		nx = x - 1
		ny = y
		if isWithinMap(nx, ny) {
			switch hikingMap[ny][nx] {
			case ".", "S", "E", "^", "v", "<", ">":
				res = append(res, Point{nx, ny, 0})
			}
		}
	}
	// right
	{
		nx = x + 1
		ny = y
		if isWithinMap(nx, ny) {
			switch hikingMap[ny][nx] {
			case ".", "S", "E", "^", "v", "<", ">":
				res = append(res, Point{nx, ny, 0})
			}
		}
	}

	return res
}

func isWithinMap(x, y int) bool {
	return 0 <= x && x < len(hikingMap[0]) && 0 <= y && y < len(hikingMap)
}