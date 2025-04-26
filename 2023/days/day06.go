package days

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func calcPartOne(line string) int {
	canWin := 0
	numberOfWays := make([]int, 4)
	result := 1

	t := strings.Split(string(line), "\n")[0]
	d := strings.Split(string(line), "\n")[1]
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
		result *= n
	}

	return result
}

func calcPartTwo(line string) int {
	result := 0

	t := strings.Split(string(line), "\n")[0]
	d := strings.Split(string(line), "\n")[1]
	regex := regexp.MustCompile("[0-9]+")
	times := strings.Join(regex.FindAllString(t, -1), "")
	distances := strings.Join(regex.FindAllString(d, -1), "")

	length, _ := strconv.Atoi(times)
	distance, _ := strconv.Atoi(distances)
	for i := 1; i <= length; i++ {
		milliseconds := length - i
		millimeter := milliseconds * i
		if millimeter > distance {
			result += 1
		}
	}

	return result
}

func Day06(input string) {
	partOne := calcPartOne(input)
	fmt.Printf("part one: %v\n", partOne)

	partTwo := calcPartTwo(input)
	fmt.Printf("part two: %v\n", partTwo)
}
