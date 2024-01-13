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

var instrs []string

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	for _, line := range inputAsLines {
		if line == "" {
			continue
		}
		for _, instr := range strings.Split(line, ",") {
			instrs = append(instrs, instr)
		}
	}

	sum := 0
	for _, instr := range instrs {
		sum += hash(instr)
	}
	fmt.Printf("sum = %v\n", sum)
}

func hash(s string) int {
	res := 0
	for _, c := range s {
		res += int(c)
		res *= 17
		res %= 256
	}
	return res
}