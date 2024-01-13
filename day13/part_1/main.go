package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

type Case struct {
	rows []string
	cols []string
	rowMirror int
	colMirror int
	newRowMirror int
	newColMirror int
}

const (
	mirrorDebug = true
)

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")
	var cases []*Case

	c := &Case{}
	for _, line := range inputAsLines {
		if line == "" {
			cases = append(cases, c)
			c = &Case{}
			continue
		}

		if c.rows == nil {
			for i := 0; i < len(line); i++ {
				c.cols = append(c.cols, "")
			}
		}

		c.rows = append(c.rows, line)

		for i, letter := range strings.Split(line, "") {
			c.cols[i] = c.cols[i] + letter
		}

	}

	cases = append(cases, c)

	sum := int64(0)
	for _, c := range cases {
		c.LookForMirroredLine()
		c.Print()

		if c.rowMirror > 0 {
			sum = sum + int64(100 * c.rowMirror)
		}

		if c.colMirror > 0 {
			sum = sum + int64(c.colMirror)
		}
	}

	fmt.Printf("sum = %v\n", sum)
}

func (c *Case) Print() {
	fmt.Printf("rows = %v, cols = %v, rowMirror = %v, colMirror = %v\n", c.rows, c.cols, c.rowMirror, c.colMirror)
}

func (c *Case) LookForMirroredLine() {
	var l int
	c.rowMirror = -1
	c.colMirror = -1

	l = len(c.rows)
	for i := 0; i < l; i++ {
		if i < l - i {
			if c.IsMirror(c.rows[:2*i]) {
				c.rowMirror = i
				return
			}
		} else {
			if c.IsMirror(c.rows[i-(l-i):]) {
				c.rowMirror = i
				return
			}
		}
	}

	l = len(c.cols)	
	for i := 0; i < l; i++ {
		if i < l - i {
			if c.IsMirror(c.cols[:2*i]) {
				c.colMirror = i
				return
			}
		} else {
			if c.IsMirror(c.cols[i-(l-i):]) {
				c.colMirror = i
				return
			}
		}
	}
}

func (c *Case) IsMirror(rows []string) bool {
	if mirrorDebug {
		fmt.Printf("IsMirror(%v)\n", rows)
	}

	if len(rows) == 0 {
		if mirrorDebug {
			fmt.Printf("IsMirror - false\n")
		}
		return false
	}

	i := 0
	j := len(rows) - 1

	for i < j {
		if rows[i] != rows[j] {
			if mirrorDebug {
				fmt.Printf("IsMirror - false\n")
			}
			return false
		}
		i++
		j--
	}

	if i == j {
		if mirrorDebug {
			fmt.Printf("IsMirror - false\n")
		}
		return false
	}

	if mirrorDebug {
		fmt.Printf("IsMirror - true\n")
	}
	return true
}