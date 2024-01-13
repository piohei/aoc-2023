package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

type Case struct {
	rows []string
	groups []int
}

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")
	var cases []Case

	for _, line := range inputAsLines {
		a := strings.Split(line, " ")

		c := Case{}

		for _, char := range a[0] {
			c.rows = append(c.rows, string(char))
		}
		for _, rawNumber := range strings.Split(a[1], ",") {
			val, _ := strconv.Atoi(rawNumber)
			c.groups = append(c.groups, val)
		}

		cases = append(cases, c)
	}

	fmt.Printf("%v\n", cases)

	sum := int64(0)
	for _, c := range cases {
		v, options := calculatePossibleOptions(c.rows, c.groups, "", 0)
		fmt.Printf("c = %v, options = %v\n%v\n", c, v, options)
		uniqOptions := uniq(options)
		if len(options) != len(uniqOptions) || len(options) != v {
			fmt.Printf("zzzzzzz----------------------------------------------------------------------\n")
			fmt.Printf("c = %v, options = %v\n%v\n", c, v, options)
			fmt.Printf("-----------------------------------------------------------------------------\n")
		}
		for _, opt := range options {
			if !validate(c.rows, strings.Split(opt, ""), c.groups) {
				fmt.Printf("error %v %v\n", opt, c.groups)
			}
		}
		sum = sum + int64(v)
	}
	fmt.Printf("sum = %v\n", sum)
}

func calculatePossibleOptions(rows []string, groups []int, prefix string, depth int) (int, []string) {
	debug := false
	for i := 0; i < depth; i++ {
		if debug {
			fmt.Printf("  ")
		}
	}
	if debug {
		fmt.Printf("calculatePossibleOptions(%v, %v, %v) -> ", rows, groups, prefix)
	}
	if len(rows) == 0 {
		if len(groups) > 0 {
			if debug {
				fmt.Printf("[1] return 0\n")
			}
			return 0, []string{}
		}

		if debug {
			fmt.Printf("[2] return 1\n")
		}
		return 1, []string{prefix}
	}
	if len(groups) == 0 {
		option := prefix
		for i := 0; i < len(rows); i++ {
			if rows[i] != "." && rows[i] != "?" {
				if debug {
					fmt.Printf("[3] return 0\n")
				}
				return 0, []string{}
			}
			option = option + "."
		}
		if debug {
			fmt.Printf("[4] return 1\n")
		}
		return 1, []string{option}
	}

	firstGroup := groups[0]
	switch {
	case rows[0] == "#":
		option := prefix
		for i := 0; i < firstGroup; i++ {
			if i >= len(rows) {
				if debug {
					fmt.Printf("[5] return 0\n")
				}
				return 0, []string{}
			}
			if rows[i] != "#" && rows[i] != "?" {
				if debug {
					fmt.Printf("[6] return 0\n")
				}
				return 0, []string{}
			}
			option = option + "#"
		}
		if firstGroup < len(rows) {
			if rows[firstGroup] != "." && rows[firstGroup] != "?" {
				if debug {
					fmt.Printf("[7] return 0\n")
				}
				return 0, []string{}
			}
			option = option + "."
		} else {
			if len(groups) != 1 {
				if debug {
					fmt.Printf("[8] return 0\n")
				}
				return 0, []string{}
			}
			if debug {
				fmt.Printf("[9] return 1\n")
			}
			return 1, []string{option}
		}
		if debug {
			fmt.Printf("[10] return calc\n")
		}
		return calculatePossibleOptions(rows[(firstGroup+1):], groups[1:], option, depth + 1)
	case rows[0] == "?":
		groupFitBeg := true

		option := prefix
		for i := 0; i < firstGroup; i++ {
			if i >= len(rows) {
				groupFitBeg = false
				break
			}
			if rows[i] != "#" && rows[i] != "?" {
				groupFitBeg = false
				break
			}
			option = option + "#"
		}
		if firstGroup < len(rows) {
			if rows[firstGroup] != "." && rows[firstGroup] != "?" {
				groupFitBeg = false
			}
			option = option + "."
		} else if firstGroup > len(rows) {
			if debug {
				fmt.Printf("[16] return 0\n")
			}
			return 0, []string{}
		} else {
			if len(groups) != 1 || !groupFitBeg{
				if debug {
					fmt.Printf("[11] return 0\n")
				}
				return 0, []string{}
			}
			if debug {
				fmt.Printf("[12] return 1\n")
			}
			return 1, []string{option}
		}

		if groupFitBeg {
			if debug {
				fmt.Printf("[13] return calc + calc\n")
			}
			res1, options1 := calculatePossibleOptions(rows[1:], groups, prefix + ".", depth + 1)
			res2, options2 := calculatePossibleOptions(rows[(firstGroup+1):], groups[1:], option, depth + 1)
			return res1 + res2, append(options1, options2...)
		} else {
			if debug {
				fmt.Printf("[14] return calc\n")
			}
			return calculatePossibleOptions(rows[1:], groups, prefix + ".", depth + 1)
		}
	case rows[0] == ".":
		if debug {
			fmt.Printf("[15] return calc\n")
		}
		return calculatePossibleOptions(rows[1:], groups, prefix + ".", depth + 1)
	}

	return 0, []string{}
}

func uniq(list []string) []string {
	var res []string
	for _, v1 := range list {
		found := false
		for _, v2 := range res {
			if v1 == v2 {
				found = true
				break
			}
		}
		if !found {
			res = append(res, v1)
		}
	}
	return res
}

func validate(original []string, option []string, groups []int) bool {
	if len(original) != len(option) {
		// fmt.Printf("length error, original = %v, option = %v\n", original, option)
		return false
	}

	for i := 0; i < len(original); i ++ {
		switch {
		case original[i] == "#" && option[i] != "#":
				// fmt.Printf("#1\n")
			return false
		case original[i] == "." && option[i] != ".":
				// fmt.Printf("#2\n")
			return false
		}
	}

	j := 0
	for i := 0; i < len(option); i++ {
		if option[i] == "." {
			continue
		}


		// fmt.Printf("found start at i= %v\n", i)
		for k := 0; k < groups[j]; k++ {
			if i + k >= len(option) {
				// fmt.Printf("#3\n")
				// fmt.Printf("i = %v, k = %v, len(option) = %v\n", i, k, len(option))
				return false
			}
			if option[i + k] != "#" {
				// fmt.Printf("#4\n")
				// fmt.Printf("i = %v, k = %v, len(option) = %v\n", i, k, len(option))
				return false
			}
		}

		i = i + groups[j]
		// fmt.Printf("moved to i= %v\n", i)
		j++

		if i >= len(option) {
			break
		}

		if option[i] != "." {
				// fmt.Printf("#5\n")
			return false
		}
	}

	if j < len(groups) {
				// fmt.Printf("#6\n")
		return false
	}

	return true
}