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


	lb := NewLensBox()
	for _, instr := range instrs {
		if strings.Contains(instr, "-") {
			label := strings.ReplaceAll(instr, "-", "")
			lb.remove(label)
		} else {
			splitted := strings.Split(instr, "=")
			label := splitted[0]
			focalLength, _ := strconv.Atoi(splitted[1])
			lb.put(label, focalLength)
		}
	}

	sum := 0
	for i, box := range lb.lenses {
		for j, lense := range box {
			sum += (i + 1) * (j + 1) * lense.focalLength
		}
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

func NewLensBox() *LensBox {
	return &LensBox{
		lenses: make([][]LabeledLens, 256),
	}
}

type LensBox struct {
	lenses [][]LabeledLens
}

type LabeledLens struct {
	label string
	focalLength int
}

func (lb *LensBox) remove(label string) {
	hash := hash(label)
	for i := 0; i < len(lb.lenses[hash]); i++ {
		if lb.lenses[hash][i].label == label {
			newLenses := lb.lenses[hash][:i]
			if i + 1 < len(lb.lenses[hash]) {
				newLenses = append(newLenses, lb.lenses[hash][i+1:]...)
			}
			lb.lenses[hash] = newLenses
		}
	}
}

func (lb *LensBox) put(label string, focalLength int) {
	hash := hash(label)
	for i := 0; i < len(lb.lenses[hash]); i++ {
		if lb.lenses[hash][i].label == label {
			lb.lenses[hash][i] = LabeledLens{label, focalLength}
			return
		}
	}
	lb.lenses[hash] = append(lb.lenses[hash], LabeledLens{label, focalLength})
}