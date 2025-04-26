package days

import (
	"fmt"
	"strings"
)

var index int = 0
var copies = [194]int{}

func checkWinningGames(line string) int {
	points := 0
	matching := 0
	line1 := strings.Split(line, "|")[0]
	line2 := strings.Split(line, "|")[1]
	winningNumbers := strings.Fields(line1)
	numbersYouHave := strings.Fields(line2)
	for _, wn := range winningNumbers {
		for _, nh := range numbersYouHave {
			if wn == nh {
				matching += 1
				if points == 0 {
					points += 1
				} else {
					points = points * 2
				}
			}
		}
	}
	for j := 0; j < copies[index]+1; j++ {
		for i := index; i < index+matching; i++ {
			copies[i+1] = copies[i+1] + 1
		}
	}
	index += 1
	return points
}

func Day04(input string) {
	total := 0
	cards := 0

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		l := strings.Split(line, ":")[1]
		total += checkWinningGames(l)
	}

	for _, value := range copies {
		cards = cards + (value + 1)
	}

	fmt.Println("part one: ", total)
	fmt.Println("part two: ", cards)
}
