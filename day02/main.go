package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

type GameInfo struct {
	id int
	draws []GameInfoDraw
	minRed, minGreen, minBlue int
}

type GameInfoDraw struct {
	red, green, blue int
}

const (
	maxRed = 12
	maxGreen = 13
	maxBlue = 14
)

func main() {
	rawGames := strings.Split(input, "\n")

	var games []GameInfo
	for _, rawGame := range rawGames {
		a := strings.Split(rawGame, ":")
		rawId := strings.Replace(a[0], "Game ", "", -1)
		id, _ := strconv.Atoi(rawId)
		gi := GameInfo{id: id}
		for _, rawDraw := range strings.Split(a[1], ";") {
			for _, rawGroup := range strings.Split(rawDraw, ",") {
				red, green, blue := 0, 0 ,0
				if strings.HasSuffix(rawGroup, "red") {
					red, _ = strconv.Atoi(strings.Replace(strings.Replace(rawGroup, "red", "", -1), " ", "", -1))
				}
				if strings.HasSuffix(rawGroup, "green") {
					green, _ = strconv.Atoi(strings.Replace(strings.Replace(rawGroup, "green", "", -1), " ", "", -1))
				}
				if strings.HasSuffix(rawGroup, "blue") {
					blue, _ = strconv.Atoi(strings.Replace(strings.Replace(rawGroup, "blue", "", -1), " ", "", -1))
				}
				gi.draws = append(gi.draws, GameInfoDraw{red: red, green: green, blue: blue})
			}
		}
		games = append(games, gi)
	}


	// sum := 0
	// for _, game := range games {
	// 	possible := true
	// 	for _, draw := range game.draws {
	// 		if draw.red > maxRed {
	// 			possible = false
	// 		}
	// 		if draw.green > maxGreen {
	// 			possible = false
	// 		}
	// 		if draw.blue > maxBlue {
	// 			possible = false
	// 		}
	// 	}
	// 	if possible {
	// 		sum = sum + game.id
	// 	}
	// }


	sum := 0
	for _, game := range games {
		for _, draw := range game.draws {
			if draw.red > game.minRed {
				game.minRed = draw.red
			}
			if draw.green > game.minGreen {
				game.minGreen = draw.green
			}
			if draw.blue > game.minBlue {
				game.minBlue = draw.blue
			}
		}
		fmt.Printf("%v\n", sum)
		sum = sum + game.minRed * game.minGreen * game.minBlue
	}

	fmt.Printf("%v\n", sum)
}