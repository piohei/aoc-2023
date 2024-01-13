package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed test
var input string

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")


	var time []int64
	var distance []int64

	for _, line := range inputAsLines {
		if strings.HasPrefix(line, "Time: ") {
			s := strings.Replace(line, "Time: ", "", -1)
			rawNumber := strings.Replace(s, " ", "", -1)
			if rawNumber == "" {
				continue
			}
			number, _ := strconv.ParseInt(rawNumber, 10, 64)
			time = append(time, number)
		}
		if strings.HasPrefix(line, "Distance: ") {
			s := strings.Replace(line, "Distance: ", "", -1)
			rawNumber := strings.Replace(s, " ", "", -1)
			if rawNumber == "" {
				continue
			}
			number, _ := strconv.ParseInt(rawNumber, 10, 64)
			distance = append(distance, number)
		}
	}

	fmt.Printf("time = %v\n", time)
	fmt.Printf("distance = %v\n", distance)

	res := int64(1)
	for i := 0; i < len(time); i++ {
		t := time[i]
		d := distance[i]

		minWorking := binSerachMinWorking(t, d)
		fmt.Printf("minWorking %v\n", minWorking)
		maxWorking := binSerachMaxWorking(t, d)
		fmt.Printf("maxWorking %v\n", maxWorking)

		res = res * (maxWorking - minWorking + 1)
	}

	fmt.Printf("res = %v\n", res)
}

func binSerachMinWorking(time, distance int64) int64 {
	p, q := int64(0), time
	for ; p < q; {
		i := (p + q) / 2
		// fmt.Printf("p = %v, q = %v, i = %v\n", p, q, i)
		d := calcDist(i, time)
		// fmt.Printf("i = %v, time = %v, dist = %v\n", i, time, d)
		if d <= distance {
			if p == i {
				p = i + 1
			} else {
				p = i
			}
		} else {
			q = i
		}
	}
	return p
}

func binSerachMaxWorking(time, distance int64) int64 {
	p, q := int64(0), time
	for ; p < q; {
		i := (p + q + 1) / 2
		// fmt.Printf("p = %v, q = %v, i = %v\n", p, q, i)
		d := calcDist(i, time)
		// fmt.Printf("i = %v, time = %v, dist = %v\n", i, time, d)
		if d > distance {
			p = i
		} else {
			if q == i {
				q = i - 1
			} else {
				q = i
			}
		}
	}
	return p
}

func calcDist(time, totalTime int64) int64 {
	return (totalTime - time) * time
}
