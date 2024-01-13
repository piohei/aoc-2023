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
)

var lever [][]string
var tilledLever [][]string

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		var leverLine []string
		for _, c := range line {
			leverLine = append(leverLine, string(c))
		}
		lever = append(lever, leverLine)
	}

	printLever()
	tiltNorth()
	printLever()

	sum := 0
	maxWeight := len(lever)
	for x := 0; x < len(lever[0]); x++ {

		for y := 0; y < len(lever); y++ {
			if lever[y][x] == "O" {
				sum = sum + (maxWeight - y)
			}
		}
	}
	fmt.Printf("sum = %v\n", sum)
}

func printLever() {
	fmt.Printf("----------------------------------------\n")
	for _, line := range lever {
		fmt.Printf("%v\n", strings.Join(line, ""))
	}
	fmt.Printf("----------------------------------------\n")
}

func tiltNorth() {
	for x := 0; x < len(lever[0]); x++ {
		ys := 0
		ye := 0

		for ye < len(lever) {
			if lever[ye][x] == "O" {
				if debug {
					fmt.Printf("found stone x=%v, ys=%v, ye=%v\n", x, ys, ye)
				}
				for ys < ye {
					if lever[ys][x] == "." {
						break
					}
					ys++
				}
				if lever[ys][x] == "." {
					if debug {
						fmt.Printf("swaping x=%v, ys=%v, ye=%v\n", x, ys, ye)
					}
					lever[ys][x], lever[ye][x] = lever[ye][x], lever[ys][x]
				}
			}
			if lever[ye][x] == "#" {
				ys = ye
			}
			ye++
		}
	}
}