package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
    "sort"
)

//go:embed input
var input string

type Hand struct {
	cards string
	bid int64
	rank int
}

var letterToNumber map[byte]int = map[byte]int{
	byte('J'): 1,
	byte('2'): 2,
	byte('3'): 3,
	byte('4'): 4,
	byte('5'): 5,
	byte('6'): 6,
	byte('7'): 7,
	byte('8'): 8,
	byte('9'): 9,
	byte('T'): 10,
	byte('Q'): 12,
	byte('K'): 13,
	byte('A'): 14,
}


func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	var hands []*Hand

	for _, line := range inputAsLines {
		s := strings.Split(line, " ")
		bid, _ := strconv.ParseInt(s[1], 10, 64)
		hand := &Hand{cards: s[0], bid: bid}
		rankHand(hand)
		hands = append(hands, hand)
	}

	printHands(hands)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rank < hands[j].rank {
			return true
		}
		if hands[i].rank > hands[j].rank {
			return false
		}


		for idx := 0; idx<5; idx++ {
			if letterToNumber[hands[i].cards[idx]] < letterToNumber[hands[j].cards[idx]] {
				return true
			}
			if letterToNumber[hands[i].cards[idx]] > letterToNumber[hands[j].cards[idx]] {
				return false
			}
		}
		return false
	})

	printHands(hands)

	res := int64(0)
	for k, v := range hands {
		res += int64(k + 1) * v.bid
	}
	fmt.Printf("res = %v\n", res)
}

func rankHand(hand *Hand) {
	counts := initMap()
	for _, c := range hand.cards {
		counts[byte(c)] = counts[byte(c)] + 1
	}

	five := 0
	four := 0
	three := 0
	two := 0
	jokers := 0
	for k, v := range counts {
		if k == byte('J') {
			jokers = v
			continue
		} 
		if v == 5 {
			five = five + 1
		}
		if v == 4 {
			four = four + 1
		}
		if v == 3 {
			three = three + 1
		}
		if v == 2 {
			two = two + 1
		}
	}

	fmt.Printf("hand = %v, five = %v, four = %v, three = %v, two = %v, jokers = %v\n", *hand, five, four, three, two, jokers)

	switch {
	case five == 1:
		hand.rank = 6
	case four == 1:
		if jokers > 0 {
			hand.rank = 6
		} else {
			hand.rank = 5
		}
	case three == 1 && two == 1:
		hand.rank = 4
	case three == 1:
		if jokers == 2 {
			hand.rank = 6
		} else if jokers == 1 {
			hand.rank = 5
		} else {
			hand.rank = 3
		}
	case two == 2:
		if jokers == 1 {
			hand.rank = 4
		} else {
			hand.rank = 2
		}
	case two == 1:
		if jokers == 3 {
			hand.rank = 6
		} else if jokers == 2 {
			hand.rank = 5
		} else if jokers == 1 {
			hand.rank = 3
		} else {
			hand.rank = 1
		}
	default:
		if jokers == 5 {
			hand.rank = 6
		} else if jokers == 4 {
			hand.rank = 6
		} else if jokers == 3 {
			hand.rank = 5
		} else if jokers == 2 {
			hand.rank = 3
		} else if jokers == 1 {
			hand.rank = 1
		} else {
			hand.rank = 0
		}
	}

	fmt.Printf("hand = %v, counts = %v, jokers = %v\n", *hand, counts, jokers)
}

func allSame(str string) bool {
	if len(str) == 0 {
		return true
	}
	if len(str) == 1 {
		return true
	}
	for i := 1; i < len(str) - 1; i++ {
		if str[0] != str[i] {
			return false
		}
	}
	return true
}

func initMap() map[byte]int {
	res := make(map[byte]int)
	for _, c := range "23456789TJQKA" {
		res[byte(c)] = 0
	}
	return res
}

func printHands(hands []*Hand) {
	fmt.Printf("[")
	for _, v := range hands {
		fmt.Printf("%v, ", *v)
	}
	fmt.Printf("]\n")
}