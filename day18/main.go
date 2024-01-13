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

var xs []int64
var ys []int64

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	cX, cY := int64(0), int64(0)
	lengthSum := int64(0)
	for _, line := range inputAsLines {
		s := strings.Split(line, " ")

		color := strings.ReplaceAll(strings.ReplaceAll(s[2], "(", ""), ")", "")
		length, _ := strconv.ParseInt(string(color[1:6]), 16, 64)
		var direction string
		switch string(color[6]) {
		case "3":
			direction = "U"
		case "1":
			direction = "D"
		case "2":
			direction = "L"
		case "0":
			direction = "R"
		}

		fmt.Printf("%v = %v %v\n", color, direction, length)

		var nX, nY int64
		switch string(color[6]) {
		case "3": // "U":
			nX, nY = cX, cY - length
		case "1": // "D":
			nX, nY = cX, cY + length
		case "2": // "L":
			nX, nY = cX - length, cY
		case "0": // "R":
			nX, nY = cX + length, cY
		}

		lengthSum += length
		xs = append(xs, nX)
		ys = append(ys, nY)
		cX, cY = nX, nY
	}

	fmt.Printf("xs = %v\nys = %v\n", xs, ys)

	res := area() + lengthSum / 2 + 1

	fmt.Printf("res = %v\n", res)
}

func area() int64 {
	area := int64(0)
	for i := 1; i < len(xs); i++ {
		j := i - 1
		area += (xs[i] + xs[j]) * (ys[i] - ys[j])
	}
	return area / 2
}