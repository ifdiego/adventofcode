package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var left []int
	var right []int

	file, err := os.Open("day01/input.txt")
	if err != nil {
		log.Println("ERR:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		lvalue, _ := strconv.Atoi(fields[0])
		rvalue, _ := strconv.Atoi(fields[1])
		left = append(left, lvalue)
		right = append(right, rvalue)
	}

	distance := 0
	for i := range left {
		score := 0
		for j := range right {
			if left[i] == right[j] {
				score += 1
			}
		}
		distance += left[i] * score
	}

	fmt.Println(distance)
}
