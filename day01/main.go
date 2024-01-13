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
	res := strings.Split(input, "\n")
	sum := 0

	for _, origLine := range res {
		line := origLine

		i, j := 0, len(line) - 1
		var val1, val2 int
		var err error

		for i <= j {
			val1, err = strconv.Atoi(string(line[i]))
			if err == nil && val1 >= 0 && val1 <= 9 {
				break;
			}
			if strings.HasPrefix(line[i:], "one") {
				val1 = 1
				break;
			}
			if strings.HasPrefix(line[i:], "two") {
				val1 = 2
				break;
			}
			if strings.HasPrefix(line[i:], "three") {
				val1 = 3
				break;
			}
			if strings.HasPrefix(line[i:], "four") {
				val1 = 4
				break;
			}
			if strings.HasPrefix(line[i:], "five") {
				val1 = 5
				break;
			}
			if strings.HasPrefix(line[i:], "six") {
				val1 = 6
				break;
			}
			if strings.HasPrefix(line[i:], "seven") {
				val1 = 7
				break;
			}
			if strings.HasPrefix(line[i:], "eight") {
				val1 = 8
				break;
			}
			if strings.HasPrefix(line[i:], "nine") {
				val1 = 9
				break;
			}
			i++;
		}
		for i <= j {
			val2, err = strconv.Atoi(string(line[j]))
			if err == nil && val2 >= 0 && val2 <= 9 {
				break;
			}
			if strings.HasPrefix(line[j:], "one") {
				val2 = 1
				break;
			}
			if strings.HasPrefix(line[j:], "two") {
				val2 = 2
				break;
			}
			if strings.HasPrefix(line[j:], "three") {
				val2 = 3
				break;
			}
			if strings.HasPrefix(line[j:], "four") {
				val2 = 4
				break;
			}
			if strings.HasPrefix(line[j:], "five") {
				val2 = 5
				break;
			}
			if strings.HasPrefix(line[j:], "six") {
				val2 = 6
				break;
			}
			if strings.HasPrefix(line[j:], "seven") {
				val2 = 7
				break;
			}
			if strings.HasPrefix(line[j:], "eight") {
				val2 = 8
				break;
			}
			if strings.HasPrefix(line[j:], "nine") {
				val2 = 9
				break;
			}
			j--;
		}

		fmt.Printf("%s, %s, %d, %d\n", origLine, line, val1, val2)

		sum = sum + 10 * val1 + val2
	}
	fmt.Printf("%d\n", sum)
}