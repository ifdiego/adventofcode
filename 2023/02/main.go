package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	re := regexp.MustCompile(`(\d+) (\w+)`)

	part1, part2 := 0, 0
	for i, s := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		fewest := map[string]int{}

		for _, m := range re.FindAllStringSubmatch(s, -1) {
			n, _ := strconv.Atoi(m[1])
			fewest[m[2]] = max(fewest[m[2]], n)
		}

		if fewest["red"] <= 12 && fewest["green"] <= 13 && fewest["blue"] <= 14 {
			part1 += i + 1
		}
		part2 += fewest["red"] * fewest["green"] * fewest["blue"]
	}

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}
