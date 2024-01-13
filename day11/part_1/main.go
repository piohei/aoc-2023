package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

var galaxy [][]string
var expandedGalaxy [][]string

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
		expandedGalaxy = append(expandedGalaxy, galaxy[y][:])
		if allSpace {
			expandedGalaxy = append(expandedGalaxy, galaxy[y][:])
		}
	}

	{
		x := 0
		for {
			if x >= len(expandedGalaxy[0]) {
				break
			}
			allSpace := true
			for y := 0; y < len(expandedGalaxy); y++ {
				if expandedGalaxy[y][x] != "." {
					allSpace = false
				}
			}
			for y := 0; y < len(expandedGalaxy); y++ {
				if allSpace {
					a1 := expandedGalaxy[y][:x]
					a2 := expandedGalaxy[y][x]
					a := append(a1, a2)
					a = append(a, expandedGalaxy[y][x:]...)
					expandedGalaxy[y] = a
				}
			}
			if allSpace {
				x++
			}
			x++
		}
	}

	printExpandedGalaxy()

	var points []Point
	for y := 0; y < len(expandedGalaxy); y++ {
		for x := 0; x < len(expandedGalaxy[0]); x++ {
			if expandedGalaxy[y][x] != "." {
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

func printExpandedGalaxy() {
	printLine()
	for y := 0; y < len(expandedGalaxy); y++ {
		for x := 0; x < len(expandedGalaxy[0]); x++ {
			fmt.Printf("%v ", expandedGalaxy[y][x])
		}
		fmt.Printf("\n")
	}
	printLine()
}

func dist(a, b Point) int64 {
	res := int64(0)
	if a.x < b.x {
		res = res + int64(b.x - a.x)
	} else {
		res = res + int64(a.x - b.x)
	}
	if a.y < b.y {
		res = res + int64(b.y - a.y)
	} else {
		res = res + int64(a.y - b.y)
	}
	return res
}