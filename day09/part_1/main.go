package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	var histories [][]int64

	for _, line := range inputAsLines {
		var history []int64
		for _, number := range strings.Split(line, " ") {
			val, _ := strconv.ParseInt(number, 10, 64)
			history = append(history, val)
		}
		histories = append(histories, history)
	}

	fmt.Printf("histories = %v\n", histories)

	sum := int64(0)
	for _, history := range histories {
		val := calcNext(history)
		fmt.Printf("history = %v, next = %v\n", history, val)
		sum = sum + val
	}

	fmt.Printf("sum = %v\n", sum)
}

func calcNext(in []int64) int64 {
	var rows [][]int64

	rows = append(rows, in)
	rowId := 0
	for {
		var next []int64
		for i := 1; i < len(rows[rowId]); i++ {
			next = append(next, rows[rowId][i] - rows[rowId][i - 1])
		}
		rows = append(rows, next)
		rowId++
		if allZero(rows[rowId]) {
			break
		}
	}

	fmt.Printf("rows = %v\n", rows)

	rows[rowId] = append(rows[rowId], 0)
	rowId--
	for rowId >= 0 {
		rows[rowId] = append(rows[rowId], rows[rowId][len(rows[rowId]) - 1] + rows[rowId + 1][len(rows[rowId + 1]) - 1])
		rowId--
	}

	fmt.Printf("rows = %v\n", rows)
	return rows[0][len(rows[0]) - 1]
}

func allZero(in []int64) bool {
	for _, v := range in {
		if v != 0 {
			return false
		}
	}
	return true
}