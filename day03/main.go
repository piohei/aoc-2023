package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

type Point struct {
	x, y int
}
type Found struct {
	val int
	gears []Point
}

func main() {
	fmt.Println(input)
	schema := strings.Split(input, "\n")

	var found []Found

	for i, row := range schema {
		numberStartIndex := -1
		for j := 0; j < len(row); j++ {
			if checkIfDigit(string(row[j])) {
				// fmt.Printf("found digit: %v\n", string(row[j]))
				if numberStartIndex < 0 {
					// fmt.Printf("found start at (%d, %d). vale = %s\n", i, j, string(row[j]))
					numberStartIndex = j
				}
			} else {
				// fmt.Printf("found nondigit: %v\n", string(row[j]))
				if numberStartIndex >= 0 {
					// fmt.Printf("found end at (%d, %d). vale = %s\n", i, j - 1, string(row[j - 1]))
					if res := lookForSymbolAround(schema, i, numberStartIndex, j - 1); res != nil {
						fmt.Printf("adding: %v\n", row[numberStartIndex:j])
						val, _ := strconv.Atoi(row[numberStartIndex:j])
						found = append(found, Found{val: val, gears: res})
					}
					numberStartIndex = -1
				}
			}
		}

		if numberStartIndex > 0 {
			if res := lookForSymbolAround(schema, i, numberStartIndex, len(row) - 1); res != nil {
				fmt.Printf("adding: %v\n", row[numberStartIndex:len(row)])
				val, _ := strconv.Atoi(row[numberStartIndex:len(row)])
				found = append(found, Found{val: val, gears: res})
			}
			numberStartIndex = -1
		}
	}

	// sum := 0
	l := make(map[Point][]int)
	for _, f := range found {
		for _, gear := range f.gears {
			l[gear] = append(l[gear], f.val)
		}
	}

	var sum int64 = 0
	for _, v := range l {
		if len(v) == 2 {
			sum = sum + int64(v[0]) * int64(v[1])
		}
	}

	fmt.Printf("%v\n", sum)
}

func lookForSymbolAround(schema []string, row, start, end int) []Point {
	fmt.Printf("checking for %v:\n", schema[row][start:end + 1])

	var res []Point

	s := start - 1
	if s < 0 {
		s = 0
	} else {
		fmt.Printf("looking at (%v, %v) %v\n", row, s, string(schema[row][s]))
		if checkIfSymbol(string(schema[row][s])) {
			res = append(res, Point{x: row, y: s})
		}
	}
	e := end + 1
	if e > len(schema[row]) - 1 {
		e = len(schema[row]) - 1
	} else {
		fmt.Printf("looking at (%v, %v) %v\n", row, e, string(schema[row][e]))
		if checkIfSymbol(string(schema[row][e])) {
			res = append(res, Point{x: row, y: e})
		}
	}


	prevRow := row - 1
	if prevRow >= 0 {
		for i := s; i <= e; i++ {
			fmt.Printf("looking at (%v, %v) %v\n", prevRow, i, string(schema[prevRow][i]))
			if checkIfSymbol(string(schema[prevRow][i])) {
				res = append(res, Point{x: prevRow, y: i})
			}
		}
	}
	nextRow := row + 1
	if nextRow < len(schema) {
		for i := s; i <= e; i++ {
			fmt.Printf("looking at (%v, %v) %v\n", nextRow, i, string(schema[nextRow][i]))
			if checkIfSymbol(string(schema[nextRow][i])) {
				res = append(res, Point{x: nextRow, y: i})
			}
		}
	}

	fmt.Printf("found gears at %v:\n", res)

	return res
}

func checkIfSymbol(c string) bool {
	// return c != "." && c != "0" && c != "1" && c != "2" && c != "3" && c != "4" && c != "5" && c != "6" && c != "7" && c != "8" && c != "9"
	return c == "*"
}

func checkIfDigit(c string) bool {
	return c == "0" || c == "1" || c == "2" || c == "3" || c == "4" || c == "5" || c == "6" || c == "7" || c == "8" || c == "9"
}