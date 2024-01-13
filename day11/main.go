package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

var galaxy [][]string
var doubledX []int
var doubledY []int

var expansion int = 1_000_000

type Point struct {
	x, y int
}

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		var galaxyLine []string
		for _, c := range line {
			galaxyLine = append(galaxyLine, string(c))
		}
		galaxy = append(galaxy, galaxyLine)
	}

	printGalaxy()

	for y := 0; y < len(galaxy); y++ {
		allSpace := true
		for x := 0; x < len(galaxy[0]); x++ {
			if galaxy[y][x] != "." {
				allSpace = false
			}
		}
		if allSpace {
			doubledY = append(doubledY, y)
		}
	}

	for x := 0; x < len(galaxy[0]); x++ {
		allSpace := true
		for y := 0; y < len(galaxy); y++ {
			if galaxy[y][x] != "." {
				allSpace = false
			}
		}
		if allSpace {
			doubledX = append(doubledX, x)
		}
	}

	var points []Point
	for y := 0; y < len(galaxy); y++ {
		for x := 0; x < len(galaxy[0]); x++ {
			if galaxy[y][x] != "." {
				points = append(points, Point{x, y})
			}
		}
	}

	fmt.Printf("points = %v\n", points)

	res := int64(0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			res = res + dist(points[i], points[j])
		}
	}

	fmt.Printf("res = %v\n", res)
}

func printLine() {
	for x := 0; x < 80; x++ {
		fmt.Printf("-",)
	}
	fmt.Printf("\n")
}

func printGalaxy() {
	printLine()
	for y := 0; y < len(galaxy); y++ {
		for x := 0; x < len(galaxy[0]); x++ {
			fmt.Printf("%v ", galaxy[y][x])
		}
		fmt.Printf("\n")
	}
	printLine()
}

func dist(a, b Point) int64 {
	res := int64(0)
	if a.x < b.x {
		doubled := numOfDoubledBetweenX(a.x, b.x)
		res = res + int64(b.x - a.x) + int64(doubled) * int64(expansion - 1)
	} else {
		doubled := numOfDoubledBetweenX(b.x, a.x)
		res = res + int64(a.x - b.x) + int64(doubled) * int64(expansion - 1)
	}
	if a.y < b.y {
		doubled := numOfDoubledBetweenY(a.y, b.y)
		res = res + int64(b.y - a.y) + int64(doubled) * int64(expansion - 1)
	} else {
		doubled := numOfDoubledBetweenY(b.y, a.y)
		res = res + int64(a.y - b.y) + int64(doubled) * int64(expansion - 1)
	}
	return res
}

func numOfDoubledBetweenX(x1, x2 int) int {
	id1 := findFirstBigger(doubledX, x1)
	id2 := findFirstLower(doubledX, x2)

	if id1 < 0 || id2 < 0 {
		return 0
	}

	if id1 <= id2 {
		return id2 - id1 + 1
	}

	return 0
}

func numOfDoubledBetweenY(y1, y2 int) int {
	id1 := findFirstBigger(doubledY, y1)
	id2 := findFirstLower(doubledY, y2)

	if id1 < 0 || id2 < 0 {
		return 0
	}

	if id1 <= id2 {
		return id2 - id1 + 1
	}

	return 0
}

func findFirstBigger(s []int, v int) int {
	for i, cv := range s {
		if cv > v {
			return i
		}
	}

	return -1
}

func findFirstLower(s []int, v int) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] < v {
			return i
		}
	}

	return -1
}