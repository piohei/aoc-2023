package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

const (
	debug = true

	MAX_INT = 2147483647
)

var wiringMap map[string][]string
var edges [][]string

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	wiringMap = make(map[string][]string)

	for _, line := range inputAsLines {
		l := strings.Split(line, ":")
		n1 := strings.TrimSpace(l[0])
		for _, n2 := range strings.Split(l[1], " ") {
			if n2 == "" {
				continue
			}
			wiringMap[n1] = append(wiringMap[n1], n2)
			wiringMap[n2] = append(wiringMap[n2], n1)
			edges = append(edges, []string{n1, n2})
		}
	}
	fmt.Printf("wiringMap = %v\n", wiringMap)

	res := -1

	for i := 0; i < len(edges); i++ {
		fmt.Printf("i: %v/%v\n", i + 1, len(edges))
		for j := i + 1; j < len(edges); j++ {
			fmt.Printf("j: %v/%v\n", j + 1, len(edges))
			for k := j + 1; k < len(edges); k++ {
				fmt.Printf("k: %v/%v\n", k + 1, len(edges))
				e1 := edges[i]
				e2 := edges[j]
				e3 := edges[k]

				wm := copyWiringMap()
				wm[e1[0]] = remove(wm[e1[0]], e1[1])
				wm[e1[1]] = remove(wm[e1[1]], e1[0])
				wm[e2[0]] = remove(wm[e2[0]], e2[1])
				wm[e2[1]] = remove(wm[e2[1]], e2[0])
				wm[e3[0]] = remove(wm[e3[0]], e3[1])
				wm[e3[1]] = remove(wm[e3[1]], e3[0])

				groups := defineConnectedGraphs(wm)

				maxGroup := -1
				for _, v := range groups {
					if v > maxGroup {
						maxGroup = v
					}
				}

				if maxGroup == 1 {
					countZero := 0
					countOne := 0
					for _, v := range groups {
						if v == 0 {
							countZero++
						}
						if v == 1 {
							countOne++
						}
					}

					res = countZero * countOne
				}
			}
		}
	}

	fmt.Printf("res = %v\n", res)
}

func copyWiringMap() map[string][]string {
	wiringMapCopy := make(map[string][]string)

	for k, v := range wiringMap {
		for _, lv := range v {
			wiringMapCopy[k] = append(wiringMapCopy[k], lv)
		}
	}

	return wiringMapCopy
}

func defineConnectedGraphs(wp map[string][]string) map[string]int {
	var groups map[string]int
	groups = make(map[string]int)

	for k, _ := range wp {
		groups[k] = -1
	}

	group := 0
	for k, _ := range wp {
		if groups[k] >= 0 {
			continue
		}

		setGroup(wp, k, group, groups)
		group++
	}

	return groups
}

func setGroup(wp map[string][]string, n string, group int, groups map[string]int) {
	groups[n] = group
	for _, n2 := range wp[n] {
		if groups[n2] < 0 {
			setGroup(wp, n2, group, groups)
		}
	}
}

func remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
		    return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}