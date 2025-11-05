package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	canWin := 0
	numberOfWays := make([]int, 4)
	result := 1
	result2 := 0

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	t := strings.Split(string(file), "\n")[0]
	d := strings.Split(string(file), "\n")[1]
	regex := regexp.MustCompile("[0-9]+")
	times := regex.FindAllString(t, -1)
	distances := regex.FindAllString(d, -1)

	for index, value := range times {
		length, _ := strconv.Atoi(value)
		distance, _ := strconv.Atoi(distances[index])
		for i := 1; i <= length; i++ {
			milliseconds := length - i
			millimeter := milliseconds * i
			if millimeter > distance {
				canWin += 1
			}
		}
		numberOfWays[index] = canWin
		canWin = 0
	}

	for _, n := range numberOfWays {
		result = result * n
	}

	times2 := strings.Join(regex.FindAllString(t, -1), "")
	distances2 := strings.Join(regex.FindAllString(d, -1), "")

	length, _ := strconv.Atoi(times2)
	distance, _ := strconv.Atoi(distances2)
	for i := 1; i <= length; i++ {
		milliseconds := length - i
		millimeter := milliseconds * i
		if millimeter > distance {
			result2 += 1
		}
	}

	fmt.Println("part one: ", result)
	fmt.Println("part two: ", result2)
}
