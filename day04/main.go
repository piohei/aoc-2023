package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

type Card struct {
	id int
	winningNumbers []int
	numbers []int
	weight int
}

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	var cards []Card

	for i, line := range inputAsLines {
		a := strings.Split(line, ":")
		b := strings.Split(a[1], "|")
		card := Card{id: i + 1, weight: 1}
		for _, rawNumber := range strings.Split(b[0], " ") {
			if rawNumber == "" {
				continue
			}
			val, _ := strconv.Atoi(rawNumber)
			card.winningNumbers = append(card.winningNumbers, val)
		}
		for _, rawNumber := range strings.Split(b[1], " ") {
			if rawNumber == "" {
				continue
			}
			val, _ := strconv.Atoi(rawNumber)
			card.numbers = append(card.numbers, val)
		}
		cards = append(cards, card)
	}

	for index, card := range cards {
		found := 0
		for _, num := range card.numbers {
			if hasNumber(card.winningNumbers, num) {
				found++
			}
		}
		for i := 1; i <= found; i++ {
			if index + i < len(cards) {
				cards[index + i].weight = cards[index + i].weight + card.weight
			}
		}
	}

	sum := 0
	for _, card := range cards {
		sum = sum + card.weight
	}


	fmt.Printf("%v\n", cards)
	fmt.Printf("%v\n", sum)
}

func hasNumber(list []int, n int) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}