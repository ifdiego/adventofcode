package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func partOne() int {
	canWin := 0
	numberOfWays := make([]int, 4)
	result := 1

	file, err := ioutil.ReadFile("input")
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
	return result
}

func partTwo() int {
	result := 0

	file, err := ioutil.ReadFile("input")
	if err != nil {
		log.Println("ERROR:", err)
		os.Exit(1)
	}
	t := strings.Split(string(file), "\n")[0]
	d := strings.Split(string(file), "\n")[1]
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

func main() {
	resultPartOne := partOne()
	fmt.Println("Part One:", resultPartOne)
	resultPartTwo := partTwo()
	fmt.Println("Part Two:", resultPartTwo)
}
