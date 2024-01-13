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
var leverToRound map[string]int

func main() {
	// fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		var leverLine []string
		for _, c := range line {
			leverLine = append(leverLine, string(c))
		}
		lever = append(lever, leverLine)
	}

	leverToRound = make(map[string]int)
	cyclesNumber := 1000000000

	printLever()
	jumpDone := false
	for i := 1; i <= cyclesNumber; i++ {
		tiltRound()
		// fmt.Printf("#%v\n", i)
		// printLever()

		if v, ok := leverToRound[leverAsString()]; ok && !jumpDone {
			fmt.Printf("orig i=%v, v=%v\n", i, v)
			diff := i - v
			fmt.Printf("diff=%v\n", diff)
			mul := (cyclesNumber - i) / diff
			i += diff * mul
			fmt.Printf("jump to i=%v\n", i)
			jumpDone = true
			continue
		}
		leverToRound[leverAsString()] = i
	}

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

func leverAsString() string {
	res := ""
	for _, line := range lever {
		res += strings.Join(line, "")
	}
	return res
}

func tiltRound() {
	tiltNorth()
	tiltWest()
	tiltSouth()
	tiltEast()
}

func tiltNorth() {
	for x := 0; x < len(lever[0]); x++ {
		ys := 0
		ye := ys

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

func tiltSouth() {
	for x := 0; x < len(lever[0]); x++ {
		ys := len(lever) - 1
		ye := ys

		for ye >= 0 {
			if lever[ye][x] == "O" {
				if debug {
					fmt.Printf("found stone x=%v, ys=%v, ye=%v\n", x, ys, ye)
				}
				for ys > ye {
					if lever[ys][x] == "." {
						break
					}
					ys--
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
			ye--
		}
	}
}

func tiltWest() {
	for y := 0; y < len(lever[0]); y++ {
		xs := 0
		xe := xs

		for xe < len(lever[0]) {
			if lever[y][xe] == "O" {
				if debug {
					fmt.Printf("found stone y=%v, xs=%v, xe=%v\n", y, xs, xe)
				}
				for xs < xe {
					if lever[y][xs] == "." {
						break
					}
					xs++
				}
				if lever[y][xs] == "." {
					if debug {
						fmt.Printf("swaping y=%v, xs=%v, xe=%v\n", y, xs, xe)
					}
					lever[y][xs], lever[y][xe] = lever[y][xe], lever[y][xs]
				}
			}
			if lever[y][xe] == "#" {
				xs = xe
			}
			xe++
		}
	}
}

func tiltEast() {
	for y := 0; y < len(lever[0]); y++ {
		xs := len(lever[0]) - 1
		xe := xs

		for xe >= 0 {
			if lever[y][xe] == "O" {
				if debug {
					fmt.Printf("found stone y=%v, xs=%v, xe=%v\n", y, xs, xe)
				}
				for xs > xe {
					if lever[y][xs] == "." {
						break
					}
					xs--
				}
				if lever[y][xs] == "." {
					if debug {
						fmt.Printf("swaping y=%v, xs=%v, xe=%v\n", y, xs, xe)
					}
					lever[y][xs], lever[y][xe] = lever[y][xe], lever[y][xs]
				}
			}
			if lever[y][xe] == "#" {
				xs = xe
			}
			xe--
		}
	}
}