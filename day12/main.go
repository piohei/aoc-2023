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
	rows Rows
	groups Groups

	cache map[string]map[string]int64
}

type Rows []string
type Groups []int

const (
	cacheDebug = false
)

func main() {
	// fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")
	var cases []*Case
	var extendedCasees []*Case

	for _, line := range inputAsLines {
		a := strings.Split(line, " ")

		c := &Case{}

		for _, char := range a[0] {
			c.rows = append(c.rows, string(char))
		}
		for _, rawNumber := range strings.Split(a[1], ",") {
			val, _ := strconv.Atoi(rawNumber)
			c.groups = append(c.groups, val)
		}

		cases = append(cases, c)
	}

	for _, c := range cases {
		newRows := c.rows[:]
		newGroups := c.groups[:]
		for i := 0; i < 4; i++ {
			newRows = append(newRows, "?")
			newRows = append(newRows, c.rows...)
			newGroups = append(newGroups, c.groups...)
		}
		extendedCasees = append(extendedCasees, &Case{rows: newRows, groups: newGroups})
	}

	sum := int64(0)
	for _, c := range extendedCasees {
		v := c.calculate()
		// fmt.Printf("Case(rows = %v, groups = %v) -> options %v\n", c.rows, c.groups, v)
		sum += v
	}

	fmt.Printf("sum = %v\n", sum)

}

func (c *Case) calculate() int64 {
	return c.calc(c.rows, c.groups)
}

func (c *Case) calc(rows Rows, g Groups) int64 {
	if v, ok := c.cacheGet(rows, g); ok {
		if cacheDebug {
			fmt.Printf("found in cache\n")
		}
		return v
	}

	if len(g) == 0 && len(rows) == 0 {
		c.cachePut(rows, g, 1)
		return 1
	}

	if len(g) == 0 {
		for i := 0; i < len(rows); i++ {
			if rows[i] == "#" {
				c.cachePut(rows, g, 0)
				return 0
			}
		}

		c.cachePut(rows, g, 1)
		return 1
	}

	if len(rows) == 0 {
		c.cachePut(rows, g, 0)
		return 0 // as len(groups) > 0 so there are unassigned groups
	}

	res := int64(0)
	if rows[0] == "#" || rows[0] == "?" {
		group := g[0]
		if group > len(rows) {
			// c.cachePut(rows, g, 0)
			// return 0 // won't fit
			goto next
		}

		for i := 0; i < group; i ++ {
			if rows[i] != "#" && rows[i] != "?" {
				// c.cachePut(rows, g, 0)
				// return 0
				goto next
			}
		}

		if group == len(rows) {
			// c.cachePut(rows, g, 1)
			// return 1
			if len(g) == 1 {
				res = res + 1
			}
			goto next
		}

		if group < len(rows) {
			if rows[group] != "." && rows[group] != "?" {
				// c.cachePut(rows, g, 0)
				// return 0
				goto next
			}
		}

		res = res + c.calc(Rows(rows[group + 1:]), Groups(g[1:]))
	}

next:
	if rows[0] == "." || rows[0] == "?" {
		res = res + c.calc(Rows(rows[1:]), g)
	}

	c.cachePut(rows, g, res)
	return res
}

func (c *Case) cachePut(r Rows, g Groups, v int64) {
	if cacheDebug {
		fmt.Printf("cachePut(rows = %v, groups = %v, options = %v)\n", r, g, v)
	}
	if c.cache == nil {
		c.cache = make(map[string]map[string]int64)
	}
	if _, ok := c.cache[r.Key()]; !ok {
		c.cache[r.Key()] = make(map[string]int64)
	}
	c.cache[r.Key()][g.Key()] = v
}

func (c *Case) cacheGet(r Rows, g Groups) (int64, bool) {
	if c.cache == nil {
		return -1, false
	}
	if _, ok := c.cache[r.Key()]; !ok {
		return -1, false
	}
	v, ok := c.cache[r.Key()][g.Key()]
	return v, ok
}

func (r Rows) Key() string {
	return strings.Join(r, "")
}

func (r Rows) String() string {
	return "[" + strings.Join(r, ", ") + "]"
}

func (g Groups) Key() string {
	mapper := []string {
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f",
	}
	str := ""
	for _, v := range g {
		str = str + mapper[v]
	}
	return str
}

func (g Groups) String() string {
	var res []string
	for _, v := range g {
		res = append(res, strconv.Itoa(v))
	}
	return "[" + strings.Join(res, ", ") + "]"
}