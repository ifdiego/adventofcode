package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var fewest = map[string]int{
	"red":   0,
	"green": 0,
	"blue":  0,
}

func getPower() int {
	result := 1
	for _, value := range fewest {
		result = result * value
	}
	return result
}

func setPower(number int, color string) {
	if number > fewest[color] {
		fewest[color] = number
	}
}

func isPossible(number int, color string) bool {
	if color == "red" && number > 12 {
		return false
	}
	if color == "green" && number > 13 {
		return false
	}
	if color == "blue" && number > 14 {
		return false
	}
	return true
}

func splitData(line string) int {
	isGame := true
	id := strings.Split(line, ":")[0]
	id = strings.Split(id, "Game ")[1]

	entries := strings.Split(line, ":")[1]
	games := strings.Split(entries, ";")
	for _, game := range games {
		rounds := strings.Split(game, ",")
		for _, round := range rounds {
			sets := strings.TrimSpace(round)
			set := strings.Split(sets, " ")
			n, _ := strconv.Atoi(set[0])
			possible := isPossible(n, set[1])
			if !possible {
				isGame = false
			}
			setPower(n, set[1])
		}
	}

	if isGame {
		r, _ := strconv.Atoi(id)
		return r
	} else {
		return 0
	}
}

func main() {
	total := 0
	power := 0

	file, err := os.Open("day02/input.txt")
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total = total + splitData(scanner.Text())
		power = power + getPower()
		fewest["red"] = 0
		fewest["green"] = 0
		fewest["blue"] = 0
	}

	fmt.Println(total)
	fmt.Println(power)
}
