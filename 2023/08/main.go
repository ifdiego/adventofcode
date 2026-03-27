package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	split := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

	network := map[string]map[rune]string{}
	for _, m := range re.FindAllStringSubmatch(split[1], -1) {
		network[m[1]] = map[rune]string{'L': m[2], 'R': m[3]}
	}

	walk := func(start, end string) int {
		result := 1
		for n := range network {
			if !strings.HasSuffix(n, start) {
				continue
			}
			steps := 0
			for !strings.HasSuffix(n, end) {
				n = network[n][rune(split[0][steps%len(split[0])])]
				steps++
			}

			a, b := result, steps
			x, y := a, b
			for y != 0 {
				x, y = y, x%y
			}
			result = a * b / x
		}
		return result
	}

	fmt.Println("part one:", walk("AAA", "ZZZ"))
	fmt.Println("part two:", walk("A", "Z"))
}
